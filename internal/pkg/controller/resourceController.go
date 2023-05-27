package controller

import (
	db "crudApplication/internal/pkg/database"
	models "crudApplication/internal/pkg/model"
)

type ResourceControllerIntfc interface {
	GetResource() ([]models.Book, error)
	CreateResource(book models.Book) (models.Book, error)
	GetResourceByID(id string) (models.Book, error)
	UpdateResource(newBook models.Book) (models.Book, error)
	DeleteResource(id string) error
}

type ResourceController struct {
	dbResource db.DBHandlerResourceIntfc
}

func NewController(dbResource db.DBHandler) *ResourceController {
	return &ResourceController{
		dbResource: dbResource,
	}
}

func (rsrcController *ResourceController) GetResource() ([]models.Book, error) {
	books, err := rsrcController.dbResource.Get()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (rsrcController *ResourceController) CreateResource(book models.Book) (models.Book, error) {
	books, err := rsrcController.dbResource.Create(book)
	if err != nil {
		return models.Book{}, err
	}

	return books, nil
}

func (rsrcController *ResourceController) GetResourceByID(id string) (models.Book, error) {
	books, err := rsrcController.dbResource.GetByID(id)
	if err != nil {
		return models.Book{}, err
	}

	return books, nil
}

func (rsrcController *ResourceController) UpdateResource(newBook models.Book) (models.Book, error) {
	books, err := rsrcController.dbResource.Update(newBook)
	if err != nil {
		return models.Book{}, err
	}

	return books, nil

}

func (rsrcController *ResourceController) DeleteResource(id string) error {
	err := rsrcController.dbResource.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
