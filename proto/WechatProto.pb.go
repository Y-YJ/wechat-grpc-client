// Code generated by protoc-gen-go. DO NOT EDIT.
// source: meta/WechatProto.proto

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

type WechatMsg struct {
	BaseMsg              *BaseMsg `protobuf:"bytes,1,opt,name=baseMsg,proto3" json:"baseMsg,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	TimeStamp            int32    `protobuf:"varint,4,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	Ip                   string   `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WechatMsg) Reset()         { *m = WechatMsg{} }
func (m *WechatMsg) String() string { return proto.CompactTextString(m) }
func (*WechatMsg) ProtoMessage()    {}
func (*WechatMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_WechatProto_c4b8821f4ce8a24f, []int{0}
}
func (m *WechatMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WechatMsg.Unmarshal(m, b)
}
func (m *WechatMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WechatMsg.Marshal(b, m, deterministic)
}
func (dst *WechatMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WechatMsg.Merge(dst, src)
}
func (m *WechatMsg) XXX_Size() int {
	return xxx_messageInfo_WechatMsg.Size(m)
}
func (m *WechatMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_WechatMsg.DiscardUnknown(m)
}

var xxx_messageInfo_WechatMsg proto.InternalMessageInfo

func (m *WechatMsg) GetBaseMsg() *BaseMsg {
	if m != nil {
		return m.BaseMsg
	}
	return nil
}

func (m *WechatMsg) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *WechatMsg) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *WechatMsg) GetTimeStamp() int32 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *WechatMsg) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type BaseMsg struct {
	Ret                  int32    `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Cmd                  int32    `protobuf:"varint,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
	CmdUrl               string   `protobuf:"bytes,3,opt,name=cmdUrl,proto3" json:"cmdUrl,omitempty"`
	ShortHost            string   `protobuf:"bytes,4,opt,name=shortHost,proto3" json:"shortHost,omitempty"`
	LongHost             string   `protobuf:"bytes,5,opt,name=longHost,proto3" json:"longHost,omitempty"`
	User                 *User    `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	LongHead             []byte   `protobuf:"bytes,7,opt,name=longHead,proto3" json:"longHead,omitempty"`
	Payloads             []byte   `protobuf:"bytes,8,opt,name=payloads,proto3" json:"payloads,omitempty"`
	Playloadextend       []byte   `protobuf:"bytes,9,opt,name=playloadextend,proto3" json:"playloadextend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseMsg) Reset()         { *m = BaseMsg{} }
func (m *BaseMsg) String() string { return proto.CompactTextString(m) }
func (*BaseMsg) ProtoMessage()    {}
func (*BaseMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_WechatProto_c4b8821f4ce8a24f, []int{1}
}
func (m *BaseMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseMsg.Unmarshal(m, b)
}
func (m *BaseMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseMsg.Marshal(b, m, deterministic)
}
func (dst *BaseMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseMsg.Merge(dst, src)
}
func (m *BaseMsg) XXX_Size() int {
	return xxx_messageInfo_BaseMsg.Size(m)
}
func (m *BaseMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseMsg.DiscardUnknown(m)
}

var xxx_messageInfo_BaseMsg proto.InternalMessageInfo

func (m *BaseMsg) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *BaseMsg) GetCmd() int32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *BaseMsg) GetCmdUrl() string {
	if m != nil {
		return m.CmdUrl
	}
	return ""
}

func (m *BaseMsg) GetShortHost() string {
	if m != nil {
		return m.ShortHost
	}
	return ""
}

func (m *BaseMsg) GetLongHost() string {
	if m != nil {
		return m.LongHost
	}
	return ""
}

func (m *BaseMsg) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *BaseMsg) GetLongHead() []byte {
	if m != nil {
		return m.LongHead
	}
	return nil
}

func (m *BaseMsg) GetPayloads() []byte {
	if m != nil {
		return m.Payloads
	}
	return nil
}

func (m *BaseMsg) GetPlayloadextend() []byte {
	if m != nil {
		return m.Playloadextend
	}
	return nil
}

