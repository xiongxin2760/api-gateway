package upstreamServer

import (
	upstreamserver "api-gateway/business/upstreamServer"
	"api-gateway/library/types"
	"api-gateway/pkg/app"
	"api-gateway/pkg/e"
	"api-gateway/pkg/utils"
	"errors"
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
	err := upstreamserver.Register(c, req)
	if err != nil {
		logger.WithError(err).Error("Get Msg fail")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.Error, err.Error())
		return
	}
	app.ResponseSuccess(c, nil)
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
	apiServer, err := upstreamserver.Search(c, req.ID)
	if err != nil {
		logger.WithError(err).Error("Get Msg fail")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.Error, err.Error())
		return
	}
	app.ResponseSuccess(c, apiServer)
}

func Update(c *gin.Context) {
	logger := app.GetGlobalLogger(c)
	// 参数解析
	var req upstreamserver.APIServer
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.WithError(err).Error("invalid json data")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.InvalidParams, err.Error())
		return
	}
	if req.ID == 0 {
		err := errors.New("id can't be zero")
		logger.WithError(err).Error("Get Msg fail")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.Error, err.Error())
		return
	}
	logger.Info(utils.ObjectToLogStr(req))
	err := upstreamserver.Update(c, req)
	if err != nil {
		logger.WithError(err).Error("Get Msg fail")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.Error, err.Error())
		return
	}
	app.ResponseSuccess(c, nil)
}
