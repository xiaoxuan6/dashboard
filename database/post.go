package database

import "time"

var now string

func init() {
    now = time.Now().Format("2006-01-02 03:04:05")
}

func GetNow() string {
    return now
}
