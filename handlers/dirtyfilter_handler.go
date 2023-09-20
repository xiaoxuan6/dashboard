package handlers

import (
    "dashboard/pkg/dirtryfilter"
    "github.com/gin-gonic/gin"
    validation "github.com/go-ozzo/ozzo-validation/v4"
    "io/ioutil"
    "net/http"
    "net/url"
)

var DirtryfilterHandler = newDritryfilter()

type dirtryfilterHandler struct {
}

func newDritryfilter() *dirtryfilterHandler {
    return new(dirtryfilterHandler)
}

func (d dirtryfilterHandler) Filter(c *gin.Context) {
    b, _ := ioutil.ReadAll(c.Request.Body)

    decodedString, err := url.QueryUnescape(string(b))
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status": http.StatusBadRequest,
            "data":   "",
            "msg":    err.Error(),
        })
        return
    }

    values, _ := url.ParseQuery(decodedString)
    keyword := values.Get("keyword")
    err = validation.Validate(keyword, validation.Required.Error("keyword not empty"))
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status": http.StatusBadRequest,
            "data":   "",
            "msg":    err.Error(),
        })
        return
    }

    result := dirtryfilter.Filter(keyword)
    c.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "data":   result,
        "msg":    "ok",
    })
}
