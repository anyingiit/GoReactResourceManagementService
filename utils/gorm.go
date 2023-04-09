package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateNewGormConnection(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
