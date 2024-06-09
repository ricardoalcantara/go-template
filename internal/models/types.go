package models

type TableType struct {
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(255);not null"`
}

type BarTypeId uint8

const (
	BarTypeGeneric  BarTypeId = 1
	BarTypeAbstract BarTypeId = 2
)

type BarType struct {
	ID BarTypeId `gorm:"type:smallint;primary_key"`
	TableType
}
