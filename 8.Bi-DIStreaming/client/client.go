package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	conn , err := grpc.Dial(":50052",grpc.WithInsecure(),grpc.WithBlock())

	CheckError(err)
	defer conn.Close()

	c:=messagepb.NewConversationClient(conn)
	fu(c)


}

func fu (c messagepb.ConversationClient){



fmt.Println("starting bi-di Streaming rpc ...")

	stream,err := c.Speaking(context.Background())
	CheckError(err )


	requests := messagepb.LongRequest{
		ClientString: "sk shahriar ahmed raka",
	}
	waitc := make(chan struct{})
	// we send a bunch of message to the client (go routine)
	go func(){
		//funtion to send a bunch of messages
		for i :=0 ;i <10 ;i++{
			fmt.Println("sending message : ",requests.ClientString)
			err =stream.Send(&requests)
			CheckError(err)
			time.Sleep(2*time.Second)
		}
		stream.CloseSend()
	}()

	// we receive bunch of message from server (go routine )
	go func() {
		// function to receive a bunch of message
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Fatalln("error io.EOF : ", err)
				break
			}
			if err != nil {
				log.Fatalln("err res, err := stream.Recv() : ", err)
				break
			}
			fmt.Println("received : ", res.GetServerString())
		}
		close(waitc)

	}()
	//block until everything is done
	<-waitc









}