package common

import (
    "fmt"
    "os"
)

var Token = fmt.Sprintf("%s%s", os.Getenv("CACHE_PREFIX"), "token")
var Menu = fmt.Sprintf("%s%s", os.Getenv("CACHE_PREFIX"), "menu")
var Data = fmt.Sprintf("%s%s", os.Getenv("CACHE_PREFIX"), "data")
