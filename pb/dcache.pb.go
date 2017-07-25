// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dcache.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	dcache.proto

It has these top-level messages:
	Stat
	GetReq
	GetRes
	SetReq
	SetRes
	DelReq
	DelRes
	RegisterReq
	RegisterRes
	UnregisterReq
	UnregisterRes
	SyncSrvGroupReq
	SyncSrvGroupRes
	PingReq
	PingRes
*/
package pb

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

// Statistics Data of key
type Stat struct {
	Rps float64 `protobuf:"fixed64,1,opt,name=rps" json:"rps,omitempty"`
}

func (m *Stat) Reset()                    { *m = Stat{} }
func (m *Stat) String() string            { return proto.CompactTextString(m) }
func (*Stat) ProtoMessage()               {}
func (*Stat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Stat) GetRps() float64 {
	if m != nil {
		return m.Rps
	}
	return 0
}

type GetReq struct {
	Group string `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
}

func (m *GetReq) Reset()                    { *m = GetReq{} }
func (m *GetReq) String() string            { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()               {}
func (*GetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetReq) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *GetReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetRes struct {
	Status bool   `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Value  []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *GetRes) Reset()                    { *m = GetRes{} }
func (m *GetRes) String() string            { return proto.CompactTextString(m) }
func (*GetRes) ProtoMessage()               {}
func (*GetRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetRes) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *GetRes) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type SetReq struct {
	Group string `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *SetReq) Reset()                    { *m = SetReq{} }
func (m *SetReq) String() string            { return proto.CompactTextString(m) }
func (*SetReq) ProtoMessage()               {}
func (*SetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SetReq) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *SetReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetReq) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type SetRes struct {
	Status bool `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *SetRes) Reset()                    { *m = SetRes{} }
func (m *SetRes) String() string            { return proto.CompactTextString(m) }
func (*SetRes) ProtoMessage()               {}
func (*SetRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SetRes) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type DelReq struct {
	Group string `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
}

func (m *DelReq) Reset()                    { *m = DelReq{} }
func (m *DelReq) String() string            { return proto.CompactTextString(m) }
func (*DelReq) ProtoMessage()               {}
func (*DelReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DelReq) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *DelReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DelRes struct {
	Status bool   `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Value  []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *DelRes) Reset()                    { *m = DelRes{} }
func (m *DelRes) String() string            { return proto.CompactTextString(m) }
func (*DelRes) ProtoMessage()               {}
func (*DelRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DelRes) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *DelRes) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type RegisterReq struct {
	Group string `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	Addr  string `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty"`
}

func (m *RegisterReq) Reset()                    { *m = RegisterReq{} }
func (m *RegisterReq) String() string            { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()               {}
func (*RegisterReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RegisterReq) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *RegisterReq) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type RegisterRes struct {
	Status bool `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *RegisterRes) Reset()                    { *m = RegisterRes{} }
func (m *RegisterRes) String() string            { return proto.CompactTextString(m) }
func (*RegisterRes) ProtoMessage()               {}
func (*RegisterRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RegisterRes) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type UnregisterReq struct {
	Group string `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	Addr  string `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty"`
}

func (m *UnregisterReq) Reset()                    { *m = UnregisterReq{} }
func (m *UnregisterReq) String() string            { return proto.CompactTextString(m) }
func (*UnregisterReq) ProtoMessage()               {}
func (*UnregisterReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UnregisterReq) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *UnregisterReq) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type UnregisterRes struct {
	Status bool `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *UnregisterRes) Reset()                    { *m = UnregisterRes{} }
func (m *UnregisterRes) String() string            { return proto.CompactTextString(m) }
func (*UnregisterRes) ProtoMessage()               {}
func (*UnregisterRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UnregisterRes) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type SyncSrvGroupReq struct {
	SrvGroup []byte `protobuf:"bytes,1,opt,name=srv_group,json=srvGroup,proto3" json:"srv_group,omitempty"`
}

func (m *SyncSrvGroupReq) Reset()                    { *m = SyncSrvGroupReq{} }
func (m *SyncSrvGroupReq) String() string            { return proto.CompactTextString(m) }
func (*SyncSrvGroupReq) ProtoMessage()               {}
func (*SyncSrvGroupReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *SyncSrvGroupReq) GetSrvGroup() []byte {
	if m != nil {
		return m.SrvGroup
	}
	return nil
}

type SyncSrvGroupRes struct {
	Status    bool   `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Condition int32  `protobuf:"varint,2,opt,name=condition" json:"condition,omitempty"`
	SrvGroup  []byte `protobuf:"bytes,3,opt,name=srv_group,json=srvGroup,proto3" json:"srv_group,omitempty"`
}

