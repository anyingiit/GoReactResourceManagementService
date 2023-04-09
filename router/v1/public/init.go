package public

import (
	publicController "github.com/anyingiit/GoReactResourceManagement/controller/public"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitPublic(routerGroup *gin.RouterGroup, db *gorm.DB) {
	projectController := publicController.NewProjectController(db)
	tokenController := publicController.NewTokenController(db)

	public := routerGroup.Group("/public")
	{
		public.POST("/project", projectController.Post)
		public.POST("/token", tokenController.Post)
	}
}
