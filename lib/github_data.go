package lib

import (
    "dashboard/common"
    "dashboard/pkg/cache"
    "dashboard/pkg/github"
)

func LoadData() {
    github.Init()
    github.FetchMenus()
    for _, val := range github.Menus {
        wg.Add(1)
        github.FetchContent(val, &wg)
    }

    wg.Done()

    Save(github.Menus, github.Data)
}

func Reload() {
    status := true
    menu, ok := cache.Cache.Get(common.Menu)
    if !ok {
        status = false
    }

    data, ok := cache.Cache.Get(common.Data)
    if !ok {
        status = false
    }

    if status {
        github.Menus = menu.([]string)
        github.Data = data.(map[string][]github.Item)
    } else {
        LoadData()
    }
}
