package bootstrap

import (
	"fmt"

	"github.com/VENI-VIDIVICI/gohub/pkg/config"
	"github.com/VENI-VIDIVICI/gohub/pkg/redis"
)

func SetupRedis() {
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
