package proxy

import (
	apimanage "api-gateway/business/apiManage"
	upstreamserver "api-gateway/business/upstreamServer"

	"github.com/gin-gonic/gin"
)

func RunProxy(c *gin.Context, serverID string) error {
	// 获取serverID
	server, err := upstreamserver.ServerDiscoveryHandle.Search(c, serverID)
	if err != nil {
		return err
	}
	serverURL, err := server.GetTargetURL(c)
	if err != nil {
		return err
	}
	api, err := apimanage.NewAPIManage(c, server.ID, serverURL, server.Retry, server.Timeout)
	if err != nil {
		return err
	}
	// 中间件执行
	for _, plugin := range server.Plugins {
		api, err = plugin.Process(c, api)
		if err != nil {
			return err
		}
	}
	err = api.ParamReset()
	if err != nil {
		return err
	}
	api.Run()
	return nil
}
