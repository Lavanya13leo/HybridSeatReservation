package meetingroom

import (
	"context"
	"fmt"
	pb "hybridseatreservation/reservation"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedHybridReservationServiceServer
}

var sv *server

func init() {
	sv = &server{}
}

func Server() *server {
	return sv
}

func (s *server) Authenticate(ctx context.Context, areq *pb.AuthRequest) (ares *pb.AuthResponse, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("Failed to get metadata: %v", err)
	}
	username := md.Get("username")
	password := md.Get("password")

	if len(username) == 0 || len(password) == 0 {
		return nil, fmt.Errorf("Missing credentials")
	}

	fmt.Println("Recieved username:%s password: %s ", username, password)
	return &pb.AuthResponse{Employeeid: 1}, nil
}

func (s *server) MeetingRoomReservation(ctx context.Context, req *pb.MrRequest) (ares *pb.MrResponse, err error) {
	fmt.Println("bglnum = ", req.GetBlgNumber())
	fmt.Println("floor number = ", req.GetFloorNumber())
	fmt.Println("date= ", req.GetDate())
	fmt.Println("end = ", req.GetEndTime())
	fmt.Println("start = ", req.GetStartTime())
	return nil, status.Errorf(codes.Unimplemented, "method is implemented")
}
