package main

import (
	"log"
	"net"

	pb "github.com/amogh2019/dummy_go_service/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serverAddress = "0.0.0.0:50052"

type CalculatorServiceServerImpl struct {
	pb.CalculatorServiceServer
}

func main() {

	conn, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatal("cannot open tcp conn", err)
	}
	defer conn.Close()

	server := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(server, &CalculatorServiceServerImpl{})

	reflection.Register(server) // enabling reflection to get service definitions by grpc clients

	if err := server.Serve(conn); err != nil {
		log.Fatal("cannot server at addr", serverAddress, err)
	}
}
