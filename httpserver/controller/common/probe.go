package common

import (
	"api-gateway/pkg/app"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

// Probe 探活
func Probe(c *gin.Context) {
	buf := make([]byte, 0, 1024*1024)
	tempBuf := make([]byte, 1024)
	for {
		n, err := c.Request.Body.Read(tempBuf)
		if err != nil {
			if err != io.EOF {
				log.Printf("An error occurred: %v", err)
			}
			buf = append(buf, tempBuf[:n]...)
			break
		}
		buf = append(buf, tempBuf[:n]...)
	}
	resMap := map[string]any{
		"param:": c.Request.URL.RawQuery,
		"header": c.Request.Header,
		"body":   string(buf),
	}
	app.ResponseSuccess(c, resMap)
}
