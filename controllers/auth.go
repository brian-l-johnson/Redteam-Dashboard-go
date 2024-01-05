package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/db"
	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

var lr = new(models.LoginReq)
var rr = new(models.RegisterReq)

// Login godoc
//
//	@Summary		Login
//	@Description	Login a user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			login	body		models.LoginReq	true	"Login Data"
//	@Success		200		{string}	result
//	@Router			/auth/login [post]
func (a AuthController) Login(c *gin.Context) {

	if err := c.BindJSON(&lr); err != nil {
		return
	}
	if lr.Password != "abc123" {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "login failure"})
		return
	}
	session := sessions.Default(c)
	session.Set("user", lr.User)
	session.Save()
	c.IndentedJSON(http.StatusOK, gin.H{"message": "login success"})
}

// status godoc
//
//	@Summary		Auth Status
//	@Description	Check login status
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200		{string}	result
//	@Router			/auth/status [get]
func (a AuthController) Status(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")

	if user == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "not logged in"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "logged in", "user": user})
}

// register godoc
//
// @Summary		Register User
// @Description	Register a user
// @Tags		user
// @Accept		json
// @Produces	json
// @Param		register	body		models.RegisterReq	true	"Login Data"
// @Success		200	{string} result
// @Router		/auth/register [post]
func (a AuthController) Register(c *gin.Context) {
	regreq := new(models.RegisterReq)
	db := db.GetDB()

	if err := c.BindJSON(&regreq); err != nil {
		return
	}
	var user models.User
	result := db.First(&user, "name=?", regreq.Name)

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println(("some other errer"))
			fmt.Println(result.Error)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"state": "error", "message": "db error"})
			return
		}
	}
	if result.RowsAffected != 0 {
		fmt.Println("user already exits")
		c.IndentedJSON(http.StatusOK, gin.H{"state": "error", "message": "user already exists"})
		return
	}

	fmt.Println("checked if user exists")
	user.Name = regreq.Name
	bytes, hasherr := bcrypt.GenerateFromPassword([]byte(regreq.Password), 14)
	if hasherr != nil {
		panic("error hashing password")
	}
	user.PasswordHash = string(bytes)
	user.Active = false
	result = db.Create(&user)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "error": result.Error})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "user created"})
}
