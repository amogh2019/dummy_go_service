package main

import (
	"io"
	"log"

	pb "github.com/amogh2019/dummy_go_service/calculator/proto"
)

func (serverImpl *CalculatorServiceServerImpl) Average(messageInputStream pb.CalculatorService_AverageServer) error {

	log.Println("operation received do average ")

	sum := 0
	total := 0
	for {
		req, err := messageInputStream.Recv()

		if err == io.EOF {
			var avg float64 = 0
			if total != 0 {
				avg = float64(sum) / float64(total)
			}
			log.Println("replying with  average :", avg)
			return messageInputStream.SendAndClose(&pb.AverageResponse{
				Result: avg,
			})
		} else if err != nil {
			log.Println("error in receiving input to average :", err)
			return err
		}

		log.Println("received to average :", req.Number)
		sum += int(req.Number)
		total++

	}

}
