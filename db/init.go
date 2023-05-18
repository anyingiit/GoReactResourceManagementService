package db

import (
	"fmt"

	"github.com/anyingiit/GoReactResourceManagement/globalVars"
	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/anyingiit/GoReactResourceManagement/structs"
	"github.com/anyingiit/GoReactResourceManagement/utils"
	"gorm.io/gorm"
)

func InitDB(config *structs.DatabaseConfig) (*gorm.DB, error) {
	// connect to database
	db, err := utils.CreateNewGormConnection(utils.GenerationMysqlDsn(config.Username, config.Password, config.Host, fmt.Sprintf("%d", config.Port), config.Database))
	if err != nil {
		return nil, fmt.Errorf("failed to connect database, %v", err)
	}

	// set global DB var
	err = globalVars.Db.Set(db)
	if err != nil {
		return nil, fmt.Errorf("failed to set db, %v", err)
	}

	tx := db.Begin()
	// 自动同步数据库结构
	err = tx.AutoMigrate(
		// Sys
		&models.Sys{},
		// User
		&models.Role{},
		&models.User{},
		// Client
		&models.Client{},
		&models.InvateClient{},
		&models.ClientSession{},
		// Task
		&models.Task{},

		// Queue Task
		&models.TaskQueue{},
		&models.TaskQueueResult{},

		// Service
		&models.Service{},
		&models.WebServiceType{},
		&models.WebService{},
		&models.InternalService{},
	)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to auto migrate sys, %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to commit, %v", err)
	}

	return db, nil
}
