package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/kapralovs/grpc-databus-service/databus"
	"github.com/kapralovs/grpc-databus-service/server"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &server.GRPCServer{}
	databus.RegisterDatabusServiceServer(s, srv)

	if os.Args[1] != "" {
		srv.Port = fmt.Sprintf(":%s", os.Args[1])
	} else {
		log.Fatal("port not configured")
	}

	l, err := net.Listen("tcp", srv.Port)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
