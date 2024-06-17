package captcha

import (
	"errors"
	"fmt"
	"time"

	"github.com/VENI-VIDIVICI/gohub/pkg/app"
	"github.com/VENI-VIDIVICI/gohub/pkg/config"
	"github.com/VENI-VIDIVICI/gohub/pkg/redis"
)

type StoreRedisCliient struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func (src *StoreRedisCliient) Set(key string, value string) error {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_time"))
	}
	if ok := src.RedisClient.Set(fmt.Sprint("%v%v", src.KeyPrefix, key), value, ExpireTime); ok {
		return errors.New("无法存储图片验证码答案")
	}
	return nil
}

func (s *StoreRedisCliient) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

func (s *StoreRedisCliient) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
