/**
 * @Author: sk shahriar ahemd raka <ahmed>
 * @Date:   1970-01-01T06:00:00+06:00
 * @Email:  skshahriarahmedraka@gmail.com
 * @Filename: main.go
 * @Last modified by:   ahmed
 * @Last modified time: 2021-08-28T15:55:25+06:00
 */

package main

import (
	//"context"
	"fmt"
	//"log"
	"net"
	"strconv"
	"time"

	messagepb "server/messagepb"

	"google.golang.org/grpc"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

type server struct {
	//	messagepb.UnimplementedConversationServer //optional line

}

//
//func (s *server) Speaking(ctx context.Context, in *messagepb.SpeakRequest)(*messagepb.SpeakResponse,error){
//	log.Printf("Received : %v ",in.GetMyRequest())
//	return &messagepb.SpeakResponse{MyResponse : "hello "+in.GetMyRequest()},nil
//}

func (s *server) Speaking(req *messagepb.ClientRequest, srv messagepb.Conversation_SpeakingServer) error {
	fmt.Println("server recieved request : ", req)
	m1 := req.GetClientMessage()
	for i := 0; i < 10; i++ {
		result := "hello " + m1 + "number" + strconv.Itoa(i)
		res := &messagepb.ServerResponse{
			ServerMessage: result,
		}
		srv.Send(res)
		time.Sleep(2 * time.Second)
	}
	return nil
}

func main() {

	// lis ,err := net.Listen("tcp",":50051")
	// Panicking("error : ", err)

	// MyServer := grpc.NewServer()

	// messagepb.RegisterConversationServer(MyServer,&server{})

	// err = MyServer.Serve(lis)

	// CheckError("error : ",err)

	lis, err := net.Listen("tcp", ":50051")
	CheckError(err)

	myserver := grpc.NewServer()

	messagepb.RegisterConversationServer(myserver, &server{})
	

	fmt.Println("server listening localhost:50051  ...")
	err = myserver.Serve(lis)

	CheckError(err)

}
