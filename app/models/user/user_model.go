package user

import "github.com/VENI-VIDIVICI/gohub/app/models"

type User struct {
	models.BaseModel

	Name string `json:"name,omitempty"`
	// json:"-" ，这是在指示 JSON 解析器忽略字段
	Email    string `json:"-"`
	Phone    string `json:"-"`
	PassWord string `json:"-"`

	models.CommonTimestamosField
}
