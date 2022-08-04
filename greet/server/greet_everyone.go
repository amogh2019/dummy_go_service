package main

import (
	"io"
	"log"
	"strconv"

	pb "github.com/amogh2019/dummy_go_service/greet/proto"
)

func (s *Server) GreetEveryone(bistream pb.GreetService_GreetEveryoneServer) error {
	log.Println("Greet Everyone was invoked")

	msg := "\n"
	i := 1
	for {
		req, err := bistream.Recv()

		if err == io.EOF {
			return nil
		} else if err != nil {
			log.Fatal("err in receving input", err)
			return err
		}
		msg = "Yolo! " + req.FirstName + " : " + strconv.Itoa(i) + "\n"
		log.Println("GreetEveryone : going to greet :", req)
		err = bistream.Send(&pb.GreetResponse{
			Result: msg,
		})
		if err != nil {
			log.Fatal("err in sending response for input", err)
			return err
		}

		i++

	}

}
