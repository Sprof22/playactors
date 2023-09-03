package main

import (
	"playactors/initializer"
	"playactors/models"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	initializer.DB.AutoMigrate(&models.Actors{})
}
