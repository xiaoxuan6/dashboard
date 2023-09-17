package lib

import (
    "encoding/json"
    "fmt"
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
    res := struct {
        Token  string `json:"token"`
        Token1 string `json:"token_1"`
        Token2 string `json:"token_2"`
        Token3 string `json:"token_3"`
    }{}

    _ = r.ParseForm()
    res.Token = r.FormValue("token")

    res.Token1 = r.URL.Query().Get("token")

    b, _ := ioutil.ReadAll(r.Body)
    res.Token2 = string(b)

    var request TokenRequest
    err := json.Unmarshal(b, &request)
    if err != nil {
        res.Token3 = fmt.Sprintf("json 解析请求惨错误：%s", err)
    } else {
        res.Token3 = request.Token
    }

    return SuccessWithData(res)

    //expirationAt := os.Getenv("CACHE_TOKEN_EXPIRATION_AT")
    //i, _ := strconv.Atoi(expirationAt)
    //d := time.Duration(i) * 7
    //cache.Cache.Set(common.Token, token, d)

    //return SuccessWithData(token)

    LoadData()

    return Success()
}
