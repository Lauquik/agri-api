package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shop struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name               string             `bson:"name" json:"name"`
	Location           GeoJSON            `bson:"location" json:"location"`
	ProductsInveontroy []ProductInventory `bson:"products" json:"products"`
}

type GeoJSON struct {
	Type        string    `bson:"type" json:"type"`
	Coordinates []float64 `bson:"coordinates" json:"coordinates"`
}

func (m *Shop) Id() string {
	return m.ID.Hex()
}
