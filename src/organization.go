package main

import (
	"gorm.io/gorm"
)

type Organization struct {
    gorm.Model
    Name string
    Id int
    FreeTrial bool
}