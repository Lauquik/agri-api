package main

import (
	"github.com/SaiNageswarS/go-api-boot/server"
	"github.com/mongo-tut/pb"
	"github.com/rs/cors"
)

const grpcport string = ":50051"
const webport string = ":8081"

func main() {
	server.LoadSecretsIntoEnv(false)
	inject := NewInject()
	corsConfig := cors.New(
		cors.Options{
			AllowedHeaders: []string{"*"},
		})

	bootServer := server.NewGoApiBoot(corsConfig)

	pb.RegisterProductServiceServer(bootServer.GrpcServer, inject.ProductService)
	pb.RegisterShopServiceServer(bootServer.GrpcServer, inject.ShopService)
	pb.RegisterUserServiceServer(bootServer.GrpcServer, inject.UserService)

	bootServer.Start(grpcport, webport)
}
