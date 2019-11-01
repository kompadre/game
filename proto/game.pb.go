// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SessionRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionRequest) Reset()         { *m = SessionRequest{} }
func (m *SessionRequest) String() string { return proto.CompactTextString(m) }
func (*SessionRequest) ProtoMessage()    {}
func (*SessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_b123c129239c4919, []int{0}
}
func (m *SessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionRequest.Unmarshal(m, b)
}
func (m *SessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionRequest.Marshal(b, m, deterministic)
}
func (dst *SessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionRequest.Merge(dst, src)
}
func (m *SessionRequest) XXX_Size() int {
	return xxx_messageInfo_SessionRequest.Size(m)
}
func (m *SessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SessionRequest proto.InternalMessageInfo

func (m *SessionRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SessionRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SessionGrant struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionGrant) Reset()         { *m = SessionGrant{} }
func (m *SessionGrant) String() string { return proto.CompactTextString(m) }
func (*SessionGrant) ProtoMessage()    {}
func (*SessionGrant) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_b123c129239c4919, []int{1}
}
func (m *SessionGrant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionGrant.Unmarshal(m, b)
}
func (m *SessionGrant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionGrant.Marshal(b, m, deterministic)
}
func (dst *SessionGrant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionGrant.Merge(dst, src)
}
func (m *SessionGrant) XXX_Size() int {
	return xxx_messageInfo_SessionGrant.Size(m)
}
func (m *SessionGrant) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionGrant.DiscardUnknown(m)
}

var xxx_messageInfo_SessionGrant proto.InternalMessageInfo

