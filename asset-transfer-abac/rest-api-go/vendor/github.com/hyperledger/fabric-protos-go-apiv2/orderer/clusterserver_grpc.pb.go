// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: orderer/clusterserver.proto

package orderer

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
	ClusterNodeService_Step_FullMethodName = "/orderer.ClusterNodeService/Step"
)

// ClusterNodeServiceClient is the client API for ClusterNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterNodeServiceClient interface {
	// Step passes an implementation-specific message to another cluster member.
	Step(ctx context.Context, opts ...grpc.CallOption) (ClusterNodeService_StepClient, error)
}

type clusterNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterNodeServiceClient(cc grpc.ClientConnInterface) ClusterNodeServiceClient {
	return &clusterNodeServiceClient{cc}
}

func (c *clusterNodeServiceClient) Step(ctx context.Context, opts ...grpc.CallOption) (ClusterNodeService_StepClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClusterNodeService_ServiceDesc.Streams[0], ClusterNodeService_Step_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &clusterNodeServiceStepClient{stream}
	return x, nil
}

type ClusterNodeService_StepClient interface {
	Send(*ClusterNodeServiceStepRequest) error
	Recv() (*ClusterNodeServiceStepResponse, error)
	grpc.ClientStream
}

type clusterNodeServiceStepClient struct {
	grpc.ClientStream
}

func (x *clusterNodeServiceStepClient) Send(m *ClusterNodeServiceStepRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clusterNodeServiceStepClient) Recv() (*ClusterNodeServiceStepResponse, error) {
	m := new(ClusterNodeServiceStepResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClusterNodeServiceServer is the server API for ClusterNodeService service.
// All implementations should embed UnimplementedClusterNodeServiceServer
// for forward compatibility
type ClusterNodeServiceServer interface {
	// Step passes an implementation-specific message to another cluster member.
	Step(ClusterNodeService_StepServer) error
}

// UnimplementedClusterNodeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClusterNodeServiceServer struct {
}

func (UnimplementedClusterNodeServiceServer) Step(ClusterNodeService_StepServer) error {
	return status.Errorf(codes.Unimplemented, "method Step not implemented")
}

// UnsafeClusterNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterNodeServiceServer will
// result in compilation errors.
type UnsafeClusterNodeServiceServer interface {
	mustEmbedUnimplementedClusterNodeServiceServer()
}

func RegisterClusterNodeServiceServer(s grpc.ServiceRegistrar, srv ClusterNodeServiceServer) {
	s.RegisterService(&ClusterNodeService_ServiceDesc, srv)
}

func _ClusterNodeService_Step_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClusterNodeServiceServer).Step(&clusterNodeServiceStepServer{stream})
}

type ClusterNodeService_StepServer interface {
	Send(*ClusterNodeServiceStepResponse) error
	Recv() (*ClusterNodeServiceStepRequest, error)
	grpc.ServerStream
}

type clusterNodeServiceStepServer struct {
	grpc.ServerStream
}

func (x *clusterNodeServiceStepServer) Send(m *ClusterNodeServiceStepResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clusterNodeServiceStepServer) Recv() (*ClusterNodeServiceStepRequest, error) {
	m := new(ClusterNodeServiceStepRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClusterNodeService_ServiceDesc is the grpc.ServiceDesc for ClusterNodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterNodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orderer.ClusterNodeService",
	HandlerType: (*ClusterNodeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Step",
			Handler:       _ClusterNodeService_Step_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "orderer/clusterserver.proto",
}
