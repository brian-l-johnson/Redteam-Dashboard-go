package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/db"
	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/jobmanager"
	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JobController struct{}

// get new job godoc
//
// @Summary delete team
// @Tags jobs
// @Accept json
// @Produces json
// @Param jobtype	path	string	true	"Job Type"
// @Success	200	{string} result
// @Router /jobs/{jobtype}/next [get]
func (j JobController) NewJob(c *gin.Context) {
	jm := jobmanager.GetJobManager()
	jm.Test()
	db := db.GetDB()
	if !slices.Contains(jm.Types, c.Param("jobtype")) {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "error", "message": "unknown job type"})
		return
	} else {
		j, err := jm.GetNextJob()
		if err != nil {
			db.Create(&j)
			c.IndentedJSON(http.StatusOK, gin.H{"status": "error", "message": "no available teams"})
			return
		} else {
			c.IndentedJSON(http.StatusOK, j)
			return
		}

	}
}

// get job list godoc
//
// @Summary get jobs
// @Accept json
// @Produces json
// @Success 200 {string} result
// @Router /jobs [get]
func (j JobController) GetJobs(c *gin.Context) {
	db := db.GetDB()
	var jobs []models.Job
	result := db.Find(&jobs)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error})
		return
	} else {
		c.IndentedJSON(http.StatusOK, jobs)
	}

}

// nmap scan godoc
//
//	@Summary		nmap scan
//	@Description	upload new nmap scan
//	@Tags			scan
//	@Accept			json
//	@Produce		json
//	@Param			scan	body		models.Scan	true	"nmap scan data"
//	@Success		200		{string}	result
//	@Router			/jobs/nmap/{jid} [post]
func (n JobController) UploadScan(c *gin.Context) {
	var scan = new(models.Scan)
	err := c.BindJSON(&scan)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "unable to bind requesst"})
		return
	}

	db := db.GetDB()

	jid := c.Param("jid")
	var job models.Job
	jr := db.First(&job, "jid=?", jid)
	if jr.Error != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "error", "message": "job not found"})
		return
	}

	for i, host := range scan.Hosts {
		fmt.Printf("host %v name: %v\n", i, host.Hostname)
		h := new(models.Host)
		lookupresult := db.First(&h, "IP=?", host.IP)
		if lookupresult.Error != nil {
			if errors.Is(lookupresult.Error, gorm.ErrRecordNotFound) {
				host.TeamID = job.TID
				result := db.Save(&host)
				if result.Error != nil {
					c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error})
					return
				}
			} else {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": lookupresult.Error})
				return
			}
		} else {
			h.Ports = host.Ports
			db.Save(&h)
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": "success"})
}
