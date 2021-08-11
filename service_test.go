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
			gotResponse, err := client.C2BSingleAsync(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("C2BSingleAsync() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("C2BSingleAsync() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
