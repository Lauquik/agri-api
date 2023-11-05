package grpcserver

import (
	"context"
	"fmt"
	"log"

	"github.com/mongo-tut/pkg/pb"
)

func (u *UserServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (s *UserServer) GetNearestUser(ctx context.Context, req *pb.Userreq) (*pb.Userres, error) {
	nearuser, err := s.Userrepo.GetNearUserodm(req.GetEmail())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nearuser)
	return &pb.Userres{
		Name:   nearuser.Name,
		Shopid: nearuser.ShopID.String(),
		Email:  nearuser.Email,
	}, nil

}
