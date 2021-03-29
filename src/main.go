package main

import (
    "net/http"
    "log"
)

func main() {
    InitRoutes()
    log.Fatal(http.ListenAndServe(":8000", router))
}