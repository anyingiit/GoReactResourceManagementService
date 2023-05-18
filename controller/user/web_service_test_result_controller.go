package user

import (
	"github.com/anyingiit/GoReactResourceManagement/dao"
	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/anyingiit/GoReactResourceManagement/structs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WebServiceTestResultController struct {
	*BaseController
}

func NewWebServiceTestResult(db *gorm.DB) *WebServiceTestResultController {
	return &WebServiceTestResultController{NewBaseController(db)}
}

// require RESTful API

// GET /user
func (i *WebServiceTestResultController) GetAll(c *gin.Context) {
	standeredQueryParsed, _ := c.MustGet("standeredQueryParsed").(*structs.StanderedQueryParsed)

	var webServices []models.WebService
	if err := dao.Find(i.Db, &webServices, dao.Preload("WebServiceType"), dao.Preload("Service")); err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		c.Abort()
		return
	}

	// 不同的webServiceType对应不同的数组
	webServiceTestResult, count, err := func() (result []gin.H, count int64, err error) {
		result = make([]gin.H, 0)

		webServiceServiceIds := func() []interface{} {
			result := make([]interface{}, 0)
			for _, webService := range webServices {
				result = append(result, webService.ServiceID)
			}

			return result
		}()

		// var serviceTaskQueueResults []models.TaskQueueResult
		// db := i.Db.Debug()
		// if err := dao.Find(db,
		// 	&serviceTaskQueueResults,
		// 	// dao.Preload("TaskQueue"),
		// 	dao.Joins("TaskQueue"),
		// 	dao.FieldIn("TaskQueue.service_id", webServiceServiceIds)); err != nil {
		// 	return nil, err
		// }

		//

		var serviceTaskQueueResults []models.TaskQueueResult
		if err := dao.StanderedQueryFind(i.Db, &serviceTaskQueueResults, &count, standeredQueryParsed,
			dao.Joins("TaskQueue"),
			dao.FieldIn("TaskQueue.service_id", webServiceServiceIds)); err != nil {

			c.Error(err).SetType(gin.ErrorTypePrivate)
			c.Abort()
			return nil, 0, err
		}
		//

		for _, webServiceTaskQueueResult := range serviceTaskQueueResults {
			result = append(result, gin.H{
				"id":         webServiceTaskQueueResult.ID,
				"service_id": webServiceTaskQueueResult.TaskQueue.ServiceID,
				"succeed":    webServiceTaskQueueResult.Succeed,
				"created_at": webServiceTaskQueueResult.CreatedAt,
			})
		}

		return result, count, nil
	}()
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		c.Abort()
		return
	}

	i.SetContentRange(c, standeredQueryParsed.PaginationPage, standeredQueryParsed.PaginationPerPage, count)
	c.JSON(200, gin.H{
		"data": webServiceTestResult,
	})
}
