package lib

import (
    "dashboard/common"
    "dashboard/pkg/cache"
    "dashboard/pkg/github"
    cache2 "github.com/patrickmn/go-cache"
)

func Save(menu []string, data map[string][]github.Item) {
    cache.Cache.Set(common.Menu, menu, cache2.DefaultExpiration)
    cache.Cache.Set(common.Data, data, cache2.DefaultExpiration)
}
