package auth

import (
	v1 "github.com/VENI-VIDIVICI/gohub/app/http/controllers/api/v1"
	"github.com/VENI-VIDIVICI/gohub/pkg/captcha"
	"github.com/VENI-VIDIVICI/gohub/pkg/logger"
	"github.com/VENI-VIDIVICI/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

type VerifyCodeController struct {
	v1.BaseAPIController
}

func (sc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// 生成验证码
	id, b64s, _, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志，因为验证码是用户的入口，出错时应该记 error 等级的日志
	logger.LogIf(err)
	// 返回给用户
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
