package lib

import (
    "errors"
    "net/http"
)

type Response struct {
    Status int         `json:"status"`
    Data   interface{} `json:"data"`
    Msg    string      `json:"msg"`
}

func success() *Response {
    return successWithData(nil)
}

func successWithData(data interface{}) *Response {
    res := &Response{
        http.StatusOK,
        data,
        "ok",
    }

    return res
}

func fail(err error) *Response {
    res := &Response{
        http.StatusBadRequest,
        nil,
        err.Error(),
    }

    return res
}

func failWithMsg(msg string) *Response {
    err := errors.New(msg)
    return fail(err)
}
