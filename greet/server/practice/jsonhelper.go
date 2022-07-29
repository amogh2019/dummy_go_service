package main

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func toJSON(pb proto.Message) string {

	msgBytes, err := protojson.Marshal(pb) // note : this is not proto.marshal // proto.marshal will serialize in the optimal proto serialized structure using low memory and ordered fields as per schema
	if err != nil {
		log.Fatal(err)
	}
	return string(msgBytes)
}

func fromJSON(jsonString string, pb proto.Message) {
	err := protojson.Unmarshal([]byte(jsonString), pb)
	if err != nil {
		log.Fatal(err)
	}
}
