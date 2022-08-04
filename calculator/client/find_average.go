package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/amogh2019/dummy_go_service/calculator/proto"
)

func callDoAverage(clientInstance pb.CalculatorServiceClient, nums []int32) {

	messageInputStream, err := clientInstance.Average(context.Background())
	if err != nil {
		log.Fatal("error in making average call to server", err)
	}

	for i := 0; i < len(nums); i++ {
		err = messageInputStream.Send(&pb.AverageRequest{
			Number: nums[i],
		})
		if err != nil {
			log.Fatal("error while sending inputs to server for average: ", nums[i], err)
		}
		fmt.Println("sending ", i, " : ", nums[i])
	}
	resp, err := messageInputStream.CloseAndRecv()
	if err != nil {
		log.Fatal("error received in averaging from server", err)
	}
	log.Println("the average is :", resp.Result)

}

func askForInputToAveage() {
	fmt.Println()
	fmt.Println("Enter comma separated numbers to get its Average : ")

}
