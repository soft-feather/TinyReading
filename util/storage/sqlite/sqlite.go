package sqlite

import (
	"github.com/soft-feather/TinyReading/util/common/profile"
	"github.com/soft-feather/TinyReading/util/errors"
	"github.com/soft-feather/TinyReading/util/storage/sqlite/table"
	"github.com/wonderivan/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var (
	db *gorm.DB
)

func Init() error {
	return initDB()
}

type Product struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Price float64
}

func initDB() error {
	var err error
	db, err = gorm.Open(sqlite.Open(profile.GetEtcProfile().SqliteDBFileName), &gorm.Config{})

	if err != nil {
		logger.Fatal(errors.SqliteInitError)
		os.Exit(errors.SqliteInitError)
	}
	// Migrate the schema
	logger.Debug("start migrate tables ...")
	err = db.AutoMigrate(&table.Tag{}, &table.File{}, &table.Meta{})

	if err != nil {
		logger.Fatal(errors.SqliteMigrateError)
		os.Exit(errors.SqliteMigrateError)
	}
	// db.Create(&Product{Name: "Mobile", Price: 500.50})
	return nil
}

func GetDB() *gorm.DB {
	return db
}
