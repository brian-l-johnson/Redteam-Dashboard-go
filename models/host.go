package models

import "gorm.io/gorm"

type Host struct {
	gorm.Model
	IP       string
	Hostname string
	OS       string
	Ports    []Port
}
