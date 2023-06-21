package controllers

import (
	"boilerplate-api/api/services"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/paginations"
	"boilerplate-api/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CategoryController struct
type CategoryController struct {
	logger          infrastructure.Logger
	CategoryService services.CategoryService
}

// NewCategoryController constructor
func NewCategoryController(
	logger infrastructure.Logger,
	CategoryService services.CategoryService,
) CategoryController {
	return CategoryController{
		logger:          logger,
		CategoryService: CategoryService,
	}
}

// CreateCategory Create Category
// @Summary				Create Category
// @Description			Create Category
// @Param				JSON body models.Category{} true "Enter JSON"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce				application/json
// @Tags				Category
// @Success				200 {object} responses.Success "OK"
// @Failure      		400 {object} responses.Error
// @Router				/categories [post]
func (cc CategoryController) CreateCategory(c *gin.Context) {
	category := models.Category{}

	if err := c.ShouldBindJSON(&category); err != nil {
		cc.logger.Zap.Error("Error [CreateCategory] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind Category")
		responses.HandleError(c, err)
		return
	}

	if _, err := cc.CategoryService.CreateCategory(category); err != nil {
		cc.logger.Zap.Error("Error [CreateCategory] [db CreateCategory]: ", err.Error())
		err := errors.BadRequest.Wrap(err, "Failed To Create Category")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Category Created Successfully")
}

// GetAllCategory Get All Category
// @Summary				Get all Category.
// @Param				page_size query string false "10"
// @Param				page query string false "Page no" "1"
// @Param				keyword query string false "search by name"
// @Description			Return all the Category
// @Produce				application/json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags				Category
// @Success 			200 {array} responses.DataCount{data=[]models.Category}
// @Failure      		500 {object} responses.Error
// @Router				/categories [get]
func (cc CategoryController) GetAllCategory(c *gin.Context) {
	pagination := paginations.BuildPagination[*paginations.Pagination](c)
	pagination.Sort = "created_at desc"

	category, count, err := cc.CategoryService.GetAllCategory(*pagination)
	if err != nil {
		cc.logger.Zap.Error("Error finding Category records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed To Find Category")
		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, category, count)
}

// GetOneCategory Get One Category
// @Summary				Get one Category by id
// @Description			Get one Category by id
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce				application/json
// @Tags				Category
// @Success 			200 {object} responses.Data{data=models.Category}
// @Failure      		500 {object} responses.Error
// @Router				/categories/{id} [get]
func (cc CategoryController) GetOneCategory(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	category, err := cc.CategoryService.GetOneCategory(ID)
	if err != nil {
		cc.logger.Zap.Error("Error [GetOneCategory] [db GetOneCategory]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed To Find Category")
		responses.HandleError(c, err)
		return
	}

	responses.JSON(c, http.StatusOK, category)
}

// UpdateOneCategory Update One Category By Id
// @Summary				Update One Category By Id
// @Description			Update One Category By Id
// @Param               Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param				data body models.Category true "Enter JSON"
// @Produce				application/json
// @Tags				Category
// @Success 			200 {object} responses.Success
// @Failure      		400 {object} responses.Error
// @Failure      		500 {object} responses.Error
// @Router				/categories/{id} [put]
func (cc CategoryController) UpdateOneCategory(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	category := models.Category{}

	if err := c.ShouldBindJSON(&category); err != nil {
		cc.logger.Zap.Error("Error [UpdateCategory] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "failed to update category")
		responses.HandleError(c, err)
		return
	}
	category.ID = ID

	if err := cc.CategoryService.UpdateOneCategory(category); err != nil {
		cc.logger.Zap.Error("Error [UpdateCategory] [db UpdateCategory]: ", err.Error())
		err := errors.InternalError.Wrap(err, "failed to update category")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Category Updated Successfully")
}

// DeleteOneCategory Delete One Category By Id
// @Summary				Delete One Category By Id
// @Description			Delete One Category By Id
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce				application/json
// @Tags				Category
// @Success 			200 {object} responses.Success
// @Failure      		500 {object} responses.Error
// @Router				/categories/{id} [delete]
func (cc CategoryController) DeleteOneCategory(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := cc.CategoryService.DeleteOneCategory(ID)
	if err != nil {
		cc.logger.Zap.Error("Error [DeleteOneCategory] [db DeleteOneCategory]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Delete Category")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Category Deleted Successfully")
}
