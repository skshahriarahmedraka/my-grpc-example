package main

import (
	"context"
	"log"
	//"os"
	"r4/messagepb"
	"time"

	"google.golang.org/grpc"
)

func Panicking(msg string , err error){
	if err!= nil {
		log.Fatalln(msg ," ", err)
	}
}

func main () {
	conn , err := grpc.Dial(":50051",grpc.WithInsecure(),grpc.WithBlock())

	Panicking("error : ",err)
	defer conn.Close()

	c:=messagepb.NewConversationClient(conn)

	name:= "raka"
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	r ,err:= c.Speaking(ctx,&messagepb.SpeakRequest{MyRequest:name})

	Panicking("error ",err)
	log.Println("greeting :",r.GetMyResponse())




}