// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: person.proto

package rpc

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

// PersonClient is the client API for Person service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonClient interface {
	GetPersons(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*PersonList, error)
}

type personClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonClient(cc grpc.ClientConnInterface) PersonClient {
	return &personClient{cc}
}

func (c *personClient) GetPersons(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*PersonList, error) {
	out := new(PersonList)
	err := c.cc.Invoke(ctx, "/Person/GetPersons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonServer is the server API for Person service.
// All implementations must embed UnimplementedPersonServer
// for forward compatibility
type PersonServer interface {
	GetPersons(context.Context, *PersonRequest) (*PersonList, error)
	mustEmbedUnimplementedPersonServer()
}

// UnimplementedPersonServer must be embedded to have forward compatible implementations.
type UnimplementedPersonServer struct {
}

func (UnimplementedPersonServer) GetPersons(context.Context, *PersonRequest) (*PersonList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersons not implemented")
}
func (UnimplementedPersonServer) mustEmbedUnimplementedPersonServer() {}

// UnsafePersonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonServer will
// result in compilation errors.
type UnsafePersonServer interface {
	mustEmbedUnimplementedPersonServer()
}

func RegisterPersonServer(s grpc.ServiceRegistrar, srv PersonServer) {
	s.RegisterService(&Person_ServiceDesc, srv)
}

func _Person_GetPersons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServer).GetPersons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Person/GetPersons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServer).GetPersons(ctx, req.(*PersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Person_ServiceDesc is the grpc.ServiceDesc for Person service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Person_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Person",
	HandlerType: (*PersonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPersons",
			Handler:    _Person_GetPersons_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "person.proto",
}
