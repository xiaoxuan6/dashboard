package main

import (
    "dashboard/api"
    "net/http"
)

func main() {
    http.HandleFunc("/", api.Run)
    http.HandleFunc("/apis", api.Api)

    _ = http.ListenAndServe(":8080", nil)
}
