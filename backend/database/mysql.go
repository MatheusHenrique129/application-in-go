package database

import (
	"fmt"
	"log"
	"time"

	"github.com/MatheusHenrique129/application-in-go/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

var err error

func CreateDB(conf *configs.Config) *gorm.DB {
	dsn := getDsn(conf)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("A connection to the database could not be created. ", err)
	}

	config, _ := db.DB()

	config.SetMaxOpenConns(configs.DBMaxOpenConns)
	config.SetMaxIdleConns(configs.DBMaxIdleConns)
	config.SetConnMaxLifetime(time.Hour)

	return db
}

//getDsn Create a data source name for MySQL
func getDsn(conf *configs.Config) string {
	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.GetDatabaseUser(),
		conf.GetDatabasePass(),
		conf.GetDatabaseHost(),
		conf.GetDatabaseName())

	return dsn
}
