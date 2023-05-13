package models

import "gorm.io/gorm"

type WebServiceType struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null;comment:服务类型名称"`
	Protocol string `gorm:"type:varchar(100);not null;comment:服务类型协议"`
}
