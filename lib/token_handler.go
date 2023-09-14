package lib

import (
    "dashboard/pkg/jwts"
    "net/http"
)

type TokenHandler struct {
}

func (t TokenHandler) Run() *Response {
    return nil
}

func (t TokenHandler) Do(r *http.Request) *Response {
    email := r.PostFormValue("email")
    token := r.PostFormValue("token")

    if email == "" || token == "" {
        return failWithMsg("无效的 Token")
    }

    var c jwts.MyClaims
    c1, err := jwts.ParseWithClaims(token, &c)
    if err != nil {
        return failWithMsg(err.Error())
    }

    if c1.Email != email {
        return failWithMsg("无效的 Token")
    }

    return success()
}
