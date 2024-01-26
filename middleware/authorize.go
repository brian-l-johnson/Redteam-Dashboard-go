package middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthorizeHTML(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user != nil {
			c.Set("user", user)
			rolestring := session.Get("roles").(string)
			c.Set("roles", rolestring)
			roles := strings.Split(rolestring, ",")
			if role == "any" || slices.Contains(roles, role) {
				c.Next()
				return
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		} else {
			c.Redirect(http.StatusFound, "http://127.0.0.1:8080/login.html")
			c.Abort()
			return
		}
	}
}

func Authorize(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user != nil {
			rolesstring := session.Get("roles").(string)
			roles := strings.Split(rolesstring, ",")
			if role == "any" || slices.Contains(roles, role) {
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
