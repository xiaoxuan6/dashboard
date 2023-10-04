package lib

import (
    "dashboard/pkg/bleve"
    _package "dashboard/pkg/package"
    "fmt"
    "github.com/tidwall/gjson"
    "io/ioutil"
    "net/http"
    "net/url"
    "strconv"
    "sync"
)

//var (
//    wg    sync.WaitGroup
//    lock  sync.Mutex
//    items []github.Item
//)

var (
	wg sync.WaitGroup
	lock sync.Mutex
)

type SearchHandler struct {
}

func (s SearchHandler) Run() *Response {
    return nil
}

type (
    Request struct {
        Keyword string `json:"keyword"`
    }

    SearchResponse struct {
        Keyword string          `json:"keyword"`
        Posts   []_package.Post `json:"posts"`
        Tags    map[string]int  `json:"tags"`
    }
)

func (s SearchHandler) Do(r *http.Request) *Response {
    b, _ := ioutil.ReadAll(r.Body)

    // 解码URL编码的字符串
    decodedString, err := url.QueryUnescape(string(b))
    if err != nil {
        return FailWithMsg(fmt.Sprintf("解码URL编码字符串时发生错误: %s", err.Error()))
    }
    keyword := gjson.Get(decodedString, "keyword").String()

    var response SearchResponse
    response.Keyword = keyword

    if err = _package.Load(); err != nil {
        return Fail(err)
    }

    if err = bleve.Init(); err != nil {
        return Fail(err)
    }

    ids, err := bleve.Search(keyword)
    if err != nil {
        return Fail(err)
    }

    var tags = make(map[string]int, 0)
    var posts []_package.Post
    for _, val := range ids {
        wg.Add(1)

        go func(wg *sync.WaitGroup, val string) {
            defer wg.Done()
            i, _ := strconv.Atoi(val)
            post := _package.Posts[i]

            lock.Lock()
            posts = append(posts, post)
            lock.Unlock()

            if _, ok := tags[post.Tag]; !ok {
                tags[post.Tag] = 1
            } else {
                tags[post.Tag] = tags[post.Tag] + 1
            }
        }(&wg, val)
    }

    wg.Wait()
    response.Posts = posts
    response.Tags = tags

    return SuccessWithData(response)

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

//var Weeks = map[int]string{1: "星期一", 2: "星期二", 3: "星期三", 4: "星期四", 5: "星期五", 6: "星期六", 7: "星期天"}

//type dayInfoResponse struct {
//    Code int `json:"code"`
//    Type struct {
//        Week int `json:"week"`
//    } `json:"type"`
//}

//func fetchWeek() string {
//    r, _ := http.NewRequest(http.MethodGet, fmt.Sprintf(common.HOLIDAY_INFO, time.Now().Format("2006-01-02")), nil)
//    defer func() {
//        if r.Body != nil && r != nil {
//            _ = r.Body.Close()
//        }
//    }()
//
//    r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.66 Safari/537.36 Edg/103.0.1264.44")
//    response, err := http.DefaultClient.Do(r)
//    if err != nil {
//        return fmt.Sprintf("请求错误：%s", err.Error())
//    }
//    defer func() {
//        if response != nil && response.Body != nil {
//            _ = response.Body.Close()
//        }
//    }()
//
//    var res dayInfoResponse
//    err = json.NewDecoder(response.Body).Decode(&res)
//    if err != nil {
//        return fmt.Sprintf("json 解析错误：%s", err.Error())
//    }
//
//    if res.Code == 0 {
//        if val, ok := Weeks[res.Type.Week]; ok {
//            return val
//        }
//
//        return strconv.Itoa(res.Type.Week)
//    }
//
//    return "未知"
//}

//type nextDayResponse struct {
//    Code int    `json:"code"`
//    Tts  string `json:"tts"`
//}

//func fetchNextHoliday() string {
//    r, _ := http.NewRequest(http.MethodGet, common.HOLIDAY, nil)
//    defer func() {
//        if r.Body != nil && r != nil {
//            _ = r.Body.Close()
//        }
//    }()
//
//    r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.66 Safari/537.36 Edg/103.0.1264.44")
//    response, err := http.DefaultClient.Do(r)
//    if err != nil {
//        return fmt.Sprintf("请求错误：%s", err.Error())
//    }
//    defer func() {
//        if response != nil && response.Body != nil {
//            _ = response.Body.Close()
//        }
//    }()
//
//    var res nextDayResponse
//    err = json.NewDecoder(response.Body).Decode(&res)
//    if err != nil {
//        return fmt.Sprintf("json 解析错误：%s", err.Error())
//    }
//
//    if res.Code == 0 {
//        re := regexp.MustCompile(`[0-9]{2}的(.*)，`)
//        holiday := re.FindStringSubmatch(res.Tts)
//
//        matchDays := regexp.MustCompile(`\d+天`)
//        days := matchDays.FindString(res.Tts)
//        return fmt.Sprintf("距离%s还有%s", holiday[1], days)
//    }
//
//    return ""
//}

//func checkKeyword(keyword string, data github.Item) {
//    defer wg.Done()
//
//    if strings.ToLower(keyword) == "all" { // 查询所有数据
//        lock.Lock()
//        items = append(items, data)
//        lock.Unlock()
//    } else {
//        if strings.Contains(data.Title, keyword) { // 查询标题中包含关键字的数据
//            lock.Lock()
//            items = append(items, data)
//            lock.Unlock()
//        }
//    }
//}
