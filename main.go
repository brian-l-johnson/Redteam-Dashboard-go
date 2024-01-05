package main

import (
	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/db"
	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/server"
)

func main() {
	db.Init()
	server.Init()
}
