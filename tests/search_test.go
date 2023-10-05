package tests

import (
    "dashboard/lib"
    _package "dashboard/pkg/package"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "net/http"
    "os"
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

// go test -run TestSearchLoad tests/search_test.go -v
func TestSearchLoad(t *testing.T) {
    _ = os.Setenv("GITHUB_TOKEN", "xxx")
    _ = os.Setenv("GITHUB_OWNER", "xiaoxuan6")
    _ = os.Setenv("GITHUB_REPO", "resource")

    lib.Load()
    t.Log(_package.Posts)
    t.Log(len(_package.Posts))
    assert.Nil(t, nil)
}
