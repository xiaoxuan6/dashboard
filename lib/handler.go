package lib

import "net/http"

type Handler interface {
    Run() *Response
    Do(*http.Request) *Response
}
