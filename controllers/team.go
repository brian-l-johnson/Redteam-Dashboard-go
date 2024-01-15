package controllers

import (
	"errors"
	"net/http"

	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/db"
	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TeamController struct{}

// get teams godoc
//
// @Summary Get Teams
// @Description	get all teams
// @Tags	teams
// @Accept	json
// @Produces	json
// @Success	200	{string} result
// @Router	/teams [get]
func (t TeamController) GetTeams(c *gin.Context) {
	db := db.GetDB()
	var teams []models.Team

	result := db.Find(&teams)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error})
		return
	}
	c.IndentedJSON(http.StatusOK, teams)
}

// create team godoc
//
// @Summary Create team
// @Description create a team
// @Tags teams
// @Accept	json
// @Produces	json
// @Param	create	body	models.Team	true "team data"
// @Success	200 {string} result
// @Router	/teams [post]
func (t TeamController) CreateTeam(c *gin.Context) {
	teamreq := new(models.Team)
	if err := c.BindJSON(&teamreq); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "unable to bind request to team object"})
		return
	}

	teamdb := new(models.Team)
	db := db.GetDB()
	result := db.First(&teamdb, "name=?", teamreq.Name)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			t := models.MakeTeam(teamreq.Name, teamreq.IPRange)

			cr := db.Create(&t)
			if cr.Error != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": cr.Error})
				return
			} else {
				c.IndentedJSON(http.StatusOK, gin.H{"status": "success", "message": "team created"})
				return
			}
		}
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "error", "message": "team already exists"})
		return
	}

}

// delete team godoc
//
// @Summary delete team
// @Tags teams
// @Accept json
// @Produces json
// @Param tid	path	string	true	"Team ID"
// @Success	200	{string} result
// @Router /teams/{tid} [delete]
func (t TeamController) DeleteTeam(c *gin.Context) {
	db := db.GetDB()
	var team models.Team
	result := db.First(&team, "ID=?", c.Param("tid"))
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusOK, gin.H{"status": "error", "message": "team not found"})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error})
			return
		}
	} else {
		result := db.Delete(&team)
		if result.Error != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error})
			return
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"status": "success", "message": "team deleted"})
			return
		}
	}

}
