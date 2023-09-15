package main

import (
    "dashboard/api"
    "net/http"
)

func main() {
    http.HandleFunc("/", api.Run)

    _ = http.ListenAndServe(":8080", nil)
}
