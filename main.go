package main

import (
	"flag"
	"fmt"

	"github.com/VENI-VIDIVICI/gohub/bootstrap"
	btsConfig "github.com/VENI-VIDIVICI/gohub/config"
	"github.com/VENI-VIDIVICI/gohub/pkg/config"
	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.Initialize()
}
func main() {
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)
	bootstrap.SetupDB()
	// 初始化 Gin 实例
	router := gin.New()
	bootstrap.SetupRoute(router)
	// 初始化 DB
	// 运行服务
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
