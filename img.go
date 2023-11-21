package main

import (
    "dashboard/common"
    "dashboard/util"
    "fmt"
    xj "github.com/basgys/goxml2json"
    "github.com/tidwall/gjson"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)

func main() {
    //res, err := util.Get(fmt.Sprintf(common.ImageUrl, random(15, 1)))
    //if err != nil {
    //   fmt.Println(err.Error())
    //   os.Exit(1)
    //}
    //
    //uri := gjson.Parse(res).Get(fmt.Sprintf("data.records.%s.url", random(20, 0))).String()
    //fmt.Println(uri)

    res, err := util.Get(fmt.Sprintf(common.BingImageUrl, random(10, 0)))
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    json, err := xj.Convert(strings.NewReader(res))
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    url := gjson.Parse(json.String()).Get("images.image.url").String()
    fmt.Println(fmt.Sprintf("%s%s", common.BingUrl, url))
}

func random(num int, incr int) string {
    rand.Seed(time.Now().UnixNano())

    randomNum := rand.Intn(num) + incr

    return strconv.Itoa(randomNum)
}
