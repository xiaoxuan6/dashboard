package api

import (
    "dashboard/common"
    "dashboard/lib"
    middlewares2 "dashboard/middlewares"
    "dashboard/pkg/cache"
    "dashboard/pkg/github"
    "encoding/json"
    cache2 "github.com/patrickmn/go-cache"
    "net/http"
    "strings"
)

var (
    handlers map[string]lib.Handler
    // 允许执行 Do
    allowMethods = []string{common.LOGIN, common.CHECK, common.SEARCH, common.GITHUB}
    // 路由允许执行的中间件
    allowMiddlewares map[string][]string
    middleware       []middlewares2.Middleware
    // 所有中间件
    middlewares map[string]middlewares2.Middleware
)

func init() {
    handlers = map[string]lib.Handler{
        "test":        lib.TestHandler{},
        "index":       lib.IndexHandler{},
        common.LOGIN:  lib.LoginHandler{},
        common.CHECK:  lib.TokenHandler{},
        common.SEARCH: lib.SearchHandler{},
        common.GITHUB: lib.GithubHandler{},
    }

    middlewares = map[string]middlewares2.Middleware{
        common.AUTH:  middlewares2.AuthMiddleware{},
        common.CHECK: middlewares2.CheckTokenMiddleware{},
    }

    allowMiddlewares = map[string][]string{
        common.SEARCH: []string{common.AUTH, common.CHECK},
        common.GITHUB: []string{common.AUTH},
    }

    cache.Init()
}

func Run(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=utf-8")

    action := r.URL.Query().Get("action")

    var body *lib.Response
    if handler, ok := handlers[action]; ok {
        if allowMethod(action) {

            if allowMiddleware(action) {
                var response *lib.Response
                response = Handler(r)
                if response.Status != http.StatusOK {
                    res, _ := json.Marshal(response)

                    _, _ = w.Write(res)
                    return
                }
            }

            body = handler.Do(r)
        } else {
            body = handler.Run()
        }
    } else {
        handler, _ = handlers["test"]
        body = handler.Run()
    }

    // 测试代码
    if strings.Compare(action, "cache") == 0 {
        key := r.URL.Query().Get("value")
        if strings.Compare(action, "set") == 0 {
            val := r.URL.Query().Get("value")
            cache.Cache.Set(key, val, cache2.DefaultExpiration)
            _, _ = w.Write([]byte("添加成功"))
            return
        }

        val, found := cache.Cache.Get(key)
        if !found {
            _, _ = w.Write([]byte("暂无参数"))
        }

        var cacheVal interface{}
        switch valT := val.(type) {
        case string:
            cacheVal = valT
        case []string:
            cacheVal = valT
        case map[string][]github.Item:
            cacheVal = valT
        }

        b, _ := json.Marshal(cacheVal)
        _, _ = w.Write(b)
        return
    }

    middleware = nil
    res, _ := json.Marshal(body)

    _, _ = w.Write(res)
}

func allowMethod(method string) bool {
    var allow = false
    for _, val := range allowMethods {
        if strings.Compare(val, method) == 0 {
            allow = true
        }
    }

    return allow
}

func allowMiddleware(method string) bool {
    var allow = false
    for key, val := range allowMiddlewares {
        if strings.Compare(key, method) == 0 {
            allow = true

            for _, m := range val {
                mC, ok := middlewares[m]
                if !ok {
                    return allow
                }

                middleware = append(middleware, mC)
            }

        }
    }

    return allow
}

func Handler(r *http.Request) *lib.Response {
    var response *lib.Response
    for _, val := range middleware {
        if response = val.Handler(r); response.Status != http.StatusOK {
            break
        }
    }
    return response
}
