package main

import (
    "net/http"
    "log"
)

func main() {
    InitRoutes()
    InitDb()
    log.Print("Go server is running...")
    log.Fatal(http.ListenAndServe(":8080", router))
}