package tests

import (
    "dashboard/api"
    "encoding/xml"
    "github.com/stretchr/testify/assert"
    "strings"
    "testing"
    "time"
)

func TestWechat(t *testing.T) {
    re := api.TextRequest{
        ToUserName:   "xx",
        FromUserName: "xx",
        MsgType:      "text",
        CreateTime:   time.Now().Unix(),
        Content:      "test",
    }
    b, _ := xml.Marshal(re)
    body := strings.ReplaceAll(string(b), "TextRequest", "xml")
    t.Log(body)
    assert.Nil(t, nil)
}
