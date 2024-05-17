package plugin

import (
	apimanage "api-gateway/business/apiManage"
	"api-gateway/pkg/app"
	"api-gateway/pkg/utils"
	"context"

	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
)

var IPluginFactoryMap = map[int64]PluginFactory{
	1: NewSystemParamTranse,
}

// 插件的方法集合定义 -- 给外界使用的接口
type IPlugin interface {
	// 运行插件，运行插件的过程就是处理 apiManage 对象的过程
	Process(ctx context.Context, apiManage apimanage.APIManage) (apimanage.APIManage, error)
}

func IPluginFactory(id int64) (IPlugin, bool) {
	if id == 1 {
		return SystemParamTranse{}, true
	}
	return nil, false
}

// 插件，和前端交互的格式
type PluginVO struct {
	ID     int64  `json:"id"`     // 自增id
	Name   string `json:"name"`   // 插件的名称
	Config any    `json:"config"` // 配置为详细信息，每个插件都不一样
}

type PluginFactory func(ctx context.Context, id int64, name string, configMap any) (IPlugin, error)

func NewPluginConfig(ctx context.Context, configMap any, configObj any) error {
	logger := app.GetGlobalLogger(ctx)
	err := mapstructure.Decode(configMap, &configObj)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"config": utils.ObjectToLogStr(configMap),
		}).Errorln("mapstructure fail")
		return err
	}
	return nil
}
