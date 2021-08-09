package mpesa

import (
	"context"
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

func (c *Client)Send(ctx context.Context, request *Request, v interface{}) error {
	return nil
}
