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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateRequestURL(tt.args.base, tt.args.platform, tt.args.market, tt.args.endpoint); got != tt.want {
				t.Errorf("generateRequestURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
