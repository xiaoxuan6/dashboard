package handlers

import (
    "dashboard/common"
    "dashboard/util"
    "fmt"
    xj "github.com/basgys/goxml2json"
    "github.com/gin-gonic/gin"
    "github.com/tidwall/gjson"
    "net/http"
    "strings"
)

var ImageHandler = new(imageHandler)

type imageHandler struct {
}

func (i imageHandler) Random(c *gin.Context) {
    t := c.Query("type")

    switch t {
    case "1":
        randomImg(c)
    case "2":
        randomXJJ(c)
    default:
        randomImg(c)
    }

}

func randomImg(c *gin.Context) {
    res, err := util.Get(common.BingImageUrl)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status": http.StatusBadRequest,
            "data":   "",
            "msg":    err.Error(),
        })
    }

    json, err := xj.Convert(strings.NewReader(res))
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status": http.StatusBadRequest,
            "data":   "",
            "msg":    fmt.Sprintf("convert error: %s", err.Error()),
        })
    }

    url := gjson.Parse(json.String()).Get("images.image.url").String()
    c.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "data":   fmt.Sprintf("%s%s", common.BingUrl, url),
        "msg":    "ok",
    })
}

func randomXJJ(c *gin.Context) {
    res, err := util.Get(fmt.Sprintf(common.ImageUrl, util.Random(15, 1)))
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status": http.StatusBadRequest,
            "data":   "",
            "msg":    err.Error(),
        })
    }

    uri := gjson.Parse(res).Get(fmt.Sprintf("data.records.%s.url", util.Random(20, 0))).String()
    c.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "data":   uri,
        "msg":    "ok",
    })
}
