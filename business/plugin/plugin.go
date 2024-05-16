package plugin

import apimanage "api-gateway/business/apiManage"

// 定义插件，通过插件
type Plugin struct {
	ID     int64  `json:"id"`     // 自增id
	Name   string `json:"name"`   // 插件的名称
	Config any    `json:"config"` // 配置为详细信息，每个插件都不一样
}

// 插件的方法集合定义 -- 给外界使用的接口
type IPlugin interface {
	// 运行插件，运行插件的过程就是处理 apiManage 对象的过程
	Process(apiManage apimanage.APIManage) (apimanage.APIManage, error)
}
