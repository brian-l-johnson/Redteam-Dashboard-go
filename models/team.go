package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model `json:"-"`
	Name       string
	IPRange    string
	ID         string
}

func MakeTeam(name string, iprange string) Team {
	var team Team
	team.Name = name
	team.IPRange = iprange
	team.ID = uuid.New().String()

	return team
}
