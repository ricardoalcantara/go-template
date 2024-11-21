package foo

import (
	"github.com/google/uuid"
	"github.com/ricardoalcantara/go-template/internal/domain"
	"github.com/ricardoalcantara/go-template/internal/models"
)

type FooService struct {
}

func NewFooService() *FooService {
	return &FooService{}
}

func (s *FooService) List(userId uuid.UUID, pagination *models.Pagination) (*domain.ListView[FooDto], error) {
	result, err := models.ListFoo(userId, pagination)
	if err != nil {
		return nil, err
	}

	listView := domain.ListView[FooDto]{Page: pagination.Page}
	for _, foo := range result {
		listView.List = append(listView.List, NewFooDto(&foo))
	}

	return &listView, nil
}
func (s *FooService) Create(userId uuid.UUID, input FooRegister) (*FooDto, error) {
	return nil, nil
}
func (s *FooService) Get(fooId uuid.UUID, userId uuid.UUID) (*FooDto, error) {
	return nil, nil
}
func (s *FooService) Update(userId uuid.UUID, fooId uuid.UUID, input FooUpdate) (*FooDto, error) {
	return nil, nil
}
func (s *FooService) Delete(userId uuid.UUID, fooId uuid.UUID) error {
	return nil
}
