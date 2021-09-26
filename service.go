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

package mpesa

import (
	"context"
	"fmt"
	"github.com/techcraftlabs/mpesa/pkg/models"
	"net/http"
)

type Service interface {
	SessionID(ctx context.Context) (response models.SessionResponse, err error)
	C2BSingleAsync(ctx context.Context, request models.PushRequest) (models.C2BSingleStageAsyncResponse, error)
}

func (client *Client) SessionID(ctx context.Context) (response models.SessionResponse, err error) {

	token, err := client.getEncryptionKey()
	if err != nil {
		return response, err
	}
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Origin":        "*",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	request := &Request{
		Method:   http.MethodGet,
		Type:     GenerateSessionKey,
		Endpoint: defSessionKeyEndpoint,
		Payload:  nil,
		Headers:  headers,
	}
	err = client.send(ctx, request, &response)

	//save the session id
	if response.OutputErr != "" {
		err1 := fmt.Errorf("could not fetch session id: %s", response.OutputErr)
		return response, err1
	}

	return response, nil
}

func (client *Client) C2BSingleAsync(ctx context.Context, request models.PushRequest) (response models.C2BSingleStageAsyncResponse, err error) {
	sess, err := client.getSessionID()
	if err != nil {
		return response, err
	}
	token, err := generateEncryptedKey(sess, client.PublicKey)
	if err != nil {
		return response, err
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Origin":        "*",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	payload := models.C2BSingleStageReq{
		Amount:                   fmt.Sprintf("%f", request.Amount),
		Country:                  client.Market.Country,
		Currency:                 client.Market.Currency,
		CustomerMSISDN:           request.MSISDN,
		ServiceProviderCode:      client.ServiceProvideCode,
		ThirdPartyConversationID: request.ThirdPartyID,
		TransactionReference:     request.Reference,
		PurchasedItemsDesc:       request.Desc,
	}

	re := &Request{
		Method:   http.MethodPost,
		Type:     C2BSingleStage,
		Endpoint: defC2BSingleStageEndpoint,
		Payload:  payload,
		Headers:  headers,
	}
	err = client.send(ctx, re, &response)

	//save the session id
	if response.OutputErr != "" {
		err1 := fmt.Errorf("could not perform c2b single stage request: %s", response.OutputErr)
		return response, err1
	}

	return response, nil
}
