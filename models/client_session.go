package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientSession struct {
	gorm.Model
	ClientID uint
	Client   Client `gorm:"foreignKey:ClientID"`
	UUID     uuid.UUID
}
