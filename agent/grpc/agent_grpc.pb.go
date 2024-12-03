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
	Agent_CreatePod_FullMethodName        = "/agent.Agent/CreatePod"
	Agent_UpdatePod_FullMethodName        = "/agent.Agent/UpdatePod"
	Agent_DeletePod_FullMethodName        = "/agent.Agent/DeletePod"
	Agent_RebootPod_FullMethodName        = "/agent.Agent/RebootPod"
	Agent_GetPod_FullMethodName           = "/agent.Agent/GetPod"
	Agent_GetPods_FullMethodName          = "/agent.Agent/GetPods"
	Agent_GetPodStatus_FullMethodName     = "/agent.Agent/GetPodStatus"
	Agent_GetResourceQuota_FullMethodName = "/agent.Agent/GetResourceQuota"
	Agent_CreateNode_FullMethodName       = "/agent.Agent/CreateNode"
	Agent_UpdateNode_FullMethodName       = "/agent.Agent/UpdateNode"
	Agent_DeleteNode_FullMethodName       = "/agent.Agent/DeleteNode"
	Agent_GetNodeGpuUsage_FullMethodName  = "/agent.Agent/GetNodeGpuUsage"
	Agent_CreateWsToken_FullMethodName    = "/agent.Agent/CreateWsToken"
)

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	CreatePod(ctx context.Context, in *PodRequest, opts ...grpc.CallOption) (*PodResponse, error)
	UpdatePod(ctx context.Context, in *PodRequest, opts ...grpc.CallOption) (*PodResponse, error)
	DeletePod(ctx context.Context, in *PodDeleteRequest, opts ...grpc.CallOption) (*PodResponse, error)
	RebootPod(ctx context.Context, in *PodRebootRequest, opts ...grpc.CallOption) (*PodResponse, error)
	GetPod(ctx context.Context, in *PodInfoRequest, opts ...grpc.CallOption) (*PodInfoResponse, error)
	GetPods(ctx context.Context, in *PodInfoRequest, opts ...grpc.CallOption) (*PodInfoResponse, error)
	GetPodStatus(ctx context.Context, in *PodInfoRequest, opts ...grpc.CallOption) (*PodInfoResponse, error)
	GetResourceQuota(ctx context.Context, in *QuotaRequest, opts ...grpc.CallOption) (*QuotaResponse, error)
	CreateNode(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*NodeResponse, error)
	UpdateNode(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*NodeResponse, error)
	DeleteNode(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*NodeResponse, error)
	GetNodeGpuUsage(ctx context.Context, in *NodeGpuUsageRequest, opts ...grpc.CallOption) (*NodeGpuUsageResponse, error)
	CreateWsToken(ctx context.Context, in *CreateWsTokenRequest, opts ...grpc.CallOption) (*CreateWsTokenResponse, error)
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

