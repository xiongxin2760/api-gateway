package apimanage

import (
	"api-gateway/pkg/app"
	"api-gateway/pkg/utils"
	"context"
	"net/http"
	"net/url"
)

// 定义请求代理的代理逻辑，调整请求的param和header，并代理至目标接口
func director(ctx context.Context, destURL *url.URL) func(req *http.Request) {
	return func(req *http.Request) {
		tempURL, _ := url.Parse(destURL.String())
		req.URL = destURL
		req.Host = tempURL.Host
	}
}

// 定义请求代理的错误处理器，用于处理proxy前后发生的错误信息，以400返回json错误信息
// 在这里可以实现重试机制
func errorHandler(ctx context.Context) func(w http.ResponseWriter, re *http.Request, err error) {
	return func(w http.ResponseWriter, re *http.Request, err error) {
		if err != nil {
			logger := app.GetGlobalLogger(ctx)
			logger.WithError(err).Warnln("http proxy error")
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		errorJSON := utils.ToJSONString(app.Response{
			Code:  http.StatusBadRequest,
			Msg:   err.Error(),
			LogID: app.GetGlobalLogID(ctx),
		})
		_, err = w.Write([]byte(errorJSON))
		if err != nil {
			logger := app.GetGlobalLogger(ctx)
			logger.WithError(err).Warnln("http Write Err error")
		}
	}
}

// 函数响应的修改和映射
func modifyResponse(ctx context.Context) func(*http.Response) error {
	return func(res *http.Response) error {
		return nil
	}
}
