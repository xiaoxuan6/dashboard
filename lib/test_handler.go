package lib

import "net/http"

type TestHandler struct {
}

func (t TestHandler) Run() *Response {
    return successWithData("hello")
}

func (t TestHandler) Do(r *http.Request) *Response {
    return nil
}
