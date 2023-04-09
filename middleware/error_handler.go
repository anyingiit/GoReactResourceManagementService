package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errors := c.Errors
		if len(errors) > 0 {
			for _, err := range errors {
				switch err.Type {
				case gin.ErrorTypePublic:
					// 对于公开错误，可以将错误信息返回给客户端
					c.JSON(err.Meta.(int), gin.H{"message": err.Error()})
				case gin.ErrorTypePrivate:
					// 对于私有错误，可以将错误信息记录到日志中
					log.Println(err)
					c.AbortWithStatus(http.StatusInternalServerError)
				default:
					// 对于未知错误，可以返回一个 500 状态码
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}
	}
}
