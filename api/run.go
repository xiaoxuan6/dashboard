package api

import (
    "dashboard/lib"
    "encoding/json"
    "net/http"
    "strings"
)

var handlers map[string]lib.Handler

var allowMethods = []string{"login", "check_token"}

func init() {
    handlers = map[string]lib.Handler{
        "test":        lib.TestHandler{},
        "index":       lib.IndexHandler{},
        "login":       lib.LoginHandler{},
        "check_token": lib.TokenHandler{},
    }
}

func Run(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=utf-8")

    action := r.URL.Query().Get("action")

    var body *lib.Response
    if handler, ok := handlers[action]; ok {
        if allowMethod(action) {
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
