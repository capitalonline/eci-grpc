// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: agent.proto

package agent

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Agent_CreatePod_FullMethodName  = "/agent.Agent/CreatePod"
	Agent_UpdatePod_FullMethodName  = "/agent.Agent/UpdatePod"
	Agent_DeletePod_FullMethodName  = "/agent.Agent/DeletePod"
	Agent_CreateNode_FullMethodName = "/agent.Agent/CreateNode"
	Agent_UpdateNode_FullMethodName = "/agent.Agent/UpdateNode"
	Agent_DeleteNode_FullMethodName = "/agent.Agent/DeleteNode"
)

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	CreatePod(ctx context.Context, in *PodRequest, opts ...grpc.CallOption) (*PodResponse, error)
	UpdatePod(ctx context.Context, in *PodRequest, opts ...grpc.CallOption) (*PodResponse, error)
	DeletePod(ctx context.Context, in *PodDeleteRequest, opts ...grpc.CallOption) (*PodResponse, error)
	CreateNode(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*NodeResponse, error)
	UpdateNode(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*NodeResponse, error)
	DeleteNode(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*NodeResponse, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) CreatePod(ctx context.Context, in *PodRequest, opts ...grpc.CallOption) (*PodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PodResponse)
	err := c.cc.Invoke(ctx, Agent_CreatePod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) UpdatePod(ctx context.Context, in *PodRequest, opts ...grpc.CallOption) (*PodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PodResponse)
	err := c.cc.Invoke(ctx, Agent_UpdatePod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) DeletePod(ctx context.Context, in *PodDeleteRequest, opts ...grpc.CallOption) (*PodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PodResponse)
	err := c.cc.Invoke(ctx, Agent_DeletePod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) CreateNode(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*NodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NodeResponse)
	err := c.cc.Invoke(ctx, Agent_CreateNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) UpdateNode(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*NodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NodeResponse)
	err := c.cc.Invoke(ctx, Agent_UpdateNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) DeleteNode(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*NodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NodeResponse)
	err := c.cc.Invoke(ctx, Agent_DeleteNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility.
type AgentServer interface {
	CreatePod(context.Context, *PodRequest) (*PodResponse, error)
	UpdatePod(context.Context, *PodRequest) (*PodResponse, error)
	DeletePod(context.Context, *PodDeleteRequest) (*PodResponse, error)
	CreateNode(context.Context, *NodeRequest) (*NodeResponse, error)
	UpdateNode(context.Context, *NodeRequest) (*NodeResponse, error)
	DeleteNode(context.Context, *NodeRequest) (*NodeResponse, error)
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAgentServer struct{}

func (UnimplementedAgentServer) CreatePod(context.Context, *PodRequest) (*PodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePod not implemented")
}
func (UnimplementedAgentServer) UpdatePod(context.Context, *PodRequest) (*PodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePod not implemented")
}
func (UnimplementedAgentServer) DeletePod(context.Context, *PodDeleteRequest) (*PodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePod not implemented")
}
func (UnimplementedAgentServer) CreateNode(context.Context, *NodeRequest) (*NodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNode not implemented")
}
func (UnimplementedAgentServer) UpdateNode(context.Context, *NodeRequest) (*NodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNode not implemented")
}
func (UnimplementedAgentServer) DeleteNode(context.Context, *NodeRequest) (*NodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNode not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}
func (UnimplementedAgentServer) testEmbeddedByValue()               {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s grpc.ServiceRegistrar, srv AgentServer) {
	// If the following call pancis, it indicates UnimplementedAgentServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Agent_ServiceDesc, srv)
}

func _Agent_CreatePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CreatePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_CreatePod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CreatePod(ctx, req.(*PodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_UpdatePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).UpdatePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_UpdatePod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).UpdatePod(ctx, req.(*PodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_DeletePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).DeletePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_DeletePod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).DeletePod(ctx, req.(*PodDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_CreateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CreateNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_CreateNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CreateNode(ctx, req.(*NodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_UpdateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).UpdateNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_UpdateNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).UpdateNode(ctx, req.(*NodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_DeleteNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).DeleteNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_DeleteNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).DeleteNode(ctx, req.(*NodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Agent_ServiceDesc is the grpc.ServiceDesc for Agent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Agent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agent.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePod",
			Handler:    _Agent_CreatePod_Handler,
		},
		{
			MethodName: "UpdatePod",
			Handler:    _Agent_UpdatePod_Handler,
		},
		{
			MethodName: "DeletePod",
			Handler:    _Agent_DeletePod_Handler,
		},
		{
			MethodName: "CreateNode",
			Handler:    _Agent_CreateNode_Handler,
		},
		{
			MethodName: "UpdateNode",
			Handler:    _Agent_UpdateNode_Handler,
		},
		{
			MethodName: "DeleteNode",
			Handler:    _Agent_DeleteNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}
