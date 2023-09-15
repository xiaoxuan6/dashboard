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
    return SuccessWithData(nil)
}

func SuccessWithData(data interface{}) *Response {
    res := &Response{
        Status: http.StatusOK,
        Data:   data,
        Msg:    "ok",
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

func FailWithMsg(msg string) *Response {
    err := errors.New(msg)
    return fail(err)
}
