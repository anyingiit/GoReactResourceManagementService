package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientSession struct {
	gorm.Model
	UUID     uuid.UUID
	ClientID uint   `gorm:"uniqueIndex:idx_client_id"`
	Client   Client `gorm:"foreignKey:ClientID"`
}
