package main

import (
	"io"
	"log"

	pb "github.com/amogh2019/dummy_go_service/calculator/proto"
)

func getMaxNum(a int64, b int64) int64 {

	if a > b {
		return a
	} else {
		return b
	}

}

func (serverImpl *CalculatorServiceServerImpl) MaxAmongAll(bistream pb.CalculatorService_MaxAmongAllServer) error {

	log.Println("operation received do find max among all ")

	maxNum := []int64{}

	for {
		msg, err := bistream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			log.Println("error in receiving input to find max :", err)
			return err
		}

		shouldPublishNewMax := true

		log.Println("received num from stream to find max from past data :", msg.Number)
		if len(maxNum) == 0 {
			maxNum = append(maxNum, msg.Number)
		} else {
			oldMax := maxNum[0]
			maxNum[0] = getMaxNum(maxNum[0], msg.Number)
			if oldMax == maxNum[0] {
				shouldPublishNewMax = false
			}
		}

		if shouldPublishNewMax {
			log.Println("the updated current max is :", maxNum[0])
			bistream.Send(&pb.MaxAmongAllResponse{
				MaxNum: maxNum[0],
			})
		}

	}

}
