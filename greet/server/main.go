package main

import (
	"log"
	"net"

	pb "github.com/amogh2019/dummy_go_service/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50051" // localhost port 500051

type Server struct {
	pb.GreetServiceServer
}

func main() {
	// SETPS TO CREATE SERVER
	// 1. open tcp listener // if successful continue
	// 2. create a grpc server instance on that listener
	// 3. make server serve

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("cannot listen to", addr, err)
	}
	log.Println("started listening on ", addr)

	opts := []grpc.ServerOption{}
	tlsEnabled := true
	if tlsEnabled {

		cred, err := credentials.NewServerTLSFromFile("ssl/server.crt", "ssl/server.pem")
		if err != nil {
			log.Fatal("cannot load server cert", err)
		}
		opts = append(opts, grpc.Creds(cred))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{}) // registering our server type // which is an implementation of the services in proto

	if err = s.Serve(lis); err != nil {
		log.Fatal("server cannot serve on listener to address", addr, s, lis, err)
	}

}

// Do remember to do go run main.go greet.go // i.e. run both the files // else nil pointed // segmentation fault
// go run greet/server/*.go
