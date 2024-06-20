package requests

import (
	"github.com/VENI-VIDIVICI/gohub/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}
type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

// 注册
type SignupRequest struct {
	Name            string `valid:"name" json:"name"`
	PassWord        string `valid:"password" json:"password,omitempty"`
	PassWordConfirm string `valid:"password_confirm" json:"password_confirm,omitempty"`
}

func ValidateSignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}
	return validate(data, rules, messages)
}

func ValidateSignupEmailExist(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}
	messgaes := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度必须大于 4",
			"max:Email 长度必须小于 30",
			"email:Email 格式不正确，请输入正确的格式",
		},
	}
	return validate(data, rules, messgaes)
}

func SignupUsing(data interface{}, c *gin.Context) map[string][]string {
	ruler := govalidator.MapData{
		// "phone": []string{"required", "digits:11", "not_exists:users,phone"},
		// , "not_exists,name"
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
		},
	}
	_data := data.(*SignupRequest)
	errs := validate(data, ruler, messages)
	// 校验验证码和code
	errs = validators.ValidatePasswordConfirm(_data.PassWord, _data.PassWordConfirm, errs)
	return errs
	// errs := validators.Va
}
