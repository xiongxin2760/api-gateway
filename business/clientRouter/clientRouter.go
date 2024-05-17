package clientrouter

import (
	"api-gateway/repository/mysql"
	"context"
)

// 这个router貌似是多余的，反倒增加管理成本
type APIRouter struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Discription string `json:"discription"`
	LocationURL string `json:"locationUrl"` // 客户端路由
	Target      int64  `json:"target"`      // 绑定的server的id
	// Plugins     any    `json:"plugins"`     // 暂时先不要
}

func NewAPIRouterByPO(ctx context.Context, po mysql.RouterPO) (*APIRouter, error) {
	// logger := app.GetGlobalLogger(ctx)
	return &APIRouter{
		ID:          po.ID,
		Name:        po.Name,
		Discription: po.Discription,
		LocationURL: po.LocationURL,
		Target:      po.Target,
	}, nil
}
