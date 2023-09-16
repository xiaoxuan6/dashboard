package database

import "time"

var now string

func init() {
    now = time.Now().Format("2016-01-02 03:04:05")
}

func GetNow() string {
    return now
}
