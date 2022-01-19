package main

import (
	"derekshop.com/product/config"
	"derekshop.com/product/database"
	db "derekshop.com/product/database"
)

func main() {
	config.ReadConfig()
	database.InitDB()
	defer db.SqlDB.Close()
	router := InitRouter()
	router.Run(":8000")
}
