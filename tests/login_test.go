package tests

import (
    "bytes"
    "dashboard/lib"
    "encoding/json"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "net/http"
    "testing"
)

func TestLogin(t *testing.T) {
    response := login()
    t.Log(response)
    assert.Equal(t, 200, response.Status)
    data := response.Data.(*lib.LoginResponse)
    assert.NotEmpty(t, data.Token)

    form2 := &lib.LoginResponse{
        Token: data.Token,
        Email: data.Email,
    }
    b2, _ := json.Marshal(form2)

    body2 := bytes.NewReader(b2)
    request2 := &http.Request{
        Body: ioutil.NopCloser(body2),
    }
    var tokenHandler lib.TokenHandler
    response2 := tokenHandler.Do(request2)
    t.Log(response2)
    assert.Equal(t, 200, response2.Status)
}
