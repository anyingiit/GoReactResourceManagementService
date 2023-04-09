package middleware

import (
	"fmt"

	"github.com/anyingiit/GoReactResourceManagement/globalVars"
	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/gin-gonic/gin"
)

// 项目必须初始化
func ProjectMustInitialized() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := globalVars.Db.Get()
		if err != nil {
			c.Error(err).SetType(gin.ErrorTypePrivate).SetMeta(500)
			return
		}

		var count int64
		if err := db.Model(&models.Sys{}).Count(&count).Error; err != nil {
			c.Error(err).SetType(gin.ErrorTypePrivate).SetMeta(500)
			return
		}

		if count == 0 {
			c.Error(fmt.Errorf("project must initialized")).SetType(gin.ErrorTypePrivate).SetMeta(500)
			return
		}

		c.Next()
	}
}
