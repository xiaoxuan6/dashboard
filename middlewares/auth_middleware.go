package middlewares

import (
    "dashboard/lib"
    "net/http"
)

type AuthMiddleware struct {
}

var tokenHandler lib.TokenHandler

func (a AuthMiddleware) Handler(r *http.Request) *lib.Response {
    return tokenHandler.Do(r)
}
