package lib

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
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

    // 解码URL编码的字符串
    decodedString, err := url.QueryUnescape(string(b))
    if err != nil {
        return FailWithMsg(fmt.Sprintf("解码URL编码字符串时发生错误: %s", err.Error()))
    }

    keyword := Keyword{
        Keyword: decodedString,
    }

    // todo: 后续逻辑

    return SuccessWithData(keyword)
}
