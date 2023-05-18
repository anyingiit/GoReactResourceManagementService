package middleware

import (
	"encoding/json"

	"github.com/anyingiit/GoReactResourceManagement/structs"
	"github.com/gin-gonic/gin"
)

func MustParsedStandardQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params structs.StanderedQuery

		if err := c.ShouldBindQuery(&params); err != nil {
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(400)
			c.Abort()
			return
		}

		standeredQueryParsed := &structs.StanderedQueryParsed{
			PaginationPage:    params.PaginationPage,
			PaginationPerPage: params.PaginationPerPage,
			SortField:         params.SortField,
			SortOrder:         params.SortOrder,
			FilterJson:        map[string]interface{}{},
		}

		if err := json.Unmarshal([]byte(params.FilterJson), &standeredQueryParsed.FilterJson); err != nil {
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(400)
			c.Abort()
			return
		}

		c.Set("standeredQueryParsed", standeredQueryParsed)

		c.Next()
	}
}
