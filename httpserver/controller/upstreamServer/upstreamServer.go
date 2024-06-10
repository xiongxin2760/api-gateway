package upstreamServer

import (
	upstreamserver "api-gateway/business/upstreamServer"
	"api-gateway/library/types"
	"api-gateway/pkg/app"
	"api-gateway/pkg/e"
	"api-gateway/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// server相关的操作
// 注册server
func Register(c *gin.Context) {
	logger := app.GetGlobalLogger(c)
	// 参数解析
	var req upstreamserver.APIServer
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.WithError(err).Error("invalid json data")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.InvalidParams, err.Error())
		return
	}
	logger.Info(utils.ObjectToLogStr(req))
	result, err := upstreamserver.ServerDiscoveryHandle.Register(c, req)
	if err != nil {
		logger.WithError(err).Error("Get Msg fail")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.Error, err.Error())
		return
	}
	app.ResponseSuccess(c, result)
}

func Search(c *gin.Context) {
	logger := app.GetGlobalLogger(c)
	// 参数解析
	var req types.ServerAPISearch
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.WithError(err).Error("invalid json data")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.InvalidParams, err.Error())
		return
	}
	logger.Info(utils.ObjectToLogStr(req))
	apiServer, err := upstreamserver.ServerDiscoveryHandle.Search(c, req.ID)
	if err != nil {
		logger.WithError(err).Error("Get Msg fail")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.Error, err.Error())
		return
	}
	app.ResponseSuccess(c, apiServer)
}

func SearchList(c *gin.Context) {
	logger := app.GetGlobalLogger(c)
	apiServer, err := upstreamserver.ServerDiscoveryHandle.SearchList(c)
	if err != nil {
		logger.WithError(err).Error("Get Msg fail")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.Error, err.Error())
		return
	}
	app.ResponseSuccess(c, apiServer)
}

func Delete(c *gin.Context) {
	logger := app.GetGlobalLogger(c)
	// 参数解析
	var req types.ServerAPISearch
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.WithError(err).Error("invalid json data")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.InvalidParams, err.Error())
		return
	}
	logger.Info(utils.ObjectToLogStr(req))
	err := upstreamserver.ServerDiscoveryHandle.Delete(c, req.ID)
	if err != nil {
		logger.WithError(err).Error("Get Msg fail")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.Error, err.Error())
		return
	}
	app.ResponseSuccess(c, nil)
}
