package user

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func CreateUser() {
	fmt.Println("Hello World")

	elliot := &User{
		Username: "elliotName",
		Age:      22,
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("Mashalling error: ", err)
	}

	fmt.Println(data)

	newElliot := &User{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("UnMashalling error: ", err)
	}

	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetUsername())
}
