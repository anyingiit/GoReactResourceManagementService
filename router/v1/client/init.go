package client

import (
	clientController "github.com/anyingiit/GoReactResourceManagement/controller/client"
	"github.com/anyingiit/GoReactResourceManagement/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitClient(routerGroup *gin.RouterGroup, db *gorm.DB) {
	taskQueueController := clientController.NewTaskQueueController(db)
	taskQueueResultController := clientController.NewTaskResultController(db)

	client := routerGroup.Group("/client", middleware.ProjectMustInitialized(), middleware.AuthClient())
	{
		// fmt.Println(admin)
		client.GET("/queues", taskQueueController.GetAll)
		client.POST("/queues/results", taskQueueResultController.Post)
	}

}
