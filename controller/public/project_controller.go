package public

import (
	"fmt"
	"net/http"

	"github.com/anyingiit/GoReactResourceManagement/dao"
	"github.com/anyingiit/GoReactResourceManagement/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProjectController struct {
	*BaseController
}

// new ProjectController
func NewProjectController(db *gorm.DB) *ProjectController {
	return &ProjectController{
		BaseController: NewBaseController(db),
	}
}

// reuqire RESTful api

// Post /project 创建一个新项目
// 用于项目初始化，初始化完成后将不可被调用
func (p *ProjectController) Post(c *gin.Context) {

	// 只有在无法获取sys表，既无法确定项目是否初始化时才会执行：
	// 	1. 报告页面无法找到
	// 	2. 打印错误信息
	// 	3. 返回页面未找到
	// 当确定项目已初始化后，直接返回页面未找到
	// 当确定项目未初始化，出现错误直接向页面返回即可

	_, err := dao.FirstSysRecord(p.Db)
	if err == nil {
		// 当确定项目已初始化，直接返回页面未找到
		c.Status(http.StatusNotFound)
		return
	}

	if err == gorm.ErrRecordNotFound { // 项目未初始化
		setupDataResult, err := db.SetupData(p.Db, c.PostForm("new_super_admin_password"))
		if err != nil {
			// 出现错误直接向页面返回即可
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "init project success",
			"data": gin.H{
				"super_admin_username": setupDataResult.User.Username,
			},
		})
		return
	} else {
		//发生了其他错误, 此时无法确定项目是否初始化
		c.Error(fmt.Errorf("failed to get sys count, %v", err)).SetType(gin.ErrorTypePrivate)
		return
	}
}
