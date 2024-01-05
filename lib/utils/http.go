package utils

import (
	"crypto/tls"

	"github.com/go-resty/resty/v2"
)

func NewHTTPClient() *resty.Client {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetHeader("User-Agent", "server-control-api/1.0 (https://github.com/3n3a/server-control-api)")

	if IsDev() {
		// Dumps HTTP Req and Res
		client.SetDebug(true)
	}

	return client
}

func NewUnixHTTPClient(hostUrl string) *resty.Client {
	client := resty.New()
	client.SetBaseURL(hostUrl)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetHeader("User-Agent", "server-control-api/1.0 (https://github.com/3n3a/server-control-api)")

	if IsDev() {
		// Dumps HTTP Req and Res
		client.SetDebug(true)
	}

	return client
}
