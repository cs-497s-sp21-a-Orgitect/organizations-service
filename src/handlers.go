package main

import (
    "net/http"
    "encoding/json"
    "github.com/mattn/go-sqlite3"
    "database/sql"
)

var db Conn

func Exec(command string) (runner) { // TODO add an ORM (look at gorm.io)
    statement, _ := db.Prepare(command)
    runner := func(params ...string) {
        statement.Exec(params)
    }
    return
}

/*func Query(command string) (runner) {
    resultRows, _ := db.Query(command)
}*/

func InitDb() {
    db, _ = sql.Open("sqlite3", "file::memory:?cache=shared")
    db.Exec("CREATE TABLE IF NOT EXISTS organizations (id INTEGER PRIMARY KEY, name TEXT, free_trial BOOLEAN")()
}

func Index(res http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case http.MethodGet:
    case http.MethodPost:
        var org Organization
        json.NewDecoder(req.Body).Decode(&org)
        db.Exec("INSERT INTO TABLE organizations (name, free_trial) VALUES (?, ?)")(org.Name, org.FreeTrial)
    }
    result := Organization{Name: "Test Org", FreeTrial: true}
    res.Header().Set("Content-Type", "application/json")
    json.NewEncoder(res).Encode(result)
}