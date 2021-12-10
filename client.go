package mpesa

import (
	"fmt"
	"github.com/techcraftlabs/base"
	"net/http"
	"strings"
)

var (
	_ base.RequestInformer = (*requestType)(nil)
	_ market               = (*Market)(nil)
)

const (
	GhanaMarket    = Market(0)
	TanzaniaMarket = Market(1)
)

const (
	sessionID requestType = iota
	pushPay
	disburse
	queryTxn
)

const (
	SANDBOX Platform = iota
	OPENAPI
)

type (
	market interface {
		URLContextValue() string
		Country() string
		Currency() string
		Description() string
	}
	Platform    int
	Market      int
	requestType int
)

func MarketFmt(marketString string) Market {
	if strings.ToLower(marketString) == "ghana" {
		return GhanaMarket
	}

	if strings.ToLower(marketString) == "tanzania" {
		return TanzaniaMarket
	}

	return Market(-1)
}

func PlatformFmt(platformString string) Platform {
	if strings.ToLower(platformString) == "openapi" {
		return OPENAPI
	}

	if strings.ToLower(platformString) == "sandbox" {
		return SANDBOX
	}

	return Platform(-1)
}

func (p Platform) String() string {
	if p == OPENAPI {
		return "openapi"
	}

	return "sandbox"
}

func (m Market) URLContextValue() string {
	switch m {

	//ghana
	case 0:
		return "vodafoneGHA"
		//tanzania
	case 1:
		return "vodacomTZN"
	default:
		return ""
	}
}

func (m Market) Country() string {
	switch m {

	//ghana
	case 0:
		return "GHA"
		//tanzania
	case 1:
		return "TZN"
	default:
		return ""
	}
}

func (m Market) Currency() string {
	switch m {

	//ghana
	case 0:
		return "GHS"
		//tanzania
	case 1:
		return "TZS"
	default:
		return ""
	}
}

func (m Market) Description() string {
	switch m {

	//ghana
	case 0:
		return "Vodafone Ghana"
		//tanzania
	case 1:
		return "Vodacom Tanzania"
	default:
		return ""
	}
}

func (r requestType) String() string {

	return fmt.Sprintf("mno=%s:: group=%s:: request=%s:", r.MNO(), r.Group(), r.Name())
}

func (r requestType) Endpoint() string {
	switch r {

	case sessionID:
		return "/getSession/"

	case pushPay:
		return "/c2bPayment/singleStage/"

	case disburse:
		return "/b2cPayment/"

	case queryTxn:
		return "/queryTransactionStatus/"

	default:
		return ""
	}
}

func (r requestType) Method() string {
	switch r {

	case sessionID:
		return http.MethodGet

	case pushPay:
		return http.MethodPost

	default:
		return http.MethodPost

	}
}

func (r requestType) Name() string {
	return []string{"get session id", "ussd push",
		"disbursement"}[r]
}

func (r requestType) MNO() string {
	return "vodacom"
}

func (r requestType) Group() string {
	switch r {
	case sessionID:
		return "Authorization"

	case pushPay:
		return "collection"

	case disburse:
		return "disbursement"

	default:
		return ""
	}
}
