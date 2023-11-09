package handlers

import (
    "bytes"
    "fmt"
    "github.com/OwO-Network/gdeeplx"
    "github.com/abadojack/whatlanggo"
    "github.com/antchfx/htmlquery"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    db "github.com/xiaoxuan6/go-package-db"
    "io/ioutil"
    "net/http"
    "net/url"
    "os"
    "strings"
)

var (
    CollectHandler = new(collectHandler)
)

type collectHandler struct {
}

func (c collectHandler) Put(ctx *gin.Context) {
    b, _ := ioutil.ReadAll(ctx.Request.Body)
    decodedString, err := url.QueryUnescape(string(b))
    if err != nil {
        ctx.JSON(http.StatusOK, gin.H{
            "status": http.StatusBadRequest,
            "data":   "",
            "msg":    err.Error(),
        })
        return
    }

    logrus.Info("decodedString", decodedString)
    values, _ := url.ParseQuery(decodedString)
    auth := values.Get("auth")
    uri := values.Get("url")
    description := values.Get("description")
    language := values.Get("language")

    if auth == "" || uri == "" {
        ctx.JSON(http.StatusOK, gin.H{
            "status": 400,
            "data":   "",
            "msg":    "auth or url not empty",
        })
        return
    }

    if auth != os.Getenv("GITHUB_OWNER") {
        ctx.JSON(http.StatusOK, gin.H{
            "status": 400,
            "data":   "",
            "msg":    "auth error",
        })
        return
    }

    u, _ := url.Parse(uri)
    if u.Host != "github.com" {
        ctx.JSON(http.StatusOK, gin.H{
            "status": 400,
            "data":   "",
            "msg":    "url error",
        })
        return
    }

    db.Init(
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )
    defer db.Close()
    db.AutoMigrate()

    uri = strings.ReplaceAll(uri, "https://", "")
    if db.DB.Where("url = ?", uri).First(&db.Collect{}).RowsAffected > 0 {
        ctx.JSON(http.StatusOK, gin.H{
            "status": 400,
            "data":   "",
            "msg":    "url exists",
        })
        return
    }

    description = descriptionDo(strings.TrimSpace(description))
    language = languageDo(uri, language)

    var body bytes.Buffer
    body.WriteString(fmt.Sprintf(`{"event_type": "push", "client_payload": {"url": "%s", "description":"%s", "demo_url":"", "language": "%s"}}`, uri, description, language))

    r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.github.com/repos/%s/go-package-example/dispatches", auth), &body)
    r.Header.Set("Accept", "application/vnd.github+json")
    r.Header.Set("Authorization", fmt.Sprintf("token %s", os.Getenv("GITHUB_TOKEN")))
    r.Header.Set("X-GitHub-Api-Version", "2022-11-28")

    res, _ := http.DefaultClient.Do(r)
    if res.StatusCode != 204 {
        logrus.Error("dispatch error", body.String())
        ctx.JSON(http.StatusOK, gin.H{
            "status": 400,
            "data":   "",
            "msg":    "dispatch error",
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "data":   "",
        "msg":    "ok",
    })
}

func descriptionDo(description string) string {
    if len(description) < 1 {
        return ""
    }

    info := whatlanggo.Detect(description)
    lang := info.Lang.String()
    if lang != "" && lang != "Mandarin" {
        result, err := gdeeplx.Translate(description, "", "zh", 0)
        if err == nil {
            res := result.(map[string]interface{})
            description = strings.TrimSpace(res["data"].(string))
        }
    }

    return description
}

func languageDo(uri, language string) string {
    if len(language) > 0 {
        return language
    }

    doc, _ := htmlquery.LoadURL(uri)
    node := htmlquery.FindOne(doc, "//*[@id=\"repo-content-turbo-frame\"]/div/div/div[2]/div[2]/div/div[7]/div/ul/li[1]/a/span[1]")
    language = htmlquery.InnerText(node)

    return language
}
