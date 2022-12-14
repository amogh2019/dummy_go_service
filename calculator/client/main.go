package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	pb "github.com/amogh2019/dummy_go_service/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serverAddress = "0.0.0.0:50052"

func main() {

	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error in starting dialup", serverAddress, err) // TODO this is not breaking when server is down // check how to write correct code to block this when server is not up
	}
	defer conn.Close()

	calculatorServiceClient := pb.NewCalculatorServiceClient(conn)

	scanner := bufio.NewScanner(os.Stdin)

	askForPress()
	shouldContd := true
	for shouldContd && scanner.Scan() {
		switch scanner.Text() {
		case "5":
			askForInputSquareRoot()
			if scanner.Scan() {
				input := scanner.Text()
				vall, _ := strconv.ParseInt(input, 10, 32)
				callSquareRoot(calculatorServiceClient, int32(vall))
				askForPress()
			} else {
				shouldContd = false
			}
		case "4":
			askForInputToStreamAndFindMax()
			if scanner.Scan() {
				input := scanner.Text()
				strnums := strings.Split(input, ",")
				nums := make([]int64, len(strnums))
				for i, strnum := range strnums {
					vall, _ := strconv.ParseInt(strnum, 10, 64)
					nums[i] = vall
				}
				streamNumbersAndFindMax(calculatorServiceClient, nums)
				askForPress()
			} else {
				shouldContd = false
			}
		case "3":
			askForInputToAveage()
			if scanner.Scan() {
				input := scanner.Text()
				strnums := strings.Split(input, ",")
				nums := make([]int32, len(strnums))
				for i, strnum := range strnums {
					vall, _ := strconv.ParseInt(strnum, 10, 32)
					nums[i] = int32(vall)
				}
				callDoAverage(calculatorServiceClient, nums)
				askForPress()
			} else {
				shouldContd = false
			}
		case "2":
			askForInputPrimeFactorize()
			if scanner.Scan() {
				input := scanner.Text()
				vall, _ := strconv.ParseInt(input, 10, 32)
				callPrimeFactorize(calculatorServiceClient, int32(vall))
				askForPress()
			} else {
				shouldContd = false
			}
		case "1":
			askForInput()
			temp := int64(0)
			contd := true
			for contd && scanner.Scan() {
				input := scanner.Text()
				switch input {
				case "exit":
					fmt.Println("closing calculator")
					askForPress()
					contd = false
				case "+":
					fmt.Println(callCalculate(calculatorServiceClient, pb.CalculatorActions_ADD, temp))
				case "=":
					fmt.Println(callCalculate(calculatorServiceClient, pb.CalculatorActions_UNDEFINED_ACTION, 0))
				case "*":
					fmt.Println(callCalculate(calculatorServiceClient, pb.CalculatorActions_MULTIPLY, temp))
				case "AC":
					fmt.Println(callCalculate(calculatorServiceClient, pb.CalculatorActions_CLEAR, temp))
				default:
					temp, _ = strconv.ParseInt(input, 10, 64)
				}
			}
		case "exit":
			fmt.Println("closing calculator client")
			shouldContd = false
		default:
			fmt.Println("dhang se daalo!")
			askForPress()

		}
	}

	if scanner.Err() != nil {
		log.Fatal("error in taking input from console")
	}

}
