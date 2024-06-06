package etcdclient

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func EtcdClientTest(ctx context.Context) error {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://110.95.17.171:12379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}
	clientWatch := clientv3.NewWatcher(client)
	watchChan := clientWatch.Watch(ctx, "/xiongxin/etcd")
	for item := range watchChan {
		fmt.Println(item.Events[0].Kv)
	}
	return nil
}
