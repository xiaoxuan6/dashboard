package lib

import (
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
    //_ = r.ParseForm()
    //token := r.FormValue("token")

    token := r.URL.Query().Get("token")
    return SuccessWithData(token)

    //expirationAt := os.Getenv("CACHE_TOKEN_EXPIRATION_AT")
    //i, _ := strconv.Atoi(expirationAt)
    //d := time.Duration(i) * 7
    //cache.Cache.Set(common.Token, token, d)

    return SuccessWithData(token)

    LoadData()

    return Success()
}
