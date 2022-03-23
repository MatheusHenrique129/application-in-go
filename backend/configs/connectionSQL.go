package configs

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

var err error

func Connect() (*gorm.DB, error) {
	//Connect to db using GORM
	host := config.Config.GetString("DB.ADDRESS")
	port := config.Config.GetString("DB.PORT")
	user := config.Config.GetString("DB.USERNAME")
	password := config.Config.GetString("DB.PASSWORD")
	dbName := config.Config.GetString("DB.DATABASE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)

	newLogger := logger.New(
		log.New(os.Stdout, "r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return DB, nil
}
