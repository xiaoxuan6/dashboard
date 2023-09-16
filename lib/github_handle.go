package lib

import (
    "dashboard/common"
    "dashboard/pkg/cache"
    "encoding/json"
    "fmt"
    cache2 "github.com/patrickmn/go-cache"
    "io/ioutil"
    "net/http"
)

type GithubHandler struct {
}

func (g GithubHandler) Run() *Response {
    return nil
}

type TokenRequest struct {
    Token string `json:"token"`
}

func (g GithubHandler) Do(r *http.Request) *Response {
    b, _ := ioutil.ReadAll(r.Body)

    var request TokenRequest
    err := json.Unmarshal(b, &request)
    if err != nil {
        return FailWithMsg(fmt.Sprintf("json 解析请求惨错误：%s", err))
    }

    cache.Cache.Set(common.Token, request.Token, cache2.DefaultExpiration)

    LoadData()

    return Success()
}
