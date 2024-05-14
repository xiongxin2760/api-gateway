package httpserver

import (
	"api-gateway/library/resource"
	"context"
	"fmt"
	"net/http"
	"time"
)

func NewServer(ctx context.Context) *http.Server {
	ser := &http.Server{
		Addr:           fmt.Sprintf(":%d", resource.Config.Port),
		Handler:        httpRouter(),
		ReadTimeout:    time.Millisecond * time.Duration(resource.Config.HTTPServer.ReadTimeout),
		WriteTimeout:   time.Millisecond * time.Duration(resource.Config.HTTPServer.WriteTimeout),
		IdleTimeout:    time.Millisecond * time.Duration(resource.Config.HTTPServer.IdleTimeout),
		MaxHeaderBytes: 1 << 20,
	}

	return ser
}
