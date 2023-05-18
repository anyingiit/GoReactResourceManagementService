package client

import (
	"github.com/anyingiit/GoReactResourceManagement/controller"
	"gorm.io/gorm"
)

type BaseController struct {
	*controller.BaseController
}

// NewBaseController
func NewBaseController(db *gorm.DB) *BaseController {
	return &BaseController{
		BaseController: controller.NewBaseController(db),
	}
}
