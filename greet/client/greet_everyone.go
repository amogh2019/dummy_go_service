package main

import (
	"context"
	"io"
	"log"
	"strings"
	"time"

	pb "github.com/amogh2019/dummy_go_service/greet/proto"
)

func getGreetingsForEveryone(clientInstance pb.GreetServiceClient, names []string) {

	log.Println("invoking greet for everyone to server on this client")

	bistream, err := clientInstance.GreetEveryone(context.Background())

	if err != nil {
		log.Fatal("err in setting up things to get greeting for everyone", err)
	}

	go func() {
		for _, name := range names {
			bistream.Send(&pb.GreetRequest{
				FirstName: strings.ToUpper(name),
			})
			time.Sleep(time.Second * 1)
		}
		bistream.CloseSend()
	}()

	ch := make(chan string)

	go func() {
		for {
			res, err := bistream.Recv()

			if err == io.EOF {
				log.Println("everything read from bistream response from server")
				break
			} else if err != nil {
				log.Fatal("error in receiving message from bi directional stream")
			}
			log.Println("Response from server: ", res.Result)
		}

		ch <- "done"
	}()

	<-ch
	log.Println("All requests make and responses received for Greeting all ")
}
