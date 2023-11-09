package db

import (
	"log"

	"github.com/SaiNageswarS/go-api-boot/odm"
	"github.com/mongo-tut/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShopRepository struct {
	odm.AbstractRepository[models.Shop]
}

func (s *ShopRepository) FindShopByProductID(productid string) (*models.Shop, error) {
	productID, err := primitive.ObjectIDFromHex(productid)
	if err != nil {
		log.Fatal(err)
	}
	currchan, errchan := s.FindOne(bson.M{"products.productID": productID})
	select {
	case shop := <-currchan:
		return shop, nil
	case err := <-errchan:
		return nil, err
	}
}

func (s *ShopRepository) FindNearestShopByLocation(locationCoordinates *[]float64) (*models.Shop, error) {
	currchan, errchan := s.FindOne(bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": *locationCoordinates,
				},
			},
		},
	})
	select {
	case shop := <-currchan:
		return shop, nil
	case err := <-errchan:
		return nil, err
	}
}
