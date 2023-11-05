package grpcserver

import (
	"github.com/mongo-tut/pkg/pb"
	"go.mongodb.org/mongo-driver/mongo"
)

type Productlistserver struct {
	pb.UnimplementedProductServiceServer
	DB *mongo.Client
}

type Shopserver struct {
	pb.UnimplementedShopServiceServer
	DB *mongo.Client
}

type UserServer struct {
	pb.UnimplementedUserServiceServer
	DB *mongo.Client
}

// func main() {
// 	client, err := database.InitDb()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("DB connected")
// 	server.LoadSecretsIntoEnv(false)
// 	corsConfig := cors.New(
// 		cors.Options{
// 			AllowedHeaders: []string{"*"},
// 		})

// 	bootServer := server.NewGoApiBoot(corsConfig)
// 	pb.RegisterProductServiceServer(bootServer.GrpcServer, &Productlistserver{db: client})
// 	pb.RegisterShopServiceServer(bootServer.GrpcServer, &Shopserver{db: client})
// 	pb.RegisterUserServiceServer(bootServer.GrpcServer, &UserServer{db: client})
// 	bootServer.Start(":50051", ":8081")
// }
