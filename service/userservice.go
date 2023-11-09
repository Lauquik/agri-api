package service

import (
	"context"
	"fmt"
	"log"

	"github.com/mongo-tut/db"
	"github.com/mongo-tut/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	db *db.AgriDB
}

func NewUserService(db *db.AgriDB) *UserService {
	return &UserService{db: db}
}

func (u *UserService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (s *UserService) GetNearestUser(ctx context.Context, req *pb.Userreq) (*pb.Userres, error) {
	nearuser, err := s.db.User().FindNearestUserByEmail(req.GetEmail())
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
