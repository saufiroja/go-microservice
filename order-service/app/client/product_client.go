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

func NewProductServerClient() ProductServerClient {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	log.Println("Connected to product server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return ProductServerClient{
		Client: product.NewProductServiceClient(conn),
	}
}

func (p *ProductServerClient) FindOne(id string) (*product.FindOneResponse, error) {
	req := &product.FindOneRequest{
		Id: id,
	}

	return p.Client.FindOne(context.Background(), req)
}

func (p *ProductServerClient) DecreaseStock(productId, orderId string) (*product.DecreaseStockResponse, error) {
	req := &product.DecreaseStockRequest{
		Id:      productId,
		OrderId: orderId,
	}
	return p.Client.DecreaseStock(context.Background(), req)
}
