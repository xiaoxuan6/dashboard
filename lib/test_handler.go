package lib

type TestHandler struct {
}

func (t TestHandler) Run() *Response {
    return success()
}
