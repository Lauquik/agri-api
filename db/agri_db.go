package db

import (
	"github.com/SaiNageswarS/go-api-boot/odm"
	"github.com/mongo-tut/models"
)

type AgriDB struct{}

func (*AgriDB) Product() *ProductRepository {
	repo := odm.AbstractRepository[models.Product]{
		Database:       "agri",
		CollectionName: "products",
	}
	return &ProductRepository{repo}
}

func (*AgriDB) Shop() *ShopRepository {
	repo := odm.AbstractRepository[models.Shop]{
		Database:       "agri",
		CollectionName: "shops",
	}
	return &ShopRepository{repo}
}

func (*AgriDB) User() *UserRepository {
	repo := odm.AbstractRepository[models.Users]{
		Database:       "agri",
		CollectionName: "users",
	}
	return &UserRepository{repo}
}
