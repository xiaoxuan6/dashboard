package bleve

import (
    _package "dashboard/pkg/package"
    "fmt"
    "github.com/blevesearch/bleve/v2"
    "strconv"
    "strings"
    "sync"
)

var (
    index    bleve.Index
    Fields   = []string{"Title", "Url", "Tag"}
    lock     sync.Mutex
    Total    uint64
    pageSize = 100
)

func Init() error {
    ind, err := bleve.New("", bleve.NewIndexMapping())
    if err != nil {
        return err
    }

    for i, post := range _package.Posts {
        id := strconv.Itoa(i)
        _ = ind.Index(id, post)
    }

    index = ind
    return nil
}

func Search(keyword string, page int) ([]_package.Post, error) {
    var posts []_package.Post

    from := (page - 1) * pageSize
    searchResults, err := queryWithPageExec(keyword, pageSize, from)
    if err != nil {
        return posts, err
    }

    Total = searchResults.Total
    for _, hit := range searchResults.Hits {
        if hit.Fragments == nil {
            continue
        }

        for fragmentField, fragments := range hit.Fragments {
            // 高亮字段只处理 title
            if fragmentField != "Title" && fragmentField != "title" {
                continue
            }

            var highlightValue string
            for _, fragment := range fragments {
                highlightValue = fmt.Sprintf("%s", fragment)
            }
            id, _ := strconv.Atoi(hit.ID)
            post := _package.Posts[id]

            highlightValue = strings.ReplaceAll(highlightValue, "<mark>", "<span class=\"h-keyword\">")
            highlightValue = strings.ReplaceAll(highlightValue, "</mark>", "</span>")

            lock.Lock()
            posts = append(posts, _package.Post{
                Title: highlightValue,
                Url:   post.Url,
                Tag:   post.Tag,
            })
            lock.Unlock()
        }
    }

    if len(posts) < 1 {
        for _, hit := range searchResults.Hits {
            i, _ := strconv.Atoi(hit.ID)
            post := _package.Posts[i]

            lock.Lock()
            posts = append(posts, post)
            lock.Unlock()
        }
    }

    return posts, nil
}

func queryExec(keyword string) (*bleve.SearchResult, error) {
    q := bleve.NewQueryStringQuery(keyword)
    search := bleve.NewSearchRequest(q)
    search.Fields = Fields
    search.Highlight = bleve.NewHighlight()
    searchResults, err := index.Search(search)

    return searchResults, err
}

func queryWithPageExec(keyword string, size, from int) (*bleve.SearchResult, error) {
    q := bleve.NewQueryStringQuery(keyword)
    search := bleve.NewSearchRequestOptions(q, size, from, false)
    search.Fields = Fields
    search.Highlight = bleve.NewHighlight()
    searchResults, err := index.Search(search)

    return searchResults, err
}
