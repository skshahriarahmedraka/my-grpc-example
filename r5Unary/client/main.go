/**
 * @Author: sk shahriar ahemd raka <ahmed>
 * @Date:   1970-01-01T06:00:00+06:00
 * @Email:  skshahriarahmedraka@gmail.com
 * @Filename: main.go
 * @Last modified by:   ahmed
 * @Last modified time: 2021-08-28T22:58:34+06:00
 */

package main

import (
	"context"
	"fmt"
	"log"

	//"os"
	"client/messagepb"
	"time"

	"google.golang.org/grpc"
)

// func Panicking(msg string , err error){
// 	if err!= nil {
// 		log.Fatalln(msg ," ", err)
// 	}
// }

// func main () {
// 	conn , err := grpc.Dial(":50051",grpc.WithInsecure(),grpc.WithBlock())

// 	Panicking("error : ",err)
// 	defer conn.Close()

// 	c:=messagepb.NewConversationClient(conn)

// 	name:= "raka"
// 	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
// 	defer cancel()

// 	r ,err:= c.Speaking(ctx,&messagepb.SpeakRequest{MyRequest:name})

// 	Panicking("error ",err)
// 	log.Println("greeting :",r.GetMyResponse())

// }

func CheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	CheckError(err)
	defer conn.Close()
	//fmt.Printf("%T",conn)
	ClientConnection(conn)

}

func ClientConnection(conn *grpc.ClientConn) {

	c := messagepb.NewConversationClient(conn)
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Speaking(ctx, &messagepb.SpeakRequest{Client_Request: "raka"})

	CheckError(err)
	fmt.Println("server response TO client : ", r.GetServer_Response())
}
