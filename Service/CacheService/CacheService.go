package CacheService

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var goCache = cache.New(30*time.Minute, 30*time.Minute)

func SetCache(key string, value interface{}) {
	goCache.Set(key, value, cache.DefaultExpiration)
}

func GetCache(key string) interface{} {
	value, bo := goCache.Get(key)
	if !bo {
		return nil
	}
	return value
}
