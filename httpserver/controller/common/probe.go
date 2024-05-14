package common

import (
	"api-gateway/pkg/app"

	"github.com/gin-gonic/gin"
)

// Probe 探活
func Probe(c *gin.Context) {
	app.ResponseSuccess(c, c.Request.Host)
}
