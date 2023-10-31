package handlers

import (
    "fmt"
    emailverifier "github.com/AfterShip/email-verifier"
    "github.com/gin-gonic/gin"
    validation "github.com/go-ozzo/ozzo-validation/v4"
    "github.com/go-ozzo/ozzo-validation/v4/is"
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
        validation.Required.Error("email not empty"),
        validation.Length(5, 50).Error("email length 5-50"),
        is.Email,
    )
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status": http.StatusBadRequest,
            "data":   "",
            "msg":    err.Error(),
        })
        return
    }

    var verifier = emailverifier.
        NewVerifier().
        EnableAutoUpdateDisposable()

    if verifier.IsDisposable(email) {
        c.JSON(http.StatusNotFound, gin.H{
            "status": http.StatusNotFound,
            "data":   "",
            "msg":    "disposable email",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "data":   "",
        "msg":    "ok",
    })
}
