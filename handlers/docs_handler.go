package handlers

import (
    _ "embed"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/tidwall/gjson"
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

type Docs struct {
    Docs struct {
        Dirtyfilter Item `json:"dirtyfilter"`
    } `json:"docs"`
}

type Item struct {
    Url      string     `json:"url"`
    Method   string     `json:"method"`
    UrlDemo  string     `json:"url_demo"`
    Params   []params   `json:"params"`
    Response []response `json:"response"`
}

type params struct {
    Name    string `json:"name"`
    Require string `json:"require"`
    Type    string `json:"type"`
    Desc    string `json:"desc"`
}

type response struct {
    Name string `json:"name"`
    Type string `json:"type"`
    Desc string `json:"desc"`
}

func (d docsHandler) Edit(c *gin.Context) {
    id, _ := c.GetQuery("id")
    var docs Docs
    _ = json.Unmarshal(config, &docs)

    b, _ := json.Marshal(docs)
    result := gjson.Get(string(b), id).Value()
    c.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "data":   result,
        "msg":    "ok",
    })
}
