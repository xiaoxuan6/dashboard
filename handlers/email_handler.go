package handlers

import (
    "dashboard/util"
    "fmt"
    emailverifier "github.com/AfterShip/email-verifier"
    "github.com/gin-gonic/gin"
    validation "github.com/go-ozzo/ozzo-validation/v4"
    "github.com/go-ozzo/ozzo-validation/v4/is"
    "io/ioutil"
    "net/url"
)

var EmailHandler = new(emailHandler)

type emailHandler struct {
}

func (e emailHandler) Check(c *gin.Context) {
    b, _ := ioutil.ReadAll(c.Request.Body)

    decodedString, err := url.QueryUnescape(string(b))
    if err != nil {
        util.Fail(c, fmt.Sprintf("url decode error: %s", err.Error()))
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
        util.Fail(c, err.Error())
    }

    var verifier = emailverifier.
        NewVerifier().
        EnableAutoUpdateDisposable()

    if verifier.IsDisposable(email) {
        util.Fail(c, "disposable email")
    }

    util.Success(c)
}
