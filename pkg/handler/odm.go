package handler

import (
	"log"

	"github.com/SaiNageswarS/go-api-boot/odm"
	"github.com/mongo-tut/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductRepository struct {
	odm.AbstractRepository[database.Product]
}

func NewProductRepo() *ProductRepository {
	repo := odm.AbstractRepository[database.Product]{
		Database:       "agri",
		CollectionName: "products",
	}
	return &ProductRepository{repo}
}

type ShopRepository struct {
	odm.AbstractRepository[database.Shop]
}

func NewShopRepo() *ShopRepository {
	repo := odm.AbstractRepository[database.Shop]{
		Database:       "agri",
		CollectionName: "shops",
	}
	return &ShopRepository{repo}
}

type UserRepository struct {
	odm.AbstractRepository[database.Users]
}

func NewUserRepo() *UserRepository {
	repo := odm.AbstractRepository[database.Users]{
		Database:       "agri",
		CollectionName: "users",
	}
	return &UserRepository{repo}
}

func (p *ProductRepository) Getproductlistodm(serviceable bool) (*[]database.Product, error) {
	var currchan = make(chan []database.Product)
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

func (s *ShopRepository) GetByProductodm(productid string) (*database.Shop, error) {
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

func (s *ShopRepository) GetByNearestShopodm(locationCoordinates *[]float64) (*database.Shop, error) {
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

func (u *UserRepository) GetNearUserodm(useremail string) (*database.Users, error) {

	var user *database.Users
	currchan, errchan := u.FindOne(bson.M{"email": useremail})
	select {
	case finduser := <-currchan:
		user = finduser
	case err := <-errchan:
		return nil, err
	}

	var location = user.Location

	currchan, errchan = u.FindOne(
		bson.M{
			"email": bson.M{"$ne": useremail},
			"location": bson.M{
				"$nearSphere": bson.M{
					"$geometry": bson.M{
						"type":        "Point",
						"coordinates": location.Coordinates,
					},
				},
			},
		})

	select {
	case userfound := <-currchan:
		return userfound, nil
	case err := <-errchan:
		return nil, err
	}
}
