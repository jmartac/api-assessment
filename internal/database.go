package internal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	host     string
	port     string
	user     string
	password string
	dbname   string
)

func init() {
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")
}

// NewDatabase returns a new database connection
func NewDatabase() (*gorm.DB, error) {
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(mysqlInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Connected to database")
	return db, nil
}
