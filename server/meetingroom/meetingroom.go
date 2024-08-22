package meetingroom

import (
	"context"
	"fmt"
	pb "hybridseatreservation/reservation"
	dbm "hybridseatreservation/server/dbManager"

	"google.golang.org/grpc/metadata"
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
		return nil, fmt.Errorf("failed to get metadata: %v", err)
	}
	username := md.Get("username")
	password := md.Get("password")

	if len(username) == 0 || len(password) == 0 {
		return nil, fmt.Errorf("missing credentials")

	}
	empid, err := dbm.CheckCredentials(username[0], password[0])
	if err != nil {
		return nil, err
	}
	fmt.Printf("recieved username:%s password: %s ", username[0], password[0])
	return &pb.AuthResponse{Employeeid: uint64(empid)}, nil

}

func (s *server) MeetingRoomReservation(ctx context.Context, req *pb.MrRequest) (ares *pb.MrResponse, err error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get metadata: %v", err)
	}
	username := md.Get("username")
	password := md.Get("password")

	if len(username) == 0 || len(password) == 0 {
		return nil, fmt.Errorf("missing credentials")

	}
	empid, err := dbm.CheckCredentials(username[0], password[0])
	if err != nil {
		return nil, err
	}
	bldNo := req.GetBlgNumber()
	floorNo := req.GetFloorNumber()
	mroom := req.GetMeetingRoom()
	//empid = req.GetEmployeeid()
	bookingdate := req.GetDate()
	startTime := req.GetStartTime()
	endTime := req.GetEndTime()

	result, err := dbm.CheckMeetingRoomAvailability(empid, bldNo, floorNo, mroom, bookingdate, startTime, endTime)
	fmt.Println(result)
	if err != nil {
		return nil, fmt.Errorf("failed to insert record")
	}
	if result == false {
		return nil, fmt.Errorf("failed to insert record")
	}

	return &pb.MrResponse{}, nil
}
