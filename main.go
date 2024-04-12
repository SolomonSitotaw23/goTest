package main

import (
	"fmt"
	transactionviews "test_project/api/handlers/transactionViews"
	"test_project/api/routes"
	"test_project/config"
	"test_project/models"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := transactionviews.CreateTransactionViews(); err != nil {
		fmt.Println(err)
	}

	db := config.SetupDatabase()
	defer db.Close()
	db.AutoMigrate(&models.Customer{}, &models.Items{}, &models.Transaction{})

	router := gin.Default()
	router = routes.SetupRoutes(router)
	router.Run(":8088")
}
