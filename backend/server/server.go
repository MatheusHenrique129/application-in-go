package server

import (
	"log"

	"github.com/MatheusHenrique129/application-in-go/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "2022",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Println("🚀 Server running on port " + s.port)
	log.Fatal(router.Run(":" + s.port))
}
