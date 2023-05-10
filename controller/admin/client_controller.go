package admin

import (
	"fmt"
	"strconv"

	"github.com/anyingiit/GoReactResourceManagement/dao"
	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

// ClientController
type ClientController struct {
	*BaseController
}

// NewClientController
func NewClientController(db *gorm.DB) *ClientController {
	return &ClientController{NewBaseController(db)}
}

// Require RESTful api

// POST
// 参数：
//
//	@form{client_name: client name}
//	@form{client_description: client description}
func (cc *ClientController) Post(c *gin.Context) {
	var req struct {
		Name        string `form:"client_name" binding:"required"`
		Description string `form:"client_description" binding:"required"`
	}
	if err := c.ShouldBindWith(&req, binding.FormMultipart); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	// TODO: do some check

	if err := dao.Create(cc.Db, &models.Client{
		Name:        req.Name,
		Description: req.Description,
	}); err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		return
	}

	c.JSON(200, gin.H{
		"message": "create client success",
	})
}

// DELETE
// 参数：
//
//	@param{id: client id}
func (cc *ClientController) Delete(c *gin.Context) {
	clientId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(fmt.Errorf("id type error")).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	// delete
	if err := dao.Delete(cc.Db, &models.Client{}, dao.ByID(uint(clientId))); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(404)
		} else {
			c.Error(err).SetType(gin.ErrorTypePrivate)
		}
		return
	}

	c.JSON(200, gin.H{
		"message": "delete client success",
	})
}

// PUT
// 参数：
//
//	@param{id: client id}
func (cc *ClientController) Put(c *gin.Context) {
	clientId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(fmt.Errorf("id type error")).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	var req struct {
		Name        string `form:"client_name"`
		Description string `form:"client_description"`
	}
	if err := c.ShouldBindWith(&req, binding.FormMultipart); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	client := &models.Client{}
	if err := dao.First(cc.Db, client, dao.ByID(uint(clientId))); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(404)
			return
		} else {
			c.Error(err).SetType(gin.ErrorTypePrivate)
			return
		}
	}

	if req.Name != "" {
		client.Name = req.Name
	}
	if req.Description != "" {
		client.Description = req.Description
	}

	if err := dao.Update(cc.Db, client); err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		return
	}

	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		return
	}

	c.JSON(200, gin.H{
		"message": "update client success",
	})
}

// GET get all
func (cc *ClientController) GetAll(c *gin.Context) {
	clients := []models.Client{}
	err := dao.Find(cc.Db, &clients)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		return
	}

	c.JSON(200, gin.H{
		"message": "get all clients success",
		"data":    clients,
	})
}

// GET get by id
// 参数：
//
//	@param{id: client id}
func (cc *ClientController) GetById(c *gin.Context) {
	clientId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(fmt.Errorf("id type error")).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	client := &models.Client{}
	err = dao.First(cc.Db, client, dao.ByID(uint(clientId)))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(404)
			return
		} else {
			c.Error(err).SetType(gin.ErrorTypePrivate)
			return
		}
	}

	c.JSON(200, gin.H{
		"message": "get all clients success",
		"data":    client,
	})
}
