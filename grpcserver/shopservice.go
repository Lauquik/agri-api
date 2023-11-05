package grpcserver

import (
	"context"
	"log"

	"github.com/mongo-tut/internal/database"
	"github.com/mongo-tut/pkg/pb"
)

func (u *Shopserver) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (s *Shopserver) GetNearByShop(ctx context.Context, req *pb.Shopcords) (*pb.Shopres, error) {
	var shop *database.Shop
	var location []float64 = []float64{float64(req.ShopCords.Lat), float64(req.ShopCords.Long)}
	shop, err := s.Shoprepo.GetByNearestShopodm(&location)
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

func (p *Shopserver) GetByProduct(ctx context.Context, req *pb.Productid) (*pb.Productshopres, error) {
	var shop *database.Shop
	shop, err := p.Shoprepo.GetByProductodm(req.ProductID)

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
