package models

import (
	"database/sql/driver"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Roles []string

type User struct {
	gorm.Model
	Name         string
	PasswordHash string
	Active       bool
	Roles        Roles `gorm:"type:VARCHAR(255)"`
}

func (u *User) SetPassword(pw string) {
	bytes, hasherr := bcrypt.GenerateFromPassword([]byte(pw), 14)
	if hasherr != nil {
		panic("unable to hash password")
	}
	u.PasswordHash = string(bytes)
}

func (r *Roles) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("src value cannot be cast to []byte")
	}
	*r = strings.Split(string(bytes), ",")

	return nil
}

func (r Roles) Value() (driver.Value, error) {
	if len(r) == 0 {
		return nil, nil
	}
	return strings.Join(r, ","), nil
}
