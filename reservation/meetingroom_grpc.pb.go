// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: reservation/meetingroom.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	HybridReservationService_Authenticate_FullMethodName           = "/reservation.HybridReservationService/Authenticate"
	HybridReservationService_MeetingRoomReservation_FullMethodName = "/reservation.HybridReservationService/MeetingRoomReservation"
)

// HybridReservationServiceClient is the client API for HybridReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HybridReservationServiceClient interface {
	Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	MeetingRoomReservation(ctx context.Context, in *MrRequest, opts ...grpc.CallOption) (*MrResponse, error)
}

type hybridReservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHybridReservationServiceClient(cc grpc.ClientConnInterface) HybridReservationServiceClient {
	return &hybridReservationServiceClient{cc}
}

func (c *hybridReservationServiceClient) Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, HybridReservationService_Authenticate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hybridReservationServiceClient) MeetingRoomReservation(ctx context.Context, in *MrRequest, opts ...grpc.CallOption) (*MrResponse, error) {
	out := new(MrResponse)
	err := c.cc.Invoke(ctx, HybridReservationService_MeetingRoomReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HybridReservationServiceServer is the server API for HybridReservationService service.
// All implementations must embed UnimplementedHybridReservationServiceServer
// for forward compatibility
type HybridReservationServiceServer interface {
	Authenticate(context.Context, *AuthRequest) (*AuthResponse, error)
	MeetingRoomReservation(context.Context, *MrRequest) (*MrResponse, error)
	mustEmbedUnimplementedHybridReservationServiceServer()
}

// UnimplementedHybridReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHybridReservationServiceServer struct {
}

func (UnimplementedHybridReservationServiceServer) Authenticate(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedHybridReservationServiceServer) MeetingRoomReservation(context.Context, *MrRequest) (*MrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MeetingRoomReservation not implemented")
}
func (UnimplementedHybridReservationServiceServer) mustEmbedUnimplementedHybridReservationServiceServer() {
}

// UnsafeHybridReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HybridReservationServiceServer will
// result in compilation errors.
type UnsafeHybridReservationServiceServer interface {
	mustEmbedUnimplementedHybridReservationServiceServer()
}

func RegisterHybridReservationServiceServer(s grpc.ServiceRegistrar, srv HybridReservationServiceServer) {
	s.RegisterService(&HybridReservationService_ServiceDesc, srv)
}

func _HybridReservationService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HybridReservationServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HybridReservationService_Authenticate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HybridReservationServiceServer).Authenticate(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HybridReservationService_MeetingRoomReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HybridReservationServiceServer).MeetingRoomReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HybridReservationService_MeetingRoomReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HybridReservationServiceServer).MeetingRoomReservation(ctx, req.(*MrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HybridReservationService_ServiceDesc is the grpc.ServiceDesc for HybridReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HybridReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reservation.HybridReservationService",
	HandlerType: (*HybridReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _HybridReservationService_Authenticate_Handler,
		},
		{
			MethodName: "MeetingRoomReservation",
			Handler:    _HybridReservationService_MeetingRoomReservation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation/meetingroom.proto",
}
