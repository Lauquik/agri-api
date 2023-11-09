package service

import (
	"context"
	"log"

	"github.com/mongo-tut/db"
	"github.com/mongo-tut/models"
	"github.com/mongo-tut/pb"
)

type ShopService struct {
	pb.UnimplementedShopServiceServer
	db *db.AgriDB
}

func NewShopService(db *db.AgriDB) *ShopService {
	return &ShopService{db: db}
}

func (u *ShopService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (s *ShopService) GetNearByShop(ctx context.Context, req *pb.Shopcords) (*pb.Shopres, error) {
	var shop *models.Shop
	var location []float64 = []float64{float64(req.ShopCords.Lat), float64(req.ShopCords.Long)}
	shop, err := s.db.Shop().FindNearestShopByLocation(&location)
	if err != nil {
		panic(err)
	}
	return &pb.Shopres{
		ShopAttribs: &pb.Shopattribs{
			Name:   shop.Name,
			Shopid: shop.ID.String(),
		},
	}, nil

}

func (s *ShopService) GetByProduct(ctx context.Context, req *pb.Productid) (*pb.Productshopres, error) {
	var shop *models.Shop
	shop, err := s.db.Shop().FindShopByProductID(req.ProductID)

	if err != nil {
		log.Fatal(err)
	}
	return &pb.Productshopres{
		ShopAttribs: &pb.Shopattribs{
			Name:   shop.Name,
			Shopid: shop.ID.Hex(),
		},
	}, nil
}
