package app

import (
	"api-gateway/pkg/e"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SseFinishEvent = "finish"
	SseErrorEvent  = "error"
	SseMsgEvent    = "message"
	SseAckEvent    = "ack"
	SseAnswerEvent = "answer"
	SseProbeEvent  = "probe"

	SseAgentExceptionEvent = "exception"
)

// Response common响应内容
// swagger:response Response
type Response struct {
	Code  e.ErrorCode `json:"code"`
	Msg   string      `json:"msg"`
	Hint  string      `json:"hint,omitempty"`
	Data  any         `json:"data"`
	LogID string      `json:"logId"`
}

// ResponseSuccess 成功内容返回
func ResponseSuccess(c *gin.Context, data any) {
	gLogID := GetGlobalLogID(c)
	c.JSON(http.StatusOK, Response{
		Code:  e.Success,
		Msg:   e.GetMsg(e.Success),
		Data:  data,
		LogID: gLogID,
	})
}

// ResponseSuccessMsg 成功内容返回
func ResponseSuccessMsg(c *gin.Context, httpCode int, errCode e.ErrorCode, msg string, data any) {
	gLogID := GetGlobalLogID(c)
	c.JSON(httpCode, Response{
		Code:  errCode,
		Msg:   msg,
		Data:  data,
		LogID: gLogID,
	})
}

// ResponseError 错误内容返回
func ResponseError(c *gin.Context, httpCode int, errCode e.ErrorCode) {
	gLogID := GetGlobalLogID(c)
	c.JSON(httpCode, Response{
		Code:  errCode,
		Msg:   e.GetMsg(errCode),
		LogID: gLogID,
	})
}

// ResponseDetailMsg 错误内容返回+报错详情
func ResponseDetailMsg(c *gin.Context, httpCode int, errCode e.ErrorCode, errMsg string) {
	gLogID := GetGlobalLogID(c)
	c.JSON(httpCode, Response{
		Code:  errCode,
		Msg:   errMsg,
		LogID: gLogID,
	})
}

// ResponseDetailMsg 错误内容返回+报错详情
func ResponseDetailMsgHint(c *gin.Context, httpCode int, errCode e.ErrorCode, errMsg, hintMsg string) {
	gLogID := GetGlobalLogID(c)
	c.JSON(httpCode, Response{
		Code:  errCode,
		Msg:   errMsg,
		Hint:  hintMsg,
		LogID: gLogID,
	})
}

type ChatChanStruct struct {
	SseEvent string
	Data     any
}

func SseResponse(c *gin.Context, chanStream chan ChatChanStruct, clientEndFunc func(c *gin.Context)) {
	clientCloseFlag := true
	c.Writer.Header().Set("Content-Type", "text/event-stream;charset=utf-8")
	c.Writer.Header().Add("X-Accel-Buffering", "no")

	c.Stream(func(w io.Writer) bool {
		select {
		case <-c.Writer.CloseNotify():
			// 客户端主动关闭
			return false
		case msg, ok := <-chanStream:
			if ok {
				c.SSEvent(msg.SseEvent, msg.Data)
				if msg.SseEvent == SseFinishEvent {
					clientCloseFlag = false
				}
				return true
			}
			return false
		}
	})
	// client主动关闭处理
	if clientCloseFlag && clientEndFunc != nil {
		clientEndFunc(c)
	}
}
