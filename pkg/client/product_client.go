package client

import (
	"context"
	"fmt"
	"log"
	"order-micro/pkg/pb"

	"google.golang.org/grpc"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func InitProductServiceClient(url string) ProductServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		log.Fatalln("unable to connect grpc server", err.Error())
	}

	c := ProductServiceClient{
		Client: pb.NewProductServiceClient(cc),
	}
	return c
}

func (c *ProductServiceClient) FindOne(ProductId int64) (*pb.FindOneResponse, error) {
	fmt.Println("productidd", ProductId)

	req := &pb.FindOneRequest{
		Id: ProductId,
	}

	return c.Client.FindOne(context.Background(), req)
}

func (c *ProductServiceClient) DecreaseStock(ProductId int64, OrderId int64) (*pb.DecreaseStockResponse, error) {
	req := &pb.DecreaseStockRequest{
		Id:      ProductId,
		OrderId: OrderId,
	}

	return c.Client.DecreaseStock(context.Background(), req)
}
