package mpesa

import (
	"context"
	"fmt"
	"github.com/techcraftlabs/mpesa/pkg/models"
	"io"
	"net/http"
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
	defBasePath = "openapi.m-pesa.com"
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
		Method string
		Type     RequestType
		Endpoint string
		Payload  interface{}
		Headers  map[string]interface{}
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

func (client *Client)SessionKey(ctx context.Context, platform Platform, market Market)(response models.SessionResponse,err error){

	headers := map[string]interface{}{
		"Content-Type":"application/json",
	}

	url := fmt.Sprintf("https://%s",defSessionKeyEndpoint)
	request := &Request{
		Type:     0,
		Endpoint: url,
		Payload:  nil,
		Headers:  headers,
		Market:   &market,
		Platform: platform,
	}
	err = client.send(ctx,request, &response)
	return response,err
}

func (client *Client) send(ctx context.Context, request *Request, v interface{}) error {
	//url := generateRequestURL(defBasePath,request.Platform, *request.Market,request.Endpoint)
	req, err := request.newRequestWithContext(ctx)
	if err != nil{
		return err
	}
	resp, err := client.Http.Do(req)

	if err != nil{
		return err
	}
	fmt.Printf("%v\n",resp)
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
func generateEncryptedKey()(string,error){
	return "",nil
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

	return fmt.Sprintf("https://%s/%s/ipg/v2/%s/%s/",base,platformStr,marketStr,endpoint)
}

func (r *Request) newRequestWithContext(ctx context.Context) (*http.Request,error){
	return nil, nil
}
