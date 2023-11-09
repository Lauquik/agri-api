package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ShopID   primitive.ObjectID `bson:"shopID" json:"shopID"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Location GeoJSON            `bson:"location" json:"location"`
}

func (m *Users) Id() string {
	return m.ShopID.Hex()
}
