package models

import "gorm.io/gorm"

type InternalService struct {
	gorm.Model
	ServiceID uint     `gorm:"not null"`
	Service   *Service `gorm:"polymorphic:Owner;polymorphicValue:internal_services"`
}
