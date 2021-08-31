package main

import (
	"context"
	"fmt"
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

	req := &messagepb.LongRequest{
		RequestString : "sk shahriar ahmed raka",
	}
	reqStream ,err := c.Speaking(context.Background())
	CheckError(err)
	for i:=0 ;i<10 ; i++ {
		fmt.Println("sending req : ",req.RequestString)
		err := reqStream.Send(req)
		CheckError(err)

		time.Sleep(2*time.Second)
	}

	res,err:= reqStream.CloseAndRecv()
	CheckError(err)

	fmt.Println("client stream server response : ",res)






}