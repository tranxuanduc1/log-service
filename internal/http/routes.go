package httpserver

import (
	"log-service/internal/http/handlers"
	"log-service/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(logActivityHandler *handlers.LogActivityHandler) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(middleware.ErrorHandler())
	router.Use(middleware.RequestID())

	api := router.Group("/api/v1")
	{
		logActivities := api.Group("/log-activities")
		{
			logActivities.POST("", logActivityHandler.Create)
			logActivities.GET("", logActivityHandler.FindAll)
		}
	}

	return router
}
