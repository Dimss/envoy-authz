// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package trav3alpha

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

// TraServiceClient is the client API for TraService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TraServiceClient interface {
	Create(ctx context.Context, in *TraServiceRequest, opts ...grpc.CallOption) (*TraServiceResponse, error)
	Update(ctx context.Context, in *TraServiceRequest, opts ...grpc.CallOption) (*TraServiceResponse, error)
	Retrieve(ctx context.Context, in *TraServiceRequest, opts ...grpc.CallOption) (*TraServiceResponse, error)
	Delete(ctx context.Context, in *TraServiceRequest, opts ...grpc.CallOption) (*TraServiceResponse, error)
	Subscribe(ctx context.Context, in *TraServiceRequest, opts ...grpc.CallOption) (TraService_SubscribeClient, error)
}

type traServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTraServiceClient(cc grpc.ClientConnInterface) TraServiceClient {
	return &traServiceClient{cc}
}

func (c *traServiceClient) Create(ctx context.Context, in *TraServiceRequest, opts ...grpc.CallOption) (*TraServiceResponse, error) {
	out := new(TraServiceResponse)
	err := c.cc.Invoke(ctx, "/envoy.extensions.filters.network.sip_proxy.tra.v3alpha.TraService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traServiceClient) Update(ctx context.Context, in *TraServiceRequest, opts ...grpc.CallOption) (*TraServiceResponse, error) {
	out := new(TraServiceResponse)
	err := c.cc.Invoke(ctx, "/envoy.extensions.filters.network.sip_proxy.tra.v3alpha.TraService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traServiceClient) Retrieve(ctx context.Context, in *TraServiceRequest, opts ...grpc.CallOption) (*TraServiceResponse, error) {
	out := new(TraServiceResponse)
	err := c.cc.Invoke(ctx, "/envoy.extensions.filters.network.sip_proxy.tra.v3alpha.TraService/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traServiceClient) Delete(ctx context.Context, in *TraServiceRequest, opts ...grpc.CallOption) (*TraServiceResponse, error) {
	out := new(TraServiceResponse)
	err := c.cc.Invoke(ctx, "/envoy.extensions.filters.network.sip_proxy.tra.v3alpha.TraService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traServiceClient) Subscribe(ctx context.Context, in *TraServiceRequest, opts ...grpc.CallOption) (TraService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &TraService_ServiceDesc.Streams[0], "/envoy.extensions.filters.network.sip_proxy.tra.v3alpha.TraService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &traServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TraService_SubscribeClient interface {
	Recv() (*TraServiceResponse, error)
	grpc.ClientStream
}

type traServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *traServiceSubscribeClient) Recv() (*TraServiceResponse, error) {
	m := new(TraServiceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TraServiceServer is the server API for TraService service.
// All implementations should embed UnimplementedTraServiceServer
// for forward compatibility
type TraServiceServer interface {
	Create(context.Context, *TraServiceRequest) (*TraServiceResponse, error)
	Update(context.Context, *TraServiceRequest) (*TraServiceResponse, error)
	Retrieve(context.Context, *TraServiceRequest) (*TraServiceResponse, error)
	Delete(context.Context, *TraServiceRequest) (*TraServiceResponse, error)
	Subscribe(*TraServiceRequest, TraService_SubscribeServer) error
}

// UnimplementedTraServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTraServiceServer struct {
}

func (UnimplementedTraServiceServer) Create(context.Context, *TraServiceRequest) (*TraServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTraServiceServer) Update(context.Context, *TraServiceRequest) (*TraServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTraServiceServer) Retrieve(context.Context, *TraServiceRequest) (*TraServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedTraServiceServer) Delete(context.Context, *TraServiceRequest) (*TraServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTraServiceServer) Subscribe(*TraServiceRequest, TraService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

// UnsafeTraServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TraServiceServer will
// result in compilation errors.
type UnsafeTraServiceServer interface {
	mustEmbedUnimplementedTraServiceServer()
}

func RegisterTraServiceServer(s grpc.ServiceRegistrar, srv TraServiceServer) {
	s.RegisterService(&TraService_ServiceDesc, srv)
}

func _TraService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.extensions.filters.network.sip_proxy.tra.v3alpha.TraService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraServiceServer).Create(ctx, req.(*TraServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.extensions.filters.network.sip_proxy.tra.v3alpha.TraService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraServiceServer).Update(ctx, req.(*TraServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraService_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraServiceServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.extensions.filters.network.sip_proxy.tra.v3alpha.TraService/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraServiceServer).Retrieve(ctx, req.(*TraServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.extensions.filters.network.sip_proxy.tra.v3alpha.TraService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraServiceServer).Delete(ctx, req.(*TraServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TraServiceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TraServiceServer).Subscribe(m, &traServiceSubscribeServer{stream})
}

type TraService_SubscribeServer interface {
	Send(*TraServiceResponse) error
	grpc.ServerStream
}

type traServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *traServiceSubscribeServer) Send(m *TraServiceResponse) error {
	return x.ServerStream.SendMsg(m)
}

// TraService_ServiceDesc is the grpc.ServiceDesc for TraService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TraService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.extensions.filters.network.sip_proxy.tra.v3alpha.TraService",
	HandlerType: (*TraServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TraService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TraService_Update_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _TraService_Retrieve_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TraService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _TraService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "contrib/envoy/extensions/filters/network/sip_proxy/tra/v3alpha/tra.proto",
}
