package user

import "github.com/VENI-VIDIVICI/gohub/pkg/database"

// IsEmailExist 判断邮箱是否被注册

func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExit 判断手机号是否被注册

func IsPhoneExit(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}
