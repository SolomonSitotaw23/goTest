package customers

import (
	"net/http"
	"test_project/config"
	"test_project/models"

	"github.com/gin-gonic/gin"
)

func UpdateCustomer(c *gin.Context) {

	db := config.SetupDatabase()
	id := c.Param("id")
	var customer models.Customer

	// Check if the customer exists
	if err := db.Where("id = ?", id).First(&customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	// Bind the updated data from the request to the customer model
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the customer in the database
	db.Save(&customer)

	c.JSON(http.StatusOK, customer)
}
