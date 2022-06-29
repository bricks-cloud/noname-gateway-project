package http

import (
	"net"
	"net/http"

	"github.com/bricks-cloud/noname-gateway-project/common"
)

type client struct {
	client *http.Client
}

// NewClient creates a new http client that the gateway uses to send http requests
func NewClient(config common.ClientConfig) *client {
	return &client{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   config.GetDialTimeout(),
					KeepAlive: config.GetDialKeepAlive(),
				}).DialContext,
				MaxIdleConns:          config.GetMaxIdleConns(),
				MaxIdleConnsPerHost:   config.GetMaxConnsPerHost(),
				MaxConnsPerHost:       config.GetMaxConnsPerHost(),
				DisableCompression:    config.GetDisableCompression(),
				IdleConnTimeout:       config.GetIdleConnTimeout(),
				TLSHandshakeTimeout:   config.GetTLSHandshakeTimeout(),
				ExpectContinueTimeout: config.GetExpectContinueTimeout(),
			},
		},
	}
}

func (c *client) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	req.URL.Scheme = "http"
	req.RequestURI = ""
	resp, err = c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
