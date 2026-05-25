package routes

import (
	"github.com/gin-gonic/gin"

	"log-service/handlers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{

		api.POST("/logs", handlers.CreateLogActivity)
		api.GET("/logs", handlers.GetLogActivities)
	}
}
