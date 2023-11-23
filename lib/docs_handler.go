package lib

import (
    "dashboard/common"
    "net/http"
)

type DocsHandler struct{}

func (dh DocsHandler) Run() *Response {
    return SuccessWithData(common.Docs)
}

func (dh DocsHandler) Do(r *http.Request) *Response {
    return nil
}
