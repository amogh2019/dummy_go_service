package main

import (
	"io"
	"log"
	"strconv"

	pb "github.com/amogh2019/dummy_go_service/greet/proto"
)

func (s *Server) LongGreet(inputStream pb.GreetService_LongGreetServer) error {
	log.Println("Long Greet was invoked")

	msg := "\n"
	i := 1
	for {
		req, err := inputStream.Recv()

		if err == io.EOF {
			return inputStream.SendAndClose(&pb.GreetResponse{
				Result: msg,
			})
		} else if err != nil {
			log.Fatal("err in receving input", err)
			return err
		}

		log.Println("LongGreet : got message ", req)
		msg = msg + "Yolo! " + req.FirstName + " : " + strconv.Itoa(i) + "\n"
		i++

	}

}