func (m *SyncSrvGroupRes) Reset()                    { *m = SyncSrvGroupRes{} }
func (m *SyncSrvGroupRes) String() string            { return proto.CompactTextString(m) }
func (*SyncSrvGroupRes) ProtoMessage()               {}
func (*SyncSrvGroupRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *SyncSrvGroupRes) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *SyncSrvGroupRes) GetCondition() int32 {
	if m != nil {
		return m.Condition
	}
	return 0
}

func (m *SyncSrvGroupRes) GetSrvGroup() []byte {
	if m != nil {
		return m.SrvGroup
	}
	return nil
}

type PingReq struct {
	Group string `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	Addr  string `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty"`
}

func (m *PingReq) Reset()                    { *m = PingReq{} }
func (m *PingReq) String() string            { return proto.CompactTextString(m) }
func (*PingReq) ProtoMessage()               {}
func (*PingReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *PingReq) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *PingReq) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type PingRes struct {
	Status   bool   `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	SrvGroup []byte `protobuf:"bytes,2,opt,name=srv_group,json=srvGroup,proto3" json:"srv_group,omitempty"`
}

func (m *PingRes) Reset()                    { *m = PingRes{} }
func (m *PingRes) String() string            { return proto.CompactTextString(m) }
func (*PingRes) ProtoMessage()               {}
func (*PingRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *PingRes) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *PingRes) GetSrvGroup() []byte {
	if m != nil {
		return m.SrvGroup
	}
	return nil
}

func init() {
	proto.RegisterType((*Stat)(nil), "pb.Stat")
	proto.RegisterType((*GetReq)(nil), "pb.GetReq")
	proto.RegisterType((*GetRes)(nil), "pb.GetRes")
	proto.RegisterType((*SetReq)(nil), "pb.SetReq")
	proto.RegisterType((*SetRes)(nil), "pb.SetRes")
	proto.RegisterType((*DelReq)(nil), "pb.DelReq")
	proto.RegisterType((*DelRes)(nil), "pb.DelRes")
	proto.RegisterType((*RegisterReq)(nil), "pb.RegisterReq")
	proto.RegisterType((*RegisterRes)(nil), "pb.RegisterRes")
	proto.RegisterType((*UnregisterReq)(nil), "pb.UnregisterReq")
	proto.RegisterType((*UnregisterRes)(nil), "pb.UnregisterRes")
	proto.RegisterType((*SyncSrvGroupReq)(nil), "pb.SyncSrvGroupReq")
	proto.RegisterType((*SyncSrvGroupRes)(nil), "pb.SyncSrvGroupRes")
	proto.RegisterType((*PingReq)(nil), "pb.PingReq")
	proto.RegisterType((*PingRes)(nil), "pb.PingRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CacheServ service

type CacheServClient interface {
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRes, error)
	Set(ctx context.Context, in *SetReq, opts ...grpc.CallOption) (*SetRes, error)
	Del(ctx context.Context, in *DelReq, opts ...grpc.CallOption) (*DelRes, error)
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error)
	Unregister(ctx context.Context, in *UnregisterReq, opts ...grpc.CallOption) (*UnregisterRes, error)
	SyncSrvGroup(ctx context.Context, in *SyncSrvGroupReq, opts ...grpc.CallOption) (*SyncSrvGroupRes, error)
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRes, error)
}

type cacheServClient struct {
	cc *grpc.ClientConn
}

func NewCacheServClient(cc *grpc.ClientConn) CacheServClient {
	return &cacheServClient{cc}
}

func (c *cacheServClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRes, error) {
	out := new(GetRes)
	err := grpc.Invoke(ctx, "/pb.CacheServ/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServClient) Set(ctx context.Context, in *SetReq, opts ...grpc.CallOption) (*SetRes, error) {
	out := new(SetRes)
	err := grpc.Invoke(ctx, "/pb.CacheServ/Set", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServClient) Del(ctx context.Context, in *DelReq, opts ...grpc.CallOption) (*DelRes, error) {
	out := new(DelRes)
	err := grpc.Invoke(ctx, "/pb.CacheServ/Del", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error) {
	out := new(RegisterRes)
	err := grpc.Invoke(ctx, "/pb.CacheServ/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServClient) Unregister(ctx context.Context, in *UnregisterReq, opts ...grpc.CallOption) (*UnregisterRes, error) {
	out := new(UnregisterRes)
	err := grpc.Invoke(ctx, "/pb.CacheServ/Unregister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServClient) SyncSrvGroup(ctx context.Context, in *SyncSrvGroupReq, opts ...grpc.CallOption) (*SyncSrvGroupRes, error) {
	out := new(SyncSrvGroupRes)
	err := grpc.Invoke(ctx, "/pb.CacheServ/SyncSrvGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRes, error) {
	out := new(PingRes)
	err := grpc.Invoke(ctx, "/pb.CacheServ/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CacheServ service

type CacheServServer interface {
	Get(context.Context, *GetReq) (*GetRes, error)
	Set(context.Context, *SetReq) (*SetRes, error)
	Del(context.Context, *DelReq) (*DelRes, error)
	Register(context.Context, *RegisterReq) (*RegisterRes, error)
	Unregister(context.Context, *UnregisterReq) (*UnregisterRes, error)
	SyncSrvGroup(context.Context, *SyncSrvGroupReq) (*SyncSrvGroupRes, error)
	Ping(context.Context, *PingReq) (*PingRes, error)
}

func RegisterCacheServServer(s *grpc.Server, srv CacheServServer) {
	s.RegisterService(&_CacheServ_serviceDesc, srv)
}

func _CacheServ_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CacheServ/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheServ_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CacheServ/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServServer).Set(ctx, req.(*SetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheServ_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CacheServ/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServServer).Del(ctx, req.(*DelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheServ_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CacheServ/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheServ_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CacheServ/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServServer).Unregister(ctx, req.(*UnregisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheServ_SyncSrvGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncSrvGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServServer).SyncSrvGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CacheServ/SyncSrvGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServServer).SyncSrvGroup(ctx, req.(*SyncSrvGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheServ_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CacheServ/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CacheServ_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CacheServ",
	HandlerType: (*CacheServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CacheServ_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _CacheServ_Set_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _CacheServ_Del_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _CacheServ_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _CacheServ_Unregister_Handler,
		},
		{
			MethodName: "SyncSrvGroup",
			Handler:    _CacheServ_SyncSrvGroup_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _CacheServ_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dcache.proto",
}

func init() { proto.RegisterFile("dcache.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xb5, 0x25, 0x5b, 0xb5, 0xc6, 0x2a, 0x6e, 0xb7, 0xa5, 0x18, 0xb5, 0x50, 0xb3, 0x50, 0xda,
	0x93, 0x28, 0x75, 0x49, 0x48, 0x0e, 0xb9, 0xc4, 0xc4, 0xd7, 0xb0, 0x4b, 0xce, 0x41, 0x96, 0x16,
	0x47, 0xc4, 0x48, 0xca, 0xee, 0x5a, 0xe0, 0xff, 0x95, 0x1f, 0x18, 0x66, 0x25, 0x23, 0xc9, 0x89,
	0x0e, 0xca, 0x6d, 0x3e, 0xde, 0x9b, 0x37, 0xa3, 0x99, 0x15, 0x78, 0x71, 0x14, 0x46, 0x0f, 0x22,
	0xc8, 0x65, 0xa6, 0x33, 0x62, 0xe5, 0x1b, 0x3a, 0x87, 0x11, 0xd7, 0xa1, 0x26, 0x9f, 0xc0, 0x96,
	0xb9, 0x9a, 0x0f, 0x17, 0xc3, 0x3f, 0x43, 0x86, 0x26, 0xfd, 0x0b, 0xce, 0x5a, 0x68, 0x26, 0x9e,
	0xc8, 0x57, 0x18, 0x6f, 0x65, 0xb6, 0xcf, 0x4d, 0xd6, 0x65, 0xa5, 0x83, 0x8c, 0x47, 0x71, 0x98,
	0x5b, 0x26, 0x86, 0x26, 0x3d, 0xab, 0x18, 0x8a, 0x7c, 0x03, 0x47, 0xe9, 0x50, 0xef, 0xcb, 0x82,
	0x13, 0x56, 0x79, 0x58, 0xa9, 0x08, 0x77, 0x7b, 0x61, 0x58, 0x1e, 0x2b, 0x1d, 0x7a, 0x03, 0x0e,
	0xef, 0xa5, 0x54, 0xd7, 0xb1, 0x9b, 0x75, 0x16, 0x55, 0x9d, 0x4e, 0x7d, 0x9c, 0x69, 0x25, 0x76,
	0x3d, 0x67, 0x32, 0x8c, 0xbe, 0x33, 0x9d, 0xc3, 0x94, 0x89, 0x6d, 0xa2, 0xb4, 0x90, 0xdd, 0x72,
	0x04, 0x46, 0x61, 0x1c, 0xcb, 0x4a, 0xcf, 0xd8, 0xf4, 0x57, 0x93, 0xd8, 0x3d, 0xc9, 0x05, 0x7c,
	0xbc, 0x4b, 0xe5, 0xbb, 0x14, 0x7e, 0xb7, 0xa9, 0xdd, 0x1a, 0x01, 0xcc, 0xf8, 0x21, 0x8d, 0xb8,
	0x2c, 0xd6, 0x58, 0x0c, 0x55, 0xbe, 0x83, 0xab, 0x64, 0x71, 0x5f, 0x2b, 0x79, 0x6c, 0xa2, 0xaa,
	0x3c, 0x8d, 0x4f, 0xf1, 0xdd, 0x1f, 0xed, 0x07, 0xb8, 0x51, 0x96, 0xc6, 0x89, 0x4e, 0xb2, 0xd4,
	0x34, 0x37, 0x66, 0x75, 0xa0, 0xad, 0x62, 0x9f, 0xa8, 0x2c, 0xe1, 0xc3, 0x6d, 0x92, 0x6e, 0xfb,
	0xcd, 0x7c, 0x75, 0x24, 0x75, 0xb7, 0xd4, 0x12, 0xb5, 0xda, 0xa2, 0xff, 0x9e, 0x2d, 0x70, 0xaf,
	0xf1, 0xe9, 0x70, 0x21, 0x0b, 0xf2, 0x13, 0xec, 0xb5, 0xd0, 0x04, 0x82, 0x7c, 0x13, 0x94, 0x6f,
	0xc4, 0xaf, 0x6d, 0x45, 0x07, 0x08, 0xe0, 0x47, 0x00, 0x6f, 0x00, 0x78, 0x03, 0xb0, 0x12, 0xbb,
	0x12, 0x50, 0x5e, 0xa4, 0x5f, 0xdb, 0x08, 0x08, 0x60, 0x72, 0x3c, 0x03, 0x32, 0xc3, 0x4c, 0xe3,
	0x9a, 0xfc, 0x93, 0x00, 0xe2, 0xff, 0x03, 0xd4, 0x4b, 0x25, 0x9f, 0x11, 0xd0, 0xba, 0x0f, 0xff,
	0x55, 0x08, 0x59, 0x97, 0xe0, 0x35, 0x37, 0x46, 0xbe, 0x98, 0x26, 0xdb, 0x3b, 0xf7, 0xdf, 0x08,
	0x22, 0x97, 0xc2, 0x08, 0x3f, 0x29, 0x99, 0x62, 0xba, 0xda, 0x88, 0xdf, 0x70, 0x14, 0x1d, 0x6c,
	0x1c, 0xf3, 0xa3, 0x59, 0xbe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x62, 0xa0, 0x2d, 0x61, 0x78, 0x04,
	0x00, 0x00,
}
