package main

import (
	"api-gateway/bootstrap"
	etcdclient "api-gateway/etcdClient"
	"context"
	"fmt"
)

func main() {
	ctx := context.TODO()
	bootstrap.MustLoadAppConfig("./conf/")

	etcdclient.EtcdClientTest(ctx)
	fmt.Println("ook")
	fmt.Println("ook")
	fmt.Println("ook")
}
