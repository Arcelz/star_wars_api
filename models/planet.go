package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Planet struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name        string             `json:"name" `
	Climate     string             `json:"climate"`
	Terrain     string             `json:"terrain"`
	Appearances int                `json:"appearances"`
}
