package repository

import (
	"category-service/model"
	"log"

	"gorm.io/gorm"
)

type CategoryRepo interface {
	Migrate() error
}

type categoryRepo struct {
	DB *gorm.DB
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
