// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package parkinglot

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

// ParkingLotServiceClient is the client API for ParkingLotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParkingLotServiceClient interface {
	// CreateParkingLot creates ParkingLot record.
	CreateParkingLot(ctx context.Context, in *CreateParkingLotRequest, opts ...grpc.CallOption) (*CreateParkingLotResponse, error)
	// ListParkingLot gets all parking lot records.
	ListParkingLot(ctx context.Context, in *ListParkingLotRequest, opts ...grpc.CallOption) (*ListParkingLotResponse, error)
	// GetParkingLot gets all parking lot records.
	GetParkingLot(ctx context.Context, in *GetParkingLotRequest, opts ...grpc.CallOption) (*GetParkingLotResponse, error)
}

type parkingLotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewParkingLotServiceClient(cc grpc.ClientConnInterface) ParkingLotServiceClient {
	return &parkingLotServiceClient{cc}
}

func (c *parkingLotServiceClient) CreateParkingLot(ctx context.Context, in *CreateParkingLotRequest, opts ...grpc.CallOption) (*CreateParkingLotResponse, error) {
	out := new(CreateParkingLotResponse)
	err := c.cc.Invoke(ctx, "/parkinglot.ParkingLotService/CreateParkingLot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parkingLotServiceClient) ListParkingLot(ctx context.Context, in *ListParkingLotRequest, opts ...grpc.CallOption) (*ListParkingLotResponse, error) {
	out := new(ListParkingLotResponse)
	err := c.cc.Invoke(ctx, "/parkinglot.ParkingLotService/ListParkingLot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parkingLotServiceClient) GetParkingLot(ctx context.Context, in *GetParkingLotRequest, opts ...grpc.CallOption) (*GetParkingLotResponse, error) {
	out := new(GetParkingLotResponse)
	err := c.cc.Invoke(ctx, "/parkinglot.ParkingLotService/GetParkingLot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParkingLotServiceServer is the server API for ParkingLotService service.
// All implementations must embed UnimplementedParkingLotServiceServer
// for forward compatibility
type ParkingLotServiceServer interface {
	// CreateParkingLot creates ParkingLot record.
	CreateParkingLot(context.Context, *CreateParkingLotRequest) (*CreateParkingLotResponse, error)
	// ListParkingLot gets all parking lot records.
	ListParkingLot(context.Context, *ListParkingLotRequest) (*ListParkingLotResponse, error)
	// GetParkingLot gets all parking lot records.
	GetParkingLot(context.Context, *GetParkingLotRequest) (*GetParkingLotResponse, error)
	mustEmbedUnimplementedParkingLotServiceServer()
}

// UnimplementedParkingLotServiceServer must be embedded to have forward compatible implementations.
type UnimplementedParkingLotServiceServer struct{}

func (UnimplementedParkingLotServiceServer) CreateParkingLot(context.Context, *CreateParkingLotRequest) (*CreateParkingLotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateParkingLot not implemented")
}

func (UnimplementedParkingLotServiceServer) ListParkingLot(context.Context, *ListParkingLotRequest) (*ListParkingLotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListParkingLot not implemented")
}

func (UnimplementedParkingLotServiceServer) GetParkingLot(context.Context, *GetParkingLotRequest) (*GetParkingLotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParkingLot not implemented")
}
func (UnimplementedParkingLotServiceServer) mustEmbedUnimplementedParkingLotServiceServer() {}

// UnsafeParkingLotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParkingLotServiceServer will
// result in compilation errors.
type UnsafeParkingLotServiceServer interface {
	mustEmbedUnimplementedParkingLotServiceServer()
}

func RegisterParkingLotServiceServer(s grpc.ServiceRegistrar, srv ParkingLotServiceServer) {
	s.RegisterService(&ParkingLotService_ServiceDesc, srv)
}

func _ParkingLotService_CreateParkingLot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateParkingLotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParkingLotServiceServer).CreateParkingLot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parkinglot.ParkingLotService/CreateParkingLot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParkingLotServiceServer).CreateParkingLot(ctx, req.(*CreateParkingLotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParkingLotService_ListParkingLot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListParkingLotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParkingLotServiceServer).ListParkingLot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parkinglot.ParkingLotService/ListParkingLot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParkingLotServiceServer).ListParkingLot(ctx, req.(*ListParkingLotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParkingLotService_GetParkingLot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParkingLotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParkingLotServiceServer).GetParkingLot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parkinglot.ParkingLotService/GetParkingLot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParkingLotServiceServer).GetParkingLot(ctx, req.(*GetParkingLotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ParkingLotService_ServiceDesc is the grpc.ServiceDesc for ParkingLotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ParkingLotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "parkinglot.ParkingLotService",
	HandlerType: (*ParkingLotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateParkingLot",
			Handler:    _ParkingLotService_CreateParkingLot_Handler,
		},
		{
			MethodName: "ListParkingLot",
			Handler:    _ParkingLotService_ListParkingLot_Handler,
		},
		{
			MethodName: "GetParkingLot",
			Handler:    _ParkingLotService_GetParkingLot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parking/gunk/api/v1/parkinglot/all.proto",
}