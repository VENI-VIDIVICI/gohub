package limiter

import (
	"strings"

	"github.com/VENI-VIDIVICI/gohub/pkg/config"
	"github.com/VENI-VIDIVICI/gohub/pkg/logger"
	"github.com/VENI-VIDIVICI/gohub/pkg/redis"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

const LimitKEY = "limiter-once"

func GetKeyId(c *gin.Context) string {
	return c.ClientIP()
}

func GetKeyRouteWithId(c *gin.Context) string {
	// c.FullPath()
	return routeToKeyString(c.FullPath() + c.ClientIP())
}

func CheckRate(c *gin.Context, key string, formatted string) (limiter.Context, error) {
	var content limiter.Context
	rate, err := limiter.NewRateFromFormatted(formatted)
	if err != nil {
		logger.LogIf(err)
		return content, err
	}

	store, err := sredis.NewStoreWithOptions(redis.Redis.Client, limiter.StoreOptions{
		Prefix: config.GetString("app.name") + ":limiter",
	})
	if err != nil {
		logger.LogIf(err)
		return content, err
	}
	limiterObj := limiter.New(store, rate)
	if c.GetBool(LimitKEY) {
		return limiterObj.Peek(c, key)
	}
	c.Set(LimitKEY, true)
	return limiterObj.Get(c, key)
}

func routeToKeyString(routeName string) string {
	routeName = strings.ReplaceAll(routeName, "/", "-")
	routeName = strings.ReplaceAll(routeName, ":", "_")
	return routeName
}
