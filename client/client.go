package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/khanhduyy/shopms-product/rpc/product"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:8083", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	client := pb.NewProductAPIClient(cc)

	res, err := client.GetAllProduct(context.Background(), &pb.Empty{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(res.Products)
}
