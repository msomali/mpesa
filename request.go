package mpesa

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Request struct {
	Method      string
	Type        RequestType
	Endpoint    string
	Payload     interface{}
	Headers     map[string]string
	QueryParams map[string]string
}

func (t RequestType) name() string {
	values := map[int]string{
		0: "session key",
		1: "c2b single stage",
	}

	return values[int(t)]
}

func (r *Request) newRequestWithContext(ctx context.Context, market *Market, platform Platform) (*http.Request, error) {

	var buffer io.Reader

	url := generateRequestURL(defBasePath, platform, *market, r.Endpoint)
	if r.Payload != nil {
		buf, err := json.Marshal(r.Payload)
		if err != nil {
			return nil, err
		}

		buffer = bytes.NewBuffer(buf)
		req, err := http.NewRequestWithContext(ctx, r.Method, url, buffer)
		if err != nil {
			return nil, err
		}
		for key, value := range r.Headers {
			req.Header.Add(key, value)
		}
		return req, nil
	}

	req, err := http.NewRequestWithContext(ctx, r.Method, url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range r.Headers {
		req.Header.Add(key, value)
	}

	for name, value := range r.QueryParams {
		values := req.URL.Query()
		values.Add(name, value)
		req.URL.RawQuery = values.Encode()
	}
	return req, nil
}
