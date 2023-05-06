package dao

import (
	"github.com/anyingiit/GoReactResourceManagement/models"
	"gorm.io/gorm"
)

// ClientSessionByClientID
func ClientSessionByClientID(id uint) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("client_id = ?", id)
	}
}

// FindClientSessions
func FindClientSessions(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) ([]*models.ClientSession, error) {
	var clientSessions []*models.ClientSession
	query := db.Model(&models.ClientSession{})
	for _, scope := range scopes {
		query = scope(query)
	}
	err := query.Find(&clientSessions).Error
	return clientSessions, err
}

// FirstClientSession
func FirstClientSession(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (*models.ClientSession, error) {
	clientSession := models.ClientSession{}
	query := db.Model(&models.ClientSession{})
	for _, scope := range scopes {
		query = scope(query)
	}
	err := query.First(&clientSession).Error
	return &clientSession, err
}

// CreateClientSession
func CreateClientSession(db *gorm.DB, clientSession *models.ClientSession) error {
	return db.Create(clientSession).Error
}
