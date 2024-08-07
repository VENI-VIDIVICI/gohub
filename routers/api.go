package routers

import (
	"net/http"

	"github.com/VENI-VIDIVICI/gohub/app/http/controllers/api/v1/auth"
	"github.com/VENI-VIDIVICI/gohub/app/http/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	v1.Use(middlewares.LimitIP("200-H"))
	{
		// 注册一个路由
		v1.GET("/", func(c *gin.Context) {
			// 以 JSON 格式响应
			c.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})
		authGroup := v1.Group("/auth")
		authGroup.Use(middlewares.LimitIP("1000-H"))
		suc := new(auth.SignupController)
		verfify := new(auth.VerifyCodeController)
		lgc := new(auth.LoginController)
		// 判断手机是否已注册
		authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
		authGroup.POST("/signup/email/exist", suc.IsEmailExist)
		// SignupUsing
		authGroup.POST("/signup/using", suc.SignupUsing)
		// 获取验证码
		authGroup.POST("/verify-codes/captcha", verfify.ShowCaptcha)
		authGroup.POST("/login/using-password", lgc.LoginByPassword)

	}
}
