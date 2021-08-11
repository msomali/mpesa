package models

type (
	C2BSingleStageReq struct {
		Amount                   string `json:"input_Amount"`
		Country                  string `json:"input_Country"`
		Currency                 string `json:"input_Currency"`
		CustomerMSISDN           string `json:"input_CustomerMSISDN"`
		ServiceProviderCode      string `json:"input_ServiceProviderCode"`
		ThirdPartyConversationID string `json:"input_ThirdPartyConversationID"`
		TransactionReference     string `json:"input_TransactionReference"`
		PurchasedItemsDesc       string `json:"input_PurchasedItemsDesc"`
	}

	PushRequest struct {
		ThirdPartyID string  `json:"id"`
		Reference    string  `json:"reference"`
		Amount       float64 `json:"amount"`
		MSISDN       string  `json:"msisdn"`
		Desc         string  `json:"desc"`
	}
)
