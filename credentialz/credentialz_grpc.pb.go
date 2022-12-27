// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: github.com/openconfig/gnsi/credentialz/credentialz.proto

package credentialz

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

// CredentialzClient is the client API for Credentialz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CredentialzClient interface {
	RotateAccountCredentials(ctx context.Context, opts ...grpc.CallOption) (Credentialz_RotateAccountCredentialsClient, error)
	RotateHostCredentials(ctx context.Context, opts ...grpc.CallOption) (Credentialz_RotateHostCredentialsClient, error)
	CanGenerateKey(ctx context.Context, in *CanGenerateKeyRequest, opts ...grpc.CallOption) (*CanGenerateKeyResponse, error)
}

type credentialzClient struct {
	cc grpc.ClientConnInterface
}

func NewCredentialzClient(cc grpc.ClientConnInterface) CredentialzClient {
	return &credentialzClient{cc}
}

func (c *credentialzClient) RotateAccountCredentials(ctx context.Context, opts ...grpc.CallOption) (Credentialz_RotateAccountCredentialsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Credentialz_ServiceDesc.Streams[0], "/gnsi.credentialz.Credentialz/RotateAccountCredentials", opts...)
	if err != nil {
		return nil, err
	}
	x := &credentialzRotateAccountCredentialsClient{stream}
	return x, nil
}

type Credentialz_RotateAccountCredentialsClient interface {
	Send(*RotateAccountCredentialsRequest) error
	Recv() (*RotateAccountCredentialsResponse, error)
	grpc.ClientStream
}

type credentialzRotateAccountCredentialsClient struct {
	grpc.ClientStream
}

func (x *credentialzRotateAccountCredentialsClient) Send(m *RotateAccountCredentialsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *credentialzRotateAccountCredentialsClient) Recv() (*RotateAccountCredentialsResponse, error) {
	m := new(RotateAccountCredentialsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *credentialzClient) RotateHostCredentials(ctx context.Context, opts ...grpc.CallOption) (Credentialz_RotateHostCredentialsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Credentialz_ServiceDesc.Streams[1], "/gnsi.credentialz.Credentialz/RotateHostCredentials", opts...)
	if err != nil {
		return nil, err
	}
	x := &credentialzRotateHostCredentialsClient{stream}
	return x, nil
}

type Credentialz_RotateHostCredentialsClient interface {
	Send(*RotateHostCredentialsRequest) error
	Recv() (*RotateHostCredentialsResponse, error)
	grpc.ClientStream
}

type credentialzRotateHostCredentialsClient struct {
	grpc.ClientStream
}

func (x *credentialzRotateHostCredentialsClient) Send(m *RotateHostCredentialsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *credentialzRotateHostCredentialsClient) Recv() (*RotateHostCredentialsResponse, error) {
	m := new(RotateHostCredentialsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *credentialzClient) CanGenerateKey(ctx context.Context, in *CanGenerateKeyRequest, opts ...grpc.CallOption) (*CanGenerateKeyResponse, error) {
	out := new(CanGenerateKeyResponse)
	err := c.cc.Invoke(ctx, "/gnsi.credentialz.Credentialz/CanGenerateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CredentialzServer is the server API for Credentialz service.
// All implementations must embed UnimplementedCredentialzServer
// for forward compatibility
type CredentialzServer interface {
	RotateAccountCredentials(Credentialz_RotateAccountCredentialsServer) error
	RotateHostCredentials(Credentialz_RotateHostCredentialsServer) error
	CanGenerateKey(context.Context, *CanGenerateKeyRequest) (*CanGenerateKeyResponse, error)
	mustEmbedUnimplementedCredentialzServer()
}

// UnimplementedCredentialzServer must be embedded to have forward compatible implementations.
type UnimplementedCredentialzServer struct {
}

func (UnimplementedCredentialzServer) RotateAccountCredentials(Credentialz_RotateAccountCredentialsServer) error {
	return status.Errorf(codes.Unimplemented, "method RotateAccountCredentials not implemented")
}
func (UnimplementedCredentialzServer) RotateHostCredentials(Credentialz_RotateHostCredentialsServer) error {
	return status.Errorf(codes.Unimplemented, "method RotateHostCredentials not implemented")
}
func (UnimplementedCredentialzServer) CanGenerateKey(context.Context, *CanGenerateKeyRequest) (*CanGenerateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CanGenerateKey not implemented")
}
func (UnimplementedCredentialzServer) mustEmbedUnimplementedCredentialzServer() {}

// UnsafeCredentialzServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CredentialzServer will
// result in compilation errors.
type UnsafeCredentialzServer interface {
	mustEmbedUnimplementedCredentialzServer()
}

func RegisterCredentialzServer(s grpc.ServiceRegistrar, srv CredentialzServer) {
	s.RegisterService(&Credentialz_ServiceDesc, srv)
}

func _Credentialz_RotateAccountCredentials_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CredentialzServer).RotateAccountCredentials(&credentialzRotateAccountCredentialsServer{stream})
}

type Credentialz_RotateAccountCredentialsServer interface {
	Send(*RotateAccountCredentialsResponse) error
	Recv() (*RotateAccountCredentialsRequest, error)
	grpc.ServerStream
}

type credentialzRotateAccountCredentialsServer struct {
	grpc.ServerStream
}

func (x *credentialzRotateAccountCredentialsServer) Send(m *RotateAccountCredentialsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *credentialzRotateAccountCredentialsServer) Recv() (*RotateAccountCredentialsRequest, error) {
	m := new(RotateAccountCredentialsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Credentialz_RotateHostCredentials_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CredentialzServer).RotateHostCredentials(&credentialzRotateHostCredentialsServer{stream})
}

type Credentialz_RotateHostCredentialsServer interface {
	Send(*RotateHostCredentialsResponse) error
	Recv() (*RotateHostCredentialsRequest, error)
	grpc.ServerStream
}

type credentialzRotateHostCredentialsServer struct {
	grpc.ServerStream
}

func (x *credentialzRotateHostCredentialsServer) Send(m *RotateHostCredentialsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *credentialzRotateHostCredentialsServer) Recv() (*RotateHostCredentialsRequest, error) {
	m := new(RotateHostCredentialsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Credentialz_CanGenerateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CanGenerateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialzServer).CanGenerateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnsi.credentialz.Credentialz/CanGenerateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialzServer).CanGenerateKey(ctx, req.(*CanGenerateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Credentialz_ServiceDesc is the grpc.ServiceDesc for Credentialz service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Credentialz_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnsi.credentialz.Credentialz",
	HandlerType: (*CredentialzServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CanGenerateKey",
			Handler:    _Credentialz_CanGenerateKey_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RotateAccountCredentials",
			Handler:       _Credentialz_RotateAccountCredentials_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RotateHostCredentials",
			Handler:       _Credentialz_RotateHostCredentials_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/openconfig/gnsi/credentialz/credentialz.proto",
}
