package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupDatabase() *gorm.DB {
	vars := LoadConfig()

	db, err := gorm.Open("postgres",
		"host="+vars.DBHOST+" port="+vars.DBPORT+" user="+vars.DBUSER+" dbname="+vars.DBNAME+" password="+vars.DBPASSWORD+" sslmode=disable")
	if err != nil {
		panic("faild to connect database " + err.Error())
	}
	return db

}
