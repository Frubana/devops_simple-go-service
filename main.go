package main

import (
    "net/http"
)

func main() {
    r := RouteBuilder()
    http.ListenAndServe(":8080", r)
}