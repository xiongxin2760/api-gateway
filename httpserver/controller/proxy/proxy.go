package proxy

import (
	"api-gateway/business/proxy"
	"api-gateway/pkg/app"
	"api-gateway/pkg/e"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// 跑任意工具
func Run(c *gin.Context) {
	logger := app.GetGlobalLogger(c)
	// 获取参数
	routerID := strings.TrimLeft(c.Param("path"), "/path")
	// 将参数转换为数字
	serverID, err := strconv.Atoi(routerID)
	if err != nil {
		logger.WithError(err).Error("invalid json data")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.InvalidParams, err.Error())
		return
	}
	// 用toolID查询
	err = proxy.RunProxy(c, int64(serverID))
	if err != nil {
		logger.WithError(err).Error("invalid json data")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.InvalidParams, err.Error())
		return
	}
}
