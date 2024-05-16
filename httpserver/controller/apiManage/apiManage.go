package apimanage

import (
	apimanage "api-gateway/business/apiManage"
	"api-gateway/pkg/app"
	"api-gateway/pkg/e"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 跑任意工具
func Run(c *gin.Context) {
	logger := app.GetGlobalLogger(c)
	// 获取参数
	routerID := c.Param("path")
	// 将参数转换为数字
	toolID, err := strconv.Atoi(routerID)
	if err != nil {
		logger.WithError(err).Error("invalid json data")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.InvalidParams, err.Error())
		return
	}
	// 用toolID查询
	apimanage, err := apimanage.NewAPIManage(c, int64(toolID))
	if err != nil {
		logger.WithError(err).Error("invalid json data")
		app.ResponseDetailMsg(c, http.StatusInternalServerError, e.InvalidParams, err.Error())
		return
	}
	apimanage.Run()
}
