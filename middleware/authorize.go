package middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authorize(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user != nil {
			rolesstring := session.Get("roles").(string)
			roles := strings.Split(rolesstring, ",")
			if slices.Contains(roles, role) {
				c.Next()
			} else {
				c.IndentedJSON(http.StatusForbidden, gin.H{"message": "user does not have required role"})
				c.Abort()
				return
			}
		} else {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "not logged in"})
			c.Abort()
			return
		}

	}
}
