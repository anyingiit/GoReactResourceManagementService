package public

import (
	publicController "github.com/anyingiit/GoReactResourceManagement/controller/public"
	"github.com/anyingiit/GoReactResourceManagement/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitPublic(routerGroup *gin.RouterGroup, db *gorm.DB) {
	projectController := publicController.NewProjectController(db)
	tokenController := publicController.NewTokenController(db)
	clientSessionController := publicController.NewClientSessionController(db)

	public := routerGroup.Group("/public")
	{
		public.POST("/project", projectController.Post)
		public.POST("/token", middleware.ProjectMustInitialized(), tokenController.Post)

		public.POST("/client_session", middleware.ProjectMustInitialized(), clientSessionController.Post)
	}
}
