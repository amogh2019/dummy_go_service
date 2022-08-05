package main

import (
	"context"
	"log"
	"math"

	pb "github.com/amogh2019/dummy_go_service/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (serverImpl *CalculatorServiceServerImpl) Sqrt(ctx context.Context, req *pb.SquareRootRequest) (*pb.SquareRootResponse, error) {

	log.Println("operation received to find square root : ", req)

	if req.Number < 0 {
		errString := "A non negative value is expected."
		log.Println(codes.InvalidArgument, errString, req)
		return nil, status.Error(codes.InvalidArgument, errString)
	}

	root := math.Sqrt(float64(req.Number))
	log.Println("root for ", req.Number, " is : ", root)
	return &pb.SquareRootResponse{
		Root: root,
	}, nil
}
