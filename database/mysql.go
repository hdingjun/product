package database

import (
	"database/sql"
	"derekshop.com/product/config"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func InitDB() {
	var err error
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.AppConfig.UserName, config.AppConfig.Password, config.AppConfig.Addr, config.AppConfig.Port, config.AppConfig.Database)
	SqlDB, err = sql.Open("mysql", conn)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
