package main

import (
	"time"

	"github.com/MatheusHenrique129/application-in-go/internal/app"
	"github.com/MatheusHenrique129/application-in-go/internal/config"
	"github.com/MatheusHenrique129/application-in-go/internal/database"
	"github.com/MatheusHenrique129/application-in-go/internal/util"
	"github.com/MatheusHenrique129/application-in-go/libraries/logger"
)

func main() {
	defaultTimeHelper := util.NewTimeHelper()
	startTime := defaultTimeHelper.GetCurrentUtcDateTime()

	logger.Info("Initializing config.")

	conf := config.NewConfig()
	logger.Infof("Finished config. %d ms", time.Since(startTime).Milliseconds())

	startTime = defaultTimeHelper.GetCurrentUtcDateTime()
	logger.Info("Initializing create DB.")

	database.CreateDB(conf)
	logger.Infof("Finished create DB. %d ms", time.Since(startTime).Milliseconds())

	startTime = defaultTimeHelper.GetCurrentUtcDateTime()
	logger.Info("Initializing new app.")

	application := app.NewApplication(conf)
	logger.Infof("Finished create new app. %d ms", time.Since(startTime).Milliseconds())

	err := application.Run()
	if err != nil {
		logger.Panic("Error running the app.", err)
		panic(err)
	}
}
