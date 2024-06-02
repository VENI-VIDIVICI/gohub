package bootstrap

import (
	"net/http"
	"strings"

	"github.com/VENI-VIDIVICI/gohub/routers"
	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {
	registerGuard(router)
	//  注册 API 路由
	routers.RegisterAPIRoutes(router)

	//  配置 404 路由
	setup404Handler(router)
}

func registerGuard(router *gin.Engine) {
	router.Use(gin.Recovery(),
		gin.Logger())
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(ctx *gin.Context) {
		// 获取请求的类型 text html  json
		acceptString := ctx.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			ctx.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
