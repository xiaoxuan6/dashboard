package lib

import (
    "context"
    "dashboard/pkg/bleve"
    "dashboard/pkg/github"
    _package "dashboard/pkg/package"
    "fmt"
    github2 "github.com/google/go-github/v48/github"
    "github.com/sirupsen/logrus"
    "github.com/tidwall/gjson"
    "io/ioutil"
    "net/http"
    "net/url"
    "os"
    "regexp"
    "strings"
    "sync"
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
        Keyword   string          `json:"keyword"`
        Posts     []_package.Post `json:"posts"`
        Tags      map[string]int  `json:"tags"`
        Total     int64           `json:"total"`
        PageCount int             `json:"page_count"`
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

    if len(_package.Posts) < 1 {
        Load()
        if err = _package.Load(); err != nil {
            return Fail(err)
        }
    }

    if err = bleve.Init(); err != nil {
        return Fail(err)
    }

    posts, err := bleve.Search(keyword)
    if err != nil {
        return Fail(err)
    }

    var tags = make(map[string]int, 0)
    for _, post := range posts {
        if _, ok := tags[post.Tag]; !ok {
            tags[post.Tag] = 1
        } else {
            tags[post.Tag] = tags[post.Tag] + 1
        }
    }

    response.Posts = posts
    response.Tags = tags
    response.Total = int64(bleve.Total)
    response.PageCount = len(posts)

    return SuccessWithData(response)

}

var (
    ctx      = context.Background()
    filename = "result.txt"
    lock     sync.Mutex
    wg       sync.WaitGroup
)

func Load() {
    defer func() {
        if r := recover(); r != nil {
            logrus.Error("github load data Recovered in f", r)
            return
        }
    }()

    github.Init()
    repositoryContent, _, _, err := github.Client.Repositories.GetContents(ctx, os.Getenv("GITHUB_OWNER"), os.Getenv("GITHUB_REPO"), filename, &github2.RepositoryContentGetOptions{})
    if err != nil {
        panic(err)
    }

    content, err2 := repositoryContent.GetContent()
    if err2 != nil {
        panic(err2)
    }

    contents := strings.Split(content, "\n")
    for _, val := range contents {
        wg.Add(1)
        go func(wg *sync.WaitGroup, val string) {
            defer wg.Done()
            uri := github.RegexpUrl(val)
            title := github.RegexpTitle(val)
            tag := RegexpTag(val)
            if uri != "" && title != "" && tag != "" {
                lock.Lock()
                _package.Posts = append(_package.Posts, _package.Post{
                    Title: title,
                    Url:   uri,
                    Tag:   tag,
                })
                lock.Unlock()
            }
        }(&wg, val)
    }

    wg.Wait()
    return
}

func RegexpTag(str string) string {
    re := regexp.MustCompile(`\)\|(.*?)\|`)
    matches := re.FindStringSubmatch(str)
    if len(matches) > 1 {
        return matches[1]
    }

    return ""
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
