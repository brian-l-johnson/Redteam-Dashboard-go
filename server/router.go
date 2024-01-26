package server

import (
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/controllers"
	docs "github.com/brian-l-johnson/Redteam-Dashboard-go/v2/docs"
	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/middleware"
	"github.com/gin-contrib/sessions"

	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func getAPIBaseURL() string {
	return os.Getenv("API_BASE_URL")
}
func isAdmin(roles string) bool {
	return strings.Contains(roles, "admin")
}

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
	router.GET("/auth/logout", auth.Logout)
	router.GET("/auth/users", middleware.Authorize("admin"), auth.ListUsers)
	router.PUT("/auth/users/:uid", middleware.Authorize("admin"), auth.UpdateUser)
	router.DELETE("/auth/user/:uid", middleware.Authorize("admin"), auth.DeleteUser)

	team := new(controllers.TeamController)
	router.GET("/teams", middleware.Authorize("viewer"), team.GetTeams)
	router.POST("/teams", middleware.Authorize("admin"), team.CreateTeam)
	router.DELETE("/teams/:tid", middleware.Authorize("admin"), team.DeleteTeam)

	jobs := new(controllers.JobController)
	router.GET("/jobs/manager", middleware.Authorize("viewer"), jobs.GetJobManagerState)
	router.GET("/jobs/:jobtype/next", middleware.Authorize("scanner"), jobs.NewJob)
	router.GET("/jobs", middleware.Authorize("any"), jobs.GetJobs)
	router.POST("/jobs/nmap/:jid", middleware.Authorize("scanner"), jobs.UploadScan)

	host := new(controllers.HostController)
	router.GET("/hosts/by-team/:tid", middleware.Authorize("viewer"), host.GetHostsByTeam)
	router.GET("/hosts/by-team/", middleware.Authorize("viewer"), host.GetAllHostsByTeam)

	docs.SwaggerInfo.BasePath = "/"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//static
	router.Static("/static", "./static")

	//html

	router.SetFuncMap(template.FuncMap{
		"getAPIBaseURL": getAPIBaseURL,
		"isAdmin":       isAdmin,
	})

	router.LoadHTMLGlob("templates/*")
	router.GET("/login.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{"title": "Login"})
	})
	router.GET("/main.html", middleware.AuthorizeHTML("any"), func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		roles := session.Get("roles")
		c.HTML(http.StatusOK, "main.html", gin.H{"user": user, "roles": roles, "title": "Main"})
	})
	router.GET("/register.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{"title": "Register"})
	})
	router.GET("/users.html", middleware.AuthorizeHTML("admin"), func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		roles := session.Get("roles")
		c.HTML(http.StatusOK, "users.html", gin.H{"user": user, "roles": roles, "title": "User Admin"})
	})

	router.GET("/logout.html", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.Redirect(http.StatusFound, "/login.html")
	})

	return router
}
