package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/db"
	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NmapController struct{}

// nmap scan godoc
//
//	@Summary		nmap scan
//	@Description	upload new nmap scan
//	@Tags			scan
//	@Accept			json
//	@Produce		json
//	@Param			scan	body		models.Scan	true	"nmap scan data"
//	@Success		200		{string}	result
//	@Router			/scan/nmap [post]
func (n NmapController) UploadScan(c *gin.Context) {
	var scan = new(models.Scan)
	err := c.BindJSON(&scan)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"satatu": "error", "message": "unable to bind requesst"})
		return
	}

	db := db.GetDB()

	for i, host := range scan.Hosts {
		fmt.Printf("host %v name: %v\n", i, host.Hostname)
		h := new(models.Host)
		lookupresult := db.First(&h, "IP=?", host.IP)
		if lookupresult.Error != nil {
			if errors.Is(lookupresult.Error, gorm.ErrRecordNotFound) {
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
