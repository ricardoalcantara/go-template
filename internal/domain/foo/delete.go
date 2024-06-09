package foo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ricardoalcantara/go-template/internal/domain"
	"github.com/ricardoalcantara/go-template/internal/models"
	"github.com/ricardoalcantara/go-template/internal/utils"
)

func delete(c *gin.Context) {
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

	err = foo.Delete()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	c.JSON(http.StatusAccepted, FooView{
		Id:   foo.ID.String(),
		Name: foo.Name,
	})
}
