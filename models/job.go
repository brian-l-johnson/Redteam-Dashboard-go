package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model `json:"-"`
	JID        string
	Type       string
	IPRange    string
	Status     string
	Scanner    string
	TID        uint
}

func MakeJob(jobtype string, iprange string, tid uint) Job {
	var job Job
	job.Type = jobtype
	job.IPRange = iprange
	job.JID = uuid.New().String()
	job.TID = tid
	job.Status = "queued"

	return job
}
