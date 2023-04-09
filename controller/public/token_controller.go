package public

import (
	"errors"
	"fmt"

	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TokenController struct {
	*BaseController
}

// NewTokenController returns a new TokenController
func NewTokenController(db *gorm.DB) *TokenController {
	return &TokenController{
		BaseController: NewBaseController(db),
	}
}

// require RESTful api

// POST: 生成一个新的token
func (t *TokenController) Post(c *gin.Context) {
	// 获取用户名和密码并验证
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := models.User{Username: username}
	if t.Db.Where(&user).First(&user).Error != nil {
		c.Error(errors.New("invalid username or password")).SetType(gin.ErrorTypePublic).SetMeta(403)
		return
	}
	if !user.VaildPassword(password) {
		c.Error(errors.New("invalid username or password")).SetType(gin.ErrorTypePublic).SetMeta(403)
		return
	}

	if err := user.AccountError(); err != nil {
		if errors.Is(err, models.ErrUserMustChangePassword) {
			c.Error(errors.New("must change password")).SetType(gin.ErrorTypePublic).SetMeta(403)
			return
		} else {
			c.Error(errors.New("user is not valid")).SetType(gin.ErrorTypePublic).SetMeta(403)
			return
		}
	}

	// 生成token
	token, err := user.GenerateToken()
	if err != nil {
		fmt.Println(err) //DEBUG
		c.Error(errors.New("failed to generate token")).SetType(gin.ErrorTypePublic).SetMeta(500)
		return
	}

	// 返回token
	c.JSON(200, gin.H{
		"token": token,
	})
}
