package api

import (
    "dashboard/common"
    "dashboard/lib"
    middlewares2 "dashboard/middlewares"
    "encoding/json"
    "net/http"
    "strings"
)

var (
    handlers map[string]lib.Handler
    // 允许执行 Do (需要接受参数)
    allowMethods = []string{common.LOGIN, common.CHECK, common.SEARCH}
    // 路由允许执行的中间件
    allowMiddlewares map[string][]string
    middleware       []middlewares2.Middleware
    // 所有中间件
    middlewares map[string]middlewares2.Middleware
)

func init() {
    handlers = map[string]lib.Handler{
        // 不需要接受参数，不需要执行 Do 方法
        "index": lib.IndexHandler{},
        "rss":   lib.RssHandler{},
        "docs":  lib.DocsHandler{},

        // 需要接受参数，需要执行 Do 方法
        common.LOGIN: lib.LoginHandler{},
        common.CHECK: lib.TokenHandler{},

        // 需要执行中间件
        common.SEARCH: lib.SearchHandler{},
    }

    middlewares = map[string]middlewares2.Middleware{
        common.AUTH: middlewares2.AuthMiddleware{},
    }

    allowMiddlewares = map[string][]string{
        common.SEARCH: []string{common.AUTH},
    }
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
        body = lib.FailWithMsg("无效的请求")
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
