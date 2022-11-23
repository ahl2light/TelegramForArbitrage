// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: api/proto/telegram_message.proto

package tMessage

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

// TelegramMessageSenderClient is the client API for TelegramMessageSender service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TelegramMessageSenderClient interface {
	SendMessageToTelegram(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
}

type telegramMessageSenderClient struct {
	cc grpc.ClientConnInterface
}

func NewTelegramMessageSenderClient(cc grpc.ClientConnInterface) TelegramMessageSenderClient {
	return &telegramMessageSenderClient{cc}
}

func (c *telegramMessageSenderClient) SendMessageToTelegram(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/tMessage.TelegramMessageSender/SendMessageToTelegram", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TelegramMessageSenderServer is the server API for TelegramMessageSender service.
// All implementations must embed UnimplementedTelegramMessageSenderServer
// for forward compatibility
type TelegramMessageSenderServer interface {
	SendMessageToTelegram(context.Context, *MessageRequest) (*MessageResponse, error)
	mustEmbedUnimplementedTelegramMessageSenderServer()
}

// UnimplementedTelegramMessageSenderServer must be embedded to have forward compatible implementations.
type UnimplementedTelegramMessageSenderServer struct {
}

func (UnimplementedTelegramMessageSenderServer) SendMessageToTelegram(context.Context, *MessageRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToTelegram not implemented")
}
func (UnimplementedTelegramMessageSenderServer) mustEmbedUnimplementedTelegramMessageSenderServer() {}

// UnsafeTelegramMessageSenderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TelegramMessageSenderServer will
// result in compilation errors.
type UnsafeTelegramMessageSenderServer interface {
	mustEmbedUnimplementedTelegramMessageSenderServer()
}

func RegisterTelegramMessageSenderServer(s grpc.ServiceRegistrar, srv TelegramMessageSenderServer) {
	s.RegisterService(&TelegramMessageSender_ServiceDesc, srv)
}

func _TelegramMessageSender_SendMessageToTelegram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelegramMessageSenderServer).SendMessageToTelegram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tMessage.TelegramMessageSender/SendMessageToTelegram",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelegramMessageSenderServer).SendMessageToTelegram(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TelegramMessageSender_ServiceDesc is the grpc.ServiceDesc for TelegramMessageSender service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TelegramMessageSender_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tMessage.TelegramMessageSender",
	HandlerType: (*TelegramMessageSenderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessageToTelegram",
			Handler:    _TelegramMessageSender_SendMessageToTelegram_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/telegram_message.proto",
}