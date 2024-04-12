package transactionviews

import (
	"net/http"
	"test_project/config"
	"test_project/models"

	"github.com/gin-gonic/gin"
)

func GetTransactionViews(c *gin.Context) {

	db := config.SetupDatabase()
	var transactions []models.TransactionViews
	if err := db.Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data from TransactionViews"})
		return
	}
	c.JSON(http.StatusOK, transactions)
}
