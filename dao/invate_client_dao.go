package dao

import (
	"github.com/anyingiit/GoReactResourceManagement/models"
	"gorm.io/gorm"
)

// InvateClientByClientID
func InvateClientByClientID(id uint) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("client_id = ?", id)
	}
}

// InvateClientByInvateCode
func InvateClientByInvateCode(invateCode string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("invate_code = ?", invateCode)
	}
}

func CreateInvateClient(db *gorm.DB, invateClient *models.InvateClient) error {
	return db.Create(invateClient).Error
}

// FindInvateClients
func FindInvateClients(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) ([]*models.InvateClient, error) {
	var invateClients []*models.InvateClient
	query := db.Model(&models.InvateClient{})
	for _, scope := range scopes {
		query = scope(query)
	}
	err := query.Find(&invateClients).Error
	return invateClients, err
}

// FindInvateClient
func FirstInvateClient(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (*models.InvateClient, error) {
	invateClient := models.InvateClient{}
	query := db.Model(&models.InvateClient{})
	for _, scope := range scopes {
		query = scope(query)
	}
	err := query.First(&invateClient).Error
	return &invateClient, err
}

// DeleteInvateClient
func DeleteInvateClient(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) error {
	query := db.Model(&models.InvateClient{})
	for _, scope := range scopes {
		query = scope(query)
	}
	return query.Delete(&models.InvateClient{}).Error
}
