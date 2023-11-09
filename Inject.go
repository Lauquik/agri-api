package main

import (
	"github.com/mongo-tut/db"
	"github.com/mongo-tut/service"
)

type Inject struct {
	AgriDB *db.AgriDB

	ProductService *service.ProductService
	ShopService    *service.ShopService
	UserService    *service.UserService
}

func NewInject() *Inject {
	inj := &Inject{}
	inj.AgriDB = &db.AgriDB{}
	inj.ProductService = service.NewProductService(inj.AgriDB)
	inj.ShopService = service.NewShopService(inj.AgriDB)
	inj.UserService = service.NewUserService(inj.AgriDB)
	return inj
}
