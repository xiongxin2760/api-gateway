package plugin

import (
	apimanage "api-gateway/business/apiManage"
	"api-gateway/pkg/app"
	"context"
	"encoding/json"
	"strings"

	"github.com/sirupsen/logrus"
)

// 参数转换的插件
type ParamTranse Plugin

const (
	ParamTranseTypeRouter = "router"
	ParamTranseTypeSystem = "system"
	ParamTranseTypeDate   = "date"
)

// 参数转换的配置
type ParamTranseConfig struct {
	Params []ParamType `json:"params"`
}

type ParamType struct {
	RouterPosition string `json:"routerPosition"` // 字段来源的地方
	TargetPosition string `json:"targetPosition"` // 字段的目的转换位置
	Type           string `json:"type"`           // 值的来源类型：1. router的输入  2. 系统参数  3. 动态计算，如时间等
	Value          any    `json:"value"`          // 这里的值，可以是固定值，也可以是填充值
}

func NewParamTranse(ctx context.Context, id int64, name string, config string) (ParamTranse, error) {
	logger := app.GetGlobalLogger(ctx)
	realConfig := ParamTranseConfig{}
	err := json.Unmarshal([]byte(config), &realConfig)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"config": config,
		}).Errorln("unmarsh fail")
		return ParamTranse{}, nil
	}
	plugin := Plugin{
		ID:     id,
		Name:   name,
		Config: realConfig,
	}
	paramTranse := ParamTranse(plugin)
	return paramTranse, nil
}

func (paramTranse ParamTranse) Process(apiManage apimanage.APIManage) (apimanage.APIManage, error) {
	if rules, ok := paramTranse.Config.(ParamTranseConfig); ok {
		for _, rule := range rules.Params {
			apiManage = processRaw(apiManage, rule)
		}
	}
	return apiManage, nil
}

func processRaw(apiManage apimanage.APIManage, rule ParamType) apimanage.APIManage {
	// routerPos1, routerPos2 := resolvePos(rule.RouterPosition)
	targetPos1, targetPos2 := resolvePos(rule.TargetPosition)
	if rule.Type == ParamTranseTypeRouter {
		var val any
		// switch routerPos1 {
		// case "header":
		// 	val = apiManage.OriHeader[routerPos2]
		// case "param":
		// 	val = apiManage.OriParam[routerPos2]
		// case "body":
		// 	val = apiManage.OriBody[routerPos2]
		// }
		switch targetPos1 {
		case "header":
			apiManage.Header[targetPos2] = val
		case "param":
			apiManage.Param[targetPos2] = val
		case "body":
			apiManage.Body[targetPos2] = val
		}
	}
	return apiManage
}

func resolvePos(position string) (string, string) {
	pos1, pos2, _ := strings.Cut(position, ".")
	return pos1, pos2
}
