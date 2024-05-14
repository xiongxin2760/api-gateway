package middleware

import (
	"api-gateway/library/resource"
	"strconv"
	"strings"
	"time"

	// "icode.baidu.com/baidu/so-recsys/aichat-server/library/resource"
	// "icode.baidu.com/baidu/so-recsys/aichat-server/pkg/app"

	// staticData "icode.baidu.com/baidu/so-recsys/aichat-server/data/static_dict"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// NewAccessLogMiddleware 日志中间件
func NewAccessLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		uri := strings.Split(c.Request.RequestURI, "?")[0]
		// 设置当前请求全局 logId
		globalLogID := app.GetLogIDFromRequest(c)
		curLogger := logrus.WithFields(logrus.Fields{
			"uri":              uri,
			staticData.GcLogID: globalLogID,
		})
		app.SetGlobalLogger(c, curLogger)

		// 处理请求
		c.Next()
		// 执行时间 ms
		latencyTime := strconv.FormatInt(time.Now().Sub(startTime).Milliseconds(), 10) + "ms"
		userID := app.GetGlobalUserID(c)
		userAgent := app.GetUserAgentFromRequest(c)
		perfID := app.GetPerfIDFromRequest(c)
		// 日志格式
		resource.GinAccessLogger.WithFields(logrus.Fields{
			"method":     c.Request.Method,     // 请求方式
			"uri":        c.Request.RequestURI, // 请求路由
			"http_code":  c.Writer.Status(),    // 请求状态
			"cost_time":  latencyTime,          // 执行时间
			"ip":         c.ClientIP(),         // 请求IP
			"log_id":     globalLogID,
			"user_id":    userID,
			"user_agent": userAgent,
			"perf_id":    perfID,
		}).Infof("gin access")
	}
}
