package lib

import (
    "dashboard/pkg/jwts"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type TokenHandler struct {
}

func (t TokenHandler) Run() *Response {
    return nil
}

func (t TokenHandler) Do(r *http.Request) *Response {
    b, _ := ioutil.ReadAll(r.Body)

    var request LoginResponse
    err := json.Unmarshal(b, &request)
    if err != nil {
        return FailWithMsg(fmt.Sprintf("json 解析请求惨错误：%s", err))
    }

    if request.Email == "" || request.Token == "" {
        return FailWithMsg("无效的 Token")
    }

    var c jwts.MyClaims
    c1, err := jwts.ParseWithClaims(request.Token, &c)
    if err != nil {
        return FailWithMsg(err.Error())
    }

    if c1.Email != request.Email {
        return FailWithMsg("无效的 Token")
    }

    return success()
}
