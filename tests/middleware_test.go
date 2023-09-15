package tests

import (
    "bytes"
    "dashboard/lib"
    middlewares2 "dashboard/middlewares"
    "encoding/json"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "net/http"
    "os"
    "testing"
)

func login() *lib.Response {
    _ = os.Setenv("VERCEL_EMIAL", "123@qq.com")
    _ = os.Setenv("VERCEL_PASSWORD", "123456")

    form := &lib.LoginRequest{
        Email:  "123@qq.com",
        Passwd: "123456",
    }
    b, _ := json.Marshal(form)

    body := bytes.NewReader(b)
    request := &http.Request{
        Body: ioutil.NopCloser(body),
    }
    var LoginHandler lib.LoginHandler
    return LoginHandler.Do(request)
}

func TestMiddleware(t *testing.T) {
    response := login()
    data := response.Data.(*lib.LoginResponse)
    assert.NotEmpty(t, data.Token)

    header := http.Header{}
    header.Set("Authorization", data.Token)
    header.Set("email", data.Email)

    request2 := &http.Request{
        Header: header,
    }

    var middleware middlewares2.AuthMiddleware
    response2 := middleware.Handler(request2)
    t.Log(response2)
    assert.Equal(t, 200, response2.Status)
}
