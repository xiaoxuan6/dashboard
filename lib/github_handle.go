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

    _ = r.ParseForm()
    token := r.FormValue("token")

    b, _ := ioutil.ReadAll(r.Body)

    var request TokenRequest
    var token3 string
    err := json.Unmarshal(b, &request)
    if err != nil {
        token3 = fmt.Sprintf("json 解析请求惨错误：%s", err)
    } else {
        token3 = request.Token
    }

    res := struct {
        Token  string `json:"token"`
        Token1 string `json:"token_1"`
        Token2 string `json:"token_2"`
        Token3 string `json:"token_3"`
    }{
        Token:  token,
        Token1: r.URL.Query().Get("token"),
        Token2: string(b),
        Token3: token3,
    }

    return SuccessWithData(&res)

    //expirationAt := os.Getenv("CACHE_TOKEN_EXPIRATION_AT")
    //i, _ := strconv.Atoi(expirationAt)
    //d := time.Duration(i) * 7
    //cache.Cache.Set(common.Token, token, d)

    //return SuccessWithData(token)

    LoadData()

    return Success()
}
