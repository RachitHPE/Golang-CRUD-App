package main

import (
	"crudApplication/internal/pkg/config"
	db "crudApplication/internal/pkg/database"
	"fmt"
)

func main() {
	dbConfig, err := config.GetDatabaseConfig()
	if err != nil {
		fmt.Println("Failed initialising database config")
	}

	db := db.Init(dbConfig)
	fmt.Println(db)
}
