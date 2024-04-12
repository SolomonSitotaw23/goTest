package models

import (
	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	CustomerName string
	Balance      float64
}

type Transaction struct {
	gorm.Model
	CustomerID int
	ItemId     int
	Qty        int
	Price      float64
	Amount     float64
}
type Items struct {
	gorm.Model
	ItemName string
	Cost     float64
	Price    float64
	Sort     int
}
type TransactionViews struct {
	// gorm.Model
	CustomerId   int
	CustomerName string
	ItemId       int
	ItemName     string
	Qty          int
	Price        float64
	Amount       float64
}
