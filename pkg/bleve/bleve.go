package bleve

import (
    _package "dashboard/pkg/package"
    "fmt"
    "github.com/blevesearch/bleve/v2"
    "github.com/sirupsen/logrus"
)

var index bleve.Index

type Item struct {
    Title string `json:"title"`
    Url   string `json:"url"`
    Tag   string `json:"tag"`
}

func Init() error {
    ind, err := bleve.New("", bleve.NewIndexMapping())
    if err != nil {
        return err
    }

    for i, post := range _package.Posts {
        _ = ind.Index(fmt.Sprintf("%d", i), post)
    }

    index = ind
    return nil
}

func Search(keyword string) ([]Item, error) {
    query := bleve.NewMatchQuery(keyword)
    search := bleve.NewSearchRequest(query)
    search.Fields = []string{"Title", "Url", "Tag"}
    search.Highlight = bleve.NewHighlight()
    searchResults, err := index.Search(search)
    if err != nil {
        return nil, err
    }

    logrus.Info("searchResults.Hits", searchResults.Hits)
    var items []Item
    for _, hit := range searchResults.Hits {
        logrus.Info("hit.Fields", hit.Fields)
        if hit.Fields["Title"] != nil {
            title := hit.Fields["Title"].(string)
            items = append(items, Item{
                Title: title,
                Url:   hit.Fields["Url"].(string),
                Tag:   hit.Fields["Tag"].(string),
            })
        }
    }

    return items, nil
}