type User struct {
	Uin                  int64    `protobuf:"varint,1,opt,name=uin,proto3" json:"uin,omitempty"`
	Cookies              []byte   `protobuf:"bytes,2,opt,name=cookies,proto3" json:"cookies,omitempty"`
	SessionKey           []byte   `protobuf:"bytes,3,opt,name=sessionKey,proto3" json:"sessionKey,omitempty"`
	DeviceId             string   `protobuf:"bytes,4,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	DeviceType           string   `protobuf:"bytes,5,opt,name=deviceType,proto3" json:"deviceType,omitempty"`
	DeviceName           string   `protobuf:"bytes,6,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	CurrentsyncKey       []byte   `protobuf:"bytes,7,opt,name=currentsyncKey,proto3" json:"currentsyncKey,omitempty"`
	MaxSyncKey           []byte   `protobuf:"bytes,8,opt,name=maxSyncKey,proto3" json:"maxSyncKey,omitempty"`
	AutoAuthKey          []byte   `protobuf:"bytes,9,opt,name=autoAuthKey,proto3" json:"autoAuthKey,omitempty"`
	Userame              string   `protobuf:"bytes,10,opt,name=userame,proto3" json:"userame,omitempty"`
	Nickname             []byte   `protobuf:"bytes,11,opt,name=nickname,proto3" json:"nickname,omitempty"`
	UserExt              []byte   `protobuf:"bytes,12,opt,name=userExt,proto3" json:"userExt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_WechatProto_c4b8821f4ce8a24f, []int{2}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUin() int64 {
	if m != nil {
		return m.Uin
	}
	return 0
}

func (m *User) GetCookies() []byte {
	if m != nil {
		return m.Cookies
	}
	return nil
}

func (m *User) GetSessionKey() []byte {
	if m != nil {
		return m.SessionKey
	}
	return nil
}

func (m *User) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *User) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *User) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *User) GetCurrentsyncKey() []byte {
	if m != nil {
		return m.CurrentsyncKey
	}
	return nil
}

func (m *User) GetMaxSyncKey() []byte {
	if m != nil {
		return m.MaxSyncKey
	}
	return nil
}

func (m *User) GetAutoAuthKey() []byte {
	if m != nil {
		return m.AutoAuthKey
	}
	return nil
}

func (m *User) GetUserame() string {
	if m != nil {
		return m.Userame
	}
	return ""
}

func (m *User) GetNickname() []byte {
	if m != nil {
		return m.Nickname
	}
	return nil
}

func (m *User) GetUserExt() []byte {
	if m != nil {
		return m.UserExt
	}
	return nil
}

func init() {
	proto.RegisterType((*WechatMsg)(nil), "WechatProto.WechatMsg")
	proto.RegisterType((*BaseMsg)(nil), "WechatProto.BaseMsg")
	proto.RegisterType((*User)(nil), "WechatProto.User")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WechatClient is the client API for Wechat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WechatClient interface {
	HelloWechat(ctx context.Context, in *WechatMsg, opts ...grpc.CallOption) (*WechatMsg, error)
}

type wechatClient struct {
	cc *grpc.ClientConn
}

func NewWechatClient(cc *grpc.ClientConn) WechatClient {
	return &wechatClient{cc}
}

func (c *wechatClient) HelloWechat(ctx context.Context, in *WechatMsg, opts ...grpc.CallOption) (*WechatMsg, error) {
	out := new(WechatMsg)
	err := c.cc.Invoke(ctx, "/WechatProto.Wechat/HelloWechat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WechatServer is the server API for Wechat service.
type WechatServer interface {
	HelloWechat(context.Context, *WechatMsg) (*WechatMsg, error)
}

func RegisterWechatServer(s *grpc.Server, srv WechatServer) {
	s.RegisterService(&_Wechat_serviceDesc, srv)
}

func _Wechat_HelloWechat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WechatMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatServer).HelloWechat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WechatProto.Wechat/HelloWechat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatServer).HelloWechat(ctx, req.(*WechatMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Wechat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "WechatProto.Wechat",
	HandlerType: (*WechatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWechat",
			Handler:    _Wechat_HelloWechat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "meta/WechatProto.proto",
}

func init() { proto.RegisterFile("meta/WechatProto.proto", fileDescriptor_WechatProto_c4b8821f4ce8a24f) }

var fileDescriptor_WechatProto_c4b8821f4ce8a24f = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xd1, 0x6a, 0xdb, 0x30,
	0x14, 0x9d, 0x9d, 0x38, 0xa9, 0xaf, 0x43, 0xd9, 0x44, 0x09, 0xa2, 0x8c, 0x61, 0x02, 0x1b, 0x79,
	0xca, 0xa0, 0xfb, 0x80, 0xb1, 0xc1, 0xa0, 0x65, 0x74, 0x0c, 0x75, 0x65, 0xcf, 0xaa, 0x7d, 0x49,
	0x4c, 0x6c, 0xcb, 0x58, 0x72, 0x49, 0x7e, 0x60, 0xff, 0xb0, 0x3f, 0xdc, 0x67, 0x8c, 0x2b, 0xc9,
	0x8e, 0x5b, 0xe8, 0x4b, 0xb8, 0xe7, 0x9c, 0x7b, 0x7d, 0x74, 0x8f, 0x22, 0x58, 0x56, 0x68, 0xe4,
	0xc7, 0xdf, 0x98, 0xed, 0xa4, 0xf9, 0xd9, 0x2a, 0xa3, 0x36, 0x0d, 0xfd, 0xb2, 0x64, 0x44, 0xad,
	0xfe, 0x06, 0x10, 0x3b, 0x7c, 0xab, 0xb7, 0x6c, 0x03, 0xf3, 0x07, 0xa9, 0xf1, 0x56, 0x6f, 0x79,
	0x90, 0x06, 0xeb, 0xe4, 0xea, 0x62, 0x33, 0x9e, 0xff, 0xea, 0x34, 0xd1, 0x37, 0xb1, 0x0b, 0x88,
	0x8c, 0xda, 0x63, 0xcd, 0xc3, 0x34, 0x58, 0xc7, 0xc2, 0x01, 0xc6, 0x61, 0xfe, 0x88, 0xad, 0x2e,
	0x54, 0xcd, 0x27, 0x96, 0xef, 0x21, 0x7b, 0x0b, 0xb1, 0x29, 0x2a, 0xbc, 0x33, 0xb2, 0x6a, 0xf8,
	0x34, 0x0d, 0xd6, 0x91, 0x38, 0x11, 0xec, 0x1c, 0xc2, 0xa2, 0xe1, 0x91, 0x1d, 0x09, 0x8b, 0x66,
	0xf5, 0x27, 0x84, 0xb9, 0xb7, 0x64, 0xaf, 0x61, 0xd2, 0xa2, 0xb1, 0xa7, 0x8a, 0x04, 0x95, 0xc4,
	0x64, 0x55, 0x6e, 0x9d, 0x23, 0x41, 0x25, 0x5b, 0xc2, 0x2c, 0xab, 0xf2, 0xfb, 0xb6, 0xf4, 0xb6,
	0x1e, 0x91, 0xab, 0xde, 0xa9, 0xd6, 0x5c, 0x2b, 0x6d, 0xac, 0x6b, 0x2c, 0x4e, 0x04, 0xbb, 0x84,
	0xb3, 0x52, 0xd5, 0x5b, 0x2b, 0x3a, 0xef, 0x01, 0xb3, 0xf7, 0x30, 0xed, 0x34, 0xb6, 0x7c, 0x66,
	0xc3, 0x78, 0xf3, 0x24, 0x8c, 0x7b, 0x8d, 0xad, 0xb0, 0xf2, 0xf0, 0x09, 0x94, 0x39, 0x9f, 0xa7,
	0xc1, 0x7a, 0x21, 0x06, 0x4c, 0x5a, 0x23, 0x8f, 0xa5, 0x92, 0xb9, 0xe6, 0x67, 0x4e, 0xeb, 0x31,
	0xfb, 0x00, 0xe7, 0x4d, 0xe9, 0x00, 0x1e, 0x0c, 0xd6, 0x39, 0x8f, 0x6d, 0xc7, 0x33, 0x76, 0xf5,
	0x2f, 0x84, 0x29, 0xd9, 0xd1, 0xce, 0x5d, 0x51, 0xdb, 0x14, 0x26, 0x82, 0x4a, 0xca, 0x3a, 0x53,
	0x6a, 0x5f, 0xa0, 0xb6, 0x49, 0x2c, 0x44, 0x0f, 0xd9, 0x3b, 0x00, 0x8d, 0x9a, 0x62, 0xff, 0x8e,
	0x47, 0x9b, 0xc8, 0x42, 0x8c, 0x18, 0x3a, 0x58, 0x8e, 0x8f, 0x45, 0x86, 0x37, 0xb9, 0x0f, 0x65,
	0xc0, 0x34, 0xeb, 0xea, 0x5f, 0xc7, 0x06, 0x7d, 0x2a, 0x23, 0xe6, 0xa4, 0xff, 0x90, 0x15, 0xda,
	0x74, 0x06, 0x9d, 0x18, 0x5a, 0x2c, 0xeb, 0xda, 0x16, 0x6b, 0xa3, 0x8f, 0x75, 0x46, 0xfe, 0x2e,
	0x96, 0x67, 0x2c, 0x7d, 0xa7, 0x92, 0x87, 0x3b, 0xdf, 0xe3, 0xe2, 0x19, 0x31, 0x2c, 0x85, 0x44,
	0x76, 0x46, 0x7d, 0xe9, 0xcc, 0x8e, 0x1a, 0x5c, 0x3a, 0x63, 0x8a, 0xf6, 0xa7, 0x2b, 0xa0, 0x63,
	0x80, 0xfb, 0xaf, 0x79, 0x48, 0xfb, 0xd5, 0x45, 0xb6, 0xaf, 0x49, 0x4a, 0x5c, 0xf0, 0x3d, 0xee,
	0xa7, 0xbe, 0x1d, 0x0c, 0x5f, 0xb8, 0xd4, 0x3c, 0xbc, 0xba, 0x81, 0x99, 0xbb, 0x64, 0xf6, 0x19,
	0x92, 0x6b, 0x2c, 0x4b, 0xe5, 0xe1, 0xf2, 0xc9, 0xe5, 0x0f, 0x4f, 0xe6, 0xf2, 0x05, 0x7e, 0xf5,
	0xea, 0x61, 0x66, 0x9f, 0xdb, 0xa7, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x14, 0xdc, 0xa2, 0x93,
	0x88, 0x03, 0x00, 0x00,
}
