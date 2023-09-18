package tests

import (
    "dashboard/lib"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "net/http"
    "strings"
    "testing"
)

func TestSearch(t *testing.T) {
    r := &http.Request{
        Body: ioutil.NopCloser(strings.NewReader("keyword=test")),
    }

    var search lib.SearchHandler
    response := search.Do(r)
    t.Log(response)
    assert.Equal(t, 200, response.Status)
}
