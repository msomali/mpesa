/*
 * MIT License
 *
 * Copyright (c) 2021 TECHCRAFT TECHNOLOGIES CO LTD.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package models

type (
	SessionResponse struct {
		Code      string `json:"output_ResponseCode,omitempty"`
		Desc      string `json:"output_ResponseDesc,omitempty"`
		ID        string `json:"output_SessionID,omitempty"`
		OutputErr string `json:"output_error,omitempty"`
	}

	B2CSingleStageResponse struct {
		ConversationID           string `json:"output_ConversationID,omitempty"`
		ResponseCode             string `json:"output_ResponseCode,omitempty"`
		ResponseDesc             string `json:"output_ResponseDesc,omitempty"`
		TransactionID            string `json:"output_TransactionID,omitempty"`
		ThirdPartyConversationID string `json:"output_ThirdPartyConversationID,omitempty"`
		OutputErr                string `json:"output_error,omitempty"`
	}

	C2BSingleStageAsyncResponse struct {
		ResponseCode             string `json:"output_ResponseCode,omitempty"`
		ResponseDesc             string `json:"output_ResponseDesc,omitempty"`
		TransactionID            string `json:"output_TransactionID,omitempty"`
		ConversationID           string `json:"output_ConversationID,omitempty"`
		ThirdPartyConversationID string `json:"output_ThirdPartyConversationID,omitempty"`
		OutputErr                string `json:"output_error,omitempty"`
	}
)
