package main

import (
	"context"
	"fmt"
	"github.com/techcraftlabs/mpesa"
	"github.com/techcraftlabs/mpesa/pkg/models"
	"log"
	"os"
)

func main() {
	apiKey := "APwvV7m1UKSjHXquRZNBbW24HzC5SK31"
	//pubKey := "MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEArv9yxA69XQKBo24BaF/D+fvlqmGdYjqLQ5WtNBb5tquqGvAvG3WMFETVUSow/LizQalxj2ElMVrUmzu5mGGkxK08bWEXF7a1DEvtVJs6nppIlFJc2SnrU14AOrIrB28ogm58JjAl5BOQawOXD5dfSk7MaAA82pVHoIqEu0FxA8BOKU+RGTihRU+ptw1j4bsAJYiPbSX6i71gfPvwHPYamM0bfI4CmlsUUR3KvCG24rB6FNPcRBhM3jDuv8ae2kC33w9hEq8qNB55uw51vK7hyXoAa+U7IqP1y6nBdlN25gkxEA8yrsl1678cspeXr+3ciRyqoRgj9RD/ONbJhhxFvt1cLBh+qwK2eqISfBb06eRnNeC71oBokDm3zyCnkOtMDGl7IvnMfZfEPFCfg5QgJVk1msPpRvQxmEsrX9MQRyFVzgy2CWNIb7c+jPapyrNwoUbANlN8adU1m6yOuoX7F49x+OjiG2se0EJ6nafeKUXw/+hiJZvELUYgzKUtMAZVTNZfT8jjb58j8GVtuS+6TM2AutbejaCV84ZK58E2CRJqhmjQibEUO6KPdD7oTlEkFy52Y1uOOBXgYpqMzufNPmfdqqqSM4dU70PO8ogyKGiLAIxCetMjjm6FCMEA3Kc8K0Ig7/XtFm9By6VxTJK1Mg36TlHaZKP6VzVLXMtesJECAwEAAQ=="

	pubKey := "MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEArv9yxA69XQKBo24BaF/D+fvlqmGdYjqLQ5WtNBb5tquqGvAvG3WMFETVUSow/LizQalxj2ElMVrUmzu5mGGkxK08bWEXF7a1DEvtVJs6nppIlFJc2SnrU14AOrIrB28ogm58JjAl5BOQawOXD5dfSk7MaAA82pVHoIqEu0FxA8BOKU+RGTihRU+ptw1j4bsAJYiPbSX6i71gfPvwHPYamM0bfI4CmlsUUR3KvCG24rB6FNPcRBhM3jDuv8ae2kC33w9hEq8qNB55uw51vK7hyXoAa+U7IqP1y6nBdlN25gkxEA8yrsl1678cspeXr+3ciRyqoRgj9RD/ONbJhhxFvt1cLBh+qwK2eqISfBb06eRnNeC71oBokDm3zyCnkOtMDGl7IvnMfZfEPFCfg5QgJVk1msPpRvQxmEsrX9MQRyFVzgy2CWNIb7c+jPapyrNwoUbANlN8adU1m6yOuoX7F49x+OjiG2se0EJ6nafeKUXw/+hiJZvELUYgzKUtMAZVTNZfT8jjb58j8GVtuS+6TM2AutbejaCV84ZK58E2CRJqhmjQibEUO6KPdD7oTlEkFy52Y1uOOBXgYpqMzufNPmfdqqqSM4dU70PO8ogyKGiLAIxCetMjjm6FCMEA3Kc8K0Ig7/XtFm9By6VxTJK1Mg36TlHaZKP6VzVLXMtesJECAwEAAQ=="

	config := &mpesa.Config{
		Name:                   "paymonga",
		Version:                "1.0",
		Description:            "mongas",
		APIKey:                 apiKey,
		PublicKey:              pubKey,
		BasePath:               "openapi.m-pesa.com",
		SessionLifetimeMinutes: 60,
		ServiceProvideCode:     "000000",
		TrustedSources:         nil,
	}
	var clientOpts []mpesa.ClientOpt

	wOpt := mpesa.WithWriter(os.Stderr)
	dOpt := mpesa.WithDebugMode(true)
	clientOpts = append(clientOpts, wOpt, dOpt)

	client, err := mpesa.NewClient(config, mpesa.TanzaniaMarket, mpesa.SANDBOX, clientOpts...)

	if err != nil {
		log.Fatalf("error %v\n", err)
	}
	fm := &mpesa.FormatAdapter{
		Next: client,
	}

	logger := log.New(os.Stderr, "MPESA", log.Ldate|log.Ltime|log.Lshortfile)
	lg := &mpesa.LoggerAdapter{
		Logger: logger,
		Next:   fm,
	}

	resp, err := lg.SessionID(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	push := models.PushRequest{
		ThirdPartyID: "djkefhjfugydefytfdtydcty",
		Reference:    "ndkndkjdnkdjndkjnd",
		Amount:       1000,
		MSISDN:       "000000000001",
		Desc:         "demo",
	}
	fmt.Printf("reponse error: %s\n", resp.OutputErr)
	resp2, err := lg.C2BSingleAsync(context.Background(), push)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("reponse2: %s\n", resp2)
}
