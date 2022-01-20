package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/tagpro/grpc-go-demo/proto"
	"google.golang.org/grpc"

	// install xDS balancers and resolovers
	_ "google.golang.org/grpc/xds"
)

var (
	server     = "xds:///localhost:8080"
	num1, num2 int64
)

func main() {

	flag.Int64Var(&num1, "num1", 0, "first number")
	flag.Int64Var(&num2, "num2", 0, "second number")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewMathServiceClient(conn)
	response, err := client.Add(context.Background(), &pb.MathRequest{Num1: num1, Num2: num2})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d + %d = %d", num1, num2, response.Value)
}
