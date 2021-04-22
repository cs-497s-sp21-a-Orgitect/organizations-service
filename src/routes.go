package main

import (
    "net/http"
)

var router = http.NewServeMux()

func InitRoutes() {
    router.HandleFunc("/organizations/", Org) // two routes are needed to handle when there is/isn't a path parameter
    router.HandleFunc("/organizations", Org)
    router.HandleFunc("/members/", Mem)
    router.HandleFunc("/members", Mem)
}