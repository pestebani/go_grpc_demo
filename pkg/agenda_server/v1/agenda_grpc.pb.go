// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: agenda.proto

package v1

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
	AgendaService_Ping_FullMethodName         = "/agenda.v1.AgendaService/Ping"
	AgendaService_CreateAgenda_FullMethodName = "/agenda.v1.AgendaService/CreateAgenda"
	AgendaService_GetAgenda_FullMethodName    = "/agenda.v1.AgendaService/GetAgenda"
	AgendaService_GetAgendas_FullMethodName   = "/agenda.v1.AgendaService/GetAgendas"
	AgendaService_UpdateAgenda_FullMethodName = "/agenda.v1.AgendaService/UpdateAgenda"
	AgendaService_DeleteAgenda_FullMethodName = "/agenda.v1.AgendaService/DeleteAgenda"
)

// AgendaServiceClient is the client API for AgendaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgendaServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	CreateAgenda(ctx context.Context, in *CreateAgendaRequest, opts ...grpc.CallOption) (*CreateAgendaResponse, error)
	GetAgenda(ctx context.Context, in *GetAgendaRequest, opts ...grpc.CallOption) (*GetAgendaResponse, error)
	GetAgendas(ctx context.Context, in *GetAgendasRequest, opts ...grpc.CallOption) (*GetAgendasResponse, error)
	UpdateAgenda(ctx context.Context, in *UpdateAgendaRequest, opts ...grpc.CallOption) (*UpdateAgendaResponse, error)
	DeleteAgenda(ctx context.Context, in *DeleteAgendaRequest, opts ...grpc.CallOption) (*DeleteAgendaResponse, error)
}

type agendaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgendaServiceClient(cc grpc.ClientConnInterface) AgendaServiceClient {
	return &agendaServiceClient{cc}
}

func (c *agendaServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, AgendaService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agendaServiceClient) CreateAgenda(ctx context.Context, in *CreateAgendaRequest, opts ...grpc.CallOption) (*CreateAgendaResponse, error) {
	out := new(CreateAgendaResponse)
	err := c.cc.Invoke(ctx, AgendaService_CreateAgenda_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agendaServiceClient) GetAgenda(ctx context.Context, in *GetAgendaRequest, opts ...grpc.CallOption) (*GetAgendaResponse, error) {
	out := new(GetAgendaResponse)
	err := c.cc.Invoke(ctx, AgendaService_GetAgenda_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agendaServiceClient) GetAgendas(ctx context.Context, in *GetAgendasRequest, opts ...grpc.CallOption) (*GetAgendasResponse, error) {
	out := new(GetAgendasResponse)
	err := c.cc.Invoke(ctx, AgendaService_GetAgendas_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agendaServiceClient) UpdateAgenda(ctx context.Context, in *UpdateAgendaRequest, opts ...grpc.CallOption) (*UpdateAgendaResponse, error) {
	out := new(UpdateAgendaResponse)
	err := c.cc.Invoke(ctx, AgendaService_UpdateAgenda_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agendaServiceClient) DeleteAgenda(ctx context.Context, in *DeleteAgendaRequest, opts ...grpc.CallOption) (*DeleteAgendaResponse, error) {
	out := new(DeleteAgendaResponse)
	err := c.cc.Invoke(ctx, AgendaService_DeleteAgenda_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgendaServiceServer is the server API for AgendaService service.
// All implementations must embed UnimplementedAgendaServiceServer
// for forward compatibility
type AgendaServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	CreateAgenda(context.Context, *CreateAgendaRequest) (*CreateAgendaResponse, error)
	GetAgenda(context.Context, *GetAgendaRequest) (*GetAgendaResponse, error)
	GetAgendas(context.Context, *GetAgendasRequest) (*GetAgendasResponse, error)
	UpdateAgenda(context.Context, *UpdateAgendaRequest) (*UpdateAgendaResponse, error)
	DeleteAgenda(context.Context, *DeleteAgendaRequest) (*DeleteAgendaResponse, error)
	mustEmbedUnimplementedAgendaServiceServer()
}

// UnimplementedAgendaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgendaServiceServer struct {
}

func (UnimplementedAgendaServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAgendaServiceServer) CreateAgenda(context.Context, *CreateAgendaRequest) (*CreateAgendaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAgenda not implemented")
}
func (UnimplementedAgendaServiceServer) GetAgenda(context.Context, *GetAgendaRequest) (*GetAgendaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgenda not implemented")
}
func (UnimplementedAgendaServiceServer) GetAgendas(context.Context, *GetAgendasRequest) (*GetAgendasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgendas not implemented")
}
func (UnimplementedAgendaServiceServer) UpdateAgenda(context.Context, *UpdateAgendaRequest) (*UpdateAgendaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAgenda not implemented")
}
func (UnimplementedAgendaServiceServer) DeleteAgenda(context.Context, *DeleteAgendaRequest) (*DeleteAgendaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAgenda not implemented")
}
func (UnimplementedAgendaServiceServer) mustEmbedUnimplementedAgendaServiceServer() {}

// UnsafeAgendaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgendaServiceServer will
// result in compilation errors.
type UnsafeAgendaServiceServer interface {
	mustEmbedUnimplementedAgendaServiceServer()
}

func RegisterAgendaServiceServer(s grpc.ServiceRegistrar, srv AgendaServiceServer) {
	s.RegisterService(&AgendaService_ServiceDesc, srv)
}

func _AgendaService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgendaServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgendaService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgendaServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgendaService_CreateAgenda_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAgendaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgendaServiceServer).CreateAgenda(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgendaService_CreateAgenda_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgendaServiceServer).CreateAgenda(ctx, req.(*CreateAgendaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgendaService_GetAgenda_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgendaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgendaServiceServer).GetAgenda(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgendaService_GetAgenda_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgendaServiceServer).GetAgenda(ctx, req.(*GetAgendaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgendaService_GetAgendas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgendasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgendaServiceServer).GetAgendas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgendaService_GetAgendas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgendaServiceServer).GetAgendas(ctx, req.(*GetAgendasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgendaService_UpdateAgenda_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAgendaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgendaServiceServer).UpdateAgenda(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgendaService_UpdateAgenda_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgendaServiceServer).UpdateAgenda(ctx, req.(*UpdateAgendaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgendaService_DeleteAgenda_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAgendaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgendaServiceServer).DeleteAgenda(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgendaService_DeleteAgenda_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgendaServiceServer).DeleteAgenda(ctx, req.(*DeleteAgendaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgendaService_ServiceDesc is the grpc.ServiceDesc for AgendaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgendaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agenda.v1.AgendaService",
	HandlerType: (*AgendaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _AgendaService_Ping_Handler,
		},
		{
			MethodName: "CreateAgenda",
			Handler:    _AgendaService_CreateAgenda_Handler,
		},
		{
			MethodName: "GetAgenda",
			Handler:    _AgendaService_GetAgenda_Handler,
		},
		{
			MethodName: "GetAgendas",
			Handler:    _AgendaService_GetAgendas_Handler,
		},
		{
			MethodName: "UpdateAgenda",
			Handler:    _AgendaService_UpdateAgenda_Handler,
		},
		{
			MethodName: "DeleteAgenda",
			Handler:    _AgendaService_DeleteAgenda_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agenda.proto",
}
