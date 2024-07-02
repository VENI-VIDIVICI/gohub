package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/VENI-VIDIVICI/gohub/app/cmd"
	"github.com/VENI-VIDIVICI/gohub/bootstrap"
	btsConfig "github.com/VENI-VIDIVICI/gohub/config"
	"github.com/VENI-VIDIVICI/gohub/pkg/config"
	"github.com/VENI-VIDIVICI/gohub/pkg/console"
	"github.com/spf13/cobra"
)

func init() {
	// 加载配置
	btsConfig.Initialize()
}
func main() {
	env := cmd.Env
	var rootCmd = &cobra.Command{
		Use:   "Gohub",
		Short: "A simple forum project",
		Long:  `Defalut will run "serve" command, you can use "-h" flag to see all subcommands`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			flag.Parse()
			config.InitConfig(env)
			bootstrap.SetupLogger()
			bootstrap.SetupDB()
		},
	}
	rootCmd.AddCommand(cmd.CmdServe)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}

}
