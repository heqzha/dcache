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
	GetIfExistReq
	GetIfExistRes
	GetAllReq
	GetAllRes
	SetReq
	SetRes
	SetWithExpireReq
	SetWithExpireRes
	DelReq
	DelRes
	ClearReq
	ClearRes
	KeysReq
	KeysRes
	LenReq
	LenRes
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

type GetIfExistReq struct {
	Group string `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
}

func (m *GetIfExistReq) Reset()                    { *m = GetIfExistReq{} }
func (m *GetIfExistReq) String() string            { return proto.CompactTextString(m) }
func (*GetIfExistReq) ProtoMessage()               {}
func (*GetIfExistReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetIfExistReq) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *GetIfExistReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetIfExistRes struct {
	Status bool   `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Value  []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *GetIfExistRes) Reset()                    { *m = GetIfExistRes{} }
func (m *GetIfExistRes) String() string            { return proto.CompactTextString(m) }
func (*GetIfExistRes) ProtoMessage()               {}
func (*GetIfExistRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetIfExistRes) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *GetIfExistRes) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type GetAllReq struct {
	Group string `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
}

func (m *GetAllReq) Reset()                    { *m = GetAllReq{} }
func (m *GetAllReq) String() string            { return proto.CompactTextString(m) }
func (*GetAllReq) ProtoMessage()               {}
func (*GetAllReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetAllReq) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

type GetAllRes struct {
	Status bool   `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Value  []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *GetAllRes) Reset()                    { *m = GetAllRes{} }
