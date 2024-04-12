package items

import (
	"net/http"
	"test_project/config"
	"test_project/models"

	"github.com/gin-gonic/gin"
)

func GetItem(c *gin.Context) {

	db := config.SetupDatabase()
	id := c.Param("id")
	var item models.Items
	db.First(&item, id)
	c.JSON(http.StatusOK, item)
}
