package routes

import (
	"test_project/api/handlers/customers"
	"test_project/api/handlers/items"
	transactionviews "test_project/api/handlers/transactionViews"
	"test_project/api/handlers/transactions"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) *gin.Engine {
	customersGroup := router.Group("/customers")
	{

		customersGroup.POST("", customers.NewCustomer)
		customersGroup.DELETE("/:id", customers.DeleteCustomer)
		customersGroup.PUT("/:id", customers.UpdateCustomer)
		customersGroup.GET("/:id", customers.GetCustomer)
	}

	transactionsGroup := router.Group("/transactions")
	{
		transactionsGroup.POST("", transactions.CreateTransaction)
		transactionsGroup.GET("/:id", transactions.GetTransaction)
		transactionsGroup.DELETE("/:id", transactions.DeleteTransaction)
		transactionsGroup.PUT("/:id", transactions.UpdateTransaction)
		transactionsGroup.GET("/filtered", transactionviews.GetFilteredTransactionViews)
	}

	itemsGroup := router.Group("/items")
	{
		itemsGroup.POST("", items.CreateItem)
		itemsGroup.GET("/:id", items.GetItem)
		itemsGroup.DELETE("/:id", items.DeleteItem)
		itemsGroup.PUT("/:id", items.UpdateItem)
	}

	router.GET("/transactionViews", transactionviews.GetTransactionViews)

	return router
}
