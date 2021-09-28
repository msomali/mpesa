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
	"errors"
	"fmt"
	"github.com/techcraftlabs/mpesa/internal"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)


var _ Service = (*Client)(nil)
var defClientLogger = os.Stderr
var defHttpClient = http.DefaultClient

// Config contains details initialize in mpesa portal
// Applications require the following details:
//•	Application Name – human-readable name of the application
//•	Version – version number of the application, allowing changes in API products to be managed in different versions
//•	Description – Free text field to describe the use of the application
//•	APIKey – Unique authorisation key used to authenticate the application on the first call. API Keys need to be encrypted in the first “Generate Session API Call” to create a valid session key to be used as an access token for future calls. Encrypting the API Key is documented in the GENERATE SESSION API page
//•	SessionLifetime – The session key has a finite lifetime of availability that can be configured. Once a session key has expired, the session is no longer usable, and the caller will need to authenticate again.
//•	TrustedSources – the originating caller can be limited to specific IP address(es) as an additional security measure.
//•	Products / Scope / Limits – the required API products for the application can be enabled and limits defined for each call.
type Config struct {
	Name                   string
	Version                string
	Description            string
	BasePath               string
	APIKey                 string
	PublicKey              string
	SessionLifetimeMinutes int64
	ServiceProvideCode     string
	TrustedSources         []string
}

type Client struct {
	*Config
	Http              *internal.BaseClient
	DebugMode         bool
	Market            *Market
	Platform          Platform
	encryptedApiKey   *string
	sessionID         *string
	sessionExpiration time.Time
}

type ClientOpt func(client *Client)

func NewClient(config *Config, market *Market, platform Platform, opts ...ClientOpt) (*Client, error) {

	encryptedKeyStr, err := generateEncryptedKey(config.APIKey, config.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("could not generate encrypted api key: %w", err)
	}

	if config.SessionLifetimeMinutes <= 0 {
		return nil, fmt.Errorf("session lifetime (in minutes) can not be zero")
	}

	up := time.Duration(config.SessionLifetimeMinutes) * time.Minute

	expiration := time.Now().Add(up)

	client := &Client{
		Config:            config,
		Http:              defHttpClient,
		DebugMode:         false,
		Market:            market,
		Platform:          platform,
		encryptedApiKey:   &encryptedKeyStr,
		sessionID:         nil,
		sessionExpiration: expiration,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}

func WithDebugMode(debugMode bool) ClientOpt {
	return func(client *Client) {
		client.DebugMode = debugMode
	}
}

func WithHTTPClient(httpd *http.Client) ClientOpt {
	return func(client *Client) {
		client.Http = httpd
	}
}

func WithWriter(writer io.Writer) ClientOpt {
	return func(client *Client) {
		client.Logger = writer
	}
}

func (client *Client) send(ctx context.Context, request *Request, v interface{}) error {
	var req *http.Request
	var res *http.Response

	var reqBodyBytes []byte
	var resBodyBytes []byte
	defer func(debug bool) {
		if debug {
			req.Body = io.NopCloser(bytes.NewBuffer(reqBodyBytes))
			res.Body = io.NopCloser(bytes.NewBuffer(resBodyBytes))
			name := strings.ToUpper(request.Type.name())
			client.logOut(name, req, res)
		}
	}(client.DebugMode)

	//creates http request with context
	req, err := request.newRequestWithContext(ctx, client.Market, client.Platform)

	if err != nil {
		return err
	}

	if req.Body != nil {
		reqBodyBytes, _ = io.ReadAll(req.Body)
	}

	if v == nil {
		return errors.New("v interface can not be empty")
	}

	req.Body = io.NopCloser(bytes.NewBuffer(reqBodyBytes))
	res, err = client.Http.Do(req)

	if err != nil {
		return err
	}

	if res.Body != nil {
		resBodyBytes, _ = io.ReadAll(res.Body)
	}

	contentType := res.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		if err := json.NewDecoder(bytes.NewBuffer(resBodyBytes)).Decode(v); err != nil {
			if err != io.EOF {
				return err
			}
		}
	}

	return nil
}

func (client *Client) getEncryptionKey() (string, error) {
	isAvailable := client.encryptedApiKey != nil && *client.encryptedApiKey != ""
	//notExpired := client.sessionExpiration.Sub(time.Now()).Minutes() > 1
	if isAvailable {
		return *client.encryptedApiKey, nil
	}

	return generateEncryptedKey(client.APIKey, client.PublicKey)
}

// log is called to print the details of http.Request sent from Tigo during
// callback, namecheck or ussd payment. It is used for debugging purposes
func (client *Client) log(name string, request *http.Request) {

	if request != nil {
		reqDump, _ := httputil.DumpRequest(request, true)
		_, err := fmt.Fprintf(client.Logger, "%s REQUEST: %s\n", name, reqDump)
		if err != nil {
			fmt.Printf("error while logging %s request: %v\n",
				strings.ToLower(name), err)
			return
		}
		return
	}
	return
}

// logOut is like log except this is for outgoing client requests:
// http.Request that is supposed to be sent to tigo
func (client *Client) logOut(name string, request *http.Request, response *http.Response) {

	if request != nil {
		reqDump, _ := httputil.DumpRequestOut(request, true)
		_, err := fmt.Fprintf(client.Logger, "%s REQUEST: %s\n", name, reqDump)
		if err != nil {
			fmt.Printf("error while logging %s request: %v\n",
				strings.ToLower(name), err)
		}
	}

	if response != nil {
		respDump, _ := httputil.DumpResponse(response, true)
		_, err := fmt.Fprintf(client.Logger, "%s RESPONSE: %s\n", name, respDump)
		if err != nil {
			fmt.Printf("error while logging %s response: %v\n",
				strings.ToLower(name), err)
		}
	}

	return
}

//getSessionID examine if there is a session id saved as Client.sessionID
//if it is available it checks if it has already expired or have more than
//1 minute till expiration date and returns it
//if the above conditions are not fulfilled it calls Client.SessionID
//then save it and increment the expiration date
func (client *Client) getSessionID() (string, error) {
	isAvailable := client.sessionID != nil && *client.sessionID != ""
	notExpired := client.sessionExpiration.Sub(time.Now()).Minutes() > 1
	if isAvailable && notExpired {
		return *client.sessionID, nil
	}

	resp, err := client.SessionID(context.Background())
	if err != nil {
		return "", fmt.Errorf("could not fetch session id: %w", err)
	}

	if resp.OutputErr != "" {
		return "", fmt.Errorf("could not fetch session id: %s", resp.OutputErr)
	}

	up := time.Duration(client.SessionLifetimeMinutes) * time.Minute
	expiration := time.Now().Add(up)
	client.sessionExpiration = expiration
	client.sessionID = &resp.ID

	return resp.ID, err

}

