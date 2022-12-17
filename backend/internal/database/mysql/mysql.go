package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/MatheusHenrique129/application-in-go/internal/config"
	"github.com/MatheusHenrique129/application-in-go/libraries/logger"
	// Required by MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

func CreateDB(conf *config.Config) *sql.DB {
	dsn := createDataSourceName(conf)

	db, err := sql.Open(config.DriverName, dsn)
	if err != nil {
		log.Fatal("A connection to the database could not be created.", err)
	}

	db.SetMaxOpenConns(config.DBMaxOpenConns)
	db.SetMaxIdleConns(config.DBMaxIdleConns)
	db.SetConnMaxLifetime(time.Hour)

	if err := db.Ping(); err != nil {
		log.Fatal("Could NOT create the database connection.", err)
	}

	logger.Info("Successful connection to MySQL DB")
	return db
}

// createDataSourceName Create a data source name for MySQL
func createDataSourceName(conf *config.Config) string {
	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=true&loc=Local",
		conf.GetDatabaseUser(),
		conf.GetDatabasePass(),
		conf.GetDatabaseHost(),
		conf.GetDatabaseName())

	return dsn
}

//migrate -path ./migrations -database "mysql://root:Math@2109@tcp(localhost:3306)/bdgolang?charset=utf8mb4&parseTime=true&loc=Local"  -verbose up
