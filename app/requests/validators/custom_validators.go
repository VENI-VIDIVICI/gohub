package validators

// ValidatePasswordConfirm 自定义规则，检查两次密码是否正确
func ValidatePasswordConfirm(password, passwordConfirm string, errs map[string][]string) map[string][]string {
	if password != passwordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配！")
	}
	return errs
}

// ValidateVerifyCode 自定义规则，验证『手机/邮箱验证码』
// func ValidateVerifyCode(key, answer string, errs map[string][]string) map[string][]string {
// 	if ok := verifycode.NewVerifyCode().CheckAnswer(key, answer); !ok {
// 		errs["verify_code"] = append(errs["verify_code"], "验证码错误")
// 	}
// 	return errs
// }

func ValidateCaptcha(id string, answer string, errs map[string][]string) map[string][]string {
	return errs
}
