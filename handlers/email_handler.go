package handlers

import (
    "fmt"
    "github.com/gin-gonic/gin"
    validation "github.com/go-ozzo/ozzo-validation/v4"
    "github.com/go-ozzo/ozzo-validation/v4/is"
    "github.com/mcnijman/go-emailaddress"
    "io/ioutil"
    "net/http"
    "net/url"
)

var EmailHandler = new(emailHandler)

type emailHandler struct {
}

func (e emailHandler) Check(c *gin.Context) {
    b, _ := ioutil.ReadAll(c.Request.Body)

    decodedString, err := url.QueryUnescape(string(b))
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status": http.StatusBadRequest,
            "data":   "",
            "msg":    fmt.Sprintf("url decode error: %s", err.Error()),
        })
        return
    }

    values, _ := url.ParseQuery(decodedString)
    email := values.Get("email")
    err = validation.Validate(
        email,
        validation.Map(
            validation.Key("email", validation.Required.Error("email not empty")),
            validation.Key("email", validation.Length(5, 50).Error("email length 5-50")),
            validation.Key("email", validation.Required, is.Email.Error("email format error")),
        ),
    )
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status": http.StatusBadRequest,
            "data":   "",
            "msg":    err.Error(),
        })
        return
    }

    emailed, err := emailaddress.Parse(email)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status": http.StatusBadRequest,
            "data":   "",
            "msg":    fmt.Sprintf("email parse error: %s", err.Error()),
        })
        return
    }

    if err := emailed.ValidateHost(); err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status": http.StatusBadRequest,
            "data":   "",
            "msg":    "email host error",
        })
        return
    }

    if err := emailed.ValidateIcanSuffix(); err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status": http.StatusBadRequest,
            "data":   "",
            "msg":    "email suffix error",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "data":   "",
        "msg":    "ok",
    })
}
