package server

import "github.com/saufiroja/auth-service/rpc/pb/user"

func (rpc *GrpcServer) defineRoute(provider gRPCProvider) {
	user.RegisterAuthServiceServer(rpc.grpcServer, &provider.handlers.user)
}
