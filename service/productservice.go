package service

import (
	"context"
	"log"

	"github.com/mongo-tut/db"
	"github.com/mongo-tut/models"
	"github.com/mongo-tut/pb"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	db *db.AgriDB
}

func NewProductService(db *db.AgriDB) *ProductService {
	return &ProductService{db: db}
}

func (u *ProductService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (p *ProductService) GetProductList(req *pb.Productlistreq, stream pb.ProductService_GetProductListServer) error {
	var listproducts *[]models.Product
	listproducts, err := p.db.Product().FindproductsByService(req.Serviceable)

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
