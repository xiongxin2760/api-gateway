package clientrouter

import (
	clientrouter "api-gateway/business/clientRouter"
	"api-gateway/library/types"
	"api-gateway/pkg/app"
	"api-gateway/pkg/e"
	"api-gateway/pkg/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 提供给客户端的router

// 注册
func Register(c *gin.Context) {
	logger := app.GetGlobalLogger(c)
	// 参数解析
	var req clientrouter.APIRouter
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.WithError(err).Error("invalid json data")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.InvalidParams, err.Error())
		return
	}
	logger.Info(utils.ObjectToLogStr(req))
	result, err := clientrouter.Register(c, req)
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
	var req types.RouterAPISearch
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.WithError(err).Error("invalid json data")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.InvalidParams, err.Error())
		return
	}
	logger.Info(utils.ObjectToLogStr(req))
	result, err := clientrouter.Search(c, req.ID)
	if err != nil {
		logger.WithError(err).Error("Get Msg fail")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.Error, err.Error())
		return
	}
	app.ResponseSuccess(c, result)
}

func Update(c *gin.Context) {
	logger := app.GetGlobalLogger(c)
	// 参数解析
	var req clientrouter.APIRouter
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.WithError(err).Error("invalid json data")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.InvalidParams, err.Error())
		return
	}
	if req.ID == 0 {
		err := errors.New("id can't be 0")
		logger.WithError(err).Error("Get Msg fail")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.Error, err.Error())
		return
	}
	logger.Info(utils.ObjectToLogStr(req))
	err := clientrouter.Update(c, req)
	if err != nil {
		logger.WithError(err).Error("Get Msg fail")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.Error, err.Error())
		return
	}
	app.ResponseSuccess(c, nil)
}
