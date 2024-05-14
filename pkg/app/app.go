package app

import (
	"api-gateway/data/static"
	"api-gateway/pkg/utils"
	"context"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// SetGlobalLogger 对gin设置全局logger
func SetGlobalLogger(c *gin.Context, logger *log.Entry) {
	c.Set(static.GlobalLogger, logger)
}

// SetCronJobGlobalLogger 对ctx设置logger，并返回ctx
func SetCronJobGlobalLogger(ctx context.Context, logger *log.Entry) context.Context {
	return context.WithValue(ctx, static.GlobalLogger, logger)
}

// GetGlobalLogger 获取全局logger
func GetGlobalLogger(ctx context.Context) *log.Entry {
	logger, exists := ctx.Value(static.GlobalLogger).(*log.Entry)
	if exists {
		return logger
	}
	return log.NewEntry(log.StandardLogger())
}

// GetGlobalLogID 获取全局logid
func GetGlobalLogID(ctx context.Context) string {
	logger, exists := ctx.Value(static.GlobalLogger).(*log.Entry)
	if exists && logger.Data[static.GcLogID] != nil &&
		logger.Data[static.GcLogID] != "" {
		return logger.Data[static.GcLogID].(string)
	}
	return utils.GenXID(static.LogPrefix)
}

// GetLogIDFromRequest 从gin中获取logid
func GetLogIDFromRequest(c *gin.Context) string {
	if len(c.Request.Header[static.GcLogID]) == 0 || c.Request.Header[static.GcLogID][0] == "" {
		return utils.GenXID(static.LogPrefix)
	}
	return static.LogPrefix + c.Request.Header[static.GcLogID][0]
}

// GetPerfIDFromRequest 获取用户perfID
func GetPerfIDFromRequest(c *gin.Context) string {
	if len(c.Request.Header[static.GcPerfID]) > 0 && c.Request.Header[static.GcPerfID][0] != "" {
		return c.Request.Header[static.GcPerfID][0]
	}
	return ""
}
