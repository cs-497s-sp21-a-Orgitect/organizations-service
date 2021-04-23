package main

import (
	"gorm.io/gorm"
)

type Organization struct {
    gorm.Model
    Name string `json:"name"`
    FreeTrial bool `json:"trial"`
}

type IdHolder struct {
    Id int `json:"id,string"` // the string tells decoder to convert id that might be a string into a number
}