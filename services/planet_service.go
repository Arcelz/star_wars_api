package services

import (
	"context"
	"errors"
	"github.com/Arcelz/start_wars_api/database"
	"github.com/Arcelz/start_wars_api/models"
	"github.com/Arcelz/swapi"
	. "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type PlanetService struct{}

// GetAllPlanets Retrives all planets from the db
func (service PlanetService) GetAllPlanets(page int, name string) ([]models.Planet, error) {
	filter := bson.D{}
	if name != "" {
		filter = append(filter, bson.E{Key: "name", Value: bson.D{
			{"$regex", primitive.Regex{Pattern: "^" + name + ".*", Options: "i"}},
		}})
	}
	planets := []models.Planet{}
	collection, err := database.GetMongoCollection(database.PLANETS)
	if err != nil {
		return planets, err
	}
	_, err = New(collection).Context(context.TODO()).Limit(10).Page(int64(page)).Filter(filter).Decode(&planets).Find()
	if err != nil {
		return planets, err
	}
	return planets, nil
}

func (service PlanetService) CreatePlanet(planet *models.Planet) error {
	if strings.TrimSpace(planet.Name) == "" {
		return errors.New("The planet name cannot be empty")
	} else if strings.TrimSpace(planet.Climate) == "" {
		return errors.New("The planet climate cannot be empty")
	} else if strings.TrimSpace(planet.Terrain) == "" {
		return errors.New("The planet terrain cannot be empty")
	}
	planet.Appearances = 0
	clientSwapi := swapi.DefaultClient
	if planets, err := clientSwapi.PlanetByName(planet.Name); err == nil {
		for i := range planets {
			if strings.ToLower(planets[i].Name) == strings.ToLower(planet.Name) {
				planet.Appearances = len(planets[i].FilmURLs)
			}
		}
	}
	planet.ID = primitive.NewObjectID()
	collection, err := database.GetMongoCollection(database.PLANETS)
	if err != nil {
		return err
	}
	res, err := collection.InsertOne(context.TODO(), planet)
	if err != nil {
		return err
	}
	planet.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (service PlanetService) DeletePlanet(id string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("The planet id cannot be empty")
	}
	hex, err2 := primitive.ObjectIDFromHex(id)
	if err2 != nil {
		return errors.New("Planet id incorrect format")
	}
	collection, err := database.GetMongoCollection(database.PLANETS)
	if err != nil {
		return err
	}
	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": hex})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("Planet not found")
	}
	return nil
}

func (service PlanetService) GetPlanet(id string) (*models.Planet, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("The planet id cannot be empty")
	}
	hex, err2 := primitive.ObjectIDFromHex(id)
	if err2 != nil {
		return nil, errors.New("Planet id incorrect format")
	}
	collection, err := database.GetMongoCollection(database.PLANETS)
	if err != nil {
		return nil, err
	}
	planet := models.Planet{}
	err = collection.FindOne(context.TODO(), bson.M{"_id": hex}).Decode(&planet)
	if err != nil {
		return nil, err
	}
	return &planet, nil
}
