package foo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ricardoalcantara/go-template/internal/domain"
	"github.com/ricardoalcantara/go-template/internal/models"
	"github.com/ricardoalcantara/go-template/internal/utils"
)

func post(c *gin.Context) {
	var input FooRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		out := utils.GetValidationErrors(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: "Form Validation", Details: out})
		return
	}

	foo := models.Foo{
		ID:   uuid.New(),
		Name: input.Name,
	}

	err := foo.Save()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.PrintError(err)})
		return
	}

	c.JSON(http.StatusCreated, FooView{
		Id:   foo.ID.String(),
		Name: foo.Name,
	})
}
