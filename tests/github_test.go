package tests

import (
    "bytes"
    "encoding/json"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "net/http"
    "net/url"
    "testing"
)

func TestGithub(t *testing.T) {
    body2 := bytes.NewReader([]byte("token=123456"))
    r := &http.Request{
        Body: ioutil.NopCloser(body2),
        URL: &url.URL{
            RawQuery: "token=xxxx",
        },
    }

    _ = r.ParseForm()
    token := r.FormValue("token") // xxxx
    t.Log(token)
    assert.NotEmpty(t, token)

    b, _ := ioutil.ReadAll(r.Body)
    t.Log(string(b)) // 123456
    decodedString, err := url.QueryUnescape(string(b))
    assert.Nil(t, err)
    t.Log(decodedString) // 123456

    token = r.URL.Query().Get("token") // xxxx
    t.Log(token)

    type Response struct {
        Token string `json:"token"`
    }

    // 后面直接报错
    var res Response
    //err = json.Unmarshal(b, &res)
    //if err != nil {
    //    t.Log(fmt.Sprintf("json 解析错误：%s", err.Error()))
    //}
    decoder := json.NewDecoder(bytes.NewReader(b))
    //decoder.DisallowUnknownFields()
    _ = decoder.Decode(&res)

    t.Log(res.Token)
    assert.NotEmpty(t, res.Token)
}
