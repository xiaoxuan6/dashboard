package api

import (
    "dashboard/lib"
    "net/http"
)

var handlers map[string]lib.Handler

func init() {
    handlers = map[string]lib.Handler{
        "test": lib.TestHandler{},
    }
}

func Run(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=utf-8")

    var body []byte
    if handler, ok := handlers["test"]; ok {
        body = handler.Run()
    }

    _, _ = w.Write(body)
}
