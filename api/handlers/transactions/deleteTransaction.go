package transactions

import (
	"net/http"
	"test_project/config"
	"test_project/models"

	"github.com/gin-gonic/gin"
)

func DeleteTransaction(c *gin.Context) {
	db := config.SetupDatabase()

	id := c.Param("id")
	var transaction models.Transaction
	db.Where("id = ?", id).Delete(&transaction)
	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}
