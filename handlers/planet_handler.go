package handlers

import (
	"github.com/Arcelz/start_wars_api/models"
	"github.com/Arcelz/start_wars_api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreatePlanet(c *gin.Context) {
	planet := models.Planet{}
	if err := c.ShouldBindJSON(&planet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service := services.PlanetService{}
	err := service.CreatePlanet(&planet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	m := make(map[string]string)
	m["id"] = planet.ID.Hex()
	c.JSON(200, m)
}

func GetPlanets(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	name := c.DefaultQuery("name", "")
	service := services.PlanetService{}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(400, "The page parameter is not in int")
		return
	}
	planets, err := service.GetAllPlanets(pageInt, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, planets)
}

func GetPlanet(c *gin.Context) {
	service := services.PlanetService{}
	planet, err := service.GetPlanet(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, planet)
}

func DeletePlanet(c *gin.Context) {
	service := services.PlanetService{}
	err := service.DeletePlanet(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
}
