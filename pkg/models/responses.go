package models

type (
	SessionResponse struct {
		Code string `json:"output_ResponseCode"`
		Desc string `json:"output_ResponseDesc"`
		ID   string `json:"output_SessionID"`
	}
)
