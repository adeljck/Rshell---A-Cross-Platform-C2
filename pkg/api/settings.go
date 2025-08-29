package api

import (
	"BackendTemplate/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListSettings(c *gin.Context) {
	var settings []database.Settings
	database.Engine.Find(&settings)
	c.JSON(http.StatusOK, gin.H{"status": 200, "data": settings})
}
func EditSettings(c *gin.Context) {
	var settings struct {
		Wecom string `json:"wecom"`
	}
	if err := c.BindJSON(&settings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	setting := database.Settings{
		Name:  "wecom",
		Value: settings.Wecom,
	}
	database.Engine.Where("name = ?", "wecom").Update(&setting)
	c.JSON(http.StatusOK, gin.H{"status": 200, "data": settings})
}
