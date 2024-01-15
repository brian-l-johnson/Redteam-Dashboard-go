package models

import "gorm.io/gorm"

type Host struct {
	gorm.Model `json:"-"`
	IP         string
	Hostname   string
	OS         string
	Ports      []Port
}
