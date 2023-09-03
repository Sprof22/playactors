package main

import (
	"playactors/controllers"
	"playactors/initializer"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.POST("/actors", controllers.CreateActor)
	r.GET("/actors", controllers.GetActors)
	r.GET("/actors/:id", controllers.GetSingleActor)
	r.DELETE("/actors/:id", controllers.DeleteActor)
	r.PUT("/actors/:id", controllers.UpdateActor)
	r.Run() // listen and serve on 0.0.0.0:8080
}
