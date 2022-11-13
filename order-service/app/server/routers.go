package server

import "github.com/saufiroja/order-service/rpc/pb/order"

func (rpc *GrpcServer) defineRoute(provider gRPCProvider) {
	order.RegisterOrderServiceServer(rpc.grpcServer, &provider.handlers.order)
}
