package v1

import (
	"github.com/anyingiit/GoReactResourceManagement/router/v1/admin"
	"github.com/anyingiit/GoReactResourceManagement/router/v1/client"
	"github.com/anyingiit/GoReactResourceManagement/router/v1/public"
	"github.com/anyingiit/GoReactResourceManagement/router/v1/superadmin"
	"github.com/anyingiit/GoReactResourceManagement/router/v1/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitV1(engin *gin.Engine, db *gorm.DB) {
	v1 := engin.Group("/v1")
	{
		public.InitPublic(v1, db)
		user.InitUser(v1, db)
		admin.InitAdmin(v1, db)
		superadmin.InitSuperAdmin(v1, db)
		client.InitClient(v1, db)
	}

}
