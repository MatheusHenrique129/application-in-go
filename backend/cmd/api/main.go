package main

import (
	"github.com/MatheusHenrique129/application-in-go/configs"
	"github.com/MatheusHenrique129/application-in-go/database"
	"github.com/MatheusHenrique129/application-in-go/server"
)

func main() {
	conf := configs.NewConfig()

	database.CreateDB(conf)

	server := server.NewServer()

	server.Run()
}
