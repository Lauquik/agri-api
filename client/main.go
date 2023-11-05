package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/mongo-tut/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	// productclint := pb.NewProductServiceClient(conn)
	// shopclient := pb.NewShopServiceClient(conn)
	userclient := pb.NewUserServiceClient(conn)

	// getproductlist(productclint, false)
	// getbyproduct(shopclient, "6544ef432e4d3c394344d72a")
	// getnearstshop(shopclient, &[2]float32{12.9715987, 77.5945627})
	var email = "lavquik@gamil.com"
	getnearestuser(userclient, &(email))

}

func getproductlist(client pb.ProductServiceClient, serviceable bool) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	res, err := client.GetProductList(ctx, &pb.Productlistreq{Serviceable: serviceable})

	if err != nil {
		log.Fatal(err)
	}

	for {
		productlists, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(productlists.GetProducts())
	}

}

func getbyproduct(client pb.ShopServiceClient, productID string) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	res, err := client.GetByProduct(ctx, &pb.Productid{ProductID: productID})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func getnearstshop(client pb.ShopServiceClient, location *[2]float32) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	res, err := client.GetNearByShop(ctx, &pb.Shopcords{ShopCords: &pb.Coordinates{Lat: location[0], Long: location[1]}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func getnearestuser(client pb.UserServiceClient, email *string) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	res, err := client.GetNearestUser(ctx, &pb.Userreq{Email: *email})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
