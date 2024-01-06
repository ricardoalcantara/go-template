package foo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/go-template/internal/domain"
	"github.com/ricardoalcantara/go-template/internal/models"
	"github.com/ricardoalcantara/go-template/internal/utils"
	"github.com/samber/lo"
)

func list(c *gin.Context) {
	p := models.NewPagination(c)
	sshKeys, err := models.ListFoo(p)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: utils.PrintError(err)})
		return
	}

	result := lo.Map(sshKeys, func(p models.Foo, _index int) FooView {
		return FooView{
			Id:   p.ID.String(),
			Name: p.Name,
		}
	})

	c.JSON(http.StatusOK, domain.ListView[FooView]{List: result, Page: p.Page})
}
