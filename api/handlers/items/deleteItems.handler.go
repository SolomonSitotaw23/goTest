package items

import (
	"net/http"
	"test_project/config"
	"test_project/models"

	"github.com/gin-gonic/gin"
)

func DeleteItem(c *gin.Context) {

	db := config.SetupDatabase()

	id := c.Param("id")
	var item models.Items
	// Implement logic to delete an item by ID from the database using GORM
	db.Where("id = ?", id).Delete(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}
