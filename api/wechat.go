package api

import (
    "encoding/xml"
    "github.com/sirupsen/logrus"
    "io/ioutil"
    "net/http"
    "strings"
    "time"
)

type Response struct {
    ToUserName   string `json:"ToUserName"`
    FromUserName string `json:"FromUserName"`
    CreateTime   int    `json:"CreateTime"`
    MsgType      string `json:"MsgType"`
    Event        string `json:"Event"`
}

type TextRequest struct {
    ToUserName   string `json:"ToUserName"`
    FromUserName string `json:"FromUserName"`
    CreateTime   int64  `json:"CreateTime"`
    MsgType      string `json:"MsgType"`
    Content      string `json:"Content"`
}

func Wechat(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        _, _ = w.Write([]byte(r.URL.Query().Get("echostr")))
    } else {
        b, _ := ioutil.ReadAll(r.Body)
        var response Response
        _ = xml.Unmarshal(b, &response)
        logrus.Info("request", response)

        fn := func(msg string) {
            request := &TextRequest{
                ToUserName:   response.FromUserName,
                FromUserName: response.ToUserName,
                CreateTime:   time.Now().Unix(),
                MsgType:      "text",
                Content:      msg,
            }
            b1, _ := xml.Marshal(request)
            body := strings.ReplaceAll(string(b1), "TextRequest", "xml")
            logrus.Info("msg", body)
            _, _ = w.Write([]byte(body))
        }

        if response.MsgType == "event" && response.Event == "subscribe" {
            fn("欢迎关注")
            return
        }

        if response.MsgType == "text" {
            fn("文本消息")
            return
        }

        _, _ = w.Write([]byte("success"))
    }
}
