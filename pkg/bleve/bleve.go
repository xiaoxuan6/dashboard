package bleve

import (
    _package "dashboard/pkg/package"
    "github.com/blevesearch/bleve/v2"
    "strconv"
)

var index bleve.Index

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

func Search(keyword string) ([]string, error) {
    query := bleve.NewMatchQuery(keyword)
    search := bleve.NewSearchRequest(query)
    search.Fields = []string{"Title", "Url", "Tag"}
    search.Highlight = bleve.NewHighlight()
    searchResults, err := index.Search(search)
    if err != nil {
        return nil, err
    }

    var ids []string
    for _, hit := range searchResults.Hits {
        ids = append(ids, hit.ID)
    }

    return ids, nil
}
