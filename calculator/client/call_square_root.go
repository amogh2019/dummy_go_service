package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/amogh2019/dummy_go_service/calculator/proto"
	"google.golang.org/grpc/status"
)

func callSquareRoot(clientInstance pb.CalculatorServiceClient, num int32) {

	res, err := clientInstance.Sqrt(context.Background(), &pb.SquareRootRequest{Number: num})

	if err != nil {
		grpcParsedStatus, ok := status.FromError(err)
		if ok {
			log.Println("ErrorCode : ", grpcParsedStatus.Code(), " ErrorMsg: ", grpcParsedStatus.Message())
		} else {
			log.Println("Error in finding sqaure root from server", err)
		}
		return
	}
	log.Println("square root of ", num, " is ", res.Root)

}

func askForInputSquareRoot() {
	fmt.Println()
	fmt.Println("Enter number to get its square root value : ")

}
