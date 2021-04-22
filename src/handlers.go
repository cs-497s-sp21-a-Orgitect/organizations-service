package main

import (
    "net/http"
    "encoding/json"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "errors"
)

var db *gorm.DB

func InitDb() {
    var err error
    db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{}) // database will be stored in memory
    if err != nil {
        panic("Error starting database")
    }
    // create the tables in the database
    db.AutoMigrate(&Organization{})
    db.AutoMigrate(&Member{})
}

func Org(res http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case http.MethodGet:
        var org Organization
        name := req.URL.Query().Get("name") // get the query parameter, ?name=...
        if name != "" {
            queryResult := db.First(&org, "name = ?", req.URL.Query().Get("name")) // search for organizations in the database matching this name
            if errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
                res.WriteHeader(http.StatusNotFound)
            } else {
                json.NewEncoder(res).Encode(org)
            }
        } else {
            res.WriteHeader(http.StatusBadRequest)
        }
    case http.MethodPost:
        var org Organization
        json.NewDecoder(req.Body).Decode(&org)
        db.Create(&org)
        res.WriteHeader(http.StatusCreated)
    }
    res.Header().Set("Content-Type", "application/json")
}

func Mem(res http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case http.MethodGet:
    }
    res.Header().Set("Content-Type", "application/json")
}