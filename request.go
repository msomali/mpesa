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

package mpesa

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

const (
	GenerateSessionKey RequestType = iota
	PushPaySync
	PushPayAsync
	DisbursementSync
	DisbursementAsync
)

type RequestType int

type Endpoints struct {
	SessionID string
	B2C       string
	C2B       string
}

type Request struct {
	Method      string
	Type        RequestType
	Endpoint    string
	Payload     interface{}
	Headers     map[string]string
	QueryParams map[string]string
}

func (t RequestType) HttpMethod() string {
	switch t {
	case GenerateSessionKey:
		return http.MethodGet

	case PushPayAsync,PushPaySync,DisbursementAsync,DisbursementSync:
		return http.MethodPost
	}

	return ""
}

func (t RequestType) name() string {
	values := map[int]string{
		0: "session key",
		1: "c2b single stage",
		2: "b2c single stage",
		3: "b2b single stage",
		4: "reversal",
		5: "query transaction status",
		6: "direct debit create",
		7: "direct debit payment",
	}

	return values[int(t)]
}

func (r *Request) newRequestWithContext(ctx context.Context, market *Market, platform Platform) (*http.Request, error) {

	var buffer io.Reader

	url := generateRequestURL(defBasePath, platform, *market, r.Endpoint)
	if r.Payload != nil {
		buf, err := json.Marshal(r.Payload)
		if err != nil {
			return nil, err
		}

		buffer = bytes.NewBuffer(buf)
		req, err := http.NewRequestWithContext(ctx, r.Method, url, buffer)
		if err != nil {
			return nil, err
		}
		for key, value := range r.Headers {
			req.Header.Add(key, value)
		}
		return req, nil
	}

	req, err := http.NewRequestWithContext(ctx, r.Method, url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range r.Headers {
		req.Header.Add(key, value)
	}

	for name, value := range r.QueryParams {
		values := req.URL.Query()
		values.Add(name, value)
		req.URL.RawQuery = values.Encode()
	}
	return req, nil
}
