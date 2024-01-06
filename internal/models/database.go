package models

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	db_url := os.Getenv("DB_URL")
	var err error
	DB, err = gorm.Open(mysql.Open(db_url), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		logrus.Fatal("connection error:", err)
	} else {
		logrus.Debug("Db Connected")
	}

	if value, ok := os.LookupEnv("AUTO_MIGRATE"); ok && value == "true" {
		migrate()
		createTypes()
	}
}

func migrate() {
	DB.AutoMigrate(&BarType{})
	DB.AutoMigrate(&Foo{})
}

func createTypes() {
	if err := DB.Take(&BarType{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		var err error

		err = DB.Transaction(func(tx *gorm.DB) error {
			for _, vt := range []BarType{
				{ID: BarTypeGeneric, TableType: TableType{Name: "BarTypeGeneric", Description: "Generic Bar"}},
				{ID: BarTypeAbstract, TableType: TableType{Name: "BarTypeAbstract", Description: "Abstract Bar"}},
			} {
				if err = DB.Create(&vt).Error; err != nil {
					return err
				}
			}

			return nil
		})

		if err != nil {
			panic(err)
		}
	}
}
