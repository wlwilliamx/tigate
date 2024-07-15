// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: pkg/messaging/proto/message.proto

package proto

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

// MessageCenterClient is the client API for MessageCenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageCenterClient interface {
	// The clients call this method to build a event channel from client to server.
	SendEvents(ctx context.Context, in *Message, opts ...grpc.CallOption) (MessageCenter_SendEventsClient, error)
	SendCommands(ctx context.Context, in *Message, opts ...grpc.CallOption) (MessageCenter_SendCommandsClient, error)
}

type messageCenterClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageCenterClient(cc grpc.ClientConnInterface) MessageCenterClient {
	return &messageCenterClient{cc}
}

func (c *messageCenterClient) SendEvents(ctx context.Context, in *Message, opts ...grpc.CallOption) (MessageCenter_SendEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessageCenter_ServiceDesc.Streams[0], "/proto.MessageCenter/sendEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageCenterSendEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessageCenter_SendEventsClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type messageCenterSendEventsClient struct {
	grpc.ClientStream
}

func (x *messageCenterSendEventsClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageCenterClient) SendCommands(ctx context.Context, in *Message, opts ...grpc.CallOption) (MessageCenter_SendCommandsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessageCenter_ServiceDesc.Streams[1], "/proto.MessageCenter/sendCommands", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageCenterSendCommandsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessageCenter_SendCommandsClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type messageCenterSendCommandsClient struct {
	grpc.ClientStream
}

func (x *messageCenterSendCommandsClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageCenterServer is the server API for MessageCenter service.
// All implementations must embed UnimplementedMessageCenterServer
// for forward compatibility
type MessageCenterServer interface {
	// The clients call this method to build a event channel from client to server.
	SendEvents(*Message, MessageCenter_SendEventsServer) error
	SendCommands(*Message, MessageCenter_SendCommandsServer) error
	mustEmbedUnimplementedMessageCenterServer()
}

// UnimplementedMessageCenterServer must be embedded to have forward compatible implementations.
type UnimplementedMessageCenterServer struct {
}

func (UnimplementedMessageCenterServer) SendEvents(*Message, MessageCenter_SendEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method SendEvents not implemented")
}
func (UnimplementedMessageCenterServer) SendCommands(*Message, MessageCenter_SendCommandsServer) error {
	return status.Errorf(codes.Unimplemented, "method SendCommands not implemented")
}
func (UnimplementedMessageCenterServer) mustEmbedUnimplementedMessageCenterServer() {}

// UnsafeMessageCenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageCenterServer will
// result in compilation errors.
type UnsafeMessageCenterServer interface {
	mustEmbedUnimplementedMessageCenterServer()
}

func RegisterMessageCenterServer(s grpc.ServiceRegistrar, srv MessageCenterServer) {
	s.RegisterService(&MessageCenter_ServiceDesc, srv)
}

func _MessageCenter_SendEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Message)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageCenterServer).SendEvents(m, &messageCenterSendEventsServer{stream})
}

type MessageCenter_SendEventsServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type messageCenterSendEventsServer struct {
	grpc.ServerStream
}

func (x *messageCenterSendEventsServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _MessageCenter_SendCommands_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Message)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageCenterServer).SendCommands(m, &messageCenterSendCommandsServer{stream})
}

type MessageCenter_SendCommandsServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type messageCenterSendCommandsServer struct {
	grpc.ServerStream
}

func (x *messageCenterSendCommandsServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

// MessageCenter_ServiceDesc is the grpc.ServiceDesc for MessageCenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageCenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MessageCenter",
	HandlerType: (*MessageCenterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "sendEvents",
			Handler:       _MessageCenter_SendEvents_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "sendCommands",
			Handler:       _MessageCenter_SendCommands_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/messaging/proto/message.proto",
}
