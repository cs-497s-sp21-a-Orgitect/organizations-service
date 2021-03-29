package main

import (
    "net/http"
    "encoding/json"
)

func Index(res http.ResponseWriter, req *http.Request) {
    result := Organization{Name: "Test Org", FreeTrial: true}
    res.Header().Set("Content-Type", "application/json")
    json.NewEncoder(res).Encode(result)
}