package main

import (
    "net/http"
    "log"
)

func main() {
    InitRoutes()
    log.Print("Server is running...")
    log.Fatal(http.ListenAndServe(":8000", router))
}