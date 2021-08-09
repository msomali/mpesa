package main

import (
	"github.com/techcraftlabs/mpesa"
	"net/http"
)

func main() {
	client := &mpesa.Client{
		Config:    nil,
		Http:      http.DefaultClient,
		BasePath:  "",
		DebugMode: false,
		Logger:    nil,
	}

	req := &mpesa.Request{
		Type:     mpesa.C2BSingleStage,
		Endpoint: "",
		Payload:  nil,
		Headers:  nil,
		Market:   mpesa.GhanaMarket,
		Platform: mpesa.OPENAPI,
	}
}
