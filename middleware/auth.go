package middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/anyingiit/GoReactResourceManagement/dao"
	"github.com/anyingiit/GoReactResourceManagement/globalVars"
	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AuthUserBase(f func(user *models.User)) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if len(tokenString) == 0 {
			c.Error(fmt.Errorf("token is empty")).SetType(gin.ErrorTypePublic).SetMeta(401)
			c.Abort()
			return
		}

		user := &models.User{}
		err := user.ParseToken(tokenString)
		if err != nil {
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(403)
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

		if err := dao.First(db, user, dao.Preload("Role")); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 无法找到用户
				c.Error(fmt.Errorf("用户已删除")).SetType(gin.ErrorTypePublic).SetMeta(403)
			} else {
				c.Error(err).SetType(gin.ErrorTypePrivate)
			}
			c.Abort()
			return
		}

		if user.AccountError() != nil {
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(403)
			c.Abort()
			return
		}

		c.Set("user", user)

		f(user)
	}
}

func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthUserBase(func(user *models.User) {
			if user.Role.Name != "Admin" && user.Role.Name != "SuperAdmin" {
				c.Error(fmt.Errorf("permission denied")).SetType(gin.ErrorTypePublic).SetMeta(403)
				c.Abort()
				return
			}

			c.Next()
		})(c)
	}
}

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthUserBase(func(user *models.User) {
			if user.Role.Name != "Admin" && user.Role.Name != "SuperAdmin" && user.Role.Name != "User" {
				c.Error(fmt.Errorf("permission denied")).SetType(gin.ErrorTypePublic).SetMeta(403)
				c.Abort()
				return
			}

			c.Next()
		})(c)
	}
}

func AuthSuperAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthUserBase(func(user *models.User) {
			if user.Role.Name != "SuperAdmin" {
				c.Error(fmt.Errorf("permission denied")).SetType(gin.ErrorTypePublic).SetMeta(403)
				c.Abort()
				return
			}

			c.Next()
		})(c)
	}
}

func AuthClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Client-UUID")
		if len(authHeader) == 0 {
			c.Error(fmt.Errorf("uuid is empty")).SetType(gin.ErrorTypePublic).SetMeta(401)
			c.Abort()
			return
		}
		uuid, err := uuid.Parse(authHeader)
		if err != nil {
			c.Error(fmt.Errorf("uuid type error")).SetType(gin.ErrorTypePublic).SetMeta(401)
			c.Abort()
			return
		}

		db, err := globalVars.Db.Get()
		if err != nil {
			c.Error(err).SetType(gin.ErrorTypePrivate)
			c.Abort()
			return
		}

		clientSession := &models.ClientSession{
			UUID: uuid,
		}
		err = dao.First(db, clientSession)
		if err != nil {
			c.Error(err).SetType(gin.ErrorTypePrivate)
			c.Abort()
			return
		}

		client := &models.Client{}
		err = dao.First(db, client, dao.ByField("id", clientSession.ClientID))
		if err != nil {
			c.Error(err).SetType(gin.ErrorTypePrivate)
			c.Abort()
			return
		}

		c.Set("client", client)

		c.Next()
	}
}
