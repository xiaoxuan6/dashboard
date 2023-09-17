package api

import (
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
    handlers         map[string]lib.Handler
    allowMethods     = []string{"login_do", "check_token", "search_do", "github_token"}
    allowMiddlewares map[string][]string
    middleware       []middlewares2.Middleware
    middlewares      map[string]middlewares2.Middleware
)

func init() {
    handlers = map[string]lib.Handler{
        "test":         lib.TestHandler{},
        "index":        lib.IndexHandler{},
        "login_do":     lib.LoginHandler{},
        "check_token":  lib.TokenHandler{},
        "search_do":    lib.SearchHandler{},
        "github_token": lib.GithubHandler{},
    }

    middlewares = map[string]middlewares2.Middleware{
        "auth":        middlewares2.AuthMiddleware{},
        "check_token": middlewares2.CheckTokenMiddleware{},
    }

    allowMiddlewares = map[string][]string{
        "search_do":    []string{"auth", "check_token"},
        "github_token": []string{"auth"},
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
