package api

import (
    "dashboard/lib"
    middlewares2 "dashboard/middlewares"
    "encoding/json"
    "net/http"
    "strings"
)

var handlers map[string]lib.Handler

var allowMethods = []string{"login_do", "check_token", "search_do"}

var middlewares map[string]middlewares2.Middleware

var allowMiddlewares map[string]string

var middleware middlewares2.Middleware

func init() {
    handlers = map[string]lib.Handler{
        "test":        lib.TestHandler{},
        "index":       lib.IndexHandler{},
        "login_do":    lib.LoginHandler{},
        "check_token": lib.TokenHandler{},
        "search_do":   lib.SearchHandler{},
    }

    middlewares = map[string]middlewares2.Middleware{
        "auth": middlewares2.AuthMiddleware{},
    }

    allowMiddlewares = map[string]string{
        "search_do": "auth",
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
        handler, _ = handlers["test"]
        body = handler.Run()
    }

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
