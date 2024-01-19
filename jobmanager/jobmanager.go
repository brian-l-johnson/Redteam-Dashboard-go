package jobmanager

import (
	"errors"
	"fmt"

	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/db"
	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/models"
)

var manager *JobManager

type JobManager struct {
	Types   []string
	Teams   []models.Team
	teamptr int
}

func Init() {
	manager = &JobManager{}
	manager.Types = append(manager.Types, "nmap")

	db := db.GetDB()
	var teams []models.Team
	result := db.Find(&teams)
	if result.Error != nil {
		panic("db error on load")
	}
	for _, t := range teams {
		manager.Teams = append(manager.Teams, t)
	}
	manager.teamptr = 0

}

func GetJobManager() *JobManager {
	return manager
}

func (m JobManager) Test() {
	fmt.Println("hi there")
}

func (m JobManager) GetNextJob() (models.Job, error) {
	if len(m.Teams) > 0 {
		job := models.MakeJob(m.Types[0], m.Teams[m.teamptr].IPRange, m.Teams[m.teamptr].ID)
		m.teamptr = (m.teamptr + 1) % len(m.Teams)

		return job, nil
	} else {
		return models.Job{}, errors.New("no available teams")
	}

}
