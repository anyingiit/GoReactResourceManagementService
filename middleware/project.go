package middleware

import (
	"fmt"

	"github.com/anyingiit/GoReactResourceManagement/dao"
	"github.com/anyingiit/GoReactResourceManagement/globalVars"
	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 项目必须初始化
func ProjectMustInitialized() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := globalVars.Db.Get()
		if err != nil {
			c.Error(err).SetType(gin.ErrorTypePrivate).SetMeta(500)
			return
		}

		if err := dao.First(db, &models.Sys{}); err != nil {
			if err == gorm.ErrRecordNotFound {
				// 告知项目未初始化
				c.Error(fmt.Errorf("project must initialized")).SetType(gin.ErrorTypePublic).SetMeta(500)
			} else {
				c.Error(err).SetType(gin.ErrorTypePrivate).SetMeta(500)
			}
			c.Abort()
			return
		}

		c.Next()
	}
}
