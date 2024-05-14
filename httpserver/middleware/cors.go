package middleware

import (
	"api-gateway/pkg/app"
	"api-gateway/pkg/utils"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

var validHost = utils.Set{}

func init() {
	validHost.AddList([]string{
		"wikigray.baidu-int.com",                         // 知识库灰度环境
		"ku.baidu-int.com",                               // 知识库线上环境
		"cloud-test.baidu-int.com",                       // icafe测试环境
		"console.cloud-sandbox.baidu-int.com",            // icafe测试环境
		"localhost:8788",                                 // fe本地前端
		"console.cloud.baidu-int.com",                    // icafe线上环境
		"localhost:8899",                                 // fe本地前端
		"zhanglei.rl_doc_beta.dev.weiyun.baidu.com:8080", // 知识库联调地址——本地
		"zhanglei.rl_doc_beta.dev.weiyun.baidu.com",      // 知识库联调地址——本地
		"ku.dev.weiyun.baidu.com",                        // 知识库联调地址——rd
		"rlkb-sandbox.dev.weiyun.baidu.com",              // 知识库联调地址——qa1
		"ku-qa.dev.weiyun.baidu.com",                     // 知识库联调地址——qa2
		"ku-qa3.dev.weiyun.baidu.com",                    // 知识库联调地址——qa3
		"ku-qa4.dev.weiyun.baidu.com",                    // 知识库联调地址——qa4
		"infoflow.static.dev.weiyun.baidu.com",           // 端fe前端测试环境
		"note.im.baidu.com",                              // 端fe前端线上环境
		"workflow-preonline.dev.weiyun.baidu.com",        // 工作卡测试环境http
		"nginx-test.dev.weiyun.baidu.com",                // 工作卡测试环境https
		"nginx-beta.dev.weiyun.baidu.com",                // 工作卡测试环境http
		"uflow.baidu-int.com",                            // 工作卡灰度环境
		"uflow-gray.baidu-int.com",                       // 工作卡线上环境
		"eos.baidu-int.com",                              // bep online
		"eosgraytest.baidu-int.com",                      // preonline
		"toos-liantiao2.bcc-bdbl.baidu.com:8048",         // dev
		"e.im.baidu.com:8899",                            // 工作卡测试
		"172.18.178.93:8200",
		"bep-qa.dev.weiyun.baidu.com",
		"bep.baidu-int.com",
		"bep-qa.local.weiyun.baidu.com:8200",
		"itab-test.dev.weiyun.baidu.com",                // itab
		"itab-beta.dev.weiyun.baidu.com",                // itab
		"itab.baidu-int.com",                            // itab
		"itab-preoline.dev.weiyun.baidu.com",            // itab
		"itab-gray.baidu-int.com",                       // itab
		"nxstp.baidu-int.com",                           // meg online
		"xstp-offline-test0001.bcc-bdbl.baidu.com:8086", // meg dev
		"ff.baidu-int.com",
		"dev.imis.baijiahao.baidu.com:8100",
		"tianmu.baidu-int.com",            // tianmu
		"imis.baijiahao.baidu.com",        // tianmu
		"testone-test.bcc-szzj.baidu.com", // testone meg
		"testone.baidu-int.com",           // testone meg
	})
}

func NewCorsMiddleware() gin.HandlerFunc {
	return cors()
}

// Cors 处理跨域请求,支持options访问
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		gLogID := app.GetGlobalLogID(c)
		logger := app.GetGlobalLogger(c)
		// 全局url跨域通过
		// requestURLs := c.Request.Header["Referer"]
		// var requestURL string
		// if len(requestURLs) > 0 {
		// 	u, err := url.Parse(requestURLs[0])
		// 	if err == nil && u != nil {
		// 		requestURL = u.Scheme + "://" + u.Host
		// 	}
		// }

		requestURLs := c.Request.Header["Referer"]
		var allowOrigin string
		if len(requestURLs) > 0 {
			host, err := url.Parse(requestURLs[0])
			if err != nil {
				logger.WithError(err).Warningf("parse url fail, url=%s", requestURLs[0])
				for _, host := range validHost.StrList() {
					if strings.HasPrefix(requestURLs[0], host) {
						allowOrigin = host
						goto PASS
					}
				}
			}
			if validHost.Has(host.Host) {
				allowOrigin = host.Scheme + "://" + host.Host
				goto PASS
			}
		}
	PASS:
		c.Header("G-LogID", gLogID)
		c.Header("Access-Control-Allow-Origin", allowOrigin)
		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token, Adapter-data, X-Scene, X-Sessionid, x-xsrf-token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, DELETE, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		method := c.Request.Method
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}
