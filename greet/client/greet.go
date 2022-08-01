package main

import (
	"context"
	"fmt"
	"strings"

	pb "github.com/amogh2019/dummy_go_service/greet/proto"
)

func doGreetToServer(clientInstance pb.GreetServiceClient, yourname string) {

	fmt.Println("invoking greet to server on this cliend")

	res, err := clientInstance.Greet(context.Background(), &pb.GreetRequest{
		FirstName: strings.ToUpper(yourname),
	})
	if err != nil {
		fmt.Println("could not greet")
	}

	fmt.Println("Response from server: ", res)
}

func askForInput() {
	fmt.Println()
	fmt.Println("Enter your name and then press enter to send your name to server!")

}
