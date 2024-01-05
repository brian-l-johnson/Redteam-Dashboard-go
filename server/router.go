package server

import (
	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/controllers"
	docs "github.com/brian-l-johnson/Redteam-Dashboard-go/v2/docs"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	store := memstore.NewStore([]byte("badbadbad"))
	router.Use(sessions.Sessions("session", store))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	auth := new(controllers.AuthController)

	router.POST("/auth/login", auth.Login)
	router.GET("/auth/status", auth.Status)
	router.POST("/auth/register", auth.Register)

	docs.SwaggerInfo.BasePath = "/"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
