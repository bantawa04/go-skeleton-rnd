package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

// CategoryRoutes struct
type CategoryRoutes struct {
	logger                    infrastructure.Logger
	router                    infrastructure.Router
	categoryController controllers.CategoryController
	middleware                middlewares.FirebaseAuthMiddleware
}

// NewCategoryRoutes creates new Category controller
func NewCategoryRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	categoryController controllers.CategoryController,
	middleware middlewares.FirebaseAuthMiddleware,
) CategoryRoutes {
	return CategoryRoutes{
		router:                    router,
		logger:                    logger,
		categoryController: categoryController,
		middleware:                middleware,
	}
}

// Setup category routes
func (c CategoryRoutes) Setup() {
	c.logger.Zap.Info(" Setting up Category routes")
	category := c.router.Gin.Group("/categories")
	{
		category.POST("", c.categoryController.CreateCategory)
		category.GET("", c.categoryController.GetAllCategory)
		category.GET("/:id", c.categoryController.GetOneCategory)
		category.PUT("/:id", c.categoryController.UpdateOneCategory)
		category.DELETE("/:id", c.categoryController.DeleteOneCategory)
	}
}
