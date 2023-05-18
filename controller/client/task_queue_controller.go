package client

import (
	"github.com/anyingiit/GoReactResourceManagement/dao"
	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskQueueController struct {
	*BaseController
}

// NewTaskController
func NewTaskQueueController(db *gorm.DB) *TaskController {
	return &TaskController{
		BaseController: NewBaseController(db),
	}
}

// GetAll
// context中必须包含类型为*models.Client的client
func (t *TaskController) GetAll(c *gin.Context) {

	client, _ := c.MustGet("client").(*models.Client)
	var taskQueues []models.TaskQueue

	if err := dao.Find(t.Db, &taskQueues,
		dao.Preload("Task"),
		dao.Preload("Service"),
		dao.ByField("client_id", client.ID)); err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		c.Abort()
		return
	}

	webServiceIds := make([]uint, 0)
	for _, val := range taskQueues {
		if val.Service.IsOwnerType(models.OwnerTypeWebService) {
			webServiceIds = append(webServiceIds, val.Service.OwnerID)

		}
		// if val.Service.OwnerType == models.OwnerTypeWebService {
		// }
	}

	webServices, err := func() (result map[uint]models.WebService, err error) {
		var webServices []models.WebService
		err = dao.Find(t.Db, &webServices, dao.Preload("WebServiceType"), dao.IDsIn(webServiceIds))
		if err != nil {
			return nil, err
		}

		result = make(map[uint]models.WebService)
		for _, webService := range webServices {
			result[webService.ServiceID] = webService
		}

		return result, nil
	}()

	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		c.Abort()
		return
	}

	var data []gin.H
	for _, taskQueue := range taskQueues {
		data = append(data, gin.H{
			"queue_id":  taskQueue.ID,
			"task_name": taskQueue.Task.Name,
			"sequence":  taskQueue.Sequence,
			"info": func() *models.WebService {
				if webService, ok := webServices[taskQueue.ServiceID]; ok {
					return &webService
				}

				return nil
			}(),
		})
	}

	c.JSON(200, gin.H{
		"data": data,
	})
}
