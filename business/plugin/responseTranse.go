package plugin

import (
	apimanage "api-gateway/business/apiManage"
	"api-gateway/pkg/utils"
	"context"
)

// 输出参数转换
type ResponseTranse struct {
	ID     int64          `json:"id"`     // 自增id
	Name   string         `json:"name"`   // 插件的名称
	Config ResponseConfig `json:"config"` // 配置为详细信息，每个插件都不一样
}

// 参数转换的配置
type ResponseConfig struct {
	Params []ResponseParamType `json:"params"`
}

type ResponseParamType struct {
	SysPath string `json:"sysPath"` // 在系统参数中的路径
	TarPos  string `json:"tarPos"`  // 目的位置
	TarPath string `json:"tarPath"` // 目的路径
	Type    string `json:"type"`    // 值的来源类型：1. router的输入  2. 系统参数  3. 动态计算，如时间等
	Value   any    `json:"value"`   // 这里的值，可以是固定值，也可以是填充值
}

func NewResponseTranse(ctx context.Context, id int64, name string, configMap any) (IPlugin, error) {
	realConfig := ResponseConfig{}
	err := NewPluginConfig(ctx, configMap, &realConfig)
	return ResponseTranse{
		ID:     id,
		Name:   name,
		Config: realConfig,
	}, err
}

// 处理函数
func (responseTranse ResponseTranse) Process(ctx context.Context, apiManage apimanage.APIManage) (apimanage.APIManage, error) {
	rules := responseTranse.Config
	for _, rule := range rules.Params {
		sysVal := apiManage.SystemParam[rule.SysPath]
		if utils.IsNil(sysVal) {
			continue
		}
		switch rule.TarPos {
		case "body":
			apiManage.ResBody[rule.TarPath] = sysVal
		}
	}
	return apiManage, nil
}
