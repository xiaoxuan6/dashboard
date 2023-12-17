package handlers

import (
    "dashboard/pkg/dirtryfilter"
    "dashboard/util"
    "github.com/gin-gonic/gin"
    validation "github.com/go-ozzo/ozzo-validation/v4"
    "io/ioutil"
    "net/url"
)

var DirtryfilterHandler = new(dirtryfilterHandler)

type dirtryfilterHandler struct {
}

func (d dirtryfilterHandler) Filter(c *gin.Context) {
    b, _ := ioutil.ReadAll(c.Request.Body)

    decodedString, err := url.QueryUnescape(string(b))
    if err != nil {
        util.Fail(c, err.Error())
        return
    }

    values, _ := url.ParseQuery(decodedString)
    keyword := values.Get("keyword")
    err = validation.Validate(keyword, validation.Required.Error("keyword not empty"))
    if err != nil {
        util.Fail(c, err.Error())
        return
    }

    result := dirtryfilter.Filter(keyword)
    util.SuccessWithData(c, result)
}
