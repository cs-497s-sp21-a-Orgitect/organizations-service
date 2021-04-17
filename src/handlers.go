package main

import (
    "net/http"
    "encoding/json"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

var db *gorm.DB

func InitDb() {
    var err error
    db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&Organization{})
}

func Index(res http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case http.MethodGet:
        var org Organization
        db.First(&org, "name = ?", req.URL.Query().Get("name"))
        json.NewEncoder(res).Encode(org)
    case http.MethodPost:
        var org Organization
        json.NewDecoder(req.Body).Decode(&org)
        db.Create(&org)
        res.WriteHeader(http.StatusCreated)
    }
    res.Header().Set("Content-Type", "application/json")
}