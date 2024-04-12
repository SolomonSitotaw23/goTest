package transactions

import (
	"net/http"
	"test_project/config"
	"test_project/models"

	"github.com/gin-gonic/gin"
)

func GetTransaction(c *gin.Context) {
	db := config.SetupDatabase()
	id := c.Param("id")
	var transaction models.Transaction
	db.First(&transaction, id)
	c.JSON(http.StatusOK, transaction)
}
