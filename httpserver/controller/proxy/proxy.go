package proxy

import (
	"api-gateway/business/proxy"
	"api-gateway/pkg/app"
	"api-gateway/pkg/e"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 跑任意工具
func Run(c *gin.Context) {
	logger := app.GetGlobalLogger(c)
	// 获取参数
	serverID := strings.TrimLeft(c.Param("path"), "/path")
	// 用toolID查询
	err := proxy.RunProxy(c, serverID)
	if err != nil {
		logger.WithError(err).Error("invalid json data")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.InvalidParams, err.Error())
		return
	}
}
