package lib

type Handler interface {
    Run() *Response
}
