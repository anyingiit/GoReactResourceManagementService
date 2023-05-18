package user

import (
	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserInfoController struct {
	*BaseController
}

func NewUserInfoController(db *gorm.DB) *UserInfoController {
	return &UserInfoController{NewBaseController(db)}
}

// require RESTful API

// GET /user
func (i *UserInfoController) GetSelf(c *gin.Context) {
	user, _ := c.MustGet("user").(*models.User)

	c.JSON(200, gin.H{
		"data": gin.H{
			"username": user.Username,
			"name":     user.Name,
			"age":      user.Age,
			"role":     user.Role.Name,
		},
	})
}
