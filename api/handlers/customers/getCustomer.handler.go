package customers

import (
	"net/http"
	"test_project/config"
	"test_project/models"

	"github.com/gin-gonic/gin"
)

func GetCustomer(c *gin.Context) {

	db := config.SetupDatabase()

	id := c.Param("id")
	var customer models.Customer

	// Retrieve the customer by ID from the database
	if err := db.Where("id = ?", id).First(&customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}