func (m *SessionGrant) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *SessionGrant) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type LookAroundRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookAroundRequest) Reset()         { *m = LookAroundRequest{} }
func (m *LookAroundRequest) String() string { return proto.CompactTextString(m) }
func (*LookAroundRequest) ProtoMessage()    {}
func (*LookAroundRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_b123c129239c4919, []int{2}
}
func (m *LookAroundRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookAroundRequest.Unmarshal(m, b)
}
func (m *LookAroundRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookAroundRequest.Marshal(b, m, deterministic)
}
func (dst *LookAroundRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookAroundRequest.Merge(dst, src)
}
func (m *LookAroundRequest) XXX_Size() int {
	return xxx_messageInfo_LookAroundRequest.Size(m)
}
func (m *LookAroundRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LookAroundRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LookAroundRequest proto.InternalMessageInfo

type LookAroundAnswer struct {
	Results              []*Object `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LookAroundAnswer) Reset()         { *m = LookAroundAnswer{} }
func (m *LookAroundAnswer) String() string { return proto.CompactTextString(m) }
func (*LookAroundAnswer) ProtoMessage()    {}
func (*LookAroundAnswer) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_b123c129239c4919, []int{3}
}
func (m *LookAroundAnswer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookAroundAnswer.Unmarshal(m, b)
}
func (m *LookAroundAnswer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookAroundAnswer.Marshal(b, m, deterministic)
}
func (dst *LookAroundAnswer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookAroundAnswer.Merge(dst, src)
}
func (m *LookAroundAnswer) XXX_Size() int {
	return xxx_messageInfo_LookAroundAnswer.Size(m)
}
func (m *LookAroundAnswer) XXX_DiscardUnknown() {
	xxx_messageInfo_LookAroundAnswer.DiscardUnknown(m)
}

var xxx_messageInfo_LookAroundAnswer proto.InternalMessageInfo

func (m *LookAroundAnswer) GetResults() []*Object {
	if m != nil {
		return m.Results
	}
	return nil
}

type Object struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	X                    int32    `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_b123c129239c4919, []int{4}
}
func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (dst *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(dst, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Object) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Object) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func init() {
	proto.RegisterType((*SessionRequest)(nil), "gameproto.v1.SessionRequest")
	proto.RegisterType((*SessionGrant)(nil), "gameproto.v1.SessionGrant")
	proto.RegisterType((*LookAroundRequest)(nil), "gameproto.v1.LookAroundRequest")
	proto.RegisterType((*LookAroundAnswer)(nil), "gameproto.v1.LookAroundAnswer")
	proto.RegisterType((*Object)(nil), "gameproto.v1.Object")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SessionClient is the client API for Session service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionClient interface {
	NewSession(ctx context.Context, opts ...grpc.CallOption) (Session_NewSessionClient, error)
	LookAround(ctx context.Context, in *LookAroundRequest, opts ...grpc.CallOption) (*LookAroundAnswer, error)
}

type sessionClient struct {
	cc *grpc.ClientConn
}

func NewSessionClient(cc *grpc.ClientConn) SessionClient {
	return &sessionClient{cc}
}

func (c *sessionClient) NewSession(ctx context.Context, opts ...grpc.CallOption) (Session_NewSessionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Session_serviceDesc.Streams[0], "/gameproto.v1.Session/NewSession", opts...)
	if err != nil {
		return nil, err
	}
	x := &sessionNewSessionClient{stream}
	return x, nil
}

type Session_NewSessionClient interface {
	Send(*SessionRequest) error
	Recv() (*SessionGrant, error)
	grpc.ClientStream
}

type sessionNewSessionClient struct {
	grpc.ClientStream
}

func (x *sessionNewSessionClient) Send(m *SessionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sessionNewSessionClient) Recv() (*SessionGrant, error) {
	m := new(SessionGrant)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sessionClient) LookAround(ctx context.Context, in *LookAroundRequest, opts ...grpc.CallOption) (*LookAroundAnswer, error) {
	out := new(LookAroundAnswer)
	err := c.cc.Invoke(ctx, "/gameproto.v1.Session/LookAround", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServer is the server API for Session service.
type SessionServer interface {
	NewSession(Session_NewSessionServer) error
	LookAround(context.Context, *LookAroundRequest) (*LookAroundAnswer, error)
}

func RegisterSessionServer(s *grpc.Server, srv SessionServer) {
	s.RegisterService(&_Session_serviceDesc, srv)
}

func _Session_NewSession_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SessionServer).NewSession(&sessionNewSessionServer{stream})
}

type Session_NewSessionServer interface {
	Send(*SessionGrant) error
	Recv() (*SessionRequest, error)
	grpc.ServerStream
}

type sessionNewSessionServer struct {
	grpc.ServerStream
}

func (x *sessionNewSessionServer) Send(m *SessionGrant) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sessionNewSessionServer) Recv() (*SessionRequest, error) {
	m := new(SessionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Session_LookAround_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookAroundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).LookAround(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gameproto.v1.Session/LookAround",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).LookAround(ctx, req.(*LookAroundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Session_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gameproto.v1.Session",
	HandlerType: (*SessionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LookAround",
			Handler:    _Session_LookAround_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NewSession",
			Handler:       _Session_NewSession_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "game.proto",
}

func init() { proto.RegisterFile("game.proto", fileDescriptor_game_b123c129239c4919) }

var fileDescriptor_game_b123c129239c4919 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4e, 0xeb, 0x30,
	0x10, 0x85, 0xe5, 0xdb, 0xdb, 0x04, 0x86, 0x08, 0x81, 0x41, 0x28, 0x8a, 0x10, 0x54, 0x59, 0x65,
	0x15, 0x41, 0xd9, 0x20, 0x76, 0xed, 0x06, 0x84, 0xf8, 0x91, 0xc2, 0x8e, 0x9d, 0x4b, 0x46, 0xa8,
	0x40, 0xed, 0xe2, 0xb1, 0x49, 0xfb, 0x3e, 0x3c, 0x28, 0x8a, 0xe3, 0x50, 0x22, 0xca, 0x2a, 0xf3,
	0xcd, 0x99, 0x8c, 0xe7, 0x1c, 0x80, 0x67, 0x31, 0xc3, 0x7c, 0xae, 0x95, 0x51, 0x3c, 0xaa, 0x6b,
	0x57, 0xe6, 0x1f, 0xa7, 0xe9, 0x15, 0x6c, 0x3f, 0x20, 0xd1, 0x54, 0xc9, 0x02, 0xdf, 0x2d, 0x92,
	0xe1, 0x09, 0x6c, 0x58, 0x42, 0x2d, 0xc5, 0x0c, 0x63, 0x36, 0x60, 0xd9, 0x66, 0xf1, 0xcd, 0xb5,
	0x36, 0x17, 0x44, 0x95, 0xd2, 0x65, 0xfc, 0xaf, 0xd1, 0x5a, 0x4e, 0x2f, 0x20, 0xf2, 0x9b, 0x2e,
	0xb5, 0x90, 0x86, 0x73, 0xf8, 0x6f, 0xed, 0xb4, 0xf4, 0x3b, 0x5c, 0xcd, 0x0f, 0x20, 0xd0, 0x28,
	0x48, 0x49, 0xff, 0xb7, 0xa7, 0x74, 0x0f, 0x76, 0x6f, 0x94, 0x7a, 0x1d, 0x69, 0x65, 0x65, 0xe9,
	0x0f, 0x49, 0xc7, 0xb0, 0xb3, 0x6a, 0x8e, 0x24, 0x55, 0xa8, 0x79, 0x0e, 0xa1, 0x46, 0xb2, 0x6f,
	0x86, 0x62, 0x36, 0xe8, 0x65, 0x5b, 0xc3, 0xfd, 0xfc, 0xa7, 0x9d, 0xfc, 0x7e, 0xf2, 0x82, 0x4f,
	0xa6, 0x68, 0x87, 0xd2, 0x73, 0x08, 0x9a, 0xd6, 0xda, 0x73, 0x22, 0x60, 0x0b, 0x77, 0x49, 0xbf,
	0x60, 0x8b, 0x9a, 0x96, 0x71, 0xaf, 0xa1, 0xe5, 0xf0, 0x93, 0x41, 0xe8, 0xfd, 0xf0, 0x6b, 0x80,
	0x3b, 0xac, 0x5a, 0x3a, 0xec, 0x3e, 0xd9, 0x8d, 0x2f, 0x49, 0xd6, 0xaa, 0x2e, 0x92, 0x8c, 0x9d,
	0x30, 0x7e, 0x0b, 0xb0, 0x72, 0xc5, 0x8f, 0xbb, 0xd3, 0xbf, 0x42, 0x48, 0x8e, 0xfe, 0x1a, 0x68,
	0x02, 0x19, 0x87, 0x8f, 0x7d, 0x27, 0x4e, 0x02, 0xf7, 0x39, 0xfb, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0x98, 0x1e, 0x50, 0x76, 0xeb, 0x01, 0x00, 0x00,
}
