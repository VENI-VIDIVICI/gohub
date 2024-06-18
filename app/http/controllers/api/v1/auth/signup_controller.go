package auth

import (
	v1 "github.com/VENI-VIDIVICI/gohub/app/http/controllers/api/v1"
	"github.com/VENI-VIDIVICI/gohub/app/models/user"
	"github.com/VENI-VIDIVICI/gohub/app/requests"
	"github.com/VENI-VIDIVICI/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	request := requests.SignupPhoneExistRequest{}
	ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist)
	if ok == false {
		return
	}
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExit(request.Phone),
	})
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}

	ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist)
	if ok == false {
		return
	}

	//  检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
