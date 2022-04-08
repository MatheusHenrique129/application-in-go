package migrations

import (
	"github.com/MatheusHenrique129/application-in-go/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Product{})

}
