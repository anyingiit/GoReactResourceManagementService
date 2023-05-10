package dao

// ClientSessionByClientID
func ClientSessionByClientID(id uint) ScopeFunc {
	return ByField("client_id", id)
}

/* 使用dao通用函数替换了如下专有函数
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
*/
