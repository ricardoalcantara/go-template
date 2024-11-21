package foo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ricardoalcantara/go-template/internal/domain"
	"github.com/ricardoalcantara/go-template/internal/middlewares"
	"github.com/ricardoalcantara/go-template/internal/models"
	"github.com/ricardoalcantara/go-template/internal/utils"
)

type FooController struct {
	FooService FooService
}

func NewFooController() *FooController {
	return &FooController{
		FooService: *NewFooService(),
	}
}

func (controller *FooController) list(c *gin.Context) {
	userId, err := models.GetUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}
	p := models.NewPagination(c)
	result, err := controller.FooService.List(userId, p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (controller *FooController) post(c *gin.Context) {
	var input FooRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		out := utils.GetValidationErrors(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: "Form Validation", Details: out})
		return
	}
	userId, err := models.GetUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	view, err := controller.FooService.Create(userId, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	c.JSON(http.StatusCreated, view)
}

func (controller *FooController) get(c *gin.Context) {
	fooId, err := uuid.Parse(c.Param("fooId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}
	userId, err := models.GetUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	view, err := controller.FooService.Get(fooId, userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	c.JSON(http.StatusOK, view)
}

func (controller *FooController) put(c *gin.Context) {
	var input FooUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		out := utils.GetValidationErrors(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: "Form Validation", Details: out})
		return
	}

	fooId, err := uuid.Parse(c.Param("fooId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	userId, err := models.GetUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	view, err := controller.FooService.Update(userId, fooId, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	c.JSON(http.StatusAccepted, view)
}

func (controller *FooController) delete(c *gin.Context) {
	fooId, err := uuid.Parse(c.Param("fooId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}
	userId, err := models.GetUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	err = controller.FooService.Delete(userId, fooId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	c.Status(http.StatusAccepted)
}

func RegisterRoutes(r *gin.Engine) {

	controller := NewFooController()

	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	protected.GET("/foo", controller.list)
	protected.POST("/foo", controller.post)
	protected.GET("/foo/:fooId", controller.get)
	protected.PUT("/foo/:fooId", controller.put)
	protected.DELETE("/foo/:fooId", controller.delete)
}
