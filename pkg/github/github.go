package github

import (
    "context"
    "github.com/google/go-github/v48/github"
    "golang.org/x/oauth2"
    "os"
    "regexp"
    "strings"
    "sync"
)

var (
    Github *github.Client
    one    sync.Once
    lock   sync.Mutex
    ctx    = context.Background()
    Menus  []string
    Data   = make(map[string][]Item) // {"test1":[{"title": "xxx", "url": "xxx"}, {"title": "xxx", "url": "xxx"}]}
)

func Init() {
    token := os.Getenv("GITHUB_TOKEN")
    one.Do(func() {
        oauth := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
            AccessToken: token,
        }))

        Github = github.NewClient(oauth)
    })
}

type Item struct {
    Title string `json:"title"`
    Url   string `json:"url"`
}

func FetchMenus() {
    _, directoryContent, _, _ := Github.Repositories.GetContents(ctx, os.Getenv("GITHUB_OWNER"), os.Getenv("GITHUB_REPO"), "", &github.RepositoryContentGetOptions{})

    for _, val := range directoryContent {
        if strings.HasSuffix(val.GetName(), ".md") {
            Menus = append(Menus, val.GetName())
        }
    }

    return
}

func FetchContent(filename string, wg *sync.WaitGroup) {
    defer wg.Done()
    repositoryContent, _, _, err := Github.Repositories.GetContents(ctx, os.Getenv("GITHUB_OWNER"), os.Getenv("GITHUB_REPO"), filename, &github.RepositoryContentGetOptions{})
    if err != nil {
        return
    }

    content, err2 := repositoryContent.GetContent()
    if err2 != nil {
        return
    }

    var items []Item
    contents := strings.Split(content, "\n")
    for _, val := range contents {
        url := regexpUrl(val)
        if url == "" {
            continue
        }

        title := regexpTitle(val)
        items = append(items, Item{
            Title: title,
            Url:   url,
        })
    }

    lock.Lock()
    menu := strings.ReplaceAll(filename, ".md", "")
    Data[menu] = items
    lock.Unlock()
}

func regexpTitle(str string) string {
    re := regexp.MustCompile(`\[(.*?)\]`)
    matches := re.FindStringSubmatch(str)
    if len(matches) > 1 {
        return matches[1]
    }

    return ""
}

func regexpUrl(str string) string {
    re := regexp.MustCompile(`\((.*?)\)`)
    matches := re.FindStringSubmatch(str)
    if len(matches) > 1 {
        return matches[1]
    }
    return ""
}
