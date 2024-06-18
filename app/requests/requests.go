package requests

import (
	"fmt"

	"github.com/VENI-VIDIVICI/gohub/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// ValidatorFunc 验证函数类型
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func validate(data interface{}, rules, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}
	return govalidator.New(opts).ValidateStruct()
}

func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	if err := c.ShouldBind(obj); err != nil {
		// 解析失败，返回 422 状态码和错误信息
		response.BadRequest(c, err)
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return false
	}
	// 表单验证
	errs := handler(obj, c)
	// errs 返回长度等于零即通过，大于 0 即有错误发生
	if len(errs) > 0 {
		// 验证失败，返回 422 状态码和错误信息
		response.ValidationError(c, errs)
		return false
	}
	return true
}
