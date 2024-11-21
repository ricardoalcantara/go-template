package foo

import "github.com/ricardoalcantara/go-template/internal/models"

type FooRegister struct {
	Name string `json:"name" binding:"required"`
}

type FooUpdate struct {
	Name string `json:"key" binding:"required"`
}

type FooDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewFooDto(foo *models.Foo) FooDto {
	return FooDto{
		Id:   foo.ID.String(),
		Name: foo.Name,
	}
}
