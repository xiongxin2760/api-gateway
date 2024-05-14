package main

import (
	"api-gateway/bootstrap"
	"context"
	"flag"
)

var appConfig = flag.String("conf", "./conf/", "app config file")

func main() {
	flag.Parse()

	// 加载配置文件
	bootstrap.MustLoadAppConfig(*appConfig)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bootstrap.MustInit(ctx)

	httpServer := bootstrap.StartServers(ctx)

	// 执行退出前的回调，如日志关闭落盘、其他资源清理工作
	defer bootstrap.BeforeShutdown()

	httpServer.ListenAndServe()
}
