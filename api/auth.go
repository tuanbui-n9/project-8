package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (server *Server) Auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		server.AuthorizationError(c, errors.New("authorization header is required"))
		return
	}
	const bearerPrefix = "Bearer "

	if !strings.HasPrefix(tokenString, bearerPrefix) {
		server.AuthorizationError(c, errors.New("authorization header wrong format"))
		return
	}
	tokenString = strings.TrimPrefix(tokenString, bearerPrefix)

	authToken, err := server.firebaseAdmin.VerifyToken(tokenString)
	if err != nil {
		server.AuthorizationError(c, err)
		return
	}

	user, err := server.firebaseAdmin.GetUser(authToken.UID)
	if err != nil {
		server.AuthorizationError(c, err)
		return
	}

	customToken, err := server.firebaseAdmin.CreateCustomToken(user.UID, user.CustomClaims)
	if err != nil {
		server.AuthorizationError(c, err)
		return
	}

	c.SetCookie("sid", user.UID, 60*60*24*14, "/", server.config.AuthRootDomain, true, true)

	c.JSON(http.StatusOK, gin.H{
		"customClaims": user.CustomClaims,
		"customToken":  customToken,
		"uid":          user.UID,
	})
}

func (server *Server) CheckSession(c *gin.Context) {
	_, err := c.Request.Cookie("sid")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized1",
		})
		return
	}

	user, err := server.firebaseAdmin.GetUser("M4qvn8AOGRfoHYohRsFvg00ZHQQ2")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized3",
		})
		return
	}

	customToken, err := server.firebaseAdmin.CreateCustomToken(user.UID, user.CustomClaims)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized4",
		})
		return
	}

	c.SetCookie("sid", user.UID, 60*60*24*14, "/", server.config.AuthRootDomain, true, true)

	c.JSON(http.StatusOK, gin.H{
		"customClaims": user.CustomClaims,
		"uid":          user.UID,
		"customToken":  customToken,
	})
}

func (server *Server) ClearSession(c *gin.Context) {
	c.SetCookie("sid", "", -1, "/", server.config.AuthRootDomain, true, true)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
