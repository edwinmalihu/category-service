package controller

import (
	"category-service/repository"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	AddCategory(*gin.Context)
}

type categoryController struct {
	custRepo repository.CategoryRepo
}

// AddCategory implements CategoryController.
func (categoryController) AddCategory(*gin.Context) {
	panic("unimplemented")
}

func CategoryNewController(repo repository.CategoryRepo) CategoryController {
	return categoryController{
		custRepo: repo,
	}
}
