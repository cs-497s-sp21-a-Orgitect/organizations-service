package main

import (
    "net/http"
)

var router = http.NewServeMux()

func InitRoutes() {
    router.HandleFunc("/", Index)
}