package user

import (
	"github.com/VENI-VIDIVICI/gohub/pkg/hash"
	"gorm.io/gorm"
)

func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(userModel.PassWord) {
		userModel.PassWord = hash.BcryptHash(userModel.PassWord)
	}
	return
}
