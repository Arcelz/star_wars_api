package test

import (
	"context"
	"github.com/Arcelz/start_wars_api/database"
	"github.com/Arcelz/start_wars_api/models"
	"github.com/Arcelz/start_wars_api/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setup() {
	database.DB = "star_wars_test_db"
}

func clear(t *testing.T) {
	client, err := database.GetMongoClient()
	if err != nil {
		t.Fatal()
	}
	err = client.Database(database.DB).Drop(context.TODO())
	if err != nil {
		t.Fatal()
	}
}

func TestCreate(t *testing.T) {
	setup()
	defer clear(t)
	assert := assert.New(t)
	planet := models.Planet{
		Name:    "Test Planet 5158841195233984",
		Climate: "Arid",
		Terrain: "Dessert",
	}
	service := services.PlanetService{}
	err := service.CreatePlanet(&planet)
	assert.Nil(err, "Não deve retornar erro ao criar planeta")
	assert.Zero(planet.Appearances, "O planeta com este nome não existe no mundo star wars e deve ter 0 aparições")
	assert.NotEqualValues("000000000000000000000000", planet.ID.Hex(), "Erro ao criar planeta o id foi salvo incorretamente.")
}

func TestExistingPlanetCreate(t *testing.T) {
	setup()
	defer clear(t)
	assert := assert.New(t)
	planet := models.Planet{
		Name:    "Tatooine",
		Climate: "Arid",
		Terrain: "Dessert",
	}
	service := services.PlanetService{}
	err := service.CreatePlanet(&planet)
	assert.Nil(err, "Não deve retornar erro ao criar planeta")
	assert.NotZero(planet.Appearances, "O planeta deve conter aparições nos filmes")
	assert.NotEqualValues("000000000000000000000000", planet.ID.Hex(), "Erro ao criar planeta o id foi salvo incorretamente.")
}

func TestGetById(t *testing.T) {
	setup()
	defer clear(t)
	assert := assert.New(t)
	planet := models.Planet{
		Name:    "Tatooine",
		Climate: "Arid",
		Terrain: "Dessert",
	}
	service := services.PlanetService{}
	_ = service.CreatePlanet(&planet)
	planetFromById, err := service.GetPlanet(planet.ID.Hex())
	assert.Nil(err, "Não deve retornar erro ao buscar planeta ja existente")
	assert.NotNil(planetFromById, "O planeta não deve retornar nil")
	assert.EqualValues(planetFromById.ID, planet.ID, "O id do planeta deve ser igual ao criado anteriormente")
	assert.EqualValues(planetFromById.Name, planet.Name, "O nome do planeta deve ser igual ao criado anteriormente")
}

func TestGetAll(t *testing.T) {
	setup()
	defer clear(t)
	assert := assert.New(t)
	planet1 := models.Planet{
		Name:    "Tatooine",
		Climate: "Arid",
		Terrain: "Dessert",
	}
	planet2 := models.Planet{
		Name:    "Alderaan",
		Climate: "temperate",
		Terrain: "grasslands, mountains",
	}
	planet3 := models.Planet{
		Name:    "Yavin IV",
		Climate: "temperate, tropical",
		Terrain: "jungle, rainforests",
	}
	service := services.PlanetService{}
	_ = service.CreatePlanet(&planet1)
	_ = service.CreatePlanet(&planet2)
	_ = service.CreatePlanet(&planet3)
	planets, err := service.GetAllPlanets(1, "")
	assert.Nil(err, "Não deve retornar erro ao buscar planetas ja existentes sem filtro de nome")
	assert.EqualValues(3, len(planets), "Deve retornar o mesmo tanto de planetas criados")
	planets, err = service.GetAllPlanets(1, planet1.Name)
	assert.Nil(err, "Não deve retornar erro ao buscar planetas ja existentes com filtro de nome")
	assert.EqualValues(1, len(planets), "Deve retornar 1 planeta")
	assert.EqualValues(planet1.Name, planets[0].Name, "O nome do planeta buscado deve ser igual o do resultado")
}

func TestDeleteById(t *testing.T) {
	setup()
	defer clear(t)
	assert := assert.New(t)
	planet := models.Planet{
		Name:    "Tatooine",
		Climate: "Arid",
		Terrain: "Dessert",
	}
	service := services.PlanetService{}
	_ = service.CreatePlanet(&planet)
	err := service.DeletePlanet(planet.ID.Hex())
	assert.Nil(err, "Não deve retornar erro ao deletar planeta ja existente")
	_, err = service.GetPlanet(planet.ID.Hex())
	assert.NotNil(err, "Deve retornar um erro pois o planeta ja foi deletado")
}
