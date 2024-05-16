package clientrouter

import (
	"api-gateway/repository/mysql"
	"context"
)

func (apiRouter APIRouter) ToPO(ctx context.Context) (mysql.RouterPO, error) {
	po := mysql.RouterPO{
		ID:          apiRouter.ID,
		Name:        apiRouter.Name,
		Discription: apiRouter.Discription,
		LocationURL: apiRouter.LocationURL,
		Target:      apiRouter.Target,
	}
	return po, nil
}

// 路由注册
func Register(ctx context.Context, apiRouter APIRouter) (APIRouter, error) {
	// 暂时将数据存储在内存
	po, err := apiRouter.ToPO(ctx)
	if err != nil {
		return APIRouter{}, err
	}
	id, err := mysql.CreatRouterPO(ctx, po)
	if err != nil {
		return APIRouter{}, err
	}
	apiRouter.ID = id

	return apiRouter, nil
}

// 服务查询
func Search(ctx context.Context, ID int64) (*APIRouter, error) {
	po, err := mysql.SearchRouterPO(ctx, ID)
	if err != nil {
		return nil, err
	}
	apiServer, err := NewAPIRouterByPO(ctx, *po)
	if err != nil {
		return nil, err
	}
	return apiServer, nil
}

// 更新服务
func Update(ctx context.Context, apiRouter APIRouter) error {
	po, err := apiRouter.ToPO(ctx)
	if err != nil {
		return err
	}
	err = mysql.UpdateRouterPO(ctx, po)
	if err != nil {
		return err
	}
	return nil
}

func Run(ctx context.Context) {
	// 暂时将数据存储在内存
}
