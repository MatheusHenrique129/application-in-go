package repository

import (
	"database/sql"

	"github.com/MatheusHenrique129/application-in-go/internal/config"
	"github.com/MatheusHenrique129/application-in-go/internal/util"
)

type BaseRepository struct {
	config *config.Config
	db     *sql.DB
	logger *util.Logger
}
