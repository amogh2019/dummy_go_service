package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/amogh2019/dummy_go_service/calculator/proto"
)

func callPrimeFactorize(clientInstance pb.CalculatorServiceClient, val int32) {

	streamResponse, err := clientInstance.PrimeFactorize(context.Background(), &pb.PrimeFactorizationRequeset{
		Number: val,
	})
	if err != nil {
		log.Fatal("error in making call to server", err)
	}

	fmt.Println("Prime factors of", val, "are : ")
	for i := 0; ; {
		resp, err := streamResponse.Recv()
		if err == io.EOF {
			fmt.Println("everything read from server response")
			break
		} else if err != nil {
			log.Fatal("error while reading received stream response", err)
		}
		i++
		fmt.Println("factor ", i, " : ", resp.Factor)
	}

}

func askForInputPrimeFactorize() {
	fmt.Println()
	fmt.Println("Enter number to get its prime factors : ")

}
