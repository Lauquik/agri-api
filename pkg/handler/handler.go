package handler

import (
	"context"
	"log"

	"github.com/mongo-tut/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	err error = nil
)

func GetProductList(client *mongo.Client, serviceable bool) (*[]database.Product, error) {
	collection := client.Database("agri").Collection("products")
	var results []database.Product
	var cur *mongo.Cursor
	if serviceable {
		cur, err = collection.Find(context.TODO(), bson.M{"warranty": bson.M{"$ne": 0}})
	} else {
		cur, err = collection.Find(context.TODO(), bson.D{})
	}
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elm database.Product
		err = cur.Decode(&elm)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elm)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}

func GetByProduct(client *mongo.Client, productid string) (*database.Shop, error) {
	productID, err := primitive.ObjectIDFromHex(productid)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("agri").Collection("shops")
	var shop database.Shop
	err = collection.FindOne(context.TODO(), bson.M{"products.productID": productID}).Decode(&shop)
	if err != nil {
		log.Fatal(err)
	}
	return &shop, nil
}

func GetByNearestShop(client *mongo.Client, locationCoordinates *[]float64) (*database.Shop, error) {
	collection := client.Database("agri").Collection("shops")

	// _, err := collection.Indexes().CreateOne(
	// 	context.TODO(),
	// 	mongo.IndexModel{
	// 		Keys: bson.M{
	// 			"location": "2dsphere",
	// 		},
	// 	},
	// )

	// if err != nil {
	// 	log.Fatal(err)
	// }

	var shop database.Shop
	err := collection.FindOne(context.TODO(), bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": *locationCoordinates,
				},
			},
		},
	}).Decode(&shop)
	if err != nil {
		log.Fatal(err)
	}
	return &shop, nil
}

func GetByServicableProduct(client *mongo.Client) (*[]database.Product, error) {
	collection := client.Database("agri").Collection("products")
	var product []database.Product
	cur, err := collection.Find(context.TODO(), bson.M{"warranty": bson.M{"$ne": 0}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elm database.Product
		err = cur.Decode(&elm)
		if err != nil {
			log.Fatal(err)
		}

		product = append(product, elm)
	}
	return &product, nil
}

func GetNearUser(client *mongo.Client, useremail string) *database.Users {
	collection := client.Database("agri").Collection("users")

	var user database.Users

	err := collection.FindOne(context.TODO(), bson.M{"email": useremail}).Decode(&user)

	if err != nil {
		panic(err)
	}

	var location = user.Location

	err = collection.FindOne(context.TODO(), bson.M{
		"email": bson.M{"$ne": useremail},
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": location.Coordinates,
				},
			},
		},
	}).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	return &user
}
