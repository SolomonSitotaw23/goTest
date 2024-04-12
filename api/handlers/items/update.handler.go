package items

import (
	"net/http"
	"test_project/config"
	"test_project/models"

	"github.com/gin-gonic/gin"
)

func UpdateItem(c *gin.Context) {

	db := config.SetupDatabase()
	id := c.Param("id")
	var item models.Items
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&item).Where("id = ?", id).Updates(item)
	c.JSON(http.StatusOK, item)
}
