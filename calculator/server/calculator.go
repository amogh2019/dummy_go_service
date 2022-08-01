package main

import (
	"context"
	"fmt"
	"sync"

	pb "github.com/amogh2019/dummy_go_service/calculator/proto"
)

type Calculator struct {
	lastValue int64
}

var cal = Calculator{
	lastValue: 0,
}

func (cal *Calculator) perfomOperation(req *pb.CalculatorRequest) (response int64) {
	switch req.Action {
	case pb.CalculatorActions_ADD:
		cal.lastValue += req.GetValue()
	case pb.CalculatorActions_MULTIPLY:
		cal.lastValue *= req.GetValue()
	case pb.CalculatorActions_CLEAR:
		cal.lastValue = 0
	case pb.CalculatorActions_UNDEFINED_ACTION:
		break
	}
	return cal.lastValue
}

func getCalculatorState() *Calculator {
	return &cal
}

func (serverImpl *CalculatorServiceServerImpl) Operate(ctx context.Context, req *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {

	fmt.Println("operation received : ", req)

	var m sync.Mutex
	m.Lock()
	res := getCalculatorState().perfomOperation(req)
	m.Unlock()

	return &pb.CalculatorResponse{
		ResultantValue: res,
	}, nil

}
