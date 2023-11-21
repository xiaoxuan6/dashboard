package lib

import (
    "dashboard/common"
    "dashboard/util"
    xj "github.com/basgys/goxml2json"
    "github.com/mmcdole/gofeed"
    "github.com/tidwall/gjson"
    "net/http"
    "net/url"
    "strings"
)

var (
    Items []Item
)

type Item struct {
    Title  string `json:"title"`
    Url    string `json:"url"`
    Target string `json:"target"`
}

type RssHandler struct {
}

func (rh RssHandler) Run() *Response {
    for _, val := range common.Communes {
        wg.Add(1)
        go func(url string) {
            defer wg.Done()
            rssDo(url)
        }(val)
    }

    wg.Wait()
    return SuccessWithData(Items)
}

func rssDo(uri string) {
    fn := func(title, link, target string) {
        lock.Lock()
        Items = append(Items, Item{
            Title:  title,
            Url:    link,
            Target: target,
        })
        lock.Unlock()
    }

    u, _ := url.Parse(uri)
    fp := gofeed.NewParser()
    feed, err := fp.ParseURL(uri)
    if err == nil {
        for _, item := range feed.Items {
            fn(item.Title, item.Link, u.Host)
        }
        return
    }

    res, err := util.Get(uri)
    if err != nil {
        return
    }

    json, err := xj.Convert(strings.NewReader(res))
    if err != nil {
        return
    }

    gjson.ParseBytes(json.Bytes()).Get("entry").ForEach(func(key, value gjson.Result) bool {
        fn(value.Get("title").String(), value.Get("link").String(), u.Host)
        return true
    })
}

func (rh RssHandler) Do(r *http.Request) *Response {
    return nil
}
