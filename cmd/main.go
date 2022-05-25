package main

import (
	"fmt"
	"log"
	"net"
	"order-micro/pkg/client"
	"order-micro/pkg/config"
	"order-micro/pkg/db"
	"order-micro/pkg/pb"
	"order-micro/pkg/service"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Print(c.Port)
	h := db.Init(c.Dburl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("unable to listen at port", err.Error())
	}

	fmt.Println("Order Svc on", c.Port)

	ProductSvc := client.InitProductServiceClient(c.ProductSvcUrl)

	s := service.Server{
		H:          h,
		ProductSvc: ProductSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("unable to start grpc server", err.Error())
	}
}
