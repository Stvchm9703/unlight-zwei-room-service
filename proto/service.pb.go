// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

// import "google/protobuf/empty.proto";

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4e, 0xbb, 0x40,
	0x10, 0xc6, 0xe9, 0x3f, 0x7f, 0x0d, 0x0e, 0xd6, 0xa6, 0x7b, 0xd0, 0xc8, 0x91, 0x07, 0x20, 0x8d,
	0x8d, 0xed, 0xc1, 0x8b, 0xa2, 0xc6, 0x98, 0x94, 0x84, 0x96, 0xf4, 0xd2, 0xdb, 0xba, 0x8c, 0xb8,
	0x51, 0x58, 0xdc, 0xdd, 0x9a, 0xf8, 0x0e, 0x3e, 0x89, 0x4f, 0x69, 0x16, 0xa8, 0x06, 0x68, 0x13,
	0xe3, 0x8d, 0xf9, 0xbe, 0xef, 0x37, 0xec, 0xec, 0x2c, 0xf4, 0x15, 0xca, 0x37, 0xce, 0xd0, 0x2f,
	0xa4, 0xd0, 0x82, 0xd8, 0xcb, 0xd9, 0x2a, 0x32, 0x5f, 0x6e, 0x3f, 0x43, 0xa5, 0x68, 0x5a, 0x1b,
	0xee, 0x21, 0x13, 0x59, 0x26, 0xf2, 0xaa, 0x3a, 0xfb, 0xf8, 0x0f, 0xce, 0x42, 0x88, 0x2c, 0xae,
	0x60, 0x32, 0x05, 0xb8, 0x96, 0x48, 0x35, 0x1a, 0x91, 0x9c, 0xf8, 0x9b, 0x2e, 0xbe, 0xa9, 0x6b,
	0x07, 0x5f, 0xdd, 0xa3, 0xa6, 0xe1, 0x59, 0xe4, 0x12, 0x9c, 0x3b, 0xd4, 0xa6, 0x98, 0x71, 0xa5,
	0xdb, 0x64, 0x8c, 0x54, 0xb2, 0x27, 0x43, 0x1e, 0x37, 0x0d, 0x13, 0x5e, 0xa0, 0x2a, 0x3c, 0x8b,
	0x4c, 0xbe, 0x3b, 0xdc, 0xe7, 0x8f, 0x82, 0x0c, 0x9b, 0x41, 0xc3, 0x92, 0xb6, 0x54, 0x72, 0xe7,
	0x00, 0x37, 0xf8, 0x82, 0xf5, 0x91, 0x7f, 0x8d, 0x4d, 0x01, 0x96, 0x45, 0xf2, 0x87, 0x49, 0x2f,
	0x60, 0x60, 0x6e, 0x0b, 0x65, 0x20, 0x05, 0x4d, 0x18, 0x55, 0x7a, 0xdb, 0x4f, 0x5b, 0x52, 0xa8,
	0x52, 0xcf, 0x1a, 0xf5, 0xc8, 0x18, 0x9c, 0x18, 0xf3, 0x24, 0xac, 0x56, 0x42, 0xba, 0x29, 0x77,
	0xf0, 0x23, 0xdd, 0x66, 0x85, 0x7e, 0xf7, 0x2c, 0x32, 0x02, 0x7b, 0xbe, 0xe6, 0x7a, 0xd7, 0x7c,
	0x5b, 0x88, 0x09, 0x1c, 0xcc, 0xd7, 0x9c, 0x3d, 0x47, 0x94, 0xcb, 0xdd, 0xbb, 0xe8, 0xcc, 0x16,
	0x04, 0x70, 0xca, 0x85, 0x9f, 0xca, 0x82, 0x19, 0xeb, 0x4a, 0x29, 0xd4, 0x7e, 0x19, 0x08, 0x45,
	0x12, 0xd8, 0x1b, 0x29, 0xea, 0xad, 0xf6, 0xca, 0xe7, 0xf3, 0xf9, 0x6f, 0xd8, 0x89, 0x3d, 0xec,
	0x97, 0xd6, 0xf8, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x6c, 0x6a, 0xe0, 0x91, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RoomServiceClient is the client API for RoomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoomServiceClient interface {
	CreateRoom(ctx context.Context, in *RoomCreateReq, opts ...grpc.CallOption) (*Room, error)
	GetRoomList(ctx context.Context, in *RoomSearchReq, opts ...grpc.CallOption) (*RoomListResp, error)
	GetRoomInfo(ctx context.Context, in *RoomReq, opts ...grpc.CallOption) (*RoomResp, error)
	DeleteRoom(ctx context.Context, in *RoomReq, opts ...grpc.CallOption) (*RoomResp, error)
	UpdateRoom(ctx context.Context, in *RoomCreateReq, opts ...grpc.CallOption) (*Room, error)
	ServerBroadcast(ctx context.Context, in *RoomReq, opts ...grpc.CallOption) (RoomService_ServerBroadcastClient, error)
	SendMessage(ctx context.Context, in *RoomMsg, opts ...grpc.CallOption) (*Empty, error)
	QuitRoom(ctx context.Context, in *RoomReq, opts ...grpc.CallOption) (*Empty, error)
	QuickPair(ctx context.Context, in *RoomSearchReq, opts ...grpc.CallOption) (*Room, error)
}

