package lib

import (
    "dashboard/common"
    "dashboard/pkg/cache"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strconv"
    "time"
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

    if len(request.Token) < 1 {
        return FailWithMsg("github token 不能为空！")
    }

    expirationAt := os.Getenv("CACHE_TOKEN_EXPIRATION_AT")
    i, _ := strconv.Atoi(expirationAt)
    d := time.Duration(i) * 7
    cache.Cache.Set(common.Token, request.Token, d)

    LoadData()

    return Success()
}
