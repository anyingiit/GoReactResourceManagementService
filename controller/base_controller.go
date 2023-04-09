package controller

import "gorm.io/gorm"

type BaseController struct {
	Db *gorm.DB
}

func NewBaseController(db *gorm.DB) *BaseController {
	return &BaseController{Db: db}
}
