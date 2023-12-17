package handlers

import (
    "dashboard/util"
    _ "embed"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "net/http"
)

//go:embed config.json
var config []byte

var DocsHandler = new(docsHandler)

type docsHandler struct {
}

type (
    Config struct {
        Apis []Api `json:"apis"`
    }
    Api struct {
        Url   string `json:"url"`
        Title string `json:"title"`
        Desc  string `json:"desc"`
    }
)

func (d docsHandler) Index(c *gin.Context) {
    var cs Config
    _ = json.Unmarshal(config, &cs)

    util.SuccessWithData(c, cs)
}

type (
    Docs struct {
        Docs struct {
            Dirtyfilter Item `json:"dirtyfilter"`
            EmailCheck  Item `json:"email_check"`
            RandomImg   Item `json:"random_img"`
        } `json:"docs"`
    }
    Item struct {
        Url      string     `json:"url"`
        Method   string     `json:"method"`
        UrlDemo  string     `json:"url_demo"`
        Params   []params   `json:"params"`
        Response []response `json:"response"`
        Codes    []codes    `json:"codes"`
    }
    params struct {
        Name    string `json:"name"`
        Require string `json:"require"`
        Type    string `json:"type"`
        Desc    string `json:"desc"`
    }
    response struct {
        Name string `json:"name"`
        Type string `json:"type"`
        Desc string `json:"desc"`
    }
    codes struct {
        Code string `json:"code"`
        Desc string `json:"desc"`
    }
)

func (d docsHandler) Show(c *gin.Context) {
    id := c.Param("id")
    var docs Docs
    _ = json.Unmarshal(config, &docs)

    var result interface{}
    switch id {
    case "dirtyfilter":
        result = docs.Docs.Dirtyfilter
    case "email_check":
        result = docs.Docs.EmailCheck
    case "random_img":
        result = docs.Docs.RandomImg
    }

    util.SuccessWithData(c, result)
}