type roomServiceClient struct {
	cc *grpc.ClientConn
}

func NewRoomServiceClient(cc *grpc.ClientConn) RoomServiceClient {
	return &roomServiceClient{cc}
}

func (c *roomServiceClient) CreateRoom(ctx context.Context, in *RoomCreateReq, opts ...grpc.CallOption) (*Room, error) {
	out := new(Room)
	err := c.cc.Invoke(ctx, "/ULZProto.RoomService/CreateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) GetRoomList(ctx context.Context, in *RoomSearchReq, opts ...grpc.CallOption) (*RoomListResp, error) {
	out := new(RoomListResp)
	err := c.cc.Invoke(ctx, "/ULZProto.RoomService/GetRoomList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) GetRoomInfo(ctx context.Context, in *RoomReq, opts ...grpc.CallOption) (*RoomResp, error) {
	out := new(RoomResp)
	err := c.cc.Invoke(ctx, "/ULZProto.RoomService/GetRoomInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) DeleteRoom(ctx context.Context, in *RoomReq, opts ...grpc.CallOption) (*RoomResp, error) {
	out := new(RoomResp)
	err := c.cc.Invoke(ctx, "/ULZProto.RoomService/DeleteRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) UpdateRoom(ctx context.Context, in *RoomCreateReq, opts ...grpc.CallOption) (*Room, error) {
	out := new(Room)
	err := c.cc.Invoke(ctx, "/ULZProto.RoomService/UpdateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) ServerBroadcast(ctx context.Context, in *RoomReq, opts ...grpc.CallOption) (RoomService_ServerBroadcastClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RoomService_serviceDesc.Streams[0], "/ULZProto.RoomService/ServerBroadcast", opts...)
	if err != nil {
		return nil, err
	}
	x := &roomServiceServerBroadcastClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RoomService_ServerBroadcastClient interface {
	Recv() (*RoomMsg, error)
	grpc.ClientStream
}

type roomServiceServerBroadcastClient struct {
	grpc.ClientStream
}

func (x *roomServiceServerBroadcastClient) Recv() (*RoomMsg, error) {
	m := new(RoomMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *roomServiceClient) SendMessage(ctx context.Context, in *RoomMsg, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ULZProto.RoomService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) QuitRoom(ctx context.Context, in *RoomReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ULZProto.RoomService/QuitRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) QuickPair(ctx context.Context, in *RoomSearchReq, opts ...grpc.CallOption) (*Room, error) {
	out := new(Room)
	err := c.cc.Invoke(ctx, "/ULZProto.RoomService/QuickPair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomServiceServer is the server API for RoomService service.
type RoomServiceServer interface {
	CreateRoom(context.Context, *RoomCreateReq) (*Room, error)
	GetRoomList(context.Context, *RoomSearchReq) (*RoomListResp, error)
	GetRoomInfo(context.Context, *RoomReq) (*RoomResp, error)
	DeleteRoom(context.Context, *RoomReq) (*RoomResp, error)
	UpdateRoom(context.Context, *RoomCreateReq) (*Room, error)
	ServerBroadcast(*RoomReq, RoomService_ServerBroadcastServer) error
	SendMessage(context.Context, *RoomMsg) (*Empty, error)
	QuitRoom(context.Context, *RoomReq) (*Empty, error)
	QuickPair(context.Context, *RoomSearchReq) (*Room, error)
}

// UnimplementedRoomServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRoomServiceServer struct {
}

func (*UnimplementedRoomServiceServer) CreateRoom(ctx context.Context, req *RoomCreateReq) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (*UnimplementedRoomServiceServer) GetRoomList(ctx context.Context, req *RoomSearchReq) (*RoomListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomList not implemented")
}
func (*UnimplementedRoomServiceServer) GetRoomInfo(ctx context.Context, req *RoomReq) (*RoomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomInfo not implemented")
}
func (*UnimplementedRoomServiceServer) DeleteRoom(ctx context.Context, req *RoomReq) (*RoomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
}
func (*UnimplementedRoomServiceServer) UpdateRoom(ctx context.Context, req *RoomCreateReq) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoom not implemented")
}
func (*UnimplementedRoomServiceServer) ServerBroadcast(req *RoomReq, srv RoomService_ServerBroadcastServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerBroadcast not implemented")
}
func (*UnimplementedRoomServiceServer) SendMessage(ctx context.Context, req *RoomMsg) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (*UnimplementedRoomServiceServer) QuitRoom(ctx context.Context, req *RoomReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuitRoom not implemented")
}
func (*UnimplementedRoomServiceServer) QuickPair(ctx context.Context, req *RoomSearchReq) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuickPair not implemented")
}

func RegisterRoomServiceServer(s *grpc.Server, srv RoomServiceServer) {
	s.RegisterService(&_RoomService_serviceDesc, srv)
}

func _RoomService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ULZProto.RoomService/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).CreateRoom(ctx, req.(*RoomCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_GetRoomList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).GetRoomList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ULZProto.RoomService/GetRoomList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).GetRoomList(ctx, req.(*RoomSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_GetRoomInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).GetRoomInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ULZProto.RoomService/GetRoomInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).GetRoomInfo(ctx, req.(*RoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_DeleteRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).DeleteRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ULZProto.RoomService/DeleteRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).DeleteRoom(ctx, req.(*RoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_UpdateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).UpdateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ULZProto.RoomService/UpdateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).UpdateRoom(ctx, req.(*RoomCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_ServerBroadcast_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RoomReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RoomServiceServer).ServerBroadcast(m, &roomServiceServerBroadcastServer{stream})
}

type RoomService_ServerBroadcastServer interface {
	Send(*RoomMsg) error
	grpc.ServerStream
}

type roomServiceServerBroadcastServer struct {
	grpc.ServerStream
}

func (x *roomServiceServerBroadcastServer) Send(m *RoomMsg) error {
	return x.ServerStream.SendMsg(m)
}

func _RoomService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ULZProto.RoomService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).SendMessage(ctx, req.(*RoomMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_QuitRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).QuitRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ULZProto.RoomService/QuitRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).QuitRoom(ctx, req.(*RoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_QuickPair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).QuickPair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ULZProto.RoomService/QuickPair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).QuickPair(ctx, req.(*RoomSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoomService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ULZProto.RoomService",
	HandlerType: (*RoomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _RoomService_CreateRoom_Handler,
		},
		{
			MethodName: "GetRoomList",
			Handler:    _RoomService_GetRoomList_Handler,
		},
		{
			MethodName: "GetRoomInfo",
			Handler:    _RoomService_GetRoomInfo_Handler,
		},
		{
			MethodName: "DeleteRoom",
			Handler:    _RoomService_DeleteRoom_Handler,
		},
		{
			MethodName: "UpdateRoom",
			Handler:    _RoomService_UpdateRoom_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _RoomService_SendMessage_Handler,
		},
		{
			MethodName: "QuitRoom",
			Handler:    _RoomService_QuitRoom_Handler,
		},
		{
			MethodName: "QuickPair",
			Handler:    _RoomService_QuickPair_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerBroadcast",
			Handler:       _RoomService_ServerBroadcast_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
