package api

import (
	db "project9/db/sqlc"
	"project9/utils"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config utils.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/health-check", server.HealthCheck)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
