package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseController struct {
	Db *gorm.DB
}

func NewBaseController(db *gorm.DB) *BaseController {
	return &BaseController{Db: db}
}

// Set Content-Range
func (bc *BaseController) SetContentRange(c *gin.Context, page, perPage int, total int64) {
	c.Header("Access-Control-Expose-Headers", "Content-Range")

	start := (page-1)*perPage + 1
	end := page * perPage

	c.Header("Content-Range", fmt.Sprintf("items %d-%d/%d", start, end, total))
}

// ParseParamID
func (bc *BaseController) ParseParamID(c *gin.Context) (id uint, err error) {
	uint64ID, err := strconv.ParseUint(c.Param("id"), 10, 0)

	if err != nil {
		return 0, err
	}

	return uint(uint64ID), nil
}
