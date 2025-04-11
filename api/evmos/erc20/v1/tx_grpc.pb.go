// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/silc/silc/blob/main/LICENSE)

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: silc/erc20/v1/tx.proto

package erc20v1

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
	Msg_ConvertERC20_FullMethodName     = "/silc.erc20.v1.Msg/ConvertERC20"
	Msg_UpdateParams_FullMethodName     = "/silc.erc20.v1.Msg/UpdateParams"
	Msg_RegisterERC20_FullMethodName    = "/silc.erc20.v1.Msg/RegisterERC20"
	Msg_ToggleConversion_FullMethodName = "/silc.erc20.v1.Msg/ToggleConversion"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// ConvertERC20 mints a native Cosmos coin representation of the ERC20 token
	// contract that is registered on the token mapping.
	ConvertERC20(ctx context.Context, in *MsgConvertERC20, opts ...grpc.CallOption) (*MsgConvertERC20Response, error)
	// UpdateParams defines a governance operation for updating the x/erc20 module parameters.
	// The authority is hard-coded to the Cosmos SDK x/gov module account
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	// RegisterERC20 defines a governance operation for registering a token pair for the specified erc20 contract.
	// The authority is hard-coded to the Cosmos SDK x/gov module account
	RegisterERC20(ctx context.Context, in *MsgRegisterERC20, opts ...grpc.CallOption) (*MsgRegisterERC20Response, error)
	// ToggleConversion defines a governance operation for enabling/disablen a token pair conversion.
	// The authority is hard-coded to the Cosmos SDK x/gov module account
	ToggleConversion(ctx context.Context, in *MsgToggleConversion, opts ...grpc.CallOption) (*MsgToggleConversionResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ConvertERC20(ctx context.Context, in *MsgConvertERC20, opts ...grpc.CallOption) (*MsgConvertERC20Response, error) {
	out := new(MsgConvertERC20Response)
	err := c.cc.Invoke(ctx, Msg_ConvertERC20_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RegisterERC20(ctx context.Context, in *MsgRegisterERC20, opts ...grpc.CallOption) (*MsgRegisterERC20Response, error) {
	out := new(MsgRegisterERC20Response)
	err := c.cc.Invoke(ctx, Msg_RegisterERC20_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ToggleConversion(ctx context.Context, in *MsgToggleConversion, opts ...grpc.CallOption) (*MsgToggleConversionResponse, error) {
	out := new(MsgToggleConversionResponse)
	err := c.cc.Invoke(ctx, Msg_ToggleConversion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// ConvertERC20 mints a native Cosmos coin representation of the ERC20 token
	// contract that is registered on the token mapping.
	ConvertERC20(context.Context, *MsgConvertERC20) (*MsgConvertERC20Response, error)
	// UpdateParams defines a governance operation for updating the x/erc20 module parameters.
	// The authority is hard-coded to the Cosmos SDK x/gov module account
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	// RegisterERC20 defines a governance operation for registering a token pair for the specified erc20 contract.
	// The authority is hard-coded to the Cosmos SDK x/gov module account
	RegisterERC20(context.Context, *MsgRegisterERC20) (*MsgRegisterERC20Response, error)
	// ToggleConversion defines a governance operation for enabling/disablen a token pair conversion.
	// The authority is hard-coded to the Cosmos SDK x/gov module account
	ToggleConversion(context.Context, *MsgToggleConversion) (*MsgToggleConversionResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) ConvertERC20(context.Context, *MsgConvertERC20) (*MsgConvertERC20Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertERC20 not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) RegisterERC20(context.Context, *MsgRegisterERC20) (*MsgRegisterERC20Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterERC20 not implemented")
}
func (UnimplementedMsgServer) ToggleConversion(context.Context, *MsgToggleConversion) (*MsgToggleConversionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleConversion not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_ConvertERC20_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConvertERC20)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConvertERC20(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ConvertERC20_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConvertERC20(ctx, req.(*MsgConvertERC20))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RegisterERC20_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterERC20)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterERC20(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RegisterERC20_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterERC20(ctx, req.(*MsgRegisterERC20))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ToggleConversion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgToggleConversion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ToggleConversion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ToggleConversion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ToggleConversion(ctx, req.(*MsgToggleConversion))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "silc.erc20.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvertERC20",
			Handler:    _Msg_ConvertERC20_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "RegisterERC20",
			Handler:    _Msg_RegisterERC20_Handler,
		},
		{
			MethodName: "ToggleConversion",
			Handler:    _Msg_ToggleConversion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "silc/erc20/v1/tx.proto",
}
