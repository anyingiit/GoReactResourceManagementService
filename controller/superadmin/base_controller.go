package superadmin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anyingiit/GoReactResourceManagement/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseController struct {
	*controller.BaseController
}

// new PublicController
func NewBaseController(db *gorm.DB) *BaseController {
	return &BaseController{
		BaseController: controller.NewBaseController(db),
	}
}

// GetList
func (bc *BaseController) GetList(c *gin.Context) {
	var params struct {
		PaginationPage    int    `form:"pagination_page,default=1"`
		PaginationPerPage int    `form:"pagination_per_page,default=10"`
		SortField         string `form:"sort_field,default=id"`
		SortOrder         string `form:"sort_order,default=asc"`
		FilterJson        string `form:"filter_json,default={}"`
	}

	if err := c.BindQuery(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// fmt.Println(params)
	var filter map[string]interface{}
	if err := json.Unmarshal([]byte(params.FilterJson), &filter); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DEBUG start
	// params.PaginationPage = 1
	// params.PaginationPerPage = 10
	// params.SortField = "id"
	// params.SortOrder = "asc"
	// params.Filter = []byte(`{"key": "value"}`)
	//DEBUG end

	start := (params.PaginationPage - 1) * params.PaginationPerPage
	end := params.PaginationPage * params.PaginationPerPage

	data := make([]map[string]interface{}, 0)
	var total int64
	//Sql注入风险
	if err := bc.Db.Table(c.Param("resource")).
		Count(&total).
		Limit(params.PaginationPerPage).
		Offset(start).
		Order(params.SortField + " " + params.SortOrder).
		Where(filter).
		Find(&data).
		Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 对于cors请求，默认情况下仅准许以下标准的header字段：
	// 	Accept
	// 	Accept-Language
	// 	Content-Language
	// 	Last-Event-ID
	// 	Content-Type：只限于三个值application/x-www-form-urlencoded、multipart/form-data、text/plain
	// 	除此之外的header字段，如果要使用，需要在响应中添加Access-Control-Expose-Headers头部，才能读取。
	// 	c.Header("Access-Control-Expose-Headers", "X-Header1, X-Header2")
	c.Header("Access-Control-Expose-Headers", "Content-Range")
	c.Header("Content-Range", fmt.Sprintf("items %d-%d/%d", start, end, total))
	bc.SetContentRange(c, params.PaginationPage, params.PaginationPerPage, total)
	c.JSON(http.StatusOK, gin.H{"data": data})
}

// GetOne
func (bc *BaseController) GetOne(c *gin.Context) {
	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	resource := c.Param("resource")

	data := make(map[string]interface{})
	/*
		// doesn't work
		result := map[string]interface{}{}
		db.Table("users").First(&result)

		// works with Take
		result := map[string]interface{}{}
		db.Table("users").Take(&result)
	*/
	if err := bc.Db.Table(resource).
		Where("id = ?", c.Param("id")).
		Take(&data).
		Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// GetMany
func (bc *BaseController) GetMany(c *gin.Context) {
	var params struct {
		Ids []int `form:"ids"`
	}

	if err := c.BindQuery(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := make([]map[string]interface{}, 0)
	if err := bc.Db.Table(c.Param("resource")).
		Where("id IN ?", params.Ids).
		Find(&data).
		Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// Update
func (bc *BaseController) Update(c *gin.Context) {
	data := make(map[string]interface{})
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bc.Db.Table(c.Param("resource")).
		Where("id = ?", c.Param("id")).
		Updates(data).
		Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// UpdateMany
func (bc *BaseController) UpdateMany(c *gin.Context) {
	var params struct {
		Ids []int `form:"ids"`
	}

	if err := c.BindQuery(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := make([]map[string]interface{}, 0)
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bc.Db.Table(c.Param("resource")).
		Where("id IN ?", params.Ids).
		Updates(data).
		Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// Create
func (bc *BaseController) Create(c *gin.Context) {
	data := make(map[string]interface{})
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bc.Db.Table(c.Param("resource")).
		Create(data).
		Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// Delete
func (bc *BaseController) Delete(c *gin.Context) {
	if err := bc.Db.Table(c.Param("resource")).
		Where("id = ?", c.Param("id")).
		Delete(nil).
		Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": c.Param("id"),
	})
}

// DeleteMany
func (bc *BaseController) DeleteMany(c *gin.Context) {
	var params struct {
		Ids []int `form:"ids"`
	}

	if err := c.BindQuery(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bc.Db.Table(c.Param("resource")).
		Where("id IN ?", params.Ids).
		Delete(nil).
		Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": params.Ids,
	})
}
