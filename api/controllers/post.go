package controllers

import (
	"boilerplate-api/api/services"
	"boilerplate-api/api/validators"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/paginations"
	"boilerplate-api/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PostController struct
type PostController struct {
	logger      infrastructure.Logger
	PostService services.PostService
	validator   validators.PostValidator
}

// NewPostController constructor
func NewPostController(
	logger infrastructure.Logger,
	PostService services.PostService,
	validator validators.PostValidator,
) PostController {
	return PostController{
		logger:      logger,
		PostService: PostService,
		validator:   validator,
	}
}

// CreatePost Create Post
// @Summary				Create Post
// @Description			Create Post
// @Param				JSON body models.Post{} true "Enter JSON"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce				application/json
// @Tags				Post
// @Success				200 {object} responses.Success "OK"
// @Failure      		400 {object} responses.Error
// @Router				/posts [post]
func (cc PostController) CreatePost(c *gin.Context) {
	post := models.Post{}

	if err := c.ShouldBindJSON(&post); err != nil {
		cc.logger.Zap.Error("Error [CreatePost] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind Post")
		responses.HandleError(c, err)
		return
	}

	if validationErr := cc.validator.Validate.Struct(post); validationErr != nil {
		err := errors.BadRequest.Wrap(validationErr, "Validation error")
		err = errors.SetCustomMessage(err, "Invalid input information")
		err = errors.AddErrorContextBlock(err, cc.validator.GenerateValidationResponse(validationErr))
		responses.HandleError(c, err)
		return
	}

	if _, err := cc.PostService.CreatePost(post); err != nil {
		cc.logger.Zap.Error("Error [CreatePost] [db CreatePost]: ", err.Error())
		err := errors.BadRequest.Wrap(err, "Failed To Create Post")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Post Created Successfully")
}

// GetAllPost Get All Post
// @Summary				Get all Post.
// @Param				page_size query string false "10"
// @Param				page query string false "Page no" "1"
// @Param				keyword query string false "search by name"
// @Description			Return all the Post
// @Produce				application/json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Tags				Post
// @Success 			200 {array} responses.DataCount{data=[]models.Post}
// @Failure      		500 {object} responses.Error
// @Router				/posts [get]
func (cc PostController) GetAllPost(c *gin.Context) {
	pagination := paginations.BuildPagination[*paginations.Pagination](c)
	pagination.Sort = "created_at desc"

	posts, count, err := cc.PostService.GetAllPost(*pagination)
	if err != nil {
		cc.logger.Zap.Error("Error finding Post records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed To Find Post")
		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, posts, count)
}

// GetOnePost Get One Post
// @Summary				Get one Post by id
// @Description			Get one Post by id
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce				application/json
// @Tags				Post
// @Success 			200 {object} responses.Data{data=models.Post}
// @Failure      		500 {object} responses.Error
// @Router				/posts/{id} [get]
func (cc PostController) GetOnePost(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	post, err := cc.PostService.GetOnePost(ID)
	if err != nil {
		cc.logger.Zap.Error("Error [GetOnePost] [db GetOnePost]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed To Find Post")
		responses.HandleError(c, err)
		return
	}

	responses.JSON(c, http.StatusOK, post)
}

// UpdateOnePost Update One Post By Id
// @Summary				Update One Post By Id
// @Description			Update One Post By Id
// @Param               Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param				data body models.Post true "Enter JSON"
// @Produce				application/json
// @Tags				Post
// @Success 			200 {object} responses.Success
// @Failure      		400 {object} responses.Error
// @Failure      		500 {object} responses.Error
// @Router				/posts/{id} [put]
func (cc PostController) UpdateOnePost(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	post := models.Post{}

	if err := c.ShouldBindJSON(&post); err != nil {
		cc.logger.Zap.Error("Error [UpdatePost] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "failed to update post")
		responses.HandleError(c, err)
		return
	}
	post.ID = ID

	if err := cc.PostService.UpdateOnePost(post); err != nil {
		cc.logger.Zap.Error("Error [UpdatePost] [db UpdatePost]: ", err.Error())
		err := errors.InternalError.Wrap(err, "failed to update post")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Post Updated Successfully")
}

// DeleteOnePost Delete One Post By Id
// @Summary				Delete One Post By Id
// @Description			Delete One Post By Id
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce				application/json
// @Tags				Post
// @Success 			200 {object} responses.Success
// @Failure      		500 {object} responses.Error
// @Router				/posts/{id} [delete]
func (cc PostController) DeleteOnePost(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := cc.PostService.DeleteOnePost(ID)
	if err != nil {
		cc.logger.Zap.Error("Error [DeleteOnePost] [db DeleteOnePost]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Delete Post")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Post Deleted Successfully")
}
