package main

import (
	"context"
	"log"
	"strings"
	"time"

	pb "github.com/amogh2019/dummy_go_service/greet/proto"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(clientInstance pb.GreetServiceClient, yourname string, d time.Duration) {

	log.Println("invoking  greetwithdeadline request to server from this client")

	ctxWithDeadlin, cancelFunction := context.WithTimeout(context.Background(), d)
	defer cancelFunction()

	res, err := clientInstance.GreetWithDeadline(ctxWithDeadlin, &pb.GreetRequest{
		FirstName: strings.ToUpper(yourname),
	})
	if err != nil {
		grpcParsedStatus, ok := status.FromError(err)
		if ok {
			log.Println("ErrorCode : ", grpcParsedStatus.Code(), " ErrorMsg: ", grpcParsedStatus.Message())
		} else {
			log.Fatal("Non-gRPC Error in finding sqaure root from server", err)
		}
		return
	}

	log.Println("Response from server: ", res)

}
