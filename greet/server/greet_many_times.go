package main

import (
	"fmt"
	"strconv"

	pb "github.com/amogh2019/dummy_go_service/greet/proto"
)

func (s *Server) GreetManyTimes(req *pb.GreetRequest, responseStream pb.GreetService_GreetManyTimesServer) error { // this is basically an implementation of the Greet service server we wrote in proto // in our server implementation // // GreetServiceServer is the server API for GreetService service.

	fmt.Println("Greet Many Times function was invoked", req)

	for i := 0; i < 9; i++ {
		msg := &pb.GreetResponse{
			Result: "Yolo! " + req.FirstName + " : " + strconv.Itoa(i),
		}
		responseStream.Send(msg)
	}

	return nil //nil error
}
