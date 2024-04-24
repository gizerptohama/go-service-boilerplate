package database

import (
	logrus "boilerplate/internal/utils/logger"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	return
}

func CloseDB(db *gorm.DB) {
	log := logrus.Log
	dbInstance, err := db.DB()
	if err != nil {
		log.Fatalf("unable to get database instance: %v", err)
	}
	dbInstance.Close()
}
