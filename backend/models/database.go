package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func OpenDatabase(dsn string, isDebug bool) (*gorm.DB, error) {
	logLevel := logger.Silent
	if isDebug {
		logLevel = logger.Info
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,    // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
}

func ConnectDatabase() {
	var err error
	// config.LoadConfig()

	// dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
	// 	config.AppConfig.DBHost,
	// 	config.AppConfig.DBPort,
	// 	config.AppConfig.DBUser,
	// 	config.AppConfig.DBName,
	// 	config.AppConfig.DBPassword,
	// )
	DB, err = OpenDatabase("./mydb.db", true)
	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Connected to database!")
	DB.AutoMigrate(&Product{})
	// DB.AutoMigrate(&MenuCategory{})
	// DB.AutoMigrate(&MenuDetail{})
}
