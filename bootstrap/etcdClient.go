package bootstrap

import (
	"api-gateway/library/resource"
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func initEtcdClient(_ context.Context) {
	var err error
	peerURLS := resource.Config.EtceSetting.PeerURLS
	resource.EtcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   peerURLS,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	log.Print("init etcd client success")
}
