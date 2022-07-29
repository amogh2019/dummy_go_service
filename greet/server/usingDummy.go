package main

import (
	"fmt"

	gp "github.com/amogh2019/dummy_go_service/greet/proto"
)

func getDummy() *gp.DummyMessage {
	return &gp.DummyMessage{
		Id:       22,
		IsActive: true,
		Name:     "Dummy11",
		AltNames: []string{"yolo", "dumm"},
	}
}

func main() {
	fmt.Println(getDummy())
}
