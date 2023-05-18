package admin

import (
	"github.com/anyingiit/GoReactResourceManagement/dao"
	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/anyingiit/GoReactResourceManagement/structs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	*BaseController
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{NewBaseController(db)}
}

func (u *UserController) makeResponse(data models.User) (result gin.H) {
	result = gin.H{
		"id":                   data.ID,
		"username":             data.Username,
		"age":                  data.Age,
		"role":                 data.Role.Name,
		"must_change_password": data.MustChangePassword,
	}

	return result
}

func (u *UserController) makeResponses(data []models.User) (result []gin.H) {
	for _, v := range data {
		result = append(result, u.makeResponse(v))
	}
	return result
}

// GetAll
func (u *UserController) GetAll(c *gin.Context) {
	standeredQueryParsed, _ := c.MustGet("standeredQueryParsed").(*structs.StanderedQueryParsed)

	var users []models.User
	var count int64
	if err := dao.StanderedQueryFind(u.Db, &users, &count, standeredQueryParsed); err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		c.Abort()
		return
	}

	u.SetContentRange(c, standeredQueryParsed.PaginationPage, standeredQueryParsed.PaginationPerPage, count)
	c.JSON(200, gin.H{
		"data": u.makeResponses(users),
	})
}

// GetOne

// Post
