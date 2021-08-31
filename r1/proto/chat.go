package chat

import (
	//"context"
	"log"

	"golang.org/x/net/context"
)

type Server struct {

}

func (s *Server) SayHello (ctx context.Context ,message *Message) (*Message ,error){
	log.Printf("Received message body from clint : %s \n" , message.Body)
	return &Message{Body : "Hello From Server ! "},nil
}