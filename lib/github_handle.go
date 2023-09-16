package lib

import "net/http"

type GithubHandler struct {
}

func (g GithubHandler) Run() *Response {
    return nil
}

func (g GithubHandler) Do(r *http.Request) *Response {
    //todo:: 设置 token
    return nil
}
