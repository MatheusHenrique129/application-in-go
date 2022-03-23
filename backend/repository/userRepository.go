package repository

import (
	"github.com/MatheusHenrique129/application-in-go/configs"
	"github.com/MatheusHenrique129/application-in-go/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	FindByName(name string) (*models.User, error)
}

type userDatabase struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	if configs.DB == nil {
		_, err := configs.Connect()
		if err != nil {
			log.Error(err)
		}
	}
	return &userDatabase{
		connection: configs.DB,
	}
}

func (db userDatabase) FindByName(name string) (*models.User, error) {
	var user = models.User{}
	result := db.connection.Preload(clause.Associations).Find(&user, "name = ?", name)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return &user, nil
	}
	return nil, nil
}
