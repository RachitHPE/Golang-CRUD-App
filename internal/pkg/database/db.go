package db

import (
	"fmt"

	"crudApplication/internal/pkg/config"
	models "crudApplication/internal/pkg/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBHandlerResourceIntfc interface {
	Get() ([]models.Book, error)
	Create(book models.Book) (models.Book, error)
	GetByID(id string) (models.Book, error)
	Update(newBook models.Book) (models.Book, error)
	Delete(id string) error
}

type DBHandler struct {
	dbhandler *gorm.DB
}

// NewDBHandler implements DBHandler.
func NewDBHandler(dbhandler *gorm.DB) *DBHandler {
	return &DBHandler{
		dbhandler: dbhandler,
	}
}

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

func (handler DBHandler) Get() ([]models.Book, error) {
	var books []models.Book

	if result := handler.dbhandler.Find(&books); result.Error != nil {
		return nil, fmt.Errorf("error fetching records from database %w", result.Error)
	}

	return books, nil
}

func (handler DBHandler) Create(book models.Book) (models.Book, error) {
	if result := handler.dbhandler.Create(&book); result.Error != nil {
		return models.Book{}, fmt.Errorf("error inserting records in database %w", result.Error)
	}

	return book, nil
}

func (handler DBHandler) GetByID(id string) (models.Book, error) {
	var book models.Book

	if result := handler.dbhandler.First(&book, id); result.Error != nil {
		return models.Book{}, fmt.Errorf("failed fetching book by id: %s, Error: %w", id, result.Error)
	}

	return book, nil
}

func (handler DBHandler) Update(newBook models.Book) (models.Book, error) {
	var book models.Book

	if result := handler.dbhandler.First(&book, newBook.ID); result.Error != nil {
		return models.Book{}, fmt.Errorf("failed updating book id: %d, Error: %w", newBook.ID, result.Error)
	}

	book.Title = newBook.Title
	book.Author = newBook.Author
	book.Description = newBook.Description

	handler.dbhandler.Save(&book)

	return book, nil
}

func (handler DBHandler) Delete(id string) error {
	var book models.Book

	if result := handler.dbhandler.First(&book, id); result.Error != nil {
		return fmt.Errorf("error deleting book id: %s, Error: %w", id, result.Error)
	}

	handler.dbhandler.Delete(&book)

	return nil
}
