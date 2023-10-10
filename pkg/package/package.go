package _package

import (
    "dashboard/common"
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
    "strings"
    "sync"
)

var (
    wg    sync.WaitGroup
    Posts []Post
    lock  sync.Mutex
)

type Post struct {
    Title string `json:"title"`
    Url   string `json:"url"`
    Tag   string `json:"tag"`
}

func Load() error {
    b, err2 := fileGetContent()
    if err2 != nil {
        return err2
    }

    newContent := string(b)
    replacements := []string{"# Go 开源第三方包收集和使用示例", "|分支名|包名|描述|", "|:---|:---|:---|"}
    for _, replaceOld := range replacements {
        newContent = strings.ReplaceAll(newContent, replaceOld, ``)
    }
    newContent = strings.Trim(newContent, "\n")
    contents := strings.Split(newContent, "\n")

    for _, val := range contents {
        wg.Add(1)

        go func(wg *sync.WaitGroup, val string) {
            defer wg.Done()
            regexpStr := regexpContent(val)
            if regexpStr != nil {
                lock.Lock()
                Posts = append(Posts, Post{
                    Title: regexpStr[3],
                    Url:   fmt.Sprintf("https://%s", regexpStr[2]),
                    Tag:   common.GoTags,
                })
                lock.Unlock()
            }
        }(&wg, val)
    }

    wg.Wait()

    return nil
}

func fileGetContent() (b []byte, err error) {
    response, err := http.Get(common.PackageUrl)
    if err != nil {
        return b, errors.New("请求错误：" + err.Error())
    }

    defer response.Body.Close()

    b, err = ioutil.ReadAll(response.Body)
    if err != nil {
        return b, errors.New("获取内容失败：" + err.Error())
    }

    return b, nil
}

func regexpContent(val string) []string {
    re := regexp.MustCompile(`\|(.*?)\|(.*?)\|(.*?)\|`)
    matchers := re.FindStringSubmatch(val)
    if len(matchers) < 1 {
        return nil
    }
    return matchers
}
