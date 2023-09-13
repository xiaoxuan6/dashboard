package lib

import "encoding/json"

type TestHandler struct {
}

func (t TestHandler) Run() []byte {
    response := map[string]interface{}{
        "status": 200,
        "msg":    "ok",
    }

    b, _ := json.Marshal(response)

    return b
}
