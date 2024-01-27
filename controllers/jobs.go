package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JobController struct{}

// get jobmanager state godoc
//
// @Summary get jobmanager state
// @Tags jobs
// @Accept json
// @Produces json
// @Success 200 {string} result
// @Router /jobs/manager [get]
func (j JobController) GetJobManagerState(c *gin.Context) {
	db := models.GetDB()
	var js []models.JobStatus
	results := db.Find(&js)
	if results.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": results.Error})
		return
	} else {
		c.IndentedJSON(http.StatusOK, js)
		return
	}
}

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
	db := models.GetDB()

	var js models.JobStatus
	result := db.Find(&js, "name=?", c.Param("jobtype"))
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusOK, gin.H{"status": "error", "message": "unknown job type"})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error})
			return
		}
	} else if js.Name == "" {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "error", "message": "unknown job type"})
		return
	} else {
		fmt.Printf("got JobStatus: %v\n", js)
		j, err := js.GetNextJob()
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{"status": "error", "message": "no available teams"})
			return
		} else {
			db.Create(&j)
			c.IndentedJSON(http.StatusOK, j)
			return
		}

	}
}

// get job list godoc
//
// @Summary get jobs
// @Tags jobs
// @Accept json
// @Produces json
// @Success 200 {string} result
// @Router /jobs [get]
func (j JobController) GetJobs(c *gin.Context) {
	db := models.GetDB()
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
//	@Tags			jobs
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

	db := models.GetDB()

	jid := c.Param("jid")
	var job models.Job
	jr := db.First(&job, "j_id=?", jid)
	if jr.Error != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "error", "message": "job not found"})
		return
	}

	var ipList []string
	var teamHosts []models.Host
	db.Find(&teamHosts, "team_id=?", job.TID)
	for _, h := range teamHosts {
		db.Delete(&h)
	}
	for i, host := range scan.Hosts {
		fmt.Printf("host %v name: %v\n", i, host.Hostname)
		ipList = append(ipList, host.IP)
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

	job.Status = "complete"
	db.Save(&job)

	c.IndentedJSON(http.StatusOK, gin.H{"status": "success"})
}
