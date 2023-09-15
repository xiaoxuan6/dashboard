package middlewares

import (
    "dashboard/lib"
    "dashboard/pkg/jwts"
    "net/http"
)

type AuthMiddleware struct {
}

func (a AuthMiddleware) Handler(r *http.Request) *lib.Response {
    token := r.Header.Get("Authorization")
    if len(token) < 1 {
        return lib.FailWithMsg("无效的 Token")
    }

    var c jwts.MyClaims
    c1, err := jwts.ParseWithClaims(token, &c)
    if err != nil {
        return lib.FailWithMsg(err.Error())
    }

    if c1.Email != r.Header.Get("email") {
        return lib.FailWithMsg("无效的 Token")
    }

    return lib.SuccessWithData(nil)
}