func (m *GetAllRes) String() string            { return proto.CompactTextString(m) }
func (*GetAllRes) ProtoMessage()               {}
func (*GetAllRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetAllRes) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *GetAllRes) GetValue() []byte {
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
func (*SetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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
func (*SetRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SetRes) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type SetWithExpireReq struct {
	Group  string `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value  []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Expire int64  `protobuf:"varint,4,opt,name=expire" json:"expire,omitempty"`
}

func (m *SetWithExpireReq) Reset()                    { *m = SetWithExpireReq{} }
func (m *SetWithExpireReq) String() string            { return proto.CompactTextString(m) }
func (*SetWithExpireReq) ProtoMessage()               {}
func (*SetWithExpireReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SetWithExpireReq) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *SetWithExpireReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetWithExpireReq) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SetWithExpireReq) GetExpire() int64 {
	if m != nil {
		return m.Expire
	}
	return 0
}

type SetWithExpireRes struct {
	Status bool `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *SetWithExpireRes) Reset()                    { *m = SetWithExpireRes{} }
func (m *SetWithExpireRes) String() string            { return proto.CompactTextString(m) }
func (*SetWithExpireRes) ProtoMessage()               {}
func (*SetWithExpireRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SetWithExpireRes) GetStatus() bool {
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
func (*DelReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

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
func (*DelRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

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

type ClearReq struct {
	Group string `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
}

func (m *ClearReq) Reset()                    { *m = ClearReq{} }
func (m *ClearReq) String() string            { return proto.CompactTextString(m) }
func (*ClearReq) ProtoMessage()               {}
func (*ClearReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ClearReq) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

type ClearRes struct {
	Status bool `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *ClearRes) Reset()                    { *m = ClearRes{} }
func (m *ClearRes) String() string            { return proto.CompactTextString(m) }
func (*ClearRes) ProtoMessage()               {}
func (*ClearRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ClearRes) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type KeysReq struct {
	Group string `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
}

func (m *KeysReq) Reset()                    { *m = KeysReq{} }
func (m *KeysReq) String() string            { return proto.CompactTextString(m) }
func (*KeysReq) ProtoMessage()               {}
func (*KeysReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *KeysReq) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

type KeysRes struct {
	Status bool     `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Keys   []string `protobuf:"bytes,2,rep,name=keys" json:"keys,omitempty"`
}

func (m *KeysRes) Reset()                    { *m = KeysRes{} }
func (m *KeysRes) String() string            { return proto.CompactTextString(m) }
func (*KeysRes) ProtoMessage()               {}
func (*KeysRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *KeysRes) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *KeysRes) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type LenReq struct {
	Group string `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
}

func (m *LenReq) Reset()                    { *m = LenReq{} }
func (m *LenReq) String() string            { return proto.CompactTextString(m) }
func (*LenReq) ProtoMessage()               {}
func (*LenReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *LenReq) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

type LenRes struct {
	Status bool  `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Len    int64 `protobuf:"varint,2,opt,name=len" json:"len,omitempty"`
}

func (m *LenRes) Reset()                    { *m = LenRes{} }
func (m *LenRes) String() string            { return proto.CompactTextString(m) }
func (*LenRes) ProtoMessage()               {}
func (*LenRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *LenRes) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *LenRes) GetLen() int64 {
	if m != nil {
		return m.Len
	}
	return 0
}

type RegisterReq struct {
	Group string `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	Addr  string `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty"`
}

func (m *RegisterReq) Reset()                    { *m = RegisterReq{} }
func (m *RegisterReq) String() string            { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()               {}
func (*RegisterReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

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
func (*RegisterRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

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
func (*UnregisterReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

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
func (*UnregisterRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

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
func (*SyncSrvGroupReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

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
func (*SyncSrvGroupRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{24} }

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
func (*PingReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{25} }

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
func (*PingRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{26} }

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
	proto.RegisterType((*GetIfExistReq)(nil), "pb.GetIfExistReq")
	proto.RegisterType((*GetIfExistRes)(nil), "pb.GetIfExistRes")
	proto.RegisterType((*GetAllReq)(nil), "pb.GetAllReq")
	proto.RegisterType((*GetAllRes)(nil), "pb.GetAllRes")
	proto.RegisterType((*SetReq)(nil), "pb.SetReq")
	proto.RegisterType((*SetRes)(nil), "pb.SetRes")
	proto.RegisterType((*SetWithExpireReq)(nil), "pb.SetWithExpireReq")
	proto.RegisterType((*SetWithExpireRes)(nil), "pb.SetWithExpireRes")
	proto.RegisterType((*DelReq)(nil), "pb.DelReq")
	proto.RegisterType((*DelRes)(nil), "pb.DelRes")
	proto.RegisterType((*ClearReq)(nil), "pb.ClearReq")
	proto.RegisterType((*ClearRes)(nil), "pb.ClearRes")
	proto.RegisterType((*KeysReq)(nil), "pb.KeysReq")
	proto.RegisterType((*KeysRes)(nil), "pb.KeysRes")
	proto.RegisterType((*LenReq)(nil), "pb.LenReq")
	proto.RegisterType((*LenRes)(nil), "pb.LenRes")
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
	GetIfExist(ctx context.Context, in *GetIfExistReq, opts ...grpc.CallOption) (*GetIfExistRes, error)
	Set(ctx context.Context, in *SetReq, opts ...grpc.CallOption) (*SetRes, error)
	SetWithExpire(ctx context.Context, in *SetWithExpireReq, opts ...grpc.CallOption) (*SetWithExpireRes, error)
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

func (c *cacheServClient) GetIfExist(ctx context.Context, in *GetIfExistReq, opts ...grpc.CallOption) (*GetIfExistRes, error) {
	out := new(GetIfExistRes)
	err := grpc.Invoke(ctx, "/pb.CacheServ/GetIfExist", in, out, c.cc, opts...)
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

func (c *cacheServClient) SetWithExpire(ctx context.Context, in *SetWithExpireReq, opts ...grpc.CallOption) (*SetWithExpireRes, error) {
	out := new(SetWithExpireRes)
	err := grpc.Invoke(ctx, "/pb.CacheServ/SetWithExpire", in, out, c.cc, opts...)
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
	GetIfExist(context.Context, *GetIfExistReq) (*GetIfExistRes, error)
	Set(context.Context, *SetReq) (*SetRes, error)
	SetWithExpire(context.Context, *SetWithExpireReq) (*SetWithExpireRes, error)
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

func _CacheServ_GetIfExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIfExistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServServer).GetIfExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CacheServ/GetIfExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServServer).GetIfExist(ctx, req.(*GetIfExistReq))
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

func _CacheServ_SetWithExpire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetWithExpireReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServServer).SetWithExpire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CacheServ/SetWithExpire",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServServer).SetWithExpire(ctx, req.(*SetWithExpireReq))
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
			MethodName: "GetIfExist",
			Handler:    _CacheServ_GetIfExist_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _CacheServ_Set_Handler,
		},
		{
			MethodName: "SetWithExpire",
			Handler:    _CacheServ_SetWithExpire_Handler,
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
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x51, 0x6f, 0x12, 0x41,
	0x10, 0x2e, 0x1c, 0xbd, 0x72, 0x53, 0x48, 0x71, 0x6d, 0x0c, 0x39, 0x8d, 0xe0, 0x26, 0x46, 0xe2,
	0x03, 0x69, 0x5a, 0xb5, 0xa9, 0x49, 0x4d, 0x4c, 0x5b, 0x89, 0xd1, 0x07, 0xb3, 0x17, 0xe3, 0xa3,
	0x39, 0x60, 0x84, 0x4b, 0x2f, 0xc7, 0xb9, 0xbb, 0x90, 0xf2, 0x93, 0xfc, 0x97, 0x66, 0xf7, 0x16,
	0xef, 0x0e, 0xba, 0x0f, 0x67, 0xdf, 0x66, 0x67, 0xe6, 0xfb, 0xbe, 0x99, 0x61, 0x6e, 0x80, 0xd6,
	0x74, 0x12, 0x4e, 0xe6, 0x38, 0x4c, 0xf9, 0x42, 0x2e, 0x48, 0x3d, 0x1d, 0xd3, 0x2e, 0x34, 0x02,
	0x19, 0x4a, 0xd2, 0x01, 0x87, 0xa7, 0xa2, 0x5b, 0xeb, 0xd7, 0x06, 0x35, 0xa6, 0x4c, 0x7a, 0x02,
	0xee, 0x08, 0x25, 0xc3, 0xdf, 0xe4, 0x18, 0xf6, 0x67, 0x7c, 0xb1, 0x4c, 0x75, 0xd4, 0x63, 0xd9,
	0x43, 0x21, 0x6e, 0x71, 0xdd, 0xad, 0x6b, 0x9f, 0x32, 0xe9, 0x3b, 0x83, 0x10, 0xe4, 0x09, 0xb8,
	0x42, 0x86, 0x72, 0x99, 0x11, 0x36, 0x99, 0x79, 0x29, 0xa6, 0x55, 0x18, 0x2f, 0x51, 0xa3, 0x5a,
	0x2c, 0x7b, 0xd0, 0x73, 0x68, 0x8f, 0x50, 0x7e, 0xfe, 0x75, 0x73, 0x17, 0x89, 0x4a, 0x82, 0x97,
	0x65, 0x60, 0x55, 0xdd, 0x17, 0xe0, 0x8d, 0x50, 0x7e, 0x8c, 0x63, 0xab, 0x26, 0xbd, 0xc8, 0x53,
	0xaa, 0xb2, 0x7f, 0x02, 0x37, 0xa8, 0x34, 0xbf, 0x9c, 0xc7, 0x29, 0xf2, 0xf4, 0x0d, 0x8f, 0x55,
	0x9f, 0xce, 0xa1, 0x13, 0xa0, 0xfc, 0x11, 0xc9, 0xf9, 0xcd, 0x5d, 0x1a, 0x71, 0x7c, 0xb0, 0xa6,
	0x52, 0x42, 0x4d, 0xd5, 0x6d, 0xf4, 0x6b, 0x03, 0x87, 0x99, 0x17, 0x7d, 0xbd, 0xa3, 0x64, 0xaf,
	0xea, 0x04, 0xdc, 0x6b, 0x8c, 0x2b, 0xee, 0x8f, 0x46, 0x54, 0x9d, 0x74, 0x1f, 0x9a, 0x57, 0x31,
	0x86, 0xdc, 0xfe, 0x33, 0xd2, 0x7f, 0x19, 0xf6, 0x7a, 0x7b, 0x70, 0xf0, 0x05, 0xd7, 0xc2, 0x4e,
	0xf2, 0x76, 0x93, 0x60, 0xaf, 0x8f, 0x40, 0xe3, 0x16, 0xd7, 0xa2, 0x5b, 0xef, 0x3b, 0x03, 0x8f,
	0x69, 0x9b, 0x3e, 0x07, 0xf7, 0x2b, 0x26, 0x76, 0xda, 0x53, 0x13, 0xb7, 0xb3, 0x76, 0xc0, 0x89,
	0x31, 0xd1, 0x3d, 0x3b, 0x4c, 0x99, 0xf4, 0x1c, 0x0e, 0x19, 0xce, 0x22, 0x21, 0xd1, 0xde, 0xb4,
	0x2a, 0x26, 0x9c, 0x4e, 0xb9, 0x99, 0xb0, 0xb6, 0xe9, 0xcb, 0x22, 0xd0, 0x3e, 0x8b, 0x0b, 0x68,
	0x7f, 0x4f, 0xf8, 0x7f, 0x29, 0xbc, 0x2a, 0x43, 0xed, 0x1a, 0x43, 0x38, 0x0a, 0xd6, 0xc9, 0x24,
	0xe0, 0xab, 0x91, 0x22, 0x53, 0x2a, 0x4f, 0xc1, 0x13, 0x7c, 0xf5, 0x33, 0x57, 0x6a, 0xb1, 0xa6,
	0x30, 0x71, 0x3a, 0xdd, 0xce, 0xb7, 0x0f, 0xec, 0x19, 0x78, 0x93, 0x45, 0x32, 0x8d, 0x64, 0xb4,
	0xc8, 0xc6, 0xb6, 0xcf, 0x72, 0x47, 0x59, 0xc5, 0xd9, 0x52, 0x39, 0x83, 0x83, 0x6f, 0x51, 0x32,
	0xab, 0xd6, 0xf3, 0x87, 0x0d, 0xc8, 0x5e, 0x52, 0x49, 0xb4, 0x5e, 0x16, 0x3d, 0xfd, 0xe3, 0x80,
	0x77, 0xa5, 0x0e, 0x73, 0x80, 0x7c, 0x45, 0x7a, 0xe0, 0x8c, 0x50, 0x12, 0x18, 0xa6, 0xe3, 0x61,
	0x76, 0x81, 0xfd, 0xdc, 0x16, 0x74, 0x8f, 0xbc, 0x01, 0xc8, 0xcf, 0x1e, 0x79, 0x64, 0x62, 0xf9,
	0xfd, 0xf4, 0x77, 0x5c, 0x0a, 0xd5, 0x03, 0x27, 0xd8, 0xd0, 0x06, 0x05, 0xda, 0x60, 0x43, 0x7b,
	0x09, 0xed, 0xd2, 0xc7, 0x4d, 0x8e, 0x4d, 0xb8, 0x74, 0x59, 0xfc, 0xfb, 0xbc, 0x86, 0xff, 0x1a,
	0xe3, 0x8c, 0x3f, 0xfb, 0xf0, 0xfd, 0xdc, 0x56, 0x09, 0x43, 0x68, 0x6e, 0x76, 0x8f, 0x1c, 0xa9,
	0x48, 0x61, 0x85, 0xfd, 0x2d, 0x87, 0x69, 0x33, 0xdf, 0xa4, 0xac, 0xcd, 0xd2, 0x52, 0xfa, 0x3b,
	0x2e, 0x85, 0x7a, 0x0f, 0xad, 0xe2, 0x9a, 0x90, 0xc7, 0xba, 0xdc, 0xf2, 0xa2, 0xf9, 0xf7, 0x38,
	0x15, 0x96, 0x42, 0x43, 0xfd, 0x8e, 0xe4, 0x50, 0x85, 0xcd, 0x1a, 0xf8, 0x85, 0x87, 0xa0, 0x7b,
	0x63, 0x57, 0xff, 0x77, 0x9e, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x09, 0x46, 0xa3, 0x4b,
	0x07, 0x00, 0x00,
}
