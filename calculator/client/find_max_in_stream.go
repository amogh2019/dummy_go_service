package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/amogh2019/dummy_go_service/calculator/proto"
)

func streamNumbersAndFindMax(clientInstance pb.CalculatorServiceClient, nums []int64) {

	bistream, err := clientInstance.MaxAmongAll(context.Background())
	if err != nil {
		log.Fatal("error in making find max among all call to server", err)
	}

	go func() {

		for i := 0; i < len(nums); i++ {
			err = bistream.Send(&pb.MaxAmongAllRequest{
				Number: nums[i],
			})
			if err != nil {
				log.Fatal("error while sending inputs to server for max: ", nums[i], err)
			}
			fmt.Println("sending ", i, " : ", nums[i])
			time.Sleep(time.Second * 1)
		}
		bistream.CloseSend()
	}()

	ch := make(chan struct{})
	go func() {
		for {
			resp, err := bistream.Recv()
			if err == io.EOF {
				log.Println("everything read from bistream response from server")
				break
			} else if err != nil {
				log.Fatal("error received in getting max from stream from server", err)
			}
			log.Println("the updated current max is :", resp.MaxNum)
		}

		close(ch)
	}()
	<-ch
}

func askForInputToStreamAndFindMax() {
	fmt.Println()
	fmt.Println("Enter comma separated numbers to get its stream them to server at interval of 1s and get max response : ")

}
