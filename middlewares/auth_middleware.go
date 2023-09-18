package middlewares

import (
    "dashboard/lib"
    "dashboard/pkg/jwts"
    "net/http"
    "time"
)

type AuthMiddleware struct {
}

func (a AuthMiddleware) Handler(r *http.Request) *lib.Response {
    token := r.Header.Get("Authorization")
    if len(token) < 1 {
        return lib.FailAuthWithMsg("无效的 Token")
    }

    var c jwts.MyClaims
    c1, err := jwts.ParseWithClaims(token, &c)
    if err != nil {
        return lib.FailAuth(err)
    }

    expirationTime := time.Unix(c.ExpiresAt, 0)
    if time.Now().After(expirationTime) {
        return lib.FailAuthWithMsg("token 已过期")
    }

    if c1.Email != r.Header.Get("email") {
        return lib.FailAuthWithMsg("无效的 Token")
    }

    return lib.SuccessWithData(nil)
}
