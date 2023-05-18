package client

import (
	"fmt"

	"github.com/anyingiit/GoReactResourceManagement/dao"
	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskResultController struct {
	*BaseController
}

// NewTaskResultController
func NewTaskResultController(db *gorm.DB) *TaskResultController {
	return &TaskResultController{
		BaseController: NewBaseController(db),
	}
}

// Post: 提交任务结果
func (tr *TaskResultController) Post(c *gin.Context) {
	client, ok := c.MustGet("client").(*models.Client)
	if !ok {
		c.Error(fmt.Errorf("client not found")).SetType(gin.ErrorTypePrivate)
		c.Abort()
		return
	}

	var req struct {
		Data []struct {
			Success bool   `json:"success"`
			Detail  string `json:"detail"`
			QueueID uint   `json:"queue_id"`
		} `json:"data"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	// 检查提交的任务结果是否属于该客户端
	for _, val := range req.Data {
		if ok, err := dao.Exists(tr.Db, &models.TaskQueue{}, dao.ByField("id", val.QueueID), dao.ByField("client_id", client.ID)); err != nil || !ok {
			{
				// 返回没有权限提交任务结果
				c.Error(fmt.Errorf("no permission to submit task result")).SetType(gin.ErrorTypePublic).SetMeta(403)
				c.Abort()
				return
			}
		}
	}

	var taskResults []models.TaskQueueResult
	for _, val := range req.Data {
		taskResults = append(taskResults, models.TaskQueueResult{
			Succeed:     val.Success,
			Detail:      val.Detail,
			TaskQueueID: val.QueueID,
		})
	}

	tx := tr.Db.Begin()
	if err := dao.Create(tx, &taskResults); err != nil {
		tx.Rollback()
		c.Error(err).SetType(gin.ErrorTypePrivate)
		c.Abort()
		return
	}

	err := tx.Commit().Error
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}
