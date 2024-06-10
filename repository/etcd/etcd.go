package etcd

import (
	"api-gateway/library/resource"
	"api-gateway/pkg/app"
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdData struct {
	Key string
	Val string
}

// 改变etcd的键值对
func Put(ctx context.Context, key string, val string) error {
	logger := app.GetGlobalLogger(ctx)
	client := resource.EtcdClient
	_, err := client.Put(ctx, key, val)
	if err != nil {
		logger.WithError(err).WithFields(logrus.Fields{
			"key": key,
			"val": val,
		}).Errorln("etcd put fail")
		return err
	}
	return nil
}

// etcd读取
func Get(ctx context.Context, key string) (string, error) {
	logger := app.GetGlobalLogger(ctx)
	client := resource.EtcdClient
	res, err := client.Get(ctx, key)
	if err != nil {
		logger.WithError(err).WithFields(logrus.Fields{
			"key": key,
		}).Errorln("etcd get fail")
		return "", err
	}
	val := ""
	if len(res.Kvs) > 0 {
		fmt.Println(res.Header.Revision)
		val = string(res.Kvs[0].Value)
	}
	return val, nil
}

// 通过管道监控
// 通过后台线程更新
func Watch(ctx context.Context, prefix string) chan EtcdData {
	outChannel := make(chan EtcdData)
	go func() {
		defer close(outChannel)
		client := resource.EtcdClient
		clientWatch := clientv3.NewWatcher(client)
		watchChan := clientWatch.Watch(ctx, prefix, clientv3.WithPrefix(), clientv3.WithRev(36))
		for item := range watchChan {
			for _, data := range item.Events {
				fmt.Println(string(data.Type), string(data.Kv.Key), string(data.Kv.Value))
				outChannel <- EtcdData{
					Key: string(data.Kv.Key),
					Val: string(data.Kv.Value),
				}
			}
		}
	}()
	return outChannel
}
