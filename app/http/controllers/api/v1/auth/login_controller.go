package auth

import (
	v1 "github.com/VENI-VIDIVICI/gohub/app/http/controllers/api/v1"
	"github.com/VENI-VIDIVICI/gohub/app/requests"
	"github.com/VENI-VIDIVICI/gohub/pkg/auth"
	"github.com/VENI-VIDIVICI/gohub/pkg/jwt"
	"github.com/VENI-VIDIVICI/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	v1.BaseAPIController
}

func (lc *LoginController) LoginByPassword(c *gin.Context) {
	// 验证表单
	request := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
		return
	}
	user, err := auth.Attempt(request.LoginID, request.Password)
	if err != nil {
		response.Unauthorized(c, "账号不存在或密码错误")
	} else {
		token := jwt.NewJwt().IssueToken(user.GetStringID(), user.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
