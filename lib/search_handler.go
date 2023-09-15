package lib

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Keyword struct {
    Keyword string `json:"keyword"`
}

type SearchHandler struct {
}

func (s SearchHandler) Run() *Response {
    return nil
}

func (s SearchHandler) Do(r *http.Request) *Response {
    b, _ := ioutil.ReadAll(r.Body)

    var key Keyword
    err := json.Unmarshal(b, &Keyword{})
    if err != nil {
        return FailWithMsg(fmt.Sprintf("json 解析错误：%s", err.Error()))
    }

    // todo: 后续逻辑

    return SuccessWithData(key.Keyword)
}
