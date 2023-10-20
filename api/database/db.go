package database

import (
	"kosei-web/config"
	// mysql is driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DB struct {
	DB *gorm.DB
}

// Connect to the DATABASE
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)
	if err != nil {
		return nil, err
	}
	dbs := DB{
		DB: db,
	}
	return dbs.DB, nil
}
