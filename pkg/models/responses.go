package models

type (
	SessionResponse struct {
		Code string `json:"output_ResponseCode,omitempty"`
		Desc string `json:"output_ResponseDesc,omitempty"`
		ID   string `json:"output_SessionID,omitempty"`
		OutputErr string `json:"output_error,omitempty"`
	}
)
