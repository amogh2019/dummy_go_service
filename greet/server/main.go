package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051" // localhost port 500051

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

	s := grpc.NewServer()

	if err = s.Serve(lis); err != nil {
		log.Fatal("server cannot serve on listener to address", addr, s, lis, err)
	}

}
