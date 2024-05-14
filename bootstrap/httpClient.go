package bootstrap

import (
	"api-gateway/library/resource"
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// initHTTPClient 初始化http连接池
func initHTTPClient(_ context.Context) {
	resource.HTTPClient = resty.New()
	dialer := &net.Dialer{
		// Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}
	transport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     false,
		IdleConnTimeout:       60 * time.Second,
		TLSHandshakeTimeout:   60 * time.Second,
		ResponseHeaderTimeout: 5 * 60 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   resource.Config.HTTPClient.MaxIdleConnsPerHost,
	}
	resource.HTTPClient.SetTransport(transport)
	resource.HTTPClient.SetTimeout(time.Duration(resource.Config.HTTPClient.Timeoutms) * time.Millisecond)
	log.Print("init http client success")
}
