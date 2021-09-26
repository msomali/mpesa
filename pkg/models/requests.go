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

	B2CSingleStageReq struct {
		Amount                   string `json:"input_Amount"`
		Country                  string `json:"input_Country"`
		Currency                 string `json:"input_Currency"`
		CustomerMSISDN           string `json:"input_CustomerMSISDN"`
		ServiceProviderCode      string `json:"input_ServiceProviderCode"`
		ThirdPartyConversationID string `json:"input_ThirdPartyConversationID"`
		TransactionReference     string `json:"input_TransactionReference"`
		PaymentItemsDesc         string `json:"input_PaymentItemsDesc"`
	}
	PushRequest struct {
		ThirdPartyID string  `json:"id"`
		Reference    string  `json:"reference"`
		Amount       float64 `json:"amount"`
		MSISDN       string  `json:"msisdn"`
		Desc         string  `json:"desc"`
	}
)
