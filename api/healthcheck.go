package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"time"
)

func (server *Server) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"date":   time.Now(),
	})
}
