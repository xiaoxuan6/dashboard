package handlers

import (
    _ "embed"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "net/http"
)

//go:embed config.json
var config []byte

var DocsHandler = newDocsHandler()

func newDocsHandler() *docsHandler {
    return new(docsHandler)
}

type docsHandler struct {
}

type Config struct {
    Apis []Api `json:"apis"`
}

type Api struct {
    Url   string `json:"url"`
    Title string `json:"title"`
    Desc  string `json:"desc"`
}

func (d docsHandler) Index(c *gin.Context) {
    var cs Config
    _ = json.Unmarshal(config, &cs)

    c.JSON(http.StatusOK, gin.H{
        "status": 200,
        "data":   cs,
        "msg":    "ok",
    })
}
