package models

import (
	"errors"

	"gorm.io/gorm"
)

type JobStatus struct {
	gorm.Model
	Name     string
	JobIndex int
}

func (js *JobStatus) GetNextJob() (Job, error) {
	db := GetDB()
	var teams []Team
	result := db.Find(&teams)
	if result.Error != nil {
		return Job{}, errors.New("db error loading teams")
	} else {
		if len(teams) > 0 {
			job := MakeJob(js.Name, teams[js.JobIndex].IPRange, teams[js.JobIndex].TID)
			js.JobIndex = (js.JobIndex + 1) % len(teams)
			db.Save(&js)
			return job, nil
		} else {
			return Job{}, errors.New("no available teams")
		}
	}
}
