package controller

import (
	"category-service/repository"
	"category-service/request"
	"category-service/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	AddCategory(*gin.Context)
}

type categoryController struct {
	custRepo repository.CategoryRepo
}

// AddCategory implements CategoryController.
func (c categoryController) AddCategory(ctx *gin.Context) {
	var req request.RequestAdd
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := c.custRepo.AddCategory(req)
	if err != nil {
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	res := response.ResponseSuccess{
		Category: data.Category,
		Msg:      "Data Behasil diTambahkan",
	}

	ctx.JSON(http.StatusOK, res)
}

func CategoryNewController(repo repository.CategoryRepo) CategoryController {
	return categoryController{
		custRepo: repo,
	}
}
