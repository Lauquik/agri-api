package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shop struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	Name               string             `bson:"name"`
	Location           GeoJSON            `bson:"location"`
	ProductsInveontroy []ProductInventory `bson:"products"`
}

type Users struct {
	ShopID   primitive.ObjectID `bson:"shopID"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Location GeoJSON            `bson:"location"`
}

type GeoJSON struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

type ProductInventory struct {
	ProductID primitive.ObjectID `bson:"productID"`
	Quantity  int32              `bson:"quantity"`
}

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ProductId   primitive.ObjectID `bson:"productID" json:"productID"`
	Name        string             `bson:"name" json:"name"`
	Price       int32              `bson:"price" json:"price"`
	Description string             `bson:"description" json:"description"`
	Weight      string             `bson:"weight" json:"weight"`
	Ratings     int16              `bson:"ratings" json:"ratings"`
	Warranty    int16              `bson:"warranty" json:"warranty"`
}

func (m *Product) Id() string {
	return m.ID.Hex()
}
