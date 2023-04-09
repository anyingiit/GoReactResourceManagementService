package models

import (
	"gorm.io/gorm"
)

type Sys struct {
	gorm.Model
}

func (*Sys) TableName() string {
	return "sys"
}
