package main

import (
    "fmt"
    "log"
    "net/http"
    "go-rate-limit/api"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "API Currently Running", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/docker-name", api.GetDockerName)
    log.Fatal(http.ListenAndServe(":8080", nil))
}