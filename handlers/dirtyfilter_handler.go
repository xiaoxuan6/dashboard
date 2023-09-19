package handlers

import (
    "dashboard/pkg/dirtryfilter"
    "github.com/gin-gonic/gin"
    validation "github.com/go-ozzo/ozzo-validation/v4"
    "net/http"
)

var DirtryfilterHandler = newDritryfilter()

type dirtryfilterHandler struct {
}

func newDritryfilter() *dirtryfilterHandler {
    return new(dirtryfilterHandler)
}

func (d dirtryfilterHandler) Filter(c *gin.Context) {
    keyword := c.PostForm("keyword")

    err := validation.Validate(keyword, validation.Required)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status": 400,
            "data":   "",
            "msg":    err.Error(),
        })
    }

    result := dirtryfilter.Filter(keyword)
    c.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "data":   result,
        "msg":    "ok",
    })
}
