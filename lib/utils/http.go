package utils

import (
	"crypto/tls"
	"net"
	"net/http"

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
	// Transport for Dialing Unix Socket
	transport := http.Transport{
		Dial: func(_, _ string) (net.Conn, error) {
			return net.Dial("unix", hostUrl)
		},
	}

	// From Example: https://github.com/go-resty/resty#unix-socket
	client := resty.New()
	client.SetTransport(&transport)
	client.SetScheme("http")
	client.SetBaseURL(hostUrl)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetHeader("User-Agent", "server-control-api/1.0 (https://github.com/3n3a/server-control-api)")

	if IsDev() {
		// Dumps HTTP Req and Res
		client.SetDebug(true)
	}

	return client
}
