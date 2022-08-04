package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	pb "github.com/amogh2019/dummy_go_service/greet/proto"
)

func doLongGreet(clientInstance pb.GreetServiceClient, names []string) {

	log.Println("invoking long greet to server on this client")

	messageSendingStream, err := clientInstance.LongGreet(context.Background())

	if err != nil {
		log.Fatal("err in setting up things to do long greet", err)
	}

	for _, name := range names {
		messageSendingStream.Send(&pb.GreetRequest{
			FirstName: strings.Title(name),
		})
	}

	res, err := messageSendingStream.CloseAndRecv()
	if err != nil {
		fmt.Println("could not close the message sending stream", err)
	}

	log.Println("Response from server: ", res.Result)
}

func askForInputForLongGreet() {
	fmt.Println("or Enter comma separated names and then press enter to send to server! // type exit to break")
}
