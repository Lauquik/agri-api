package db

import (
	"github.com/SaiNageswarS/go-api-boot/odm"
	"github.com/mongo-tut/models"
	"go.mongodb.org/mongo-driver/bson"
)

type ProductRepository struct {
	odm.AbstractRepository[models.Product]
}

func (p *ProductRepository) FindproductsByService(serviceable bool) (*[]models.Product, error) {
	var currchan = make(chan []models.Product)
	var errchan = make(chan error)
	if serviceable {
		currchan, errchan = p.Find(bson.M{"warranty": bson.M{"$ne": 0}}, bson.D{}, 0, 0)
	} else {
		currchan, errchan = p.Find(bson.M{}, bson.D{}, 0, 0)
	}
	select {
	case curr := <-currchan:
		return &curr, nil
	case err := <-errchan:
		return nil, err
	}

}
