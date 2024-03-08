// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package vehicle

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

// VehicleServiceClient is the client API for VehicleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VehicleServiceClient interface {
	// ParkVehicle creates Vehicle record.
	ParkVehicle(ctx context.Context, in *ParkVehicleRequest, opts ...grpc.CallOption) (*ParkVehicleResponse, error)
	// UnParkVehicle creates Vehicle record.
	UnParkVehicle(ctx context.Context, in *UnParkVehicleRequest, opts ...grpc.CallOption) (*UnParkVehicleResponse, error)
}

type vehicleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVehicleServiceClient(cc grpc.ClientConnInterface) VehicleServiceClient {
	return &vehicleServiceClient{cc}
}

func (c *vehicleServiceClient) ParkVehicle(ctx context.Context, in *ParkVehicleRequest, opts ...grpc.CallOption) (*ParkVehicleResponse, error) {
	out := new(ParkVehicleResponse)
	err := c.cc.Invoke(ctx, "/vehicle.VehicleService/ParkVehicle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleServiceClient) UnParkVehicle(ctx context.Context, in *UnParkVehicleRequest, opts ...grpc.CallOption) (*UnParkVehicleResponse, error) {
	out := new(UnParkVehicleResponse)
	err := c.cc.Invoke(ctx, "/vehicle.VehicleService/UnParkVehicle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VehicleServiceServer is the server API for VehicleService service.
// All implementations must embed UnimplementedVehicleServiceServer
// for forward compatibility
type VehicleServiceServer interface {
	// ParkVehicle creates Vehicle record.
	ParkVehicle(context.Context, *ParkVehicleRequest) (*ParkVehicleResponse, error)
	// UnParkVehicle creates Vehicle record.
	UnParkVehicle(context.Context, *UnParkVehicleRequest) (*UnParkVehicleResponse, error)
	mustEmbedUnimplementedVehicleServiceServer()
}

// UnimplementedVehicleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVehicleServiceServer struct{}

func (UnimplementedVehicleServiceServer) ParkVehicle(context.Context, *ParkVehicleRequest) (*ParkVehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParkVehicle not implemented")
}

func (UnimplementedVehicleServiceServer) UnParkVehicle(context.Context, *UnParkVehicleRequest) (*UnParkVehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnParkVehicle not implemented")
}
func (UnimplementedVehicleServiceServer) mustEmbedUnimplementedVehicleServiceServer() {}

// UnsafeVehicleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VehicleServiceServer will
// result in compilation errors.
type UnsafeVehicleServiceServer interface {
	mustEmbedUnimplementedVehicleServiceServer()
}

func RegisterVehicleServiceServer(s grpc.ServiceRegistrar, srv VehicleServiceServer) {
	s.RegisterService(&VehicleService_ServiceDesc, srv)
}

func _VehicleService_ParkVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParkVehicleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServiceServer).ParkVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vehicle.VehicleService/ParkVehicle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServiceServer).ParkVehicle(ctx, req.(*ParkVehicleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleService_UnParkVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnParkVehicleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServiceServer).UnParkVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vehicle.VehicleService/UnParkVehicle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServiceServer).UnParkVehicle(ctx, req.(*UnParkVehicleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VehicleService_ServiceDesc is the grpc.ServiceDesc for VehicleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VehicleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vehicle.VehicleService",
	HandlerType: (*VehicleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParkVehicle",
			Handler:    _VehicleService_ParkVehicle_Handler,
		},
		{
			MethodName: "UnParkVehicle",
			Handler:    _VehicleService_UnParkVehicle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parking/gunk/api/v1/vehicle/all.proto",
}