package transactionviews

import (
	"net/http"
	"test_project/config"
	"test_project/models"

	"github.com/gin-gonic/gin"
)

func GetFilteredTransactionViews(c *gin.Context) {
	db := config.SetupDatabase()

	var transactions []models.TransactionViews

	id := c.DefaultQuery("id", "")
	customerName := c.DefaultQuery("customer_name", "")
	itemName := c.DefaultQuery("item_name", "")

	query := db
	if id != "" {
		query = query.Where("id = ?", id)
	}
	if customerName != "" {
		query = query.Where("customer_name = ?", customerName)
	}
	if itemName != "" {
		query = query.Where("item_name = ?", itemName)
	}

	if err := query.Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve filtered data from TransactionViews"})
		return
	}
	c.JSON(http.StatusOK, transactions)
}
