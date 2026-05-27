package response

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, statusCode int, data any) {
	c.JSON(statusCode, gin.H{
		"data": data,
	})
}

func Error(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"error": message,
	})
}
