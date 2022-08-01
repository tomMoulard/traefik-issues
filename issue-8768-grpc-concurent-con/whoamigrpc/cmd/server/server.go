package main

import (
	"context"
	"flag"
	"net"

	log "github.com/sirupsen/logrus"
	whoamigrpc "github.com/tommoulard/whoamigrpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", "50051", "give me a port number")
}

type server struct {
	whoamigrpc.UnsafeWhoamiServer
}

func (s *server) AskWhoAmI(ctx context.Context, request *whoamigrpc.WhoAmIRequest) (*whoamigrpc.WhoAmIReply, error) {
	log.Printf("Received: %v", request.GetName())
	return &whoamigrpc.WhoAmIReply{Message: "Hello " + request.GetName()}, nil
}

func main() {
	logger := log.WithFields(log.Fields{"port": port})

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Errorf("failed to listen: %w", err)
	}

	// s := grpc.NewServer()
	// whoamigrpc.RegisterWhoamiServer(s, &server{})
	// if err := s.Serve(l); err != nil {
	// log.Fatalf("failed to serve: %v", err)
	// }

	s := server{}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	whoamigrpc.RegisterWhoamiServer(grpcServer, &s)

	logger.Info("gRPC server started")
	if err := grpcServer.Serve(l); err != nil {
		logger.Errorf("failed to serve: %w", err)
	}
}
