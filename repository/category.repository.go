package repository

import (
	"category-service/model"
	"category-service/request"
	"log"

	"gorm.io/gorm"
)

type CategoryRepo interface {
	Migrate() error
	AddCategory(request.RequestAdd) (model.Category, error)
}

type categoryRepo struct {
	DB *gorm.DB
}

// AddCategory implements CategoryRepo.
func (c categoryRepo) AddCategory(req request.RequestAdd) (data model.Category, err error) {
	data = model.Category{
		Category: req.Category,
	}

	return data, c.DB.Create(&data).Error
}

func NewCategoryRepo(db *gorm.DB) CategoryRepo {
	return categoryRepo{
		DB: db,
	}
}

func (c categoryRepo) Migrate() error {
	log.Print("[CategoryRepository]...Migrate")
	return c.DB.AutoMigrate(&model.Category{})
}
