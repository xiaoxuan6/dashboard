package dirtryfilter

import (
    "bufio"
    filter2 "github.com/antlinker/go-dirtyfilter"
    "github.com/antlinker/go-dirtyfilter/store"
    "io"
    "log"
    "net/http"
    "strings"
)

const url = "https://raw.githubusercontent.com/sunshinev/go-space-chat/master/config/words_filter.txt"

var filter *filter2.DirtyManager

func init() {
    r, _ := http.NewRequest(http.MethodGet, url, nil)
    defer func() {
        if r.Body != nil && r != nil {
            _ = r.Body.Close()
        }
    }()
    res, _ := http.DefaultClient.Do(r)
    defer func() {
        if res.Body != nil && res != nil {
            _ = res.Body.Close()
        }
    }()

    words := make([]string, 0, 1000)
    br := bufio.NewReader(res.Body)
    for {
        a, _, c := br.ReadLine()
        if c == io.EOF {
            break
        }
        words = append(words, string(a))
    }

    s, err := store.NewMemoryStore(store.MemoryConfig{
        DataSource: words,
    })

    if err != nil {
        log.Printf("NewMemoryStore err %v", err)
        return
    }

    filter = filter2.NewDirtyManager(s)
}

func Filter(keyword string) string {
    result, err := filter.Filter().Filter(keyword, '*', '@')
    if err != nil {
        return ""
    }

    if result != nil {
        for _, w := range result {
            keyword = strings.ReplaceAll(keyword, w, "*")
        }
    }

    return keyword
}
