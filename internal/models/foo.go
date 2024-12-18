package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Foo struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"type:varchar(255);not null"`

	BarTypeId BarTypeId `gorm:"type:tinyint unsigned;not null"`
	BarType   *BarType
}

func (f *Foo) Save() error {
	err := DB.Create(&f).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *Foo) Update(updates map[string]interface{}) error {
	return DB.
		Model(&f).
		Updates(updates).Error
}

func (f *Foo) Delete() error {
	return DB.
		Delete(&f).
		Error
}

func GetFoo(id uuid.UUID) (*Foo, error) {
	var err error
	f := Foo{}
	err = DB.Take(&f, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &f, nil
}

func ListFoo(userId uuid.UUID, pagination *Pagination) ([]Foo, error) {
	var err error
	var f []Foo

	if pagination == nil {
		pagination = DefaultPagination()
	}
	err = DB.
		Scopes(pagination.GetScope).
		Find(&f, "user_id = ?", userId).Error

	if err != nil {
		return nil, err
	}

	return f, nil
}
