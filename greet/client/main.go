package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/amogh2019/dummy_go_service/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {

	// TO CONNECT TO SERVER
	// 1. dialup to address
	// 2. since grpc uses ssl by default, we would need to pass default auth for the time being(till we dont setup some auth)

	tlsEnabled := true
	opts := []grpc.DialOption{}
	if tlsEnabled {
		cred, err := credentials.NewClientTLSFromFile("ssl/ca.crt", "")
		if err != nil {
			log.Fatal("error in loading CA trust certificate", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatal("error in starting dialup", addr, err) // TODO this is not breaking when server is down // check how to write correct code to block this when server is not up
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	scanner := bufio.NewScanner(os.Stdin)
	askForInput()
	askForInputForLongGreet()
	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" {
			fmt.Println("closing greetservice client")
			break
		}
		doGreetToServer(c, input)
		doGetGreetFromServerManyTimes(c, input)
		doLongGreet(c, strings.Split(input, ","))
		getGreetingsForEveryone(c, strings.Split(input, ","))

		doGreetWithDeadline(c, input, 5*time.Second)
		doGreetWithDeadline(c, input, 2*time.Second)

		askForInput()
		askForInputForLongGreet()

	}
	if scanner.Err() != nil {
		log.Fatal("error in taking input from console")
	}

}

// Do remember to do go run main.go greet.go // i.e. run both the files // else nil pointed // segmentation fault
// go run greet/client/*.go
