package lib

import (
    "dashboard/common"
    "dashboard/pkg/github"
    "encoding/json"
    "fmt"
    "net/http"
    "regexp"
    "strings"
    "sync"
    "time"
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

type SearchResponse struct {
    Keyword string `json:"keyword"`
    Date    string `json:"date"`
}

func (s SearchHandler) Do(r *http.Request) *Response {
    keyword := r.PostFormValue("keyword")

    var response SearchResponse
    response.Keyword = keyword
    response.Date = fmt.Sprintf("今日 %s %s", time.Now().Format("2006-01-02"), fetchNextDay())

    return SuccessWithData(response)
    //b, _ := ioutil.ReadAll(r.Body)
    //
    //// 解码URL编码的字符串
    //decodedString, err := url.QueryUnescape(string(b))
    //if err != nil {
    //    return FailWithMsg(fmt.Sprintf("解码URL编码字符串时发生错误: %s", err.Error()))
    //}
    //
    //return SuccessWithData(decodedString)

    //fn := func(keyword string, item []github.Item) {
    //    for _, val := range items {
    //        wg.Add(1)
    //        go checkKeyword(keyword, val)
    //    }
    //}
    //
    //// 从缓存中重新获取数据赋值
    ////Load()
    //
    //keyword := gjson.Get(decodedString, "keyword").String()
    //keywords := strings.Split(keyword, " ")
    //if len(keywords) == 2 {
    //    tag := keywords[1]
    //    for _, val := range github.Menus {
    //        if strings.Compare(val, tag) == 0 {
    //            item, _ := github.Data[tag]
    //
    //            fn(keywords[0], item)
    //            wg.Wait()
    //            return SuccessWithData(items)
    //        }
    //    }
    //}
    //
    //var dates []github.Item
    //for _, val := range github.Data {
    //    dates = append(dates, val...)
    //}
    //
    //fn(keyword, dates)
    //wg.Wait()
    //return SuccessWithData(items)
}

type nextDayResponse struct {
    Code int    `json:"code"`
    Tts  string `json:"tts"`
}

func fetchNextDay() string {
    r, _ := http.NewRequest(http.MethodGet, common.HOLIDAY, nil)
    defer func() {
        if r.Body != nil && r != nil {
            _ = r.Body.Close()
        }
    }()

    r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.66 Safari/537.36 Edg/103.0.1264.44")
    response, err := http.DefaultClient.Do(r)
    if err != nil {
        return fmt.Sprintf("请求错误：%s", err.Error())
    }
    defer func() {
        if response != nil && response.Body != nil {
            _ = response.Body.Close()
        }
    }()

    var res nextDayResponse
    err = json.NewDecoder(response.Body).Decode(&res)
    if err != nil {
        return fmt.Sprintf("json 解析错误：%s", err.Error())
    }

    if res.Code == 0 {
        re := regexp.MustCompile(`[0-9]{2}的(.*)，`)
        holiday := re.FindStringSubmatch(res.Tts)

        matchDays := regexp.MustCompile(`\d+天`)
        days := matchDays.FindString(res.Tts)
        return fmt.Sprintf("距离%s还有%s", holiday[1], days)
    }

    return ""
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
