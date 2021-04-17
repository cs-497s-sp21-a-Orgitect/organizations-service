package main

import (
    "net/http"
    "encoding/json"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "fmt"
)

var db *sql.DB

func InitDb() {
    db, _ = gorm.Open(sqlite3.Open("file::memory:?cache=shared"), &gorm.Config{})
    Exec("CREATE TABLE IF NOT EXISTS organizations (id INTEGER PRIMARY KEY, name TEXT, free_trial BOOLEAN);")()
}

func Index(res http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case http.MethodGet:
    case http.MethodPost:
        var org Organization
        json.NewDecoder(req.Body).Decode(&org)
        fmt.Printf("%+v\n", org)
        Exec("INSERT INTO organizations(name, free_trial) VALUES(?, ?);")(org.Name, org.FreeTrial)
    }
    result := Organization{Name: "Test Org", FreeTrial: true}
    res.Header().Set("Content-Type", "application/json")
    json.NewEncoder(res).Encode(result)
}