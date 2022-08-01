package main

import (
	"context"
	"fmt"

	pb "github.com/amogh2019/dummy_go_service/greet/proto"
)

func doGreetToServer(clientInstance pb.GreetServiceClient) {

	fmt.Println("invoking greet to server on this cliend")

	res, err := clientInstance.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "saitama",
	})
	if err != nil {
		fmt.Println("could not greet")
	}

	fmt.Println("Response from server: ", res)
}
