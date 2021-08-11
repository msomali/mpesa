package models

type (
	SessionResponse struct {
		Code      string `json:"output_ResponseCode,omitempty"`
		Desc      string `json:"output_ResponseDesc,omitempty"`
		ID        string `json:"output_SessionID,omitempty"`
		OutputErr string `json:"output_error,omitempty"`
	}

	C2BSingleStageAsyncResponse struct {
		ResponseCode             string `json:"output_ResponseCode,omitempty"`
		ResponseDesc             string `json:"output_ResponseDesc,omitempty"`
		ConversationID           string `json:"output_ConversationID,omitempty"`
		ThirdPartyConversationID string `json:"output_ThirdPartyConversationID,omitempty"`
		OutputErr                string `json:"output_error,omitempty"`
	}
)
