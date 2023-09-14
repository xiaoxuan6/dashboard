package lib

type LoginHandler struct {
}

func (l LoginHandler) Run() *Response {
    token := struct {
        Token string `json:"token"`
    }{
        "123456",
    }
    return successWithData(token)
}
