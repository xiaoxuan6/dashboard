package api

import (
    "dashboard/pkg/cache"
    cache2 "github.com/patrickmn/go-cache"
    "net/http"
    "strings"
)

func init() {
    cache.Init()
}

func Cache(w http.ResponseWriter, r *http.Request) {
    action := r.URL.Query().Get("action")
    key := r.URL.Query().Get("key")

    if strings.Compare(action, "set") == 0 {
        val := r.URL.Query().Get("value")
        cache.Cache.Set(key, val, cache2.DefaultExpiration)
        _, _ = w.Write([]byte("添加成功"))
    }

    val, found := cache.Cache.Get(key)
    if !found {
        _, _ = w.Write([]byte("暂无参数"))
    }
    _, _ = w.Write([]byte(val.(string)))
}
