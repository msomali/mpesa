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
	"github.com/techcraftlabs/mpesa/pkg/models"
	"io"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestClient_C2BSingleAsync(t *testing.T) {
	type fields struct {
		Config            *Config
		Http              *http.Client
		DebugMode         bool
		Logger            io.Writer
		Market            *Market
		Platform          Platform
		encryptedApiKey   *string
		sessionID         *string
		sessionExpiration time.Time
	}
	type args struct {
		ctx     context.Context
		request models.PushRequest
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse models.C2BSingleStageAsyncResponse
		wantErr      bool
	}{
		{
			name: "",
			fields: fields{
				Config:            nil,
				Http:              nil,
				DebugMode:         false,
				Logger:            nil,
				Market:            nil,
				Platform:          0,
				encryptedApiKey:   nil,
				sessionID:         nil,
				sessionExpiration: time.Time{},
			},
			args: args{
				ctx: nil,
				request: models.PushRequest{
					ThirdPartyID: "",
					Reference:    "",
					Amount:       0,
					MSISDN:       "",
					Desc:         "",
				},
			},
			wantResponse: models.C2BSingleStageAsyncResponse{
				ResponseCode:             "",
				ResponseDesc:             "",
				ConversationID:           "",
				ThirdPartyConversationID: "",
				OutputErr:                "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				Config:            tt.fields.Config,
				Http:              tt.fields.Http,
				DebugMode:         tt.fields.DebugMode,
				Logger:            tt.fields.Logger,
				Market:            tt.fields.Market,
				Platform:          tt.fields.Platform,
				encryptedApiKey:   tt.fields.encryptedApiKey,
				sessionID:         tt.fields.sessionID,
				sessionExpiration: tt.fields.sessionExpiration,
			}
			gotResponse, err := client.PushPayAsync(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("PushPayAsync() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("PushPayAsync() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
