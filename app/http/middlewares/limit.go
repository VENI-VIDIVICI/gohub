package middlewares

import (
	"net/http"

	"github.com/VENI-VIDIVICI/gohub/pkg/app"
	"github.com/VENI-VIDIVICI/gohub/pkg/limiter"
	"github.com/VENI-VIDIVICI/gohub/pkg/logger"
	"github.com/VENI-VIDIVICI/gohub/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func LimitIP(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}
	return func(ctx *gin.Context) {
		key := limiter.GetKeyId(ctx)
		if ok := limitHandler(ctx, key, limit); !ok {
			return
		}
		ctx.Next()
	}
}

func LimitPerRoute(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}
	return func(c *gin.Context) {
		// limiter.LimitKEY
		c.Set(limiter.LimitKEY, false)
		key := limiter.GetKeyRouteWithId(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}
}

func limitHandler(c *gin.Context, key string, limit string) bool {
	rate, err := limiter.CheckRate(c, key, limit)
	if err != nil {
		logger.LogIf(err)
		response.Abort500(c)
		return false
	}
	// ---- 设置标头信息-----
	// X-RateLimit-Limit :10000 最大访问次数
	// X-RateLimit-Remaining :9993 剩余的访问次数
	// X-RateLimit-Reset :1513784506 到该时间点，访问次数会重置为 X-RateLimit-Limit
	c.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
	c.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining))
	c.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))

	if rate.Reached {
		// 提示用户超额了
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "接口请求太频繁",
		})
		return false
	}
	return true
}
