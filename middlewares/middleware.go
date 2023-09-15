package middlewares

import (
    "dashboard/lib"
    "net/http"
)

type Middleware interface {
    Handler(r *http.Request) *lib.Response
}
