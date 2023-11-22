package util

import (
    "math/rand"
    "strconv"
    "strings"
    "time"
)

func Contains(s string) func(item string) bool {
    return func(item string) bool {
        return strings.Contains(s, item)
    }
}

func Random(num int, incr int) string {
    rand.Seed(time.Now().UnixNano())

    randomNum := rand.Intn(num) + incr

    return strconv.Itoa(randomNum)
}
