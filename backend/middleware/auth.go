package middleware

import (
	"../controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthRequired
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user := controllers.GetUserFromContext(c); user != nil {
			c.Next()
			return
		}
		c.String(http.StatusForbidden, "Please Login First!!")
		c.Abort()
	}
}
