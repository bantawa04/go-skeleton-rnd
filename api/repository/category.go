package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/paginations"
)

// CategoryRepository database structure
type CategoryRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

// NewCategoryRepository creates a new Category repository
func NewCategoryRepository(db infrastructure.Database, logger infrastructure.Logger) CategoryRepository {
	return CategoryRepository{
		db:     db,
		logger: logger,
	}
}

// Create Category
func (c CategoryRepository) Create(Category models.Category) (models.Category, error) {
	return Category, c.db.DB.Create(&Category).Error
}

// GetAllCategory Get All category
func (c CategoryRepository) GetAllCategory(pagination paginations.Pagination) ([]models.Category, int64, error) {
	var category []models.Category
	var totalRows int64 = 0
	queryBuilder := c.db.DB.Model(&models.Category{}).Offset(pagination.Offset).Order(pagination.Sort)

	if !pagination.All {
		queryBuilder = queryBuilder.Limit(pagination.PageSize)
	}

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where(c.db.DB.Where("`category`.`title` LIKE ?", searchQuery))
	}

	err := queryBuilder.
		Find(&category).
		Offset(-1).
		Limit(-1).
		Count(&totalRows).Error
	return category, totalRows, err
}

// GetOneCategory Get One Category By Id
func (c CategoryRepository) GetOneCategory(ID int64) (models.Category, error) {
	Category := models.Category{}
	return Category, c.db.DB.
		Where("id = ?", ID).First(&Category).Error
}

// UpdateOneCategory Update One Category By Id
func (c CategoryRepository) UpdateOneCategory(Category models.Category) error {
	return c.db.DB.Model(&models.Category{}).
		Where("id = ?", Category.ID).
		Updates(map[string]interface{}{
			"title": Category.Title,
		}).Error
}

// DeleteOneCategory Delete One Category By Id
func (c CategoryRepository) DeleteOneCategory(ID int64) error {
	return c.db.DB.
		Where("id = ?", ID).
		Delete(&models.Category{}).
		Error
}
