package main

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	OrganizationId int
	ActorId int
}