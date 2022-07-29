package main

import (
	"fmt"

	gpl3 "github.com/amogh2019/dummy_go_service/greet/proto/practice/level3"
)

func getSample1() *gpl3.SampleMessage {
	return &gpl3.SampleMessage{
		Id:       20,
		IsActive: true,
		Name:     "Dummy21",
		AltNames: []string{"hoyee", "dumm", "sdfsadf"},
	}
}

func getComplex() *gpl3.ComplexMessage {
	return &gpl3.ComplexMessage{
		Single: getSample1(),
		Multiple: []*gpl3.SampleMessage{
			{Name: "abc"},
			{Name: "pqr"},
			getSample1(),
		},
	}
}

func main() {
	fmt.Println(getSample1())
	fmt.Println(getComplex())
}
