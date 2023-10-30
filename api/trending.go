package api

import (
    "dashboard/pkg/email"
    "fmt"
    "github.com/OwO-Network/gdeeplx"
    "github.com/andygrunwald/go-trending"
    "net/http"
    "strings"
    "sync"
)

var (
    lock     sync.Mutex
    wg       sync.WaitGroup
    html     = make([]string, 0, 10)
    tagsHtml = map[string]string{"": "all", "php": "php", "go": "go"}
)

func Trending(w http.ResponseWriter, r *http.Request) {
    trend := trending.NewTrending()

    tags := []string{"", "php", "go"}
    for _, tag := range tags {
        wg.Add(1)
        go crawler(trend, tag, &wg)
    }
    wg.Wait()

    var htmls string
    for _, h := range html {
        htmls = fmt.Sprintf("%s%s", htmls, h)
    }
    _ = email.Send(email.Option{Html: []byte(htmls)})

    w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte("ok"))
}

func crawler(trend *trending.Trending, tag string, wg *sync.WaitGroup) {
    defer wg.Done()
    projects, err := trend.GetProjects(trending.TimeToday, tag)
    if err != nil {
        return
    }

    tagHtml := fmt.Sprintf("<h1>%s</h1>\n", tagsHtml[tag])
    for _, project := range projects {
        desc := project.Description
        result, err := gdeeplx.Translate(desc, "", "zh", 0)
        if err == nil {
            res := result.(map[string]interface{})
            desc = strings.TrimSpace(res["data"].(string))
        }

        if tag == "" {
            tagHtml = fmt.Sprintf("%s<a href=\"%s\">%s</a>(%s)：%s\n\n", tagHtml, project.URL, project.Name, project.Language, desc)
        } else {
            tagHtml = fmt.Sprintf("%s<a href=\"%s\">%s</a>：%s\n\n", tagHtml, project.URL, project.Name, desc)
        }
    }

    lock.Lock()
    html = append(html, tagHtml+"\n")
    lock.Unlock()
}
