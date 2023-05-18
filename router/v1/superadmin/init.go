package superadmin

import (
	superadminController "github.com/anyingiit/GoReactResourceManagement/controller/superadmin"
	"github.com/anyingiit/GoReactResourceManagement/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitSuperAdmin(routerGroup *gin.RouterGroup, db *gorm.DB) {
	baseController := superadminController.NewBaseController(db)

	superadmin := routerGroup.Group("/superadmin", middleware.AuthSuperAdmin())
	{
		// resource := superadmin.Group("/:resource", middleware.AllowAllOPTIONS())
		resource := superadmin.Group("/:resource")
		{
			resource.GET("", baseController.HandleGet)
			resource.GET("/:id", baseController.GetOne)
			// resource.GET("/:id/:association", baseController.GetMany)
			resource.PUT("/:id", baseController.Update)
			resource.PUT("", baseController.UpdateMany)
			resource.POST("", baseController.Create)
			resource.DELETE("/:id", baseController.Delete)
			resource.DELETE("", baseController.DeleteMany)

		}
		// superadmin.GET("/:resource", baseController.GetList)
		// superadmin.GET("/:resource/:id", baseController.GetOne)
		// // superadmin.GET("/:resource", baseController.GetMany)
		// superadmin.PUT("/:resource/:id", baseController.Update)
		// superadmin.PUT("/:resource", baseController.UpdateMany)
		// superadmin.POST("/:resource", baseController.Create)
		// superadmin.DELETE("/:resource/:id", baseController.Delete)
		// superadmin.DELETE("/:resource", baseController.DeleteMany)
	}
}
