package cache

import (
    cache2 "github.com/patrickmn/go-cache"
    "os"
    "strconv"
    "sync"
    "time"
)

var (
    Cache *cache2.Cache
    one   sync.Once
)

func Init() {
    one.Do(func() {
        expirationAt, _ := strconv.Atoi(os.Getenv("CACHE_EXPIRATION_AT"))
        Cache = cache2.New(time.Hour*time.Duration(expirationAt), time.Minute*time.Duration(expirationAt))
    })
}
