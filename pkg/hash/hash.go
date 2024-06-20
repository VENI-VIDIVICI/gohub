package hash

import (
	"github.com/VENI-VIDIVICI/gohub/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// 进行加密
func BcryptHash(str string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	logger.LogIf(err)
	return string(bytes)
}

// 对比

func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	return err == nil
}

func BcryptIsHashed(str string) bool {
	return len(str) == 60
}
