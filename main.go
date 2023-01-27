package main

import (
	"example.com/m/models"
	"example.com/m/route"
)

func initialMigration() {
	// Migrate the schema
	models.DB.AutoMigrate(&models.User{})
}

func main() {
	models.ConnectDatabase()
	r := route.SetupRouter()
	r.Run(":9888")

}
