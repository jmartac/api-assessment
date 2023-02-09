package internal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// TODO Move this to environment variables
const (
	host     = "localhost"
	port     = 3306
	user     = "film"
	password = "film"
	dbname   = "films"
)

// NewDatabase returns a new database connection
func NewDatabase() (*gorm.DB, error) {
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(mysqlInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Connected to database")
	return db, nil
}
