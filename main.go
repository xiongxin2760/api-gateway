package main

import (
	etcdclient "api-gateway/etcdClient"
	"context"
	"fmt"
)

func main() {
	ctx := context.TODO()
	etcdclient.EtcdClientTest(ctx)
	fmt.Println("ook")
	fmt.Println("ook")
	fmt.Println("ook")
}
