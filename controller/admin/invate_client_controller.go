package admin

import (
	"fmt"
	"strconv"

	"github.com/anyingiit/GoReactResourceManagement/dao"
	"github.com/anyingiit/GoReactResourceManagement/globalVars"
	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/anyingiit/GoReactResourceManagement/structs"
	"github.com/anyingiit/GoReactResourceManagement/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvateClientController struct {
	*BaseController
}

// NewClientController
func NewInvateClientController(db *gorm.DB) *InvateClientController {
	return &InvateClientController{NewBaseController(db)}
}

// 创建一个新的邀请码
// 1. 确保客户端ID存在
// 2. 确保该客户端未创建邀请码
// 3. 确保该客户端未被Session注册
func (i *InvateClientController) Post(c *gin.Context) {
	clientId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(fmt.Errorf("id type error")).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}
	clientIdUint := uint(clientId)

	// 确保客户端ID存在
	client, err := dao.FirstClient(i.Db, dao.ClientByID(clientIdUint))
	if err != nil {
		c.Error(fmt.Errorf("客户端不存在")).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	// 确保该客户端未创建邀请码
	_, err = dao.FirstInvateClient(i.Db, dao.InvateClientByClientID(clientIdUint))
	if err != nil && err != gorm.ErrRecordNotFound {
		c.Error(fmt.Errorf("检查邀请码是否存在失败")).SetType(gin.ErrorTypePrivate).SetMeta(400)
		return
	}
	if err == nil {
		c.Error(fmt.Errorf("该客户端已创建邀请码")).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	// 确保该客户端未被Session注册
	_, err = dao.FirstClientSession(i.Db, dao.InvateClientByClientID(clientIdUint))
	if err != nil && err != gorm.ErrRecordNotFound {
		c.Error(fmt.Errorf("检查Session是否存在失败")).SetType(gin.ErrorTypePrivate).SetMeta(400)
		return
	}
	if err == nil {
		c.Error(fmt.Errorf("该客户端已被Session注册")).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	projectConfig, err := globalVars.ProjectConfig.Get()

	if err != nil {
		c.Error(fmt.Errorf("获取项目配置失败")).SetType(gin.ErrorTypePrivate).SetMeta(400)
		return
	}

	invateClient := &models.InvateClient{
		ClientID: client.ID,
	}

	// 邀请码内包：
	// 1. 和服务器的通信方式
	// 2. 一个不含任何信息的随机代码
	err = invateClient.GenerateInvateCode(&structs.InvateClientMessage{
		ServerIP:   projectConfig.Server.PublicIp,
		ServerPort: projectConfig.Server.PublicPort,
		InvateCode: utils.BytesEncodingToHashSha256HexString([]byte(uuid.New().String())),
	})

	if err != nil {
		c.Error(fmt.Errorf("生成邀请码失败")).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	err = dao.CreateInvateClient(i.Db, invateClient)
	if err != nil {
		c.Error(fmt.Errorf("创建邀请码记录失败")).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	c.JSON(200, gin.H{
		"message": "create invate code success",
		"data": gin.H{
			"invate_code": invateClient.InvateCode,
		},
	})
}

// 获取邀请码
func (i *InvateClientController) Get(c *gin.Context) {
	clientId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(fmt.Errorf("id type error")).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}
	clientIdUint := uint(clientId)

	invateClient, err := dao.FirstInvateClient(i.Db, dao.InvateClientByClientID(clientIdUint))
	if err != nil {
		c.Error(fmt.Errorf("邀请码不存在")).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	c.JSON(200, gin.H{
		"message": "get invate code success",
		"data":    invateClient,
	})
}

// 删除邀请码
func (i *InvateClientController) Delete(c *gin.Context) {
	clientId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(fmt.Errorf("id type error")).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}
	clientIdUint := uint(clientId)

	invateClient, err := dao.FirstInvateClient(i.Db, dao.InvateClientByClientID(clientIdUint))
	if err != nil {
		c.Error(fmt.Errorf("邀请码不存在")).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	err = dao.DeleteInvateClient(i.Db, dao.InvateClientByClientID(invateClient.ClientID))
	if err != nil {
		c.Error(fmt.Errorf("删除邀请码失败")).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	c.JSON(200, gin.H{
		"message": "delete invate code success",
	})
}
