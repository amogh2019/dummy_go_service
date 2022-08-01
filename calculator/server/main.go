package main

import (
	"log"
	"net"

	pb "github.com/amogh2019/dummy_go_service/calculator/proto"
	"google.golang.org/grpc"
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

	if err := server.Serve(conn); err != nil {
		log.Fatal("cannot server at addr", serverAddress, err)
	}
}
