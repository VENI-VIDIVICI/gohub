package auth

import (
	v1 "github.com/VENI-VIDIVICI/gohub/app/http/controllers/api/v1"
	"github.com/VENI-VIDIVICI/gohub/app/models/user"
	"github.com/VENI-VIDIVICI/gohub/app/requests"
	"github.com/VENI-VIDIVICI/gohub/pkg/jwt"
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

func (sc *SignupController) SignupUsing(c *gin.Context) {
	request := requests.SignupRequest{}
	ok := requests.Validate(c, &request, requests.SignupUsing)
	if ok == false {
		return
	}

	_user := user.User{
		Name:     request.Name,
		PassWord: request.PassWord,
	}
	_user.Create()

	if _user.ID > 0 {
		token := jwt.NewJwt().IssueToken(_user.GetStringID(), _user.Name)
		response.CreatedJSON(c, gin.H{
			"data":  _user,
			"token": token,
		})
	} else {
		response.Abort500(c, "创建用户失败,请稍后尝试~")
	}

}
