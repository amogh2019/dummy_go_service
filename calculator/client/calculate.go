package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/amogh2019/dummy_go_service/calculator/proto"
)

func askForInput() {
	fmt.Println()
	fmt.Println("Enter operation/value and press enter")
	fmt.Println("operations supported [ + * AC = exit]")
}

func callCalculate(clientInstance pb.CalculatorServiceClient, action pb.CalculatorActions, val int64) int64 {

	response, err := clientInstance.Operate(context.Background(), &pb.CalculatorRequest{
		Action: action,
		Value:  val,
	})
	if err != nil {
		log.Fatal("error in making call to server", err)
	}
	return response.ResultantValue
}

func askForPress() {
	fmt.Println()
	fmt.Println("Press: \n 1 for Calculator \n 2 for prime factorization \n exit for terminate")
}
