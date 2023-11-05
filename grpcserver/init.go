package grpcserver

import (
	"github.com/mongo-tut/pkg/handler"
	"github.com/mongo-tut/pkg/pb"
)

type Productlistserver struct {
	pb.UnimplementedProductServiceServer
	Productrepo *handler.ProductRepository
}

type Shopserver struct {
	pb.UnimplementedShopServiceServer
	Shoprepo *handler.ShopRepository
}

type UserServer struct {
	pb.UnimplementedUserServiceServer
	Userrepo *handler.UserRepository
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
