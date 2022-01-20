package main

import (
	"context"
	"log"
	"net"

	pb "github.com/tagpro/grpc-go-demo/proto"

	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/xds"
)

const port = ":8000"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("unable to listen on port ", port, err)
	}
	defer func(lis net.Listener) {
		err := lis.Close()
		if err != nil {
			log.Println(err)
		}
	}(lis)
	server := xds.NewGRPCServer()
	reflection.Register(server)

	// Register API v1
	pb.RegisterMathServiceServer(server, &mathServer{})

	log.Printf("listening on port %s", port)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("Failed at: %v", err)
	}
}

// Server implementation

type mathServer struct {
	pb.UnimplementedMathServiceServer
}

func (s *mathServer) Add(ctx context.Context, in *pb.MathRequest) (*pb.MathResponse, error) {
	log.Println("Add called with ", in.Num1, in.Num2)
	return &pb.MathResponse{Value: in.Num1 + in.Num2}, nil
}
