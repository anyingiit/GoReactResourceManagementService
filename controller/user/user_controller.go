package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	*BaseController
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{NewBaseController(db)}
}

// require RESTful API

// GET /user
func (i *UserController) Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user1, user2, user3",
	})
}

// POST /user
func (i *UserController) Post(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "post",
	})
}

// PUT /user
func (i *UserController) Put(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "put",
	})
}

// DELETE /user
func (i *UserController) Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete",
	})
}
