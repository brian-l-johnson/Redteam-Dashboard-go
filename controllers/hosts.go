package controllers

import (
	"net/http"

	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/models"
	"github.com/gin-gonic/gin"
)

type HostController struct{}

// Get Hosts By Team godoc
//
// @Summary Get hosts by team
// @Tags hosts
// @Accept json
// @Produces json
// @Param tid path	string	true "Team ID"
// @Success 200 {string} result
// @Router /hosts/by-team/{tid} [get]
func (h HostController) GetHostsByTeam(c *gin.Context) {
	db := models.GetDB()
	var hosts []models.Host
	results := db.Preload("Ports").Find(&hosts, "team_id=?", c.Param("tid"))
	if results.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": results.Error})
		return
	} else {
		c.IndentedJSON(http.StatusOK, hosts)
		return
	}
}

// Get Hosts By Team godoc
//
// @Summary Get all hosts by team
// @Tags hosts
// @Accept json
// @Produces json
// @Success 200 {string} result
// @Router /hosts/by-team/ [get]
func (h HostController) GetAllHostsByTeam(c *gin.Context) {
	db := models.GetDB()
	var teams []models.Team
	results := db.Find(&teams)
	if results.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": results.Error})
		return
	} else {
		for i, t := range teams {
			var hosts []models.Host
			results := db.Preload("Ports").Find(&hosts, "team_id=?", t.TID)
			if results.Error != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": results.Error})
				return
			} else {
				teams[i].Hosts = hosts
			}
		}
	}
	c.IndentedJSON(http.StatusOK, teams)
	return
}
