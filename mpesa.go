package mpesa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
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
	defBasePath               = "openapi.m-pesa.com"
	defSessionKeyEndpoint     = "getSession"
	defC2BSingleStageEndpoint = "c2bPayment/singleStage"
)



const (
	SANDBOX Platform = 0
	OPENAPI Platform = 1
)

type (

	Market struct {
		URLContextValue string
		Country         string
		Currency        string
		Description     string
	}


	Platform int
)

//generateEncryptedKey
//To generate your Session Key for the sandbox and live environments:
//1.	log into OpenAPI with a developer account
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

	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, msg)

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


