package mpesa

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/techcraftlabs/mpesa/pkg/models"
	"io"
	"net/http"
	"net/http/httputil"
	"strings"
)

var (
	GhanaMarket = &Market{
		URLContextValue: "vodafoneGHA",
		Country:         "GHA",
		Currency:        "GHS",
		Description:     "Vodafone Ghana",
	}

	TanzaniaMarket = &Market{
		URLContextValue: "vodacomTZN",
		Country:         "TZN",
		Currency:        "TZS",
		Description:     "vodacomTZN",
	}
)

const (

	//https://openapi.m-pesa.com:443/sandbox/ipg/v2/vodacomTZN/getSession/
	defBasePath           = "openapi.m-pesa.com"
	defSessionKeyEndpoint = "getSession"
)

const (
	GenerateSessionKey RequestType = iota
	C2BSingleStage
	B2CSingleStage
	B2BSingleStage
	Reversal
	QueryTxnStatus
	DirectDebitCreate
	DirectDebitPayment
)

const (
	SANDBOX Platform = 0
	OPENAPI Platform = 1
)

type (
	// Config contains details initialize in mpesa portal
	// Applications require the following details:
	//•	Application Name – human-readable name of the application
	//•	Version – version number of the application, allowing changes in API products to be managed in different versions
	//•	Description – Free text field to describe the use of the application
	//•	APIKey – Unique authorisation key used to authenticate the application on the first call. API Keys need to be encrypted in the first “Generate Session API Call” to create a valid session key to be used as an access token for future calls. Encrypting the API Key is documented in the GENERATE SESSION API page
	//•	SessionLifetime – The session key has a finite lifetime of availability that can be configured. Once a session key has expired, the session is no longer usable, and the caller will need to authenticate again.
	//•	TrustedSources – the originating caller can be limited to specific IP address(es) as an additional security measure.
	//•	Products / Scope / Limits – the required API products for the application can be enabled and limits defined for each call.
	Config struct {
		Name            string
		Version         string
		Description     string
		APIKey          string
		PublicKey       string
		SessionLifetime int64
		TrustedSources  []string
	}
	Market struct {
		URLContextValue string
		Country         string
		Currency        string
		Description     string
	}
	RequestType int

	Platform int
	Request  struct {
		Method   string
		Type     RequestType
		Endpoint string
		Payload  interface{}
		Headers  map[string]string
		Market   *Market
		Platform Platform
	}
	Client struct {
		*Config
		Http      *http.Client
		BasePath  string
		DebugMode bool
		Logger    io.Writer
	}
)

func (t RequestType) nameValue() string {
	values := map[int]string{
		0 :"session key",
		1 :"c2b single stage",
	}

	return values[int(t)]
}

func (client *Client) SessionKey(ctx context.Context, platform Platform, market Market) (response models.SessionResponse, err error) {

	token, err := generateEncryptedKey(client.APIKey, client.PublicKey)
	if err != nil {
		return response, err
	}
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Origin":        "*",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	url := fmt.Sprintf("https://%s", defSessionKeyEndpoint)
	request := &Request{
		Type:     GenerateSessionKey,
		Endpoint: url,
		Payload:  nil,
		Headers:  headers,
		Market:   &market,
		Platform: platform,
	}
	err = client.send(ctx, request, &response)
	return response, err
}

func (client *Client) send(ctx context.Context, request *Request, v interface{}) error {

	//creates http request with context
	req, err := request.newRequestWithContext(ctx)

	if err != nil {
		return err
	}

	if v == nil {
		return errors.New("v interface can not be empty")
	}
	resp, err := client.Http.Do(req)

	if err != nil {
		return err
	}

	contentType := resp.Header.Get("Content-Type")

	if strings.Contains(contentType, "application/json") {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			if err != io.EOF {
				return err
			}
		}
	}

	return nil
}

//generateEncryptedKey
//To generate your Session Key for the sandbox and live environments:
//1.	Log into OpenAPI with a developer account
//2.	On the APPLICATION page, click Create New Application. Creating an application will generate your unique API Application Key needed to authorise and authenticate your application on the server
//3.	Type your name and version number for the application and select the products you wish to use. (The application can be configured any time). Save your new application.
//4.	Copy and save the API Key.
//5.	Copy the Public Key from the below section.
//6.	Generate a decoded Base64 string from the Public Key
//7.	Generate an instance of an RSA cipher and use the Base 64 string as the input
//8.	Encode the API Key with the RSA cipher and digest as Base64 string format
//9.	The result is your encrypted API Key.
func generateEncryptedKey(apiKey, pubKey string) (string, error) {
	decodedBase64, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return "", fmt.Errorf("could not decode pub key to Base64 string: %w", err)
	}

	publicKeyInterface, err := x509.ParsePKIXPublicKey(decodedBase64)
	if err != nil {
		return "", fmt.Errorf("could not parse encoded public key (encryption key) : %w", err)
	}

	//check if the public key is RSA public key
	publicKey, isRSAPublicKey := publicKeyInterface.(*rsa.PublicKey)
	if !isRSAPublicKey {
		return "", fmt.Errorf("public key parsed is not an RSA public key : %w", err)
	}

	msg := []byte(apiKey)

	encrypted, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, publicKey, msg, nil)

	if err != nil {
		return "", fmt.Errorf("could not encrypt api key using generated public key: %w", err)
	}

	return base64.StdEncoding.EncodeToString(encrypted), nil

}

func generateRequestURL(base string, platform Platform, market Market, endpoint string) string {
	var marketStr, platformStr string
	marketStr = market.URLContextValue
	switch platform {
	case SANDBOX:
		platformStr = "sandbox"

	case OPENAPI:
		platformStr = "openapi"
	}

	return fmt.Sprintf("https://%s/%s/ipg/v2/%s/%s/", base, platformStr, marketStr, endpoint)
}

func (r *Request) newRequestWithContext(ctx context.Context) (*http.Request, error) {

	var buffer io.Reader

	if r.Payload != nil{
		buf, err := json.Marshal(r.Payload)
		if err != nil {
			return nil, err
		}

		buffer = bytes.NewBuffer(buf)
	}

	url := generateRequestURL(defBasePath,r.Platform, *r.Market,r.Endpoint)
	req, err := http.NewRequestWithContext(ctx, r.Method, url, buffer)
	if err != nil {
		return nil, err
	}
	for key, value := range r.Headers {
		req.Header.Add(key, value)
	}

	return req, nil

}


//func (client *Client) LogPayload(t internal.PayloadType, prefix string, payload interface{}) {
//	buf, _ := internal.MarshalPayload(t, payload)
//	_, _ = client.Logger.Write([]byte(fmt.Sprintf("%s: %s\n\n", prefix, buf.String())))
//}

// Log is called to print the details of http.Request sent from Tigo during
// callback, namecheck or ussd payment. It is used for debugging purposes
func (client *Client) Log(name string, request *http.Request) {

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

// LogOut is like Log except this is for outgoing client requests:
// http.Request that is supposed to be sent to tigo
func (client *Client) LogOut(name string, request *http.Request, response *http.Response) {

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
