package customers

import (
	"test_project/config"
	"test_project/models"

	"github.com/gin-gonic/gin"
)

func NewCustomer(c *gin.Context) {

	db := config.SetupDatabase()

	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{"error": "request body is not valid" + err.Error()})
		return
	}

	db.Create(&customer)
	c.JSON(200, customer)
}
