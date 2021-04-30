package main

import (
    "net/http"
    "encoding/json"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "errors"
    "strings"
    "io/ioutil"
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
    var org Organization
    body, _ := ioutil.ReadAll(req.Body) // if the body itself is read again after the first read, it will be nil
    switch req.Method {
    case http.MethodGet:
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
        json.NewDecoder(bytes.NewReader(body)).Decode(&org)
        db.Create(&org)
        res.WriteHeader(http.StatusCreated)
    case http.MethodPatch:
        var idToChange IdHolder
        json.NewDecoder(bytes.NewReader(body)).Decode(&idToChange)
        queryResult := db.First(&org, "ID = ?", idToChange.Id)
        if errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
            res.WriteHeader(http.StatusNotFound)
        } else {
            var newOrg Organization
            json.NewDecoder(bytes.NewReader(body)).Decode(&newOrg)
            org.Name = newOrg.Name
            org.FreeTrial = newOrg.FreeTrial
            db.Save(&org)
            res.WriteHeader(http.StatusNoContent)
        }
    case http.MethodDelete:
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