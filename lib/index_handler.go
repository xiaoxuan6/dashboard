package lib

import (
    config2 "dashboard/config"
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
)

type IndexHandler struct {
}

var filename = "./config.json"

func (i IndexHandler) Run() *Response {
    b, err := ioutil.ReadFile(filename)
    if err != nil {
        return fail(errors.New(fmt.Sprintf("读取文件 %s 失败: %s", filename, err.Error())))
    }

    var config *config2.Config
    err = json.Unmarshal(b, config)
    if err != nil {
        return fail(errors.New(fmt.Sprintf("json 解析错误：%s", err.Error())))
    }

    return successWithData(config)
}
