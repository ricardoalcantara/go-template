package foo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ricardoalcantara/go-template/internal/domain"
	"github.com/ricardoalcantara/go-template/internal/models"
	"github.com/ricardoalcantara/go-template/internal/utils"
)

func put(c *gin.Context) {
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

	foo, err := models.GetFoo(fooId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	addressUpdate := make(map[string]interface{})

	addressUpdate["Name"] = input.Name

	err = foo.Update(addressUpdate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.PrintError(err)})
		return
	}

	c.JSON(http.StatusCreated, FooView{
		Id:   foo.ID.String(),
		Name: foo.Name,
	})
}
