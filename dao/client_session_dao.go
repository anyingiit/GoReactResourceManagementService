package dao

import (
	"github.com/anyingiit/GoReactResourceManagement/models"
	"gorm.io/gorm"
)

// ClientSessionByClientID
func ClientSessionByClientID(id uint) func(*gorm.DB) *gorm.DB {
	return ByField("client_id", id)
}

// FindClientSessions retrieves a list of client sessions based on given conditions.
func FindClientSessions(db *gorm.DB, scopes ...ScopeFunc) ([]*models.ClientSession, error) {
	var clientSessions []*models.ClientSession
	err := Find(db, &clientSessions, scopes...)
	return clientSessions, err
}

// FirstClientSession retrieves the first client session based on given conditions.
func FirstClientSession(db *gorm.DB, scopes ...ScopeFunc) (*models.ClientSession, error) {
	var clientSession models.ClientSession
	err := First(db, &clientSession, scopes...)
	return &clientSession, err
}

// CreateClientSession
func CreateClientSession(db *gorm.DB, clientSession *models.ClientSession) error {
	return db.Create(clientSession).Error
}
