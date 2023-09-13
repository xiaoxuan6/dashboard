package api

import (
    "dashboard/lib"
    "encoding/json"
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

    var body *lib.Response
    if handler, ok := handlers["test"]; ok {
        body = handler.Run()
    }

    res, _ := json.Marshal(body)

    _, _ = w.Write(res)
}
