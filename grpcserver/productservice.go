package grpcserver

import (
	"context"
	"log"

	"github.com/mongo-tut/internal/database"
	"github.com/mongo-tut/pkg/pb"
)

func (u *Productlistserver) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (p *Productlistserver) GetProductList(req *pb.Productlistreq, stream pb.ProductService_GetProductListServer) error {
	var listproducts *[]database.Product
	listproducts, err := p.Productrepo.Getproductlistodm(req.Serviceable)

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range *listproducts {
		res := &pb.Productlist{
			Products: &pb.Productattribs{
				Proudctid:   v.ProductId.String(),
				Name:        v.Name,
				Price:       v.Price,
				Description: v.Description,
				Weight:      v.Weight,
				Ratings:     int32(v.Ratings),
				Warranty:    int32(v.Warranty),
			},
		}
		stream.Send(res)
	}
	return nil
}
