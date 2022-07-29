package main

import (
	"fmt"
	"io/ioutil"
	"log"

	gpl3 "github.com/amogh2019/dummy_go_service/greet/proto/practice/level3"
	"google.golang.org/protobuf/proto"
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

func getSample2() *gpl3.SampleMessage2 {
	return &gpl3.SampleMessage2{
		EyeColor: gpl3.EyeColor_EYECOLOR_BLUE,
	}
}
func getSample2WithOrdinalValue() *gpl3.SampleMessage2 {
	return &gpl3.SampleMessage2{
		EyeColor: 2, // works too // set the enum ordinal value
	}
}

func getOneOfs(result interface{}) {
	switch msgType := result.(type) {
	case *gpl3.Result_Id:
		fmt.Println("result is of type ID and the value is ", result.(*gpl3.Result_Id).Id)
	case *gpl3.Result_Message:
		fmt.Println("result is of type ID and the value is ", result.(*gpl3.Result_Message).Message)
	default:
		fmt.Println("message type cannot be infered", msgType)
	}
}

func getCountry() *gpl3.Country {
	return &gpl3.Country{
		StateWiseCityList: map[string]*gpl3.CityList{
			"ka": {Cities: []string{"blr", "mys"}},
			"mh": {Cities: []string{"mum", "pun"}},
		},
	}
}

func writeToFile(fileName string, message proto.Message) { // every proto message implements this interface
	msgBytes, err := proto.Marshal(message)
	if err != nil {
		log.Fatal("cannot serialize message into bytearray", err)
	}
	if err = ioutil.WriteFile(fileName, msgBytes, 0644); err != nil {
		log.Fatal("err in writing file", err)
	}
	fmt.Println("written to file", fileName)
}

func readAndCreateMessageFromFile(fileName string) *gpl3.SampleMessage {
	dataBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("cannot read file", fileName, err)
	}
	var msg gpl3.SampleMessage
	proto.Unmarshal(dataBytes, &msg)
	return &msg
}

func main() {
	fmt.Println(getSample1())
	fmt.Println(getComplex())
	fmt.Println(getSample2())
	fmt.Println(getSample2WithOrdinalValue())
	getOneOfs(&gpl3.Result_Id{Id: 332})
	getOneOfs(&gpl3.Result_Message{Message: "done!"})
	fmt.Println(getCountry())
	fileName := "greet/server/practice/messageInSerializedForm.bin"
	writeToFile(fileName, getSample1())
	fmt.Println(readAndCreateMessageFromFile(fileName))
}
