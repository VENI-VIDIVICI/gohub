package auth

import (
	"fmt"
	"net/http"

	v1 "github.com/VENI-VIDIVICI/gohub/app/http/controllers/api/v1"
	"github.com/VENI-VIDIVICI/gohub/app/models/user"
	"github.com/VENI-VIDIVICI/gohub/app/requests"
	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	request := requests.SignupPhoneExistRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}
	errs := requests.ValidateSignupPhoneExist(&request, c)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": errs,
		})
		fmt.Println(errs)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExit(request.Phone),
	})
}
