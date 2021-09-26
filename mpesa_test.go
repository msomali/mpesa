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

import "testing"

func Test_generateRequestURL(t *testing.T) {
	type args struct {
		base     string
		platform Platform
		market   Market
		endpoint string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "get session id request url for tanzania market in sandbox",
			args: args{
				base:     defBasePath,
				platform: SANDBOX,
				market:   *TanzaniaMarket,
				endpoint: defSessionKeyEndpoint,
			},
			want: "https://openapi.m-pesa.com/sandbox/ipg/v2/vodacomTZN/getSession/",
		},
		{
			name: "get session id request url for tanzania market in openapi",
			args: args{
				base:     defBasePath,
				platform: OPENAPI,
				market:   *TanzaniaMarket,
				endpoint: defSessionKeyEndpoint,
			},
			want: "https://openapi.m-pesa.com/openapi/ipg/v2/vodacomTZN/getSession/",
		},
		{
			name: "ghana sandbox get session id url",
			args: args{
				base:     defBasePath,
				platform: SANDBOX,
				market:   *GhanaMarket,
				endpoint: defSessionKeyEndpoint,
			},
			want: "https://openapi.m-pesa.com/sandbox/ipg/v2/vodafoneGHA/getSession/",
		},
		{
			name: "test single stage endpoint",
			args: args{
				base:     defBasePath,
				platform: SANDBOX,
				market:   *TanzaniaMarket,
				endpoint: defC2BSingleStageEndpoint,
			},
			want: "https://openapi.m-pesa.com/sandbox/ipg/v2/vodacomTZN/c2bPayment/singleStage/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateRequestURL(tt.args.base, tt.args.platform, tt.args.market, tt.args.endpoint); got != tt.want {
				t.Errorf("generateRequestURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestType_nameValue(t *testing.T) {
	tests := []struct {
		name string
		t    RequestType
		want string
	}{
		{
			name: "generate session key name",
			t:    GenerateSessionKey,
			want: "session key",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.name(); got != tt.want {
				t.Errorf("name() = %v, want %v", got, tt.want)
			}
		})
	}
}
