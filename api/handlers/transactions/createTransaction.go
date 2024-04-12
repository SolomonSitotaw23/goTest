package transactions

import (
	"net/http"
	"test_project/config"
	"test_project/models"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	db := config.SetupDatabase()
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&transaction)
	c.JSON(http.StatusOK, transaction)
}
