package apihandler

import (
	"net/http"

	controller "crudApplication/internal/pkg/controller"
	models "crudApplication/internal/pkg/model"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	bookResource controller.ResourceControllerIntfc
}

// NewAPIHandler implements APIHandler.
func NewAPIHandler(bookResource *controller.ResourceController) *APIHandler {
	return &APIHandler{
		bookResource: bookResource,
	}
}

func (handler APIHandler) GetBooks(ginCtx *gin.Context) {
	result, err := handler.bookResource.GetResource()
	if err != nil {
		ginCtx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ginCtx.JSON(http.StatusOK, &result)
}

func (handler APIHandler) CreateBook(ginCtx *gin.Context) {
	var book models.Book
	if err := ginCtx.BindJSON(&book); err != nil {
		ginCtx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := handler.bookResource.CreateResource(book)
	if err != nil {
		ginCtx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ginCtx.JSON(http.StatusCreated, &result)
}

func (handler APIHandler) GetBookByID(ginCtx *gin.Context) {
	id := ginCtx.Param("id")

	result, err := handler.bookResource.GetResourceByID(id)
	if err != nil {
		ginCtx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ginCtx.JSON(http.StatusOK, &result)
}

func (handler APIHandler) UpdateBook(ginCtx *gin.Context) {
	id := ginCtx.Param("id")

	body := models.Book{}
	if err := ginCtx.BindJSON(&body); err != nil {
		ginCtx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := handler.bookResource.UpdateResource(id, body)
	if err != nil {
		ginCtx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ginCtx.JSON(http.StatusOK, &result)
}

func (handler APIHandler) DeleteBook(ginCtx *gin.Context) {
	id := ginCtx.Param("id")

	err := handler.bookResource.DeleteResource(id)
	if err != nil {
		ginCtx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ginCtx.Status(http.StatusNoContent)
}
