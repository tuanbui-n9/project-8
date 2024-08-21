package api

import (
	"net/http"
	"project8/cookies"
	db "project8/db/sqlc"
	"project8/firebaseadmin"
	"project8/utils"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config        utils.Config
	store         db.Store
	router        *gin.Engine
	cookies       cookies.Cookies
	firebaseAdmin *firebaseadmin.FirebaseAdmin
}

func NewServer(
	config utils.Config,
	store db.Store,
	cookies cookies.Cookies,
	firebaseAdmin *firebaseadmin.FirebaseAdmin,
) (*Server, error) {
	server := &Server{
		config:        config,
		store:         store,
		cookies:       cookies,
		firebaseAdmin: firebaseAdmin,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/health-check", server.HealthCheck)

	router.POST("/auth", server.Auth)
	router.POST("/check-session", server.CheckSession)
	router.POST("/clear-session", server.ClearSession)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) ErrorResponse(error string) gin.H {
	return gin.H{
		"error": error,
	}
}

func (server *Server) AuthorizationError(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": err.Error(),
	})
}
