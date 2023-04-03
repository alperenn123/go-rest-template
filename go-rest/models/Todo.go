package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Username string `json:"userName" gorm:"size:60;not null"`
	Desc     string `json:"desc" gorm:"not null"`
}
