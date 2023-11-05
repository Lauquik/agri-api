package handler

import (
	"github.com/SaiNageswarS/go-api-boot/odm"
	"github.com/mongo-tut/internal/database"
)

type ProductRepository struct {
	odm.AbstractRepository[database.Product]
}

func NewProductRepo() *ProductRepository {
	repo := odm.AbstractRepository[database.Product]{
		Database:       "agri",
		CollectionName: "Product",
	}
	return &ProductRepository{repo}
}
