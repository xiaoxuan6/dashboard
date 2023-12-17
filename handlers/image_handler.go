package handlers

import (
    "dashboard/common"
    "dashboard/util"
    "fmt"
    xj "github.com/basgys/goxml2json"
    "github.com/gin-gonic/gin"
    "github.com/tidwall/gjson"
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
        util.Fail(c, err.Error())
    }

    json, err := xj.Convert(strings.NewReader(res))
    if err != nil {
        util.Fail(c, fmt.Sprintf("convert error: %s", err.Error()))
    }

    url := gjson.Parse(json.String()).Get("images.image.url").String()
    util.SuccessWithData(c, fmt.Sprintf("%s%s", common.BingUrl, url))
}

func randomXJJ(c *gin.Context) {
    res, err := util.Get(fmt.Sprintf(common.ImageUrl, util.Random(15, 1)))
    if err != nil {
        util.Fail(c, err.Error())
    }

    uri := gjson.Parse(res).Get(fmt.Sprintf("data.records.%s.url", util.Random(20, 0))).String()
    util.SuccessWithData(c, uri)
}
