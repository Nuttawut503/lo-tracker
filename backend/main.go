package main

import (
	"app/main/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConfig := config.GetDBConfig()

	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(db:3306)/%s", dbConfig.Password, dbConfig.Path))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
