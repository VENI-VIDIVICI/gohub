package user

import (
	"github.com/VENI-VIDIVICI/gohub/app/models"
	"github.com/VENI-VIDIVICI/gohub/pkg/database"
	"github.com/VENI-VIDIVICI/gohub/pkg/hash"
)

type User struct {
	models.BaseModel

	Name string `json:"name,omitempty"`
	// json:"-" ，这是在指示 JSON 解析器忽略字段
	Email    string `json:"-"`
	Phone    string `json:"-"`
	PassWord string `json:"-"`

	models.CommonTimestamosField
}

func (useModel *User) Create() {
	database.DB.Create(&useModel)
}

func (useModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(useModel.PassWord, _password)
}
