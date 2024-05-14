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

	// ctx := context.TODO()
	// bootstrap.MustLoadAppConfig("./conf/")

	// etcdclient.EtcdClientTest(ctx)
	// fmt.Println("ook")
	// fmt.Println("ook")
	// fmt.Println("ook")
}
