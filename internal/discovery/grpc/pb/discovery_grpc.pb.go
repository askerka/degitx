// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// DiscoveryServiceClient is the client API for DiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscoveryServiceClient interface {
	// Ping - send ping request to discovery service to propagate self coordinates
	//  and receive known peers of service.
	Ping(ctx context.Context, in *NodeCoord, opts ...grpc.CallOption) (*PingResponse, error)
}

type discoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscoveryServiceClient(cc grpc.ClientConnInterface) DiscoveryServiceClient {
	return &discoveryServiceClient{cc}
}

func (c *discoveryServiceClient) Ping(ctx context.Context, in *NodeCoord, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/degitx.discovery.DiscoveryService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscoveryServiceServer is the server API for DiscoveryService service.
// All implementations must embed UnimplementedDiscoveryServiceServer
// for forward compatibility
type DiscoveryServiceServer interface {
	// Ping - send ping request to discovery service to propagate self coordinates
	//  and receive known peers of service.
	Ping(context.Context, *NodeCoord) (*PingResponse, error)
	mustEmbedUnimplementedDiscoveryServiceServer()
}

// UnimplementedDiscoveryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDiscoveryServiceServer struct {
}

func (UnimplementedDiscoveryServiceServer) Ping(context.Context, *NodeCoord) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedDiscoveryServiceServer) mustEmbedUnimplementedDiscoveryServiceServer() {}

// UnsafeDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscoveryServiceServer will
// result in compilation errors.
type UnsafeDiscoveryServiceServer interface {
	mustEmbedUnimplementedDiscoveryServiceServer()
}

func RegisterDiscoveryServiceServer(s grpc.ServiceRegistrar, srv DiscoveryServiceServer) {
	s.RegisterService(&DiscoveryService_ServiceDesc, srv)
}

func _DiscoveryService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeCoord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/degitx.discovery.DiscoveryService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceServer).Ping(ctx, req.(*NodeCoord))
	}
	return interceptor(ctx, in, info, handler)
}

// DiscoveryService_ServiceDesc is the grpc.ServiceDesc for DiscoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiscoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "degitx.discovery.DiscoveryService",
	HandlerType: (*DiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _DiscoveryService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discovery.proto",
}
