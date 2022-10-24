package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"

	"github.com/habuvo/testserverone/pkg/person"
	"github.com/habuvo/testserverone/proto"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 8071, "The port number for ServerOne")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	personServer := grpc.NewServer([]grpc.ServerOption{}...)

	proto.RegisterPersonServiceServer(personServer, person.New())
	log.Println("listen to port ", port)
	personServer.Serve(lis)

}
