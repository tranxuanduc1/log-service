
package handlers
import (
	"net/http"

	"github.com/gin-gonic/gin"

	"log-service/config"
	"log-service/models"
)
// POST /api/v1/logs
func CreateLogActivity(c *gin.Context) {
	var input models.LogActivity

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "log created",
		"data":    input,
	})
}

// GET /api/v1/logs
func GetLogActivities(c *gin.Context) {
	var logs []models.LogActivity

	if err := config.DB.Order("created_at DESC").Find(&logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  logs,
		"total": len(logs),
	})
}
