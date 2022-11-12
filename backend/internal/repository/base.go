package repository

import (
	"github.com/MatheusHenrique129/application-in-go/internal/config"
	"github.com/MatheusHenrique129/application-in-go/internal/util"
	"gorm.io/gorm"
)

type BaseRepository struct {
	config *config.Config
	db     *gorm.DB
	logger *util.Logger
}
