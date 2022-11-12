package server

import (
	"log"
	"net"

	"github.com/saufiroja/product-service/config"
	"github.com/saufiroja/product-service/utils"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	host       string
	port       string
	grpcServer *grpc.Server
}

func NewGrpcServer(host, port string) *GrpcServer {
	return &GrpcServer{
		host: host,
		port: port,
	}
}

func (rpc *GrpcServer) Start() {
	rpc.host = config.NewAppConfig().GRPC.Host
	rpc.port = config.NewAppConfig().GRPC.Port

	lis, err := net.Listen("tcp", rpc.host+":"+rpc.port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	rpc.grpcServer = grpc.NewServer()

	log.Println(utils.Color("green") + "----------------------------------------")
	log.Println(utils.Color("green")+"gRPC server is running on", rpc.host+":"+rpc.port)
	log.Println(utils.Color("green") + "----------------------------------------")

	if err := rpc.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
