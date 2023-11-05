package main

import (
	"fmt"
	"log"

	"github.com/SaiNageswarS/go-api-boot/server"
	"github.com/mongo-tut/grpcserver"
	"github.com/mongo-tut/internal/database"
	"github.com/mongo-tut/pkg/pb"
	"github.com/rs/cors"
)

const grpcport string = ":50051"
const webport string = ":8000"

func main() {
	client, err := database.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB connected")
	server.LoadSecretsIntoEnv(false)
	corsConfig := cors.New(
		cors.Options{
			AllowedHeaders: []string{"*"},
		})

	bootServer := server.NewGoApiBoot(corsConfig)
	pb.RegisterProductServiceServer(bootServer.GrpcServer, &grpcserver.Productlistserver{DB: client})
	pb.RegisterShopServiceServer(bootServer.GrpcServer, &grpcserver.Shopserver{DB: client})
	pb.RegisterUserServiceServer(bootServer.GrpcServer, &grpcserver.UserServer{DB: client})
	bootServer.Start(grpcport, webport)
}
