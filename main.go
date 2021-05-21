package main

import (
	"github.com/Arcelz/star_wars_api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/planet", handlers.CreatePlanet)
	router.GET("/planets", handlers.GetPlanets)
	router.GET("/planet/:id", handlers.GetPlanet)
	router.DELETE("/planet/:id", handlers.DeletePlanet)
	router.Run()
}
