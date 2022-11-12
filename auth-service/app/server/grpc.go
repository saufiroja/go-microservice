package server

import (
	"log"
	"net"

	"github.com/saufiroja/auth-service/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	// Register your service here
	provider := rpc.provide()

	rpc.grpcServer = grpc.NewServer()

	rpc.defineRoute(provider)

	reflection.Register(rpc.grpcServer)

	log.Println("----------------------------------------")
	log.Println("gRPC server is listening on " + rpc.host + ":" + rpc.port)
	log.Println("----------------------------------------")

	if err := rpc.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
