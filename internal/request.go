/*
 * MIT License
 *
 * Copyright (c) 2021 TECHCRAFT TECHNOLOGIES CO LTD.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package internal

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

type (

	// Request encapsulate details of a request to be sent to server
	// Endpoint is dynamic that is appended to URL
	// e.g if the url is www.server.com/users/user-id, user-id is the endpoint
	Request struct {
		Name        string
		Method      string
		URL         string
		Endpoint    string
		PayloadType PayloadType
		Payload     interface{}
		Headers     map[string]string
		QueryParams map[string]string
	}

	RequestOption func(request *Request)
)

func NewRequest(method, url string, payload interface{}, opts ...RequestOption) *Request {
	var (
		defaultRequestHeaders = map[string]string{
			"Content-Type": cTypeJson,
		}
	)

	request := &Request{
		Method:      method,
		URL:         url,
		PayloadType: JsonPayload,
		Endpoint:    "",
		Payload:     payload,
		Headers:     defaultRequestHeaders,
	}

	for _, opt := range opts {
		opt(request)
	}

	return request
}

func WithQueryParams(params map[string]string) RequestOption {
	return func(request *Request) {
		request.QueryParams = params
	}
}

func WithEndpoint(endpoint string) RequestOption {
	return func(request *Request) {
		request.Endpoint = endpoint
	}
}

// WithRequestHeaders replaces all the available headers with new ones
// WithMoreHeaders appends headers does not replace them
func WithRequestHeaders(headers map[string]string) RequestOption {
	return func(request *Request) {
		request.Headers = headers
	}
}

// WithMoreHeaders appends headers does not replace them like WithRequestHeaders
func WithMoreHeaders(headers map[string]string) RequestOption {
	return func(request *Request) {
		for key, value := range headers {
			request.Headers[key] = value
		}
	}
}

// See 2 (end of page 4) https://www.ietf.org/rfc/rfc2617.txt
// "To receive authorization, the client sends the userid and password,
// separated by a single colon (":") character, within a base64
// encoded string in the credentials."
// It is not meant to be urlencoded.
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// WithBasicAuth add password and username to request headers
func WithBasicAuth(username, password string) RequestOption {
	return func(request *Request) {
		request.Headers["Token"] = "Basic " + basicAuth(username, password)
	}
}

func (request *Request) AddHeader(key, value string) {
	request.Headers[key] = value
}

func appendEndpoint(url string, endpoint string) string {
	url, endpoint = strings.TrimSpace(url), strings.TrimSpace(endpoint)
	urlHasSuffix, endpointHasPrefix := strings.HasSuffix(url, "/"), strings.HasPrefix(endpoint, "/")

	bothTrue := urlHasSuffix == true && endpointHasPrefix == true
	bothFalse := urlHasSuffix == false && endpointHasPrefix == false
	notEqual := urlHasSuffix != endpointHasPrefix

	if notEqual {
		return fmt.Sprintf("%s%s", url, endpoint)
	}

	if bothFalse {
		return fmt.Sprintf("%s/%s", url, endpoint)
	}

	if bothTrue {
		endp := strings.TrimPrefix(endpoint, "/")
		return fmt.Sprintf("%s%s", url, endp)
	}

	return ""
}

// NewRequestWithContext takes a *Request and transform into *http.Request with a context
func NewRequestWithContext(ctx context.Context, request *Request) (req *http.Request, err error) {
	requestURL := request.URL
	requestEndpoint := request.Endpoint
	if requestEndpoint != "" {
		request.URL = appendEndpoint(requestURL, requestEndpoint)
	}
	if request.Payload == nil {
		req, err = http.NewRequestWithContext(ctx, request.Method, request.URL, nil)
		if err != nil {
			return nil, err
		}
	} else {
		buffer, err := MarshalPayload(request.PayloadType, request.Payload)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequestWithContext(ctx, request.Method, request.URL, buffer)
		if err != nil {
			return nil, err
		}
	}

	for key, value := range request.Headers {
		req.Header.Add(key, value)
	}

	for name, value := range request.QueryParams {
		values := req.URL.Query()
		values.Add(name, value)
		req.URL.RawQuery = values.Encode()
	}

	return req, nil
}
