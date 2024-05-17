package apimanage

import (
	"api-gateway/pkg/app"
	"api-gateway/pkg/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// api的代理
type APIManage struct {
	Name        string
	URL         *url.URL
	Retry       int
	Timeout     int // Ms
	SystemParam map[string]any
	Param       map[string]any
	Header      map[string]any
	Body        map[string]any
	GinCtx      *gin.Context
}

func NewAPIManage(c *gin.Context, serverID int64, serverURL *url.URL, retry int, timeOut int) (APIManage, error) {
	name := fmt.Sprintf("server_%d", serverID)
	apiManage := APIManage{
		Name:        name,
		URL:         serverURL,
		Retry:       retry,
		Timeout:     timeOut,
		SystemParam: nil,
		Param:       make(map[string]any),
		Header:      make(map[string]any),
		Body:        make(map[string]any),
		GinCtx:      c,
	}
	apiManage.paramInit()
	return apiManage, nil
}

func (api *APIManage) Run() {
	director := director(api.GinCtx, api.URL)
	// 定义错误处理器
	errhandler := errorHandler(api.GinCtx)
	responseHandler := modifyResponse(api.GinCtx)
	proxy := &httputil.ReverseProxy{Director: director, ErrorHandler: errhandler, ModifyResponse: responseHandler}

	// 超时时间
	proxy.Transport = &http.Transport{
		ResponseHeaderTimeout: time.Duration(api.Timeout) * time.Millisecond,
	}

	proxy.ServeHTTP(api.GinCtx.Writer, api.GinCtx.Request)
}

func (api *APIManage) ParamReset() error {
	// Header写回
	c := api.GinCtx
	logger := app.GetGlobalLogger(c)
	newHeader := make(http.Header)
	for k, v := range api.Header {
		if strv, ok := v.(string); ok {
			newHeader[k] = []string{strv}
		}
	}
	c.Request.Header = newHeader
	// Param 写回
	destURLQuery := api.URL.Query()
	for k, v := range api.Param {
		if strv, ok := v.(string); ok {
			destURLQuery.Add(k, strv)
		}
	}
	api.URL.RawQuery = destURLQuery.Encode()
	// body 写回
	bodyByte, err := json.Marshal(api.Body)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"api.Body": utils.ObjectToLogStr(api.Body),
		}).Errorln("marshal fail")
		return err
	}
	c.Request.Body = NewAPIReadCloser(bodyByte)
	return nil
}

// 从 gin.Context 中读取原始参数
func (api *APIManage) paramInit() error {
	// 超时时间
	if api.Timeout == 0 {
		api.Timeout = 1000
	}
	// param 获取
	c := api.GinCtx
	logger := app.GetGlobalLogger(c)
	sourceQuery, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"c.Request.URL.RawQuery": c.Request.URL.RawQuery,
		}).Errorln("get url param fail")
		return err
	}
	for k, v := range sourceQuery {
		api.Param[k] = v
	}

	// Header 读取
	for k, v := range c.Request.Header {
		api.Header[k] = v
	}

	// Body 读取
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
	err = json.Unmarshal(buf, &api.Body)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"buf": string(buf),
		}).Errorln("unmarshal fail")
		return err
	}
	// 将系统参数提取出来
	if sysParam, exist := api.Body["_systemParams"]; exist {
		if sysParamMap, ok := sysParam.(map[string]any); ok {
			api.SystemParam = sysParamMap
		}
		delete(api.Body, "_systemParams") //
	}
	return nil
}
