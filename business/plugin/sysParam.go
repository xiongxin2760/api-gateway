package plugin

import (
	apimanage "api-gateway/business/apiManage"
	"api-gateway/pkg/utils"
	"context"
)

// 系统参数填充 -- 用系统参数去覆盖请求参数
type SystemParamTranse struct {
	ID     int64                   `json:"id"`     // 自增id
	Name   string                  `json:"name"`   // 插件的名称
	Config SystemParamTranseConfig `json:"config"` // 配置为详细信息，每个插件都不一样
}

// 参数转换的配置
type SystemParamTranseConfig struct {
	Params []SystemParamType `json:"params"`
}

type SystemParamType struct {
	SysPath string `json:"sysPath"` // 在系统参数中的路径
	TarPos  string `json:"tarPos"`  // 目的位置
	TarPath string `json:"tarPath"` // 目的路径
	Type    string `json:"type"`    // 值的来源类型：1. router的输入  2. 系统参数  3. 动态计算，如时间等
	Value   any    `json:"value"`   // 这里的值，可以是固定值，也可以是填充值
}

func NewSystemParamTranse(ctx context.Context, id int64, name string, configMap any) (IPlugin, error) {
	realConfig := SystemParamTranseConfig{}
	err := NewPluginConfig(ctx, configMap, &realConfig)
	return SystemParamTranse{
		ID:     id,
		Name:   name,
		Config: realConfig,
	}, err
}

// 处理函数
func (systemParamTranse SystemParamTranse) Process(ctx context.Context, apiManage apimanage.APIManage) (apimanage.APIManage, error) {
	rules := systemParamTranse.Config
	for _, rule := range rules.Params {
		sysVal := apiManage.SystemParam[rule.SysPath]
		if utils.IsNil(sysVal) {
			continue
		}
		switch rule.TarPos {
		case "header":
			apiManage.Header[rule.TarPath] = sysVal
		case "param":
			apiManage.Param[rule.TarPath] = sysVal
		case "body":
			apiManage.Body[rule.TarPath] = sysVal
		}
	}
	return apiManage, nil
}
