package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func SetSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		token := uuid.NewV4().String()
		session.Set("patune_token", token)

		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save session"})
			c.Abort()
		}

		c.Set("patune_token", token)
	}
}

func CheckLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		if token := session.Get("patune_token"); token == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized user"})
			c.Abort()
		}
	}
}
