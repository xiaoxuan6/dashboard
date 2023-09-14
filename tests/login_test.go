package tests

import (
    "dashboard/lib"
    "fmt"
    "github.com/stretchr/testify/assert"
    "net/http"
    "net/url"
    "testing"
)

var loginHandler = lib.LoginHandler{}
var tokenHandler = lib.TokenHandler{}

func TestLogin(t *testing.T) {
    u := &url.URL{
        RawQuery: "email=123%40qq.com&passwd=123456",
    }
    request := &http.Request{
        URL: u,
    }
    response := loginHandler.Do(request)
    assert.Equal(t, 200, response.Status)
    data := response.Data.(lib.LoginRepose)
    assert.NotEmpty(t, data.Token)

    tu := &url.URL{
        RawQuery: fmt.Sprintf("email=%s&token=%s", data.Email, data.Token),
    }
    request = &http.Request{
        URL: tu,
    }
    response2 := tokenHandler.Do(request)
    assert.Equal(t, 200, response2.Status)
}
