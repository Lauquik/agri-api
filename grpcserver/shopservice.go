package grpcserver

import (
	"context"

	"github.com/mongo-tut/internal/database"
	"github.com/mongo-tut/pkg/handler"
	"github.com/mongo-tut/pkg/pb"
)

func (u *Shopserver) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (s *Shopserver) GetNearByShop(ctx context.Context, req *pb.Shopcords) (*pb.Shopres, error) {
	var shop *database.Shop
	var location []float64 = []float64{float64(req.ShopCords.Lat), float64(req.ShopCords.Long)}
	shop, err := handler.GetByNearestShop(s.DB, &location)
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
