package db

import (
	"fmt"

	"crudApplication/internal/pkg/config"
	models "crudApplication/internal/pkg/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(dbconf *config.DatabaseConfiguration) *gorm.DB {
	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		dbconf.Username,
		dbconf.Password,
		dbconf.Url,
		dbconf.Port,
		dbconf.DbName,
	)

	db, err := gorm.Open(mysql.Open(mysqlUrl), &gorm.Config{})

	if err != nil {
		fmt.Println("error in opening mysql database")
	}

	db.AutoMigrate(&models.Book{})

	return db
}
