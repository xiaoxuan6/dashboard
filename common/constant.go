package common

import (
    "fmt"
    "os"
)

var Token = fmt.Sprintf("%s%s", os.Getenv("CACHE_PREFIX"), "token")
