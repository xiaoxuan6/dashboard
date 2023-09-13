package lib

import (
    "dashboard/config"
    _ "embed"
    "encoding/json"
    "errors"
    "fmt"
)

//go:embed config.json
var settings []byte

type IndexHandler struct {
}

func (i IndexHandler) Run() *Response {
    var c config.Config
    err := json.Unmarshal(settings, &c)
    if err != nil {
        return fail(errors.New(fmt.Sprintf("json 解析错误：%s", err.Error())))
    }

    return successWithData(c)
}
