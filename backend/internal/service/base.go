package service

import (
	"github.com/MatheusHenrique129/application-in-go/internal/config"
	"github.com/MatheusHenrique129/application-in-go/internal/util"
)

type BaseService struct {
	config *config.Config
	logger *util.Logger
}

func NewBaseService(config *config.Config, logger *util.Logger) BaseService {
	return BaseService{
		config: config,
		logger: logger,
	}
}
