package router

import (
	v1 "github.com/anyingiit/GoReactResourceManagement/router/v1"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(engin *gin.Engine, db *gorm.DB) {
	v1.InitV1(engin, db)
}
