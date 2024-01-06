package models

import (
	"time"

	"gorm.io/gorm"
)

type Foo struct {
	ID        DbUUID `gorm:"primary_key;default:(UUID_TO_BIN(UUID()));"`
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

func GetFoo(id DbUUID) (*Foo, error) {
	var err error
	f := Foo{}
	err = DB.Take(&f, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &f, nil
}

func ListFoo(pagination *Pagination) ([]Foo, error) {
	var err error
	var f []Foo

	if pagination == nil {
		pagination = DefaultPagination()
	}
	err = DB.
		Scopes(pagination.GetScope).
		Find(&f).Error

	if err != nil {
		return nil, err
	}

	return f, nil
}
