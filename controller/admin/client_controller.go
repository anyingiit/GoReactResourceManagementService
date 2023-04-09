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

	err := cc.Db.Create(&models.Client{
		Name:        req.Name,
		Description: req.Description,
	}).Error

	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		return
	}

	c.JSON(200, gin.H{
		"message": "create client success",
	})
}

// DELETE
func (cc *ClientController) Delete(c *gin.Context) {
	clientId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(fmt.Errorf("id type error")).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	// delete
	err = dao.DeleteClientById(cc.Db, uint(clientId))

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
		"message": "delete client success",
	})
}

// PUT
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

	client, err := dao.FirstClient(cc.Db, dao.ClientByID(uint(clientId)))
	if err != nil {
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

	err = dao.SaveClient(cc.Db, client)

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
	clients, err := dao.FindClients(cc.Db)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		return
	}

	c.JSON(200, gin.H{
		"message": "get all clients success",
		"data":    clients,
	})
}

func (cc *ClientController) GetById(c *gin.Context) {
	clientId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(fmt.Errorf("id type error")).SetType(gin.ErrorTypePublic).SetMeta(400)
		return
	}

	client, err := dao.FirstClient(cc.Db, dao.ClientByID(uint(clientId)))
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
