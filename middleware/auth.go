package middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/anyingiit/GoReactResourceManagement/dao"
	"github.com/anyingiit/GoReactResourceManagement/globalVars"
	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthBase(f func(user *models.User)) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if len(tokenString) == 0 {
			c.Error(fmt.Errorf("token is empty")).SetType(gin.ErrorTypePublic).SetMeta(500)
			c.Abort()
			return
		}

		user := &models.User{}
		err := user.ParseToken(tokenString)
		if err != nil {
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(500)
			c.Abort()
			return
		}

		// search user in db
		db, err := globalVars.Db.Get()

		if err != nil {
			c.Error(err).SetType(gin.ErrorTypePrivate)
			c.Abort()
			return
		}

		// search user
		// 此时出现错误，Token是有效的，但是ID是无效的
		// 只会出现在用户在数据库中被删除后，但是Token未过期的情况下

		if err := dao.First(db, user); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 无法找到用户
				c.Error(fmt.Errorf("用户已删除")).SetType(gin.ErrorTypePublic).SetMeta(400)
			} else {
				c.Error(err).SetType(gin.ErrorTypePrivate)
			}
			c.Abort()
			return
		}

		if user.AccountError() != nil {
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(500)
			c.Abort()
			return
		}

		f(user)
	}
}

func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthBase(func(user *models.User) {
			if user.Role.Name != "Admin" && user.Role.Name != "SuperAdmin" {
				c.Error(fmt.Errorf("permission denied")).SetType(gin.ErrorTypePublic).SetMeta(500)
				c.Abort()
				return
			}

			c.Next()
		})(c)
	}
}

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthBase(func(user *models.User) {
			if user.Role.Name != "Admin" && user.Role.Name != "SuperAdmin" && user.Role.Name != "User" {
				c.Error(fmt.Errorf("permission denied")).SetType(gin.ErrorTypePublic).SetMeta(500)
				c.Abort()
				return
			}

			c.Next()
		})(c)
	}
}

func AuthSuperAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthBase(func(user *models.User) {
			if user.Role.Name != "SuperAdmin" {
				c.Error(fmt.Errorf("permission denied")).SetType(gin.ErrorTypePublic).SetMeta(500)
				c.Abort()
				return
			}

			c.Next()
		})(c)
	}
}
