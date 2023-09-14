package api

import (
    "dashboard/lib"
    "encoding/json"
    "net/http"
)

var handlers map[string]lib.Handler

func init() {
    handlers = map[string]lib.Handler{
        "test":  lib.TestHandler{},
        "index": lib.IndexHandler{},
        "login": lib.LoginHandler{},
    }
}

func Run(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=utf-8")

    action := r.URL.Query().Get("action")

    var body *lib.Response
    if handler, ok := handlers[action]; ok {
        body = handler.Run()
    } else {
        handler, _ = handlers["test"]
        body = handler.Run()
    }

    res, _ := json.Marshal(body)

    _, _ = w.Write(res)
}
