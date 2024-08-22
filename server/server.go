package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "hybridseatreservation/reservation"
	mr "hybridseatreservation/server/meetingroom"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)



// SayHello implements helloworld.GreeterServer


func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHybridReservationServiceServer(s, mr.Server())
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}



}
