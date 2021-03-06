// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messagepb/message.proto

package messagepb

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

type LongRequest struct {
	RequestString        string   `protobuf:"bytes,1,opt,name=RequestString,proto3" json:"RequestString,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LongRequest) Reset()         { *m = LongRequest{} }
func (m *LongRequest) String() string { return proto.CompactTextString(m) }
func (*LongRequest) ProtoMessage()    {}
func (*LongRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d4f3249a385fa58, []int{0}
}

func (m *LongRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongRequest.Unmarshal(m, b)
}
func (m *LongRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongRequest.Marshal(b, m, deterministic)
}
func (m *LongRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongRequest.Merge(m, src)
}
func (m *LongRequest) XXX_Size() int {
	return xxx_messageInfo_LongRequest.Size(m)
}
func (m *LongRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LongRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LongRequest proto.InternalMessageInfo

func (m *LongRequest) GetRequestString() string {
	if m != nil {
		return m.RequestString
	}
	return ""
}

type ShortResponse struct {
	ResponseString       string   `protobuf:"bytes,1,opt,name=ResponseString,proto3" json:"ResponseString,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShortResponse) Reset()         { *m = ShortResponse{} }
func (m *ShortResponse) String() string { return proto.CompactTextString(m) }
func (*ShortResponse) ProtoMessage()    {}
func (*ShortResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d4f3249a385fa58, []int{1}
}

func (m *ShortResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShortResponse.Unmarshal(m, b)
}
func (m *ShortResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShortResponse.Marshal(b, m, deterministic)
}
func (m *ShortResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShortResponse.Merge(m, src)
}
func (m *ShortResponse) XXX_Size() int {
	return xxx_messageInfo_ShortResponse.Size(m)
}
func (m *ShortResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShortResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShortResponse proto.InternalMessageInfo

func (m *ShortResponse) GetResponseString() string {
	if m != nil {
		return m.ResponseString
	}
	return ""
}

func init() {
	proto.RegisterType((*LongRequest)(nil), "LongRequest")
	proto.RegisterType((*ShortResponse)(nil), "ShortResponse")
}

func init() { proto.RegisterFile("messagepb/message.proto", fileDescriptor_1d4f3249a385fa58) }

var fileDescriptor_1d4f3249a385fa58 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0x48, 0xd2, 0x87, 0xb2, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x95, 0x8c,
	0xb9, 0xb8, 0x7d, 0xf2, 0xf3, 0xd2, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x54, 0xb8,
	0x78, 0xa1, 0xcc, 0xe0, 0x92, 0xa2, 0xcc, 0xbc, 0x74, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x54, 0x41, 0x25, 0x73, 0x2e, 0xde, 0xe0, 0x8c, 0xfc, 0xa2, 0x92, 0xa0, 0xd4, 0xe2, 0x82, 0xfc,
	0xbc, 0xe2, 0x54, 0x21, 0x35, 0x2e, 0x3e, 0x18, 0x1b, 0x45, 0x1f, 0x9a, 0xa8, 0x91, 0x0d, 0x17,
	0x8f, 0x73, 0x7e, 0x5e, 0x59, 0x6a, 0x51, 0x71, 0x62, 0x49, 0x66, 0x7e, 0x9e, 0x90, 0x0e, 0x17,
	0x47, 0x70, 0x41, 0x6a, 0x62, 0x76, 0x66, 0x5e, 0xba, 0x10, 0x8f, 0x1e, 0x92, 0x43, 0xa4, 0xf8,
	0xf4, 0x50, 0x6c, 0x50, 0x62, 0xd0, 0x60, 0x74, 0xe2, 0x8e, 0xe2, 0x84, 0x7b, 0x23, 0x89, 0x0d,
	0xec, 0x7e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x37, 0xa6, 0xb1, 0xda, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConversationClient is the client API for Conversation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConversationClient interface {
	Speaking(ctx context.Context, opts ...grpc.CallOption) (Conversation_SpeakingClient, error)
}

type conversationClient struct {
	cc *grpc.ClientConn
}

func NewConversationClient(cc *grpc.ClientConn) ConversationClient {
	return &conversationClient{cc}
}

func (c *conversationClient) Speaking(ctx context.Context, opts ...grpc.CallOption) (Conversation_SpeakingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Conversation_serviceDesc.Streams[0], "/Conversation/Speaking", opts...)
	if err != nil {
		return nil, err
	}
	x := &conversationSpeakingClient{stream}
	return x, nil
}

type Conversation_SpeakingClient interface {
	Send(*LongRequest) error
	CloseAndRecv() (*ShortResponse, error)
	grpc.ClientStream
}

type conversationSpeakingClient struct {
	grpc.ClientStream
}

func (x *conversationSpeakingClient) Send(m *LongRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *conversationSpeakingClient) CloseAndRecv() (*ShortResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ShortResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConversationServer is the server API for Conversation service.
type ConversationServer interface {
	Speaking(Conversation_SpeakingServer) error
}

// UnimplementedConversationServer can be embedded to have forward compatible implementations.
type UnimplementedConversationServer struct {
}

func (*UnimplementedConversationServer) Speaking(srv Conversation_SpeakingServer) error {
	return status.Errorf(codes.Unimplemented, "method Speaking not implemented")
}

func RegisterConversationServer(s *grpc.Server, srv ConversationServer) {
	s.RegisterService(&_Conversation_serviceDesc, srv)
}

func _Conversation_Speaking_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConversationServer).Speaking(&conversationSpeakingServer{stream})
}

type Conversation_SpeakingServer interface {
	SendAndClose(*ShortResponse) error
	Recv() (*LongRequest, error)
	grpc.ServerStream
}

type conversationSpeakingServer struct {
	grpc.ServerStream
}

func (x *conversationSpeakingServer) SendAndClose(m *ShortResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *conversationSpeakingServer) Recv() (*LongRequest, error) {
	m := new(LongRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Conversation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Conversation",
	HandlerType: (*ConversationServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Speaking",
			Handler:       _Conversation_Speaking_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "messagepb/message.proto",
}
