package bootstrap

import (
	"context"
	"io"
	"log"
)

func MustInit(ctx context.Context) {
	initLoggers(ctx)
	initHTTPClient(ctx)
	initEtcdClient(ctx)
	initUpServer(ctx)
}

var closeFns []func() error

func tryRegisterCloser(comp any) {
	if c, ok := comp.(io.Closer); ok {
		closeFns = append(closeFns, c.Close)
		return
	}
	if fn, ok := comp.(func() error); ok {
		closeFns = append(closeFns, fn)
	}
}

// BeforeShutdown 退出前执行，资源清理、日志落盘等
func BeforeShutdown() {
	log.Println("BeforeShutdown closing")
	for _, fn := range closeFns {
		_ = fn()
	}
}
