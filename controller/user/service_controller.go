package user

import (
	"github.com/anyingiit/GoReactResourceManagement/dao"
	"github.com/anyingiit/GoReactResourceManagement/models"
	"github.com/anyingiit/GoReactResourceManagement/structs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WebServiceController struct {
	*BaseController
}

func (i *WebServiceController) makeResponse(data models.WebService) (result gin.H) {
	result = gin.H{
		"host": data.Host,
		"port": data.Port,
	}
	if data.Service != nil {
		result["id"] = data.ID
		result["service_id"] = data.Service.ID
		result["name"] = data.Service.Name
		result["description"] = data.Service.Description
	}
	if data.WebServiceType != nil {
		result["web_service_type"] = data.WebServiceType.Name
		result["web_service_protocol"] = data.WebServiceType.Protocol
	}

	return result
}

func (i *WebServiceController) makeResponses(datas []models.WebService) (result []gin.H) {
	result = make([]gin.H, 0)
	for _, webService := range datas {
		result = append(result, i.makeResponse(webService))
	}
	return result
}

func NewWebServiceController(db *gorm.DB) *WebServiceController {
	return &WebServiceController{NewBaseController(db)}
}

func (i *WebServiceController) GetAll(c *gin.Context) {
	// user, _ := c.MustGet("user").(*models.User)
	standeredQueryParsed, _ := c.MustGet("standeredQueryParsed").(*structs.StanderedQueryParsed)

	var webServices []models.WebService
	var count int64
	if err := dao.StanderedQueryFind(i.Db, &webServices, &count, standeredQueryParsed,
		dao.Preload("WebServiceType"),
		dao.Preload("Service")); err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		c.Abort()
		return
	}
	// if err := dao.Find(i.Db, &webServices,
	// 	dao.Preload("WebServiceType"),
	// 	dao.Preload("Service"),
	// 	dao.ByFields(standeredQueryParsed.FilterJson),
	// 	dao.ByRecordRangeWithPage(standeredQueryParsed.PaginationPage, standeredQueryParsed.PaginationPerPage),
	// 	dao.CountScope(&count),
	// 	dao.Order(standeredQueryParsed.SortField, standeredQueryParsed.SortOrder)); err != nil {
	// 	c.Error(err).SetType(gin.ErrorTypePrivate)
	// 	c.Abort()
	// 	return
	// }

	// webServicedata := func() (result []gin.H) {
	// 	result = make([]gin.H, 0)
	// 	for _, webService := range webServices {
	// 		item := gin.H{
	// 			"host": webService.Host,
	// 			"port": webService.Port,
	// 		}
	// 		if webService.Service != nil {
	// 			item["id"] = webService.Service.ID
	// 			item["name"] = webService.Service.Name
	// 			item["description"] = webService.Service.Description
	// 		}
	// 		if webService.WebServiceType != nil {
	// 			item["web_service_type"] = webService.WebServiceType.Name
	// 			item["web_service_protocol"] = webService.WebServiceType.Protocol
	// 		}

	// 		result = append(result, item)
	// 	}

	// 	return result
	// }()

	i.SetContentRange(c, standeredQueryParsed.PaginationPage, standeredQueryParsed.PaginationPerPage, count)
	c.JSON(200, gin.H{
		"data": i.makeResponses(webServices),
	})
}

func (i *WebServiceController) GetOne(c *gin.Context) {
	standeredQueryParsed, _ := c.MustGet("standeredQueryParsed").(*structs.StanderedQueryParsed)

	id, err := i.ParseParamID(c)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		c.Abort()
		return
	}

	var webService models.WebService
	if err := dao.StanderedQueryFirst(i.Db, &webService, standeredQueryParsed,
		dao.Preload("WebServiceType"),
		dao.Preload("Service"),
		dao.ByID(id)); err != nil {
		c.Error(err).SetType(gin.ErrorTypePrivate)
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"data": i.makeResponse(webService),
	})
}
