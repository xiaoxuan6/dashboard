package tests

import (
    "dashboard/pkg/bleve"
    _package "dashboard/pkg/package"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestBleve(t *testing.T) {
    if err := _package.Load(); err != nil {
        assert.Error(t, err)
    }

    t.Log(len(_package.Posts))

    if err := bleve.Init(); err != nil {
        assert.Error(t, err)
    }

    ids, err := bleve.Search("图片", 1)
    if err != nil {
        assert.Error(t, err)
    }

    t.Log(ids)
    t.Log(len(ids))

    assert.Nil(t, nil)
}
