package public

import (
	"errors"
	"fmt"

	"github.com/anyingiit/GoReactResourceManagement/dao"
	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientSessionController struct {
	*BaseController
}

// NewClientController
func NewClientSessionController(db *gorm.DB) *ClientSessionController {
	return &ClientSessionController{NewBaseController(db)}
}

// 创建一个客户端会话
// 根据请求的InvateCode创建一个客户端会话
// 1. 确保InvateCode存在
// 2. 确保InvateCode未被使用
func (cs *ClientSessionController) Post(c *gin.Context) {
	var req struct {
		InvateCode string `form:"invate_code" binding:"required"`
	}
	if err := c.ShouldBindWith(&req, binding.FormMultipart); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	// 查找是否存在该邀请码
	invateClient := &models.InvateClient{}
	if err := dao.First(cs.Db, invateClient, dao.ByField("invate_code", req.InvateCode)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Error(fmt.Errorf("邀请码不存在")).SetType(gin.ErrorTypePublic).SetMeta(400)
		} else {
			c.Error(fmt.Errorf("查找邀请码失败")).SetType(gin.ErrorTypePublic).SetMeta(500)
		}
		fmt.Println(err)
		return
	}

	tx := cs.Db.Begin()

	// 1. 删除该邀请码
	// 2. 检查客户端会话是否存在
	// 3. 创建一个客户端会话

	// 删除邀请码
	if err := dao.Delete(tx, &models.InvateClient{}, dao.ByField("invate_code", req.InvateCode)); err != nil {
		tx.Rollback()
		c.Error(fmt.Errorf("删除邀请码失败")).SetType(gin.ErrorTypePublic).SetMeta(500)
		fmt.Println(err)
		return
	}

	// 检查客户端会话是否存在
	// 检查客户端会话是否已存在
	if exists, err := dao.Exists(tx, &models.ClientSession{}, dao.ByField("client_id", invateClient.ClientID)); err != nil {
		tx.Rollback()
		c.Error(fmt.Errorf("检查客户端会话失败")).SetType(gin.ErrorTypePublic).SetMeta(500)
		fmt.Println(err)
		return
	} else if exists {
		tx.Rollback()
		c.Error(fmt.Errorf("客户端会话已存在")).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	// 创建一个客户端会话
	clientSession := &models.ClientSession{
		ClientID: invateClient.ClientID,
		UUID:     uuid.New(),
	}

	// 创建客户端会话
	if err := dao.Create(tx, clientSession); err != nil {
		tx.Rollback()
		c.Error(fmt.Errorf("创建客户端会话失败")).SetType(gin.ErrorTypePublic).SetMeta(500)
		fmt.Println(err)
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.Error(fmt.Errorf("提交事务失败")).SetType(gin.ErrorTypePublic).SetMeta(500)
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"message": "create client session success",
		"data": gin.H{
			"client_session_uuid": clientSession.UUID,
		},
	})
}
