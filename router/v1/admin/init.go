package admin

import (
	adminController "github.com/anyingiit/GoReactResourceManagement/controller/admin"
	"github.com/anyingiit/GoReactResourceManagement/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAdmin(routerGroup *gin.RouterGroup, db *gorm.DB) {
	clientController := adminController.NewClientController(db)
	inviteClientController := adminController.NewInvateClientController(db)

	admin := routerGroup.Group("/admin", middleware.ProjectMustInitialized(), middleware.AuthAdmin())
	{
		admin.POST("/clients", clientController.Post)
		admin.DELETE("/clients/:id", clientController.Delete)
		admin.PUT("/clients/:id", clientController.Put)
		admin.GET("/clients", clientController.GetAll)
		admin.GET("/clients/:id", clientController.GetById)

		admin.POST("/clients/:id/invitations", inviteClientController.Post)
	}

}
