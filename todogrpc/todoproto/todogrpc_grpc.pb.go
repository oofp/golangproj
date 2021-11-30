// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package todoproto

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

// TodoGrpcClient is the client API for TodoGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoGrpcClient interface {
	CreateNewToDo(ctx context.Context, in *NewToDo, opts ...grpc.CallOption) (*ToDo, error)
	GetAllToDos(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (TodoGrpc_GetAllToDosClient, error)
}

type todoGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoGrpcClient(cc grpc.ClientConnInterface) TodoGrpcClient {
	return &todoGrpcClient{cc}
}

func (c *todoGrpcClient) CreateNewToDo(ctx context.Context, in *NewToDo, opts ...grpc.CallOption) (*ToDo, error) {
	out := new(ToDo)
	err := c.cc.Invoke(ctx, "/todoproto.TodoGrpc/CreateNewToDo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoGrpcClient) GetAllToDos(ctx context.Context, in *NoParams, opts ...grpc.CallOption) (TodoGrpc_GetAllToDosClient, error) {
	stream, err := c.cc.NewStream(ctx, &TodoGrpc_ServiceDesc.Streams[0], "/todoproto.TodoGrpc/GetAllToDos", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoGrpcGetAllToDosClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TodoGrpc_GetAllToDosClient interface {
	Recv() (*ToDo, error)
	grpc.ClientStream
}

type todoGrpcGetAllToDosClient struct {
	grpc.ClientStream
}

func (x *todoGrpcGetAllToDosClient) Recv() (*ToDo, error) {
	m := new(ToDo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TodoGrpcServer is the server API for TodoGrpc service.
// All implementations must embed UnimplementedTodoGrpcServer
// for forward compatibility
type TodoGrpcServer interface {
	CreateNewToDo(context.Context, *NewToDo) (*ToDo, error)
	GetAllToDos(*NoParams, TodoGrpc_GetAllToDosServer) error
	mustEmbedUnimplementedTodoGrpcServer()
}

// UnimplementedTodoGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedTodoGrpcServer struct {
}

func (UnimplementedTodoGrpcServer) CreateNewToDo(context.Context, *NewToDo) (*ToDo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewToDo not implemented")
}
func (UnimplementedTodoGrpcServer) GetAllToDos(*NoParams, TodoGrpc_GetAllToDosServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllToDos not implemented")
}
func (UnimplementedTodoGrpcServer) mustEmbedUnimplementedTodoGrpcServer() {}

// UnsafeTodoGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoGrpcServer will
// result in compilation errors.
type UnsafeTodoGrpcServer interface {
	mustEmbedUnimplementedTodoGrpcServer()
}

func RegisterTodoGrpcServer(s grpc.ServiceRegistrar, srv TodoGrpcServer) {
	s.RegisterService(&TodoGrpc_ServiceDesc, srv)
}

func _TodoGrpc_CreateNewToDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewToDo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoGrpcServer).CreateNewToDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todoproto.TodoGrpc/CreateNewToDo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoGrpcServer).CreateNewToDo(ctx, req.(*NewToDo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoGrpc_GetAllToDos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoGrpcServer).GetAllToDos(m, &todoGrpcGetAllToDosServer{stream})
}

type TodoGrpc_GetAllToDosServer interface {
	Send(*ToDo) error
	grpc.ServerStream
}

type todoGrpcGetAllToDosServer struct {
	grpc.ServerStream
}

func (x *todoGrpcGetAllToDosServer) Send(m *ToDo) error {
	return x.ServerStream.SendMsg(m)
}

// TodoGrpc_ServiceDesc is the grpc.ServiceDesc for TodoGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todoproto.TodoGrpc",
	HandlerType: (*TodoGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewToDo",
			Handler:    _TodoGrpc_CreateNewToDo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllToDos",
			Handler:       _TodoGrpc_GetAllToDos_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "todoproto/todogrpc.proto",
}
