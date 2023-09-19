package handlers

import (
    "dashboard/pkg/dirtryfilter"
    "github.com/gin-gonic/gin"
    validation "github.com/go-ozzo/ozzo-validation/v4"
    "github.com/sirupsen/logrus"
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
    logrus.Info("keyword", keyword)

    k, _ := c.GetPostForm("keyword")
    logrus.Info("keyword1", k)

    err := validation.Validate(keyword, validation.Required)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status": http.StatusBadRequest,
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
