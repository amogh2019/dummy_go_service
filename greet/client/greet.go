package main

import (
	"context"
	"fmt"
	"io"
	"strings"

	pb "github.com/amogh2019/dummy_go_service/greet/proto"
)

func doGreetToServer(clientInstance pb.GreetServiceClient, yourname string) {

	fmt.Println("invoking greet to server on this client")

	res, err := clientInstance.Greet(context.Background(), &pb.GreetRequest{
		FirstName: strings.ToUpper(yourname),
	})
	if err != nil {
		fmt.Println("could not greet")
	}

	fmt.Println("Response from server: ", res)
}

func doGetGreetFromServerManyTimes(clientInstance pb.GreetServiceClient, yourname string) {

	fmt.Println("invoking  greetmanytime request to server from this client")

	responseStream, err := clientInstance.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: strings.ToUpper(yourname),
	})
	if err != nil {
		fmt.Println("could not greet")
	}

	for {
		resp, err := responseStream.Recv()

		if err == io.EOF {
			fmt.Println("everything read from stream response from server")
			break
		} else if err != nil {
			fmt.Println("Broken Response from server:  ", resp)
			break
		}
		fmt.Println("Response from server: ", resp.Result)
	}

}

func askForInput() {
	fmt.Println()
	fmt.Println("Enter your name and then press enter to send your name to server! // type exit to break")

}
