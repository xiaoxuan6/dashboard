package lib

import (
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
