package upstreamserver

import (
	"api-gateway/library/resource"
	"api-gateway/pkg/app"
	"api-gateway/pkg/utils"
	"context"
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func (server *ServerDiscovery) Register(ctx context.Context, apiServer APIServer) (string, error) {
	logger := app.GetGlobalLogger(ctx)
	// 暂时将数据存储在内存
	if len(apiServer.ID) == 0 {
		apiServer.ID = utils.GenUUID()
	}
	result, err := apiServer.ToPO(ctx)
	if err != nil {
		return "", err
	}

	pobyte, err := json.Marshal(result)
	if err != nil {
		logger.WithError(err).Errorln("po marshal fail")
		return "", err
	}
	// fmt.Println(string(pobyte))

	client := resource.EtcdClient
	_, err = client.Put(ctx, server.GetKey(result.ID), string(pobyte))
	if err != nil {
		logger.WithError(err).Errorln("etcd put fail")
		return "", err
	}
	return result.ID, nil
}

func (server *ServerDiscovery) Search(ctx context.Context, ID string) (*APIServer, error) {
	if apiServer, exist := server.ServerMap[server.GetKey(ID)]; exist {
		return apiServer, nil
	}
	return nil, nil
}

func (server *ServerDiscovery) SearchList(ctx context.Context) ([]*APIServer, error) {
	resList := []*APIServer{}
	for _, apiServer := range server.ServerMap {
		resList = append(resList, apiServer)
	}
	return resList, nil
}

func (server *ServerDiscovery) Delete(ctx context.Context, ID string) error {
	logger := app.GetGlobalLogger(ctx)
	client := resource.EtcdClient
	_, err := client.Delete(ctx, server.GetKey(ID))
	if err != nil {
		logger.WithError(err).WithFields(logrus.Fields{
			"ID": ID,
		}).Errorln("etcd delete fail")
		return err
	}
	return nil
}
