package middlewares

import (
    "dashboard/common"
    "dashboard/lib"
    "dashboard/pkg/cache"
    "errors"
    "fmt"
    "net/http"
)

type CheckTokenMiddleware struct {
}

func (c CheckTokenMiddleware) Handler(r *http.Request) *lib.Response {
    val, ok := cache.Cache.Get(common.Token)
    if !ok {
        err := errors.New(fmt.Sprintf("获取缓存 %s 失败！", common.Token))
        return lib.FailAuth(err)
    }

    if len(val.(string)) < 1 {
        return lib.FailAuth(errors.New("无效的 github token"))
    }

    return lib.Success()
}
