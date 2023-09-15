package api

import (
    "dashboard/lib"
    middlewares2 "dashboard/middlewares"
    "net/http"
    "strings"
)

var middlewares map[string]middlewares2.Middleware

var allowMiddlewares map[string]string

var middleware middlewares2.Middleware

func allowMiddleware(method string) bool {
    var allow = false
    for _, val := range allowMiddlewares {
        if strings.Compare(val, method) == 0 {
            allow = true

            m, ok := middlewares[method]
            if !ok {
                return allow
            }

            middleware = m
        }
    }

    return allow
}

func Handler(r *http.Request) *lib.Response {
    return middleware.Handler(r)
}
