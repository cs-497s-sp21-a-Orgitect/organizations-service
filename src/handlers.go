package main

import (
    "net/http"
    "encoding/json"
    _ "github.com/mattn/go-sqlite3"
    "database/sql"
   // "strconv"
    "fmt"
)

var db *sql.DB

func Exec(command string) func(...interface {}) (sql.Result, error) { // TODO add an ORM (look at gorm.io)
    statement, _ := db.Prepare(command)
    return statement.Exec
}

/*func Query(command string) (runner) {
    resultRows, _ := db.Query(command)
}*/

func InitDb() {
    db, _ = sql.Open("sqlite3", "file::memory:?cache=shared")
    Exec("CREATE TABLE IF NOT EXISTS organizations (id INTEGER PRIMARY KEY, name TEXT, free_trial BOOLEAN")()
}

func Index(res http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case http.MethodGet:
    case http.MethodPost:
        var org Organization
        json.NewDecoder(req.Body).Decode(&org)
        fmt.Printf("%+v\n", org)
        Exec("INSERT INTO TABLE organizations (name, free_trial) VALUES (?, ?)")(org.Name, "false")//strconv.FormatBool(org.FreeTrial))
    }
    result := Organization{Name: "Test Org", FreeTrial: true}
    res.Header().Set("Content-Type", "application/json")
    json.NewEncoder(res).Encode(result)
}