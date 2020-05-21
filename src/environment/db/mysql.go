package db

import (
	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// CreateDBConnection returns db
func CreateDBConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(mysql-container:3306)/go_app")
	if err != nil {
		return db
	}

	db.LogMode(true)
	return db
}
