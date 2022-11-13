package server

import "github.com/saufiroja/product-service/rpc/pb/product"

func (rpc *GrpcServer) defineRoute(provider gRPCProvider) {
	product.RegisterProductServiceServer(rpc.grpcServer, &provider.handlers.product)
}
