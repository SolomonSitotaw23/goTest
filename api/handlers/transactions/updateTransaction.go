package transactions

import (
	"net/http"
	"test_project/config"
	"test_project/models"

	"github.com/gin-gonic/gin"
)

func UpdateTransaction(c *gin.Context) {
	db := config.SetupDatabase()
	id := c.Param("id")
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&transaction).Where("id = ?", id).Updates(transaction)
	c.JSON(http.StatusOK, transaction)
}
