package main

import (
    "net/http"
    "encoding/json"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "errors"
    "strings"
    "fmt"
    "bytes"
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
    res.Header().Set("Content-Type", "application/json")
    switch req.Method {
    case http.MethodGet:
        var org Organization
        var queryResult *gorm.DB
        name := req.URL.Query().Get("name") // get the query parameter, ?name=...
        id := req.URL.Query().Get("id")
        if name != "" && id != "" {
            queryResult = db.First(&org, "name = ?, id = ?", name, id) // search for organizations in the database matching this name
        } else if name != "" {
            queryResult = db.First(&org, "name = ?", name)
        } else if id != "" {
            queryResult = db.First(&org, "id = ?", id)
        } else {
            res.WriteHeader(http.StatusBadRequest)
            return
        }
        if errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
            res.WriteHeader(http.StatusNotFound)
        } else {
            json.NewEncoder(res).Encode(org)
        }
    case http.MethodPost:
        var org Organization
        json.NewDecoder(req.Body).Decode(&org)
        db.Create(&org)
        res.WriteHeader(http.StatusCreated)
    case http.MethodPatch:
        var org Organization
        var idToChange IdHolder
        json.NewDecoder(req.Body).Decode(&idToChange)
        queryResult := db.First(&org, "ID = ?", idToChange.Id)
        if errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
            res.WriteHeader(http.StatusNotFound)
        } else {
            var newOrg Organization
            json.NewDecoder(req.Body).Decode(&newOrg)
            org.Name = newOrg.Name
            org.FreeTrial = newOrg.FreeTrial
            // begin debug code
            fmt.Printf("%+v\n", org)
            fmt.Printf("%+v\n", newOrg)
            buf := new(bytes.Buffer)
            buf.ReadFrom(req.Body)
            fmt.Print(len(buf.String()))
            // end debug code
            db.Save(&org)
            res.WriteHeader(http.StatusNoContent)
        }
    case http.MethodDelete:
        var org Organization
        organizationId := strings.TrimPrefix(req.URL.Path, "/organizations/")
        if organizationId != "" {
            queryResult := db.First(&org, "ID = ?", organizationId)
            if errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
                res.WriteHeader(http.StatusNotFound)
            } else {
                db.Delete(&org)
                res.WriteHeader(http.StatusNoContent)
            }
        } else {
            res.WriteHeader(http.StatusBadRequest)
        }
    }
}

func Mem(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "application/json")
    switch req.Method {
    case http.MethodGet:
    }
}