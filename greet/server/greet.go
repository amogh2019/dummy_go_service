package main

import (
	"context"
	"fmt"

	pb "github.com/amogh2019/dummy_go_service/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) { // this is basically an implementation of the Greet service server we wrote in proto // in our server implementation // // GreetServiceServer is the server API for GreetService service.

	fmt.Println("Greet function was invoked", in)

	return &pb.GreetResponse{
		Result: "Yolo! " + in.FirstName,
	}, nil //nil error
}



