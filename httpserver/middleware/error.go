package middleware

import (
	"api-gateway/library/resource"
	"api-gateway/pkg/app"
	"api-gateway/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewDeferErrorMiddleware() gin.HandlerFunc {
	return gin.RecoveryWithWriter(resource.RecoverLogger.Out, customRecoverFunc)
}

func customRecoverFunc(c *gin.Context, err interface{}) {
	logger := resource.RecoverLogger
	logger.Errorf("panic: catch error: %+v", err)

	app.ResponseError(c, http.StatusInternalServerError, e.Error)
}
