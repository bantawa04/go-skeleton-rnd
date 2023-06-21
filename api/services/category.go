package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
	"boilerplate-api/paginations"
)

// CategoryService struct
type CategoryService struct {
	repository repository.CategoryRepository
}

// NewCategoryService creates a new CategoryService
func NewCategoryService(repository repository.CategoryRepository) CategoryService {
	return CategoryService{
		repository: repository,
	}
}

// CreateCategory call to create the Category
func (c CategoryService) CreateCategory(category models.Category) (models.Category, error) {
	return c.repository.Create(category)
}

// GetAllCategory call to create the Category
func (c CategoryService) GetAllCategory(pagination paginations.Pagination) ([]models.Category, int64, error) {
	return c.repository.GetAllCategory(pagination)
}

// GetOneCategory Get One Category By Id
func (c CategoryService) GetOneCategory(ID int64) (models.Category, error) {
	return c.repository.GetOneCategory(ID)
}

// UpdateOneCategory Update One Category By Id
func (c CategoryService) UpdateOneCategory(category models.Category) error {
	return c.repository.UpdateOneCategory(category)
}

// DeleteOneCategory Delete One Category By Id
func (c CategoryService) DeleteOneCategory(ID int64) error {
	return c.repository.DeleteOneCategory(ID)

}
