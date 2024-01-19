package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model `json:"-"`
	Name       string
	IPRange    string
	TID        string
}

func MakeTeam(name string, iprange string) Team {
	var team Team
	team.Name = name
	team.IPRange = iprange
	team.TID = uuid.New().String()

	return team
}
