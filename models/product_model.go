package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

type ProductInventory struct {
	ProductID primitive.ObjectID `bson:"productID" json:"productID"`
	Quantity  int32              `bson:"quantity" json:"quantity"`
}

type ProductCatlogue struct {
	ShopID   primitive.ObjectID `bson:"shopid" json:"shopid"`
	Products []Product          `bson:"products" json:"products"`
}

func (m *Product) Id() string {
	return m.ProductId.Hex()
}
