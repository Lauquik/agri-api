package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type insertcollection struct {
	Client     *mongo.Client
	Collection string
}

func InitCollection(c *mongo.Client, col string) *insertcollection {
	return &insertcollection{
		Client:     c,
		Collection: col,
	}
}

func (col *insertcollection) InsertInCollection(doc interface{}) {
	collection := col.Client.Database("agri").Collection(col.Collection)
	insertResult, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

}

func InitDb() (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	return client, err
}
