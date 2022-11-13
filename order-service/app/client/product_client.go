package client

import (
	"context"
	"log"

	"github.com/saufiroja/order-service/rpc/pb/product"
	"google.golang.org/grpc"
)

type ProductServerClient struct {
	Client product.ProductServiceClient
}

func InitProductServiceClient(url string) ProductServerClient {
	// use withTransportCredentials
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return ProductServerClient{
		Client: product.NewProductServiceClient(conn),
	}
}

func (c *ProductServerClient) FindOne(id string) (*product.FindOneResponse, error) {
	req := &product.FindOneRequest{
		Id: id,
	}

	return c.Client.FindOne(context.Background(), req)
}

func (c *ProductServerClient) DecreaseStock(productId, orderId string) (*product.DecreaseStockResponse, error) {
	req := &product.DecreaseStockRequest{
		Id:      productId,
		OrderId: orderId,
	}

	return c.Client.DecreaseStock(context.Background(), req)
}
