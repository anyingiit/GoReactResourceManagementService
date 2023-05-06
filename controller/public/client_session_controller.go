package public

import (
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
	invateCode, err := dao.FirstInvateClient(cs.Db, dao.InvateClientByInvateCode(req.InvateCode))
	if err != nil {
		c.Error(fmt.Errorf("查找邀请码失败")).SetType(gin.ErrorTypePublic).SetMeta(500)
		fmt.Println(err)
		return
	}

	tx := cs.Db.Begin()

	// 1. 删除该邀请码
	// 2. 检查客户端会话是否存在
	// 3. 创建一个客户端会话

	// 删除该邀请码
	if err := dao.DeleteInvateClient(tx, dao.InvateClientByInvateCode(req.InvateCode)); err != nil {
		tx.Rollback()
		c.Error(fmt.Errorf("删除邀请码失败")).SetType(gin.ErrorTypePublic).SetMeta(500)
		fmt.Println(err)
		return
	}

	// 检查客户端会话是否存在
	if _, err := dao.FirstClientSession(tx, dao.ClientSessionByClientID(invateCode.ClientID)); err == nil {
		tx.Rollback()
		c.Error(fmt.Errorf("客户端会话已存在")).SetType(gin.ErrorTypePublic).SetMeta(500)
		fmt.Println(err)
		return
	}

	// 创建一个客户端会话
	clientSession := &models.ClientSession{
		ClientID: invateCode.ClientID,
		UUID:     uuid.New(),
	}

	if err := dao.CreateClientSession(tx, clientSession); err != nil {
		tx.Rollback()
		c.Error(fmt.Errorf("创建客户端会话失败")).SetType(gin.ErrorTypePublic).SetMeta(500)
		fmt.Println(err)
		return
	}

	err = tx.Commit().Error
	if err != nil {
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
