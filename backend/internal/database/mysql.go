package database

import (
	"fmt"
	"log"
	"time"

	"github.com/MatheusHenrique129/application-in-go/internal/config"
	"github.com/MatheusHenrique129/application-in-go/libraries/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func CreateDB(conf *config.Config) *gorm.DB {
	dsn := createDataSourceName(conf)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("A connection to the database could not be created.", err)
	}

	cfg, err := db.DB()

	if err != nil {
		log.Fatal("Could NOT create the database connection.", err)
	}

	cfg.SetMaxOpenConns(config.DBMaxOpenConns)
	cfg.SetMaxIdleConns(config.DBMaxIdleConns)
	cfg.SetConnMaxLifetime(time.Hour)

	logger.Info("Successful connection to MySQL DB")

	return db
}

//createDataSourceName Create a data source name for MySQL
func createDataSourceName(conf *config.Config) string {
	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.GetDatabaseUser(),
		conf.GetDatabasePass(),
		conf.GetDatabaseHost(),
		conf.GetDatabaseName())

	return dsn
}
