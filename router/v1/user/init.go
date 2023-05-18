package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	userController "github.com/anyingiit/GoReactResourceManagement/controller/user"
	"github.com/anyingiit/GoReactResourceManagement/middleware"
)

func InitUser(routerGroup *gin.RouterGroup, db *gorm.DB) {
	webServiceController := userController.NewWebServiceController(db)
	webServiceResultController := userController.NewWebServiceTestResult(db)
	userInfoController := userController.NewUserInfoController(db)

	public := routerGroup.Group("/user", middleware.AuthUser())
	{
		public.GET("/info", userInfoController.GetSelf)
		public.GET("/services/web_service", middleware.MustParsedStandardQuery(), webServiceController.GetAll)
		public.GET("/services/web_service/:id", middleware.MustParsedStandardQuery(), webServiceController.GetOne)
		public.GET("/services/web_service/results", middleware.MustParsedStandardQuery(), webServiceResultController.GetAll)
	}
}
