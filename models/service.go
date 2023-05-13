package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null;comment:服务名称"`
	Description string `gorm:"type:varchar(100);not null;comment:服务描述"`
	OwnerID     int
	OwnerType   string
}