func (c *agentClient) RebootPod(ctx context.Context, in *PodRebootRequest, opts ...grpc.CallOption) (*PodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PodResponse)
	err := c.cc.Invoke(ctx, Agent_RebootPod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetPod(ctx context.Context, in *PodInfoRequest, opts ...grpc.CallOption) (*PodInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PodInfoResponse)
	err := c.cc.Invoke(ctx, Agent_GetPod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetPods(ctx context.Context, in *PodInfoRequest, opts ...grpc.CallOption) (*PodInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PodInfoResponse)
	err := c.cc.Invoke(ctx, Agent_GetPods_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetPodStatus(ctx context.Context, in *PodInfoRequest, opts ...grpc.CallOption) (*PodInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PodInfoResponse)
	err := c.cc.Invoke(ctx, Agent_GetPodStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetResourceQuota(ctx context.Context, in *QuotaRequest, opts ...grpc.CallOption) (*QuotaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QuotaResponse)
	err := c.cc.Invoke(ctx, Agent_GetResourceQuota_FullMethodName, in, out, cOpts...)
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

func (c *agentClient) GetNodeGpuUsage(ctx context.Context, in *NodeGpuUsageRequest, opts ...grpc.CallOption) (*NodeGpuUsageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NodeGpuUsageResponse)
	err := c.cc.Invoke(ctx, Agent_GetNodeGpuUsage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) CreateWsToken(ctx context.Context, in *CreateWsTokenRequest, opts ...grpc.CallOption) (*CreateWsTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateWsTokenResponse)
	err := c.cc.Invoke(ctx, Agent_CreateWsToken_FullMethodName, in, out, cOpts...)
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
	RebootPod(context.Context, *PodRebootRequest) (*PodResponse, error)
	GetPod(context.Context, *PodInfoRequest) (*PodInfoResponse, error)
	GetPods(context.Context, *PodInfoRequest) (*PodInfoResponse, error)
	GetPodStatus(context.Context, *PodInfoRequest) (*PodInfoResponse, error)
	GetResourceQuota(context.Context, *QuotaRequest) (*QuotaResponse, error)
	CreateNode(context.Context, *NodeRequest) (*NodeResponse, error)
	UpdateNode(context.Context, *NodeRequest) (*NodeResponse, error)
	DeleteNode(context.Context, *NodeRequest) (*NodeResponse, error)
	GetNodeGpuUsage(context.Context, *NodeGpuUsageRequest) (*NodeGpuUsageResponse, error)
	CreateWsToken(context.Context, *CreateWsTokenRequest) (*CreateWsTokenResponse, error)
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
func (UnimplementedAgentServer) RebootPod(context.Context, *PodRebootRequest) (*PodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RebootPod not implemented")
}
func (UnimplementedAgentServer) GetPod(context.Context, *PodInfoRequest) (*PodInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPod not implemented")
}
func (UnimplementedAgentServer) GetPods(context.Context, *PodInfoRequest) (*PodInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPods not implemented")
}
func (UnimplementedAgentServer) GetPodStatus(context.Context, *PodInfoRequest) (*PodInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPodStatus not implemented")
}
func (UnimplementedAgentServer) GetResourceQuota(context.Context, *QuotaRequest) (*QuotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResourceQuota not implemented")
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
func (UnimplementedAgentServer) GetNodeGpuUsage(context.Context, *NodeGpuUsageRequest) (*NodeGpuUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeGpuUsage not implemented")
}
func (UnimplementedAgentServer) CreateWsToken(context.Context, *CreateWsTokenRequest) (*CreateWsTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWsToken not implemented")
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

func _Agent_RebootPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodRebootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RebootPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_RebootPod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RebootPod(ctx, req.(*PodRebootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetPod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetPod(ctx, req.(*PodInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetPods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetPods(ctx, req.(*PodInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetPodStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetPodStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetPodStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetPodStatus(ctx, req.(*PodInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetResourceQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetResourceQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetResourceQuota_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetResourceQuota(ctx, req.(*QuotaRequest))
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

func _Agent_GetNodeGpuUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGpuUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetNodeGpuUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetNodeGpuUsage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetNodeGpuUsage(ctx, req.(*NodeGpuUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_CreateWsToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWsTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CreateWsToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_CreateWsToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CreateWsToken(ctx, req.(*CreateWsTokenRequest))
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
			MethodName: "RebootPod",
			Handler:    _Agent_RebootPod_Handler,
		},
		{
			MethodName: "GetPod",
			Handler:    _Agent_GetPod_Handler,
		},
		{
			MethodName: "GetPods",
			Handler:    _Agent_GetPods_Handler,
		},
		{
			MethodName: "GetPodStatus",
			Handler:    _Agent_GetPodStatus_Handler,
		},
		{
			MethodName: "GetResourceQuota",
			Handler:    _Agent_GetResourceQuota_Handler,
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
		{
			MethodName: "GetNodeGpuUsage",
			Handler:    _Agent_GetNodeGpuUsage_Handler,
		},
		{
			MethodName: "CreateWsToken",
			Handler:    _Agent_CreateWsToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}
