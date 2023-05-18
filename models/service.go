package models

import "gorm.io/gorm"

type OwnerType string

const (
	OwnerTypeWebService OwnerType = "web_services"
)

type Service struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(100);not null;comment:服务名称"`
	Description string    `gorm:"type:varchar(100);not null;comment:服务描述"`
	ClientID    uint      `gorm:"type:int;not null;comment:客户端ID"`
	Client      Client    `gorm:"foreignKey:ClientID"`
	OwnerID     uint      `gorm:"type:int;not null;comment:所有者ID"`
	OwnerType   OwnerType `gorm:"type:varchar(100);not null;comment:所有者类型"`
}

func (s *Service) IsOwnerType(ownerType OwnerType) bool {
	return string(s.OwnerType) == string(ownerType)
}
