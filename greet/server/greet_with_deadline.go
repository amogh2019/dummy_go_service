package main

import (
	"context"
	"log"
	"time"

	pb "github.com/amogh2019/dummy_go_service/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("Greet With Deadline")

	time.Sleep(time.Second * 4)

	if ctx.Err() == context.DeadlineExceeded {
		msg := "server will take more time than the time expected by client context"
		log.Println("req:", req, msg)
		return nil, status.Error(codes.DeadlineExceeded, msg)
	}

	return &pb.GreetResponse{Result: "Hello from server, " + req.FirstName}, nil
}
