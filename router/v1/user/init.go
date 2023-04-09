package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUser(routerGroup *gin.RouterGroup, db *gorm.DB) {
	// userController := userController.NewUserController(db)

	// user := routerGroup.Group("/user", middleware.AuthUser())
	// {
	// 	user.GET("", userController.Get)
	// 	user.POST("", userController.Post)
	// }
}
