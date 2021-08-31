package main

import (
	"fmt"
	"log"
	"r2/greetpb"

	"google.golang.org/grpc"
)
func Panicking(err error) {
	if err != nil {
		log.Fatalln("error : ",err )
	}

}

func main (){
	fmt.Println("client is running ... ")

	cc,err := grpc.Dial("localhost:50051",grpc.WithInsecure())

	Panicking(err)
	defer cc.Close()

	c:= greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Created Client : %f \n",c)
	fmt.Println("Created Client : ",c)

}