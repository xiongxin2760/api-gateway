package upstreamserver

import (
	"api-gateway/library/types"
	"api-gateway/pkg/app"
	"api-gateway/pkg/utils"
	"api-gateway/repository/mysql"
	"context"
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type APIServer struct {
	ID          int64             `json:"id"`
	Name        string            `json:"name"`
	Discription string            `json:"discription"`
	Timeout     int               `json:"timeout"`
	Retry       int               `json:"retry"`
	Balance     string            `json:"balance"`
	Service     []types.ServerAPI `json:"service"` // TODO：升级为服务发现  	// Plugins     map[string]any    `json:"plugins"` // 待定
}

func NewAPIServerByPO(ctx context.Context, po mysql.ServerPO) (*APIServer, error) {
	logger := app.GetGlobalLogger(ctx)
	service := []types.ServerAPI{}
	if len(po.Service) > 0 {
		err := json.Unmarshal([]byte(po.Service), &service)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"po.Service": po.Service,
			}).Errorln("creat mannual connect fail")
			return nil, err
		}
	}

	return &APIServer{
		Name:        po.Name,
		Discription: po.Discription,
		Timeout:     po.Timeout,
		Retry:       po.Retry,
		Balance:     po.Balance,
		Service:     service,
	}, nil
}

func (apiServer APIServer) ToPO(ctx context.Context) (mysql.ServerPO, error) {
	logger := app.GetGlobalLogger(ctx)
	serverByte, err := json.Marshal(apiServer.Service)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"apiServer.Service": utils.ObjectToLogStr(apiServer.Service),
		}).Errorln("creat mannual connect fail")
		return mysql.ServerPO{}, err
	}
	po := mysql.ServerPO{
		ID:          apiServer.ID,
		Name:        apiServer.Name,
		Discription: apiServer.Discription,
		Timeout:     apiServer.Timeout,
		Retry:       apiServer.Retry,
		Balance:     apiServer.Balance,
		Service:     string(serverByte),
	}
	return po, nil
}

// 服务注册
func Register(ctx context.Context, apiServer APIServer) error {
	// 暂时将数据存储在内存
	// apiServer := NewAPIServer(req)
	po, err := apiServer.ToPO(ctx)
	if err != nil {
		return err
	}
	id, err := mysql.CreatServerPO(ctx, po)
	if err != nil {
		return err
	}
	apiServer.ID = id

	return nil
}

// 服务查询
func Search(ctx context.Context, ID int64) (*APIServer, error) {
	po, err := mysql.SearchServerPO(ctx, ID)
	if err != nil {
		return nil, err
	}
	apiServer, err := NewAPIServerByPO(ctx, *po)
	if err != nil {
		return nil, err
	}
	return apiServer, nil
}

// 更新服务
func Update(ctx context.Context, apiServer APIServer) error {
	po, err := apiServer.ToPO(ctx)
	if err != nil {
		return err
	}
	err = mysql.UpdateServerPO(ctx, po)
	if err != nil {
		return err
	}
	return nil
}
