package db

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	fmt.Println("initing db")
	var err error
	db, err = gorm.Open(sqlite.Open("dashboard.db"), &gorm.Config{})
	if err != nil {
		panic("faled to open database file")
	}
	db.AutoMigrate(&models.User{})

	var user models.User
	result := db.First(&user, "name=?", "admin")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("Admin user does not exits, creating...")
		user.Name = "admin"
		genpw, err := GenerateRandomString(24)
		if err != nil {
			panic("unable to generate random pw")
		}
		user.SetPassword(genpw)
		user.Roles = append(user.Roles, "admin")

		result = db.Create(&user)
		if result.Error != nil {
			panic("unable to save admin user")
		}
		fmt.Printf("created 'admin user with a password of '%s'\n", genpw)
	}

}

func GetDB() *gorm.DB {
	return db
}

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret = append(ret, letters[num.Int64()])
	}

	return string(ret), nil
}
