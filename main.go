package main

import (
	"github.com/SaiNageswarS/go-api-boot/server"
	"github.com/mongo-tut/grpcserver"
	"github.com/mongo-tut/pkg/handler"
	"github.com/mongo-tut/pkg/pb"
	"github.com/rs/cors"
)

const grpcport string = ":50051"
const webport string = ":8000"

func main() {
	productrepo := handler.NewProductRepo()
	shoprepo := handler.NewShopRepo()
	userrepo := handler.NewUserRepo()
	server.LoadSecretsIntoEnv(false)
	corsConfig := cors.New(
		cors.Options{
			AllowedHeaders: []string{"*"},
		})

	bootServer := server.NewGoApiBoot(corsConfig)

	pb.RegisterProductServiceServer(bootServer.GrpcServer, &grpcserver.Productlistserver{Productrepo: productrepo})
	pb.RegisterShopServiceServer(bootServer.GrpcServer, &grpcserver.Shopserver{Shoprepo: shoprepo})
	pb.RegisterUserServiceServer(bootServer.GrpcServer, &grpcserver.UserServer{Userrepo: userrepo})

	bootServer.Start(grpcport, webport)
}
