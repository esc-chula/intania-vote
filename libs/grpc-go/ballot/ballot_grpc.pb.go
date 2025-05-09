// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: proto/ballot/ballot.proto

package ballot

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
	BallotService_CreateBallotProof_FullMethodName = "/ballot.BallotService/CreateBallotProof"
	BallotService_CreateBallot_FullMethodName      = "/ballot.BallotService/CreateBallot"
	BallotService_VerifyBallot_FullMethodName      = "/ballot.BallotService/VerifyBallot"
)

// BallotServiceClient is the client API for BallotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BallotServiceClient interface {
	CreateBallotProof(ctx context.Context, in *CreateBallotProofRequest, opts ...grpc.CallOption) (*CreateBallotProofResponse, error)
	CreateBallot(ctx context.Context, in *CreateBallotRequest, opts ...grpc.CallOption) (*CreateBallotResponse, error)
	VerifyBallot(ctx context.Context, in *VerifyBallotRequest, opts ...grpc.CallOption) (*VerifyBallotResponse, error)
}

type ballotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBallotServiceClient(cc grpc.ClientConnInterface) BallotServiceClient {
	return &ballotServiceClient{cc}
}

func (c *ballotServiceClient) CreateBallotProof(ctx context.Context, in *CreateBallotProofRequest, opts ...grpc.CallOption) (*CreateBallotProofResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBallotProofResponse)
	err := c.cc.Invoke(ctx, BallotService_CreateBallotProof_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ballotServiceClient) CreateBallot(ctx context.Context, in *CreateBallotRequest, opts ...grpc.CallOption) (*CreateBallotResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBallotResponse)
	err := c.cc.Invoke(ctx, BallotService_CreateBallot_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ballotServiceClient) VerifyBallot(ctx context.Context, in *VerifyBallotRequest, opts ...grpc.CallOption) (*VerifyBallotResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyBallotResponse)
	err := c.cc.Invoke(ctx, BallotService_VerifyBallot_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BallotServiceServer is the server API for BallotService service.
// All implementations must embed UnimplementedBallotServiceServer
// for forward compatibility.
type BallotServiceServer interface {
	CreateBallotProof(context.Context, *CreateBallotProofRequest) (*CreateBallotProofResponse, error)
	CreateBallot(context.Context, *CreateBallotRequest) (*CreateBallotResponse, error)
	VerifyBallot(context.Context, *VerifyBallotRequest) (*VerifyBallotResponse, error)
	mustEmbedUnimplementedBallotServiceServer()
}

// UnimplementedBallotServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBallotServiceServer struct{}

func (UnimplementedBallotServiceServer) CreateBallotProof(context.Context, *CreateBallotProofRequest) (*CreateBallotProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBallotProof not implemented")
}
func (UnimplementedBallotServiceServer) CreateBallot(context.Context, *CreateBallotRequest) (*CreateBallotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBallot not implemented")
}
func (UnimplementedBallotServiceServer) VerifyBallot(context.Context, *VerifyBallotRequest) (*VerifyBallotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyBallot not implemented")
}
func (UnimplementedBallotServiceServer) mustEmbedUnimplementedBallotServiceServer() {}
func (UnimplementedBallotServiceServer) testEmbeddedByValue()                       {}

// UnsafeBallotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BallotServiceServer will
// result in compilation errors.
type UnsafeBallotServiceServer interface {
	mustEmbedUnimplementedBallotServiceServer()
}

func RegisterBallotServiceServer(s grpc.ServiceRegistrar, srv BallotServiceServer) {
	// If the following call pancis, it indicates UnimplementedBallotServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BallotService_ServiceDesc, srv)
}

func _BallotService_CreateBallotProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBallotProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BallotServiceServer).CreateBallotProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BallotService_CreateBallotProof_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BallotServiceServer).CreateBallotProof(ctx, req.(*CreateBallotProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BallotService_CreateBallot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBallotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BallotServiceServer).CreateBallot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BallotService_CreateBallot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BallotServiceServer).CreateBallot(ctx, req.(*CreateBallotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BallotService_VerifyBallot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyBallotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BallotServiceServer).VerifyBallot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BallotService_VerifyBallot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BallotServiceServer).VerifyBallot(ctx, req.(*VerifyBallotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BallotService_ServiceDesc is the grpc.ServiceDesc for BallotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BallotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ballot.BallotService",
	HandlerType: (*BallotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBallotProof",
			Handler:    _BallotService_CreateBallotProof_Handler,
		},
		{
			MethodName: "CreateBallot",
			Handler:    _BallotService_CreateBallot_Handler,
		},
		{
			MethodName: "VerifyBallot",
			Handler:    _BallotService_VerifyBallot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ballot/ballot.proto",
}
