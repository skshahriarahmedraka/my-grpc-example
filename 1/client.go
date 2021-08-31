package main

import (
	//"context"
	"log"
	"r1/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)
func Panicking(err error) {
	if err != nil {
		log.Println("error : ",err)
	}
}

func main(){
	var conn *grpc.ClientConn 
	conn ,err := grpc.Dial(":9000",grpc.WithInsecure())
	Panicking(err)
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	message := chat.Message{
		Body : "Hello from Client" ,
	}

	response ,err := c.SayHello(context.Background(),&message)

	Panicking(err)
	log.Printf("Response from Server : %s ",response.Body)

}