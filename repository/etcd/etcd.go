package etcd

import (
	"api-gateway/library/resource"
	"api-gateway/pkg/app"
	"context"

	"github.com/sirupsen/logrus"
)

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
		val = string(res.Kvs[0].Value)
	}
	return val, nil
}
