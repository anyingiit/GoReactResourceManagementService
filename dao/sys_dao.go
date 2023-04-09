package dao

import (
	"github.com/anyingiit/GoReactResourceManagement/models"
	"gorm.io/gorm"
)

// FirstSysRecord
func FirstSysRecord(db *gorm.DB) (*models.Sys, error) {
	sys := models.Sys{}
	err := db.First(&sys).Error
	return &sys, err
}
