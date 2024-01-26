package models

import "gorm.io/gorm"

type Port struct {
	gorm.Model
	Number   uint16
	State    string
	Protocol string
	Service  string
	HostID   uint
}
