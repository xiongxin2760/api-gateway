package apimanage

import (
	clientrouter "api-gateway/business/clientRouter"
	upstreamserver "api-gateway/business/upstreamServer"
	"fmt"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
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

func NewAPIManage(c *gin.Context, routerID int64) (APIManage, error) {
	// 查询
	apiRouter, err := clientrouter.Search(c, routerID)
	if err != nil {
		return APIManage{}, err
	}
	server, err := upstreamserver.Search(c, apiRouter.Target)
	if err != nil {
		return APIManage{}, err
	}
	name := fmt.Sprintf("%d_%d", apiRouter.ID, server.ID)
	if err != nil {
		return APIManage{}, err
	}
	tarURL, err := server.GetTargetURL(c)
	if err != nil {
		return APIManage{}, err
	}
	apiManage := APIManage{
		Name:        name,
		URL:         tarURL,
		Retry:       server.Retry,
		Timeout:     server.Timeout,
		SystemParam: nil,
		Param:       make(map[string]any),
		Header:      make(map[string]any),
		Body:        make(map[string]any),
		GinCtx:      c,
	}
	return apiManage, nil
}

func (api *APIManage) Run() {
	director := director(api.GinCtx, api.URL)
	// 定义错误处理器
	errhandler := errorHandler(api.GinCtx)
	proxy := &httputil.ReverseProxy{Director: director, ErrorHandler: errhandler}

	// ModifyResponse: nil
	// 超时时间在这里设置
	// proxy.Transport = &http.Transport{
	// 	ResponseHeaderTimeout: 5,
	// }

	proxy.ServeHTTP(api.GinCtx.Writer, api.GinCtx.Request)
}
