package main

import (
	"context"
	"fmt"
	"io"
	"log"
	//"os"
	"client/messagepb"
	//"time"

	"google.golang.org/grpc"
)

func CheckError( err error){
	if err!= nil {
		log.Fatalln( err)
	}
}

func main () {
	conn , err := grpc.Dial(":50051",grpc.WithInsecure(),grpc.WithBlock())

	CheckError(err)
	defer conn.Close()

	c:=messagepb.NewConversationClient(conn)

	//name:= "raka"
	//ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	//defer cancel()
	//
	//r ,err:= c.Speaking(ctx,&messagepb.SpeakRequest{MyRequest:name})
	//
	//CheckError(err)
	//log.Println("greeting :",r.GetMyResponse())

	fu(c)


}

func fu (c messagepb.ConversationClient){

	fmt.Println("start to do client listening  server streaming ...")

	req := &messagepb.ClientRequest{ClientMessage: "sk shahriar ahmed raka"}
	resStream ,err := c.Speaking(context.Background(),req)
	CheckError(err)
	for {
		msg ,err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err !=nil {
			log.Fatalln("error while reading stream ",err)
		}


		log.Println("response from server : ",msg.GetServerMessage())
	}











}