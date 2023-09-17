package lib

import (
    "dashboard/pkg/github"
    "fmt"
    "github.com/tidwall/gjson"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    "sync"
)

var (
    wg    sync.WaitGroup
    lock  sync.Mutex
    items []github.Item
)

type SearchHandler struct {
}

func (s SearchHandler) Run() *Response {
    return nil
}

type Request struct {
    Keyword string `json:"keyword"`
}

func (s SearchHandler) Do(r *http.Request) *Response {
    b, _ := ioutil.ReadAll(r.Body)

    // 解码URL编码的字符串
    decodedString, err := url.QueryUnescape(string(b))
    if err != nil {
        return FailWithMsg(fmt.Sprintf("解码URL编码字符串时发生错误: %s", err.Error()))
    }

    fn := func(keyword string, item []github.Item) {
        for _, val := range items {
            wg.Add(1)
            go checkKeyword(keyword, val)
        }
    }

    // 从缓存中重新获取数据赋值
    //Load()

    keyword := gjson.Get(decodedString, "keyword").String()
    keywords := strings.Split(keyword, " ")
    if len(keywords) == 2 {
        tag := keywords[1]
        for _, val := range github.Menus {
            if strings.Compare(val, tag) == 0 {
                item, _ := github.Data[tag]

                fn(keywords[0], item)
                wg.Wait()
                return SuccessWithData(items)
            }
        }
    }

    var dates []github.Item
    for _, val := range github.Data {
        dates = append(dates, val...)
    }

    fn(keyword, dates)
    wg.Wait()
    return SuccessWithData(items)
}

func checkKeyword(keyword string, data github.Item) {
    defer wg.Done()

    if strings.ToLower(keyword) == "all" { // 查询所有数据
        lock.Lock()
        items = append(items, data)
        lock.Unlock()
    } else {
        if strings.Contains(data.Title, keyword) { // 查询标题中包含关键字的数据
            lock.Lock()
            items = append(items, data)
            lock.Unlock()
        }
    }
}
