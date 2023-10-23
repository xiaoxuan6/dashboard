package handlers

import (
    "bytes"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
)

var CollectHandler = new(collectHandler)

type collectHandler struct {
}

func (c collectHandler) Put(ctx *gin.Context) {
    auth := ctx.PostForm("auth")
    url := ctx.PostForm("url")

    if auth == "" || url == "" {
        ctx.JSON(200, gin.H{
            "status": 400,
            "data":   "",
            "msg":    "auth or url not empty",
        })
        return
    }

    if auth != os.Getenv("GITHUB_OWNER") {
        ctx.JSON(200, gin.H{
            "status": 400,
            "data":   "",
            "msg":    "auth error",
        })
        return
    }

    var body bytes.Buffer
    body.WriteString(fmt.Sprintf(`{"event_type": "push", "client_payload": {"url": "%s", "description":"", "demo_url":""}}`, url))

    r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.github.com/repos/%s/go-package-example/dispatches", auth), &body)
    r.Header.Set("Accept", "application/vnd.github+json")
    r.Header.Set("Authorization", fmt.Sprintf("token %s", os.Getenv("GITHUB_TOKEN")))
    r.Header.Set("X-GitHub-Api-Version", "2022-11-28")

    res, _ := http.DefaultClient.Do(r)
    if res.StatusCode != 204 {
        ctx.JSON(200, gin.H{
            "status": 400,
            "data":   "",
            "msg":    "dispatch error",
        })
        return
    }

    ctx.JSON(200, gin.H{
        "status": 200,
        "data":   "",
        "msg":    "ok",
    })
}
