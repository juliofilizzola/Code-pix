package db

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/juliofilizzola/Code-pix/config/env"
	"github.com/juliofilizzola/Code-pix/domain/model"
	_ "github.com/lib/pq"
	_ "gorm.io/driver/sqlite"
)

func ConnectDB() (*gorm.DB, error) {
	var err error
	var db *gorm.DB

	db, err = gorm.Open(env.DbType, env.Dns)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil, err
	}

	if env.Debug == "true" {
		db.LogMode(true)
	}

	if env.AutoMigrateDb == "true" {
		db.AutoMigrate(&model.Bank{}, &model.Account{}, &model.PixKey{}, &model.Transaction{})
	}

	return db, nil
}
