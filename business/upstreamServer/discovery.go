package upstreamserver

import (
	"api-gateway/library/resource"
	"api-gateway/pkg/app"
	"context"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var ServerDiscoveryHandle *ServerDiscovery

// 服务发现
type ServerDiscovery struct {
	lock      sync.Mutex
	ServerMap map[string]*APIServer
	Prefix    string
}

func NewServerDiscovery(ctx context.Context) *ServerDiscovery {
	return &ServerDiscovery{
		ServerMap: make(map[string]*APIServer),
		Prefix:    "/upstream/server/",
		lock:      sync.Mutex{},
	}
}

func (server *ServerDiscovery) ADD(ctx context.Context, key string, apiServer *APIServer) {
	if apiServer == nil {
		return
	}
	server.lock.Lock()
	defer server.lock.Unlock()
	server.ServerMap[key] = apiServer
}

func (server *ServerDiscovery) DEL(ctx context.Context, key string) {
	server.lock.Lock()
	defer server.lock.Unlock()
	delete(server.ServerMap, key)
}

func (server *ServerDiscovery) GetKey(ID string) string {
	return server.Prefix + ID
}

// 1. 初始化，获取全量配置
// 2. 后台监听，监控实时变化
func (server *ServerDiscovery) Init(ctx context.Context) error {
	logger := app.GetGlobalLogger(ctx)
	client := resource.EtcdClient
	resp, err := client.Get(ctx, server.Prefix, clientv3.WithPrefix())
	if err != nil {
		logger.WithError(err).Errorln("etcd init fail")
		return err
	}
	// 信息录入
	for _, ev := range resp.Kvs {
		key := string(ev.Key)
		val := string(ev.Value)
		apiServer, err := NewAPIServerByStr(ctx, val)
		if err != nil {
			continue
		}
		server.ServerMap[key] = apiServer
	}
	// 版本号
	revision := resp.Header.Revision
	// 后台监听
	go func() {
		clientWatch := clientv3.NewWatcher(client)
		watchChan := clientWatch.Watch(ctx, server.Prefix, clientv3.WithPrefix(), clientv3.WithRev(revision))
		for item := range watchChan {
			err = item.Err()
			if err != nil {
				logger.WithError(err).Errorln("etcd watch fail")
				continue
			}
			for _, data := range item.Events {
				tyep := data.Type
				key := string(data.Kv.Key)
				val := string(data.Kv.Value)
				if tyep == clientv3.EventTypePut {
					// 新增, 更新
					apiServer, err := NewAPIServerByStr(ctx, val)
					if err != nil {
						continue
					}
					server.ADD(ctx, key, apiServer)
				} else if tyep == clientv3.EventTypeDelete {
					// 删除
					server.DEL(ctx, key)
				}
			}
		}
	}()
	return nil
}
