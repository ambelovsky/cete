// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/kvs/kvs.proto

package kvs

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	raft "github.com/mosuka/cete/protobuf/raft"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KVSCommand_Type int32

const (
	KVSCommand_UNKNOWN_COMMAND       KVSCommand_Type = 0
	KVSCommand_SET_METADATA          KVSCommand_Type = 1
	KVSCommand_DELETE_METADATA       KVSCommand_Type = 2
	KVSCommand_PUT_KEY_VALUE_PAIR    KVSCommand_Type = 3
	KVSCommand_DELETE_KEY_VALUE_PAIR KVSCommand_Type = 4
)

var KVSCommand_Type_name = map[int32]string{
	0: "UNKNOWN_COMMAND",
	1: "SET_METADATA",
	2: "DELETE_METADATA",
	3: "PUT_KEY_VALUE_PAIR",
	4: "DELETE_KEY_VALUE_PAIR",
}

var KVSCommand_Type_value = map[string]int32{
	"UNKNOWN_COMMAND":       0,
	"SET_METADATA":          1,
	"DELETE_METADATA":       2,
	"PUT_KEY_VALUE_PAIR":    3,
	"DELETE_KEY_VALUE_PAIR": 4,
}

func (x KVSCommand_Type) String() string {
	return proto.EnumName(KVSCommand_Type_name, int32(x))
}

func (KVSCommand_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{1, 0}
}

type KeyValuePair struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValuePair) Reset()         { *m = KeyValuePair{} }
func (m *KeyValuePair) String() string { return proto.CompactTextString(m) }
func (*KeyValuePair) ProtoMessage()    {}
func (*KeyValuePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{0}
}

func (m *KeyValuePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValuePair.Unmarshal(m, b)
}
func (m *KeyValuePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValuePair.Marshal(b, m, deterministic)
}
func (m *KeyValuePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValuePair.Merge(m, src)
}
func (m *KeyValuePair) XXX_Size() int {
	return xxx_messageInfo_KeyValuePair.Size(m)
}
func (m *KeyValuePair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValuePair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValuePair proto.InternalMessageInfo

func (m *KeyValuePair) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyValuePair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type KVSCommand struct {
	Type                 KVSCommand_Type `protobuf:"varint,1,opt,name=type,proto3,enum=kvs.KVSCommand_Type" json:"type,omitempty"`
	Data                 *any.Any        `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *KVSCommand) Reset()         { *m = KVSCommand{} }
func (m *KVSCommand) String() string { return proto.CompactTextString(m) }
func (*KVSCommand) ProtoMessage()    {}
func (*KVSCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{1}
}

func (m *KVSCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVSCommand.Unmarshal(m, b)
}
func (m *KVSCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVSCommand.Marshal(b, m, deterministic)
}
func (m *KVSCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVSCommand.Merge(m, src)
}
func (m *KVSCommand) XXX_Size() int {
	return xxx_messageInfo_KVSCommand.Size(m)
}
func (m *KVSCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_KVSCommand.DiscardUnknown(m)
}

var xxx_messageInfo_KVSCommand proto.InternalMessageInfo

func (m *KVSCommand) GetType() KVSCommand_Type {
	if m != nil {
		return m.Type
	}
	return KVSCommand_UNKNOWN_COMMAND
}

func (m *KVSCommand) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("kvs.KVSCommand_Type", KVSCommand_Type_name, KVSCommand_Type_value)
	proto.RegisterType((*KeyValuePair)(nil), "kvs.KeyValuePair")
	proto.RegisterType((*KVSCommand)(nil), "kvs.KVSCommand")
}

func init() { proto.RegisterFile("protobuf/kvs/kvs.proto", fileDescriptor_6e9909cfc2f34163) }

var fileDescriptor_6e9909cfc2f34163 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x5f, 0x8f, 0x93, 0x40,
	0x14, 0xc5, 0xdb, 0xa5, 0xbb, 0x9a, 0x6b, 0x55, 0xbc, 0xd6, 0xa6, 0x8b, 0x2f, 0x06, 0x63, 0xb2,
	0x0f, 0x0a, 0x49, 0xcd, 0xfa, 0xe4, 0x0b, 0xb6, 0xa4, 0x51, 0x5a, 0xb6, 0x69, 0x69, 0x8d, 0xbe,
	0x90, 0xe9, 0xf6, 0x6e, 0xb7, 0xa1, 0x30, 0xa4, 0x0c, 0x24, 0x7c, 0x1d, 0xbf, 0x94, 0x5f, 0xc7,
	0xcc, 0xb0, 0x7f, 0x9a, 0x5d, 0x6b, 0xdc, 0x07, 0x08, 0x73, 0xce, 0xef, 0x0c, 0x87, 0xe1, 0x42,
	0x3b, 0xdd, 0x72, 0xc1, 0x17, 0xf9, 0x85, 0x1d, 0x15, 0x99, 0xbc, 0x2c, 0x25, 0xa0, 0x16, 0x15,
	0x99, 0x71, 0xbc, 0xe2, 0x7c, 0xb5, 0x21, 0xfb, 0x86, 0x61, 0x49, 0x59, 0xf9, 0xc6, 0xeb, 0xbb,
	0x16, 0xc5, 0xa9, 0xb8, 0x36, 0x3b, 0x37, 0xea, 0x96, 0x5d, 0x08, 0x75, 0xab, 0x1c, 0xf3, 0x13,
	0x34, 0x3d, 0x2a, 0xe7, 0x6c, 0x93, 0xd3, 0x98, 0xad, 0xb7, 0xa8, 0x83, 0x16, 0x51, 0xd9, 0xa9,
	0xbf, 0xa9, 0x9f, 0x34, 0x27, 0xf2, 0x11, 0x5b, 0x70, 0x58, 0x48, 0xbb, 0x73, 0xa0, 0xb4, 0x6a,
	0x61, 0xfe, 0xae, 0x03, 0x78, 0xf3, 0x69, 0x8f, 0xc7, 0x31, 0x4b, 0x96, 0x78, 0x02, 0x0d, 0x51,
	0xa6, 0xa4, 0x72, 0xcf, 0xba, 0x2d, 0x4b, 0xf6, 0xbe, 0xb5, 0xad, 0xa0, 0x4c, 0x69, 0xa2, 0x08,
	0x49, 0x2e, 0x99, 0x60, 0x6a, 0xb7, 0x27, 0xdd, 0x96, 0x55, 0xd5, 0xb6, 0xae, 0x0b, 0x5a, 0x4e,
	0x52, 0x4e, 0x14, 0x61, 0xe6, 0xd0, 0x90, 0x39, 0x7c, 0x09, 0xcf, 0x67, 0xbe, 0xe7, 0x9f, 0x7d,
	0xf7, 0xc3, 0xde, 0xd9, 0x68, 0xe4, 0xf8, 0x7d, 0xbd, 0x86, 0x3a, 0x34, 0xa7, 0x6e, 0x10, 0x8e,
	0xdc, 0xc0, 0xe9, 0x3b, 0x81, 0xa3, 0xd7, 0x25, 0xd6, 0x77, 0x87, 0x6e, 0xe0, 0xde, 0x8a, 0x07,
	0xd8, 0x06, 0x1c, 0xcf, 0x82, 0xd0, 0x73, 0x7f, 0x84, 0x73, 0x67, 0x38, 0x73, 0xc3, 0xb1, 0xf3,
	0x75, 0xa2, 0x6b, 0x78, 0x0c, 0xaf, 0xae, 0xe0, 0x3b, 0x56, 0xa3, 0xfb, 0x4b, 0x03, 0xcd, 0x9b,
	0x4f, 0xf1, 0x3d, 0x34, 0xbe, 0xf1, 0x75, 0x82, 0x60, 0xa9, 0xe3, 0xf2, 0xf9, 0x92, 0x8c, 0xf6,
	0xbd, 0xba, 0xae, 0x3c, 0x65, 0xb3, 0x86, 0x1f, 0xe0, 0x70, 0x48, 0xac, 0xa0, 0xff, 0xc4, 0x6d,
	0x78, 0x34, 0x20, 0x21, 0x21, 0xdc, 0x03, 0x19, 0x3b, 0x1b, 0x99, 0x35, 0x3c, 0x05, 0x18, 0x90,
	0xe8, 0x6d, 0xf2, 0x4c, 0xd0, 0x76, 0x6f, 0xe6, 0x69, 0x95, 0xb9, 0xc2, 0xcc, 0x1a, 0x7e, 0x86,
	0xc7, 0xd3, 0x84, 0xa5, 0xd9, 0x25, 0x17, 0x7b, 0x43, 0xff, 0xfa, 0x28, 0x6d, 0x40, 0x02, 0x5f,
	0x54, 0xbf, 0x73, 0x67, 0x4c, 0x8c, 0xfb, 0x92, 0x59, 0xc3, 0x2e, 0x68, 0xe3, 0xfc, 0xaf, 0xf8,
	0xfe, 0x57, 0x9c, 0xc2, 0x51, 0x9f, 0x36, 0x24, 0xe8, 0x41, 0xb1, 0x2f, 0xef, 0x7e, 0xbe, 0x5d,
	0xad, 0xc5, 0x65, 0xbe, 0xb0, 0xce, 0x79, 0x6c, 0xc7, 0x3c, 0xcb, 0x23, 0x66, 0x9f, 0x93, 0xd8,
	0x99, 0xff, 0xa8, 0xc8, 0x16, 0x47, 0x6a, 0xf5, 0xf1, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab,
	0x7c, 0xe1, 0xf6, 0x55, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KVSClient is the client API for KVS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVSClient interface {
	Join(ctx context.Context, in *raft.Node, opts ...grpc.CallOption) (*empty.Empty, error)
	Leave(ctx context.Context, in *raft.Node, opts ...grpc.CallOption) (*empty.Empty, error)
	GetNode(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*raft.Node, error)
	GetCluster(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*raft.Cluster, error)
	Snapshot(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*KeyValuePair, error)
	Put(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*empty.Empty, error)
}

type kVSClient struct {
	cc *grpc.ClientConn
}

func NewKVSClient(cc *grpc.ClientConn) KVSClient {
	return &kVSClient{cc}
}

func (c *kVSClient) Join(ctx context.Context, in *raft.Node, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) Leave(ctx context.Context, in *raft.Node, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) GetNode(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*raft.Node, error) {
	out := new(raft.Node)
	err := c.cc.Invoke(ctx, "/kvs.KVS/GetNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) GetCluster(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*raft.Cluster, error) {
	out := new(raft.Cluster)
	err := c.cc.Invoke(ctx, "/kvs.KVS/GetCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) Snapshot(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Snapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) Get(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*KeyValuePair, error) {
	out := new(KeyValuePair)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) Put(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) Delete(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVSServer is the server API for KVS service.
type KVSServer interface {
	Join(context.Context, *raft.Node) (*empty.Empty, error)
	Leave(context.Context, *raft.Node) (*empty.Empty, error)
	GetNode(context.Context, *empty.Empty) (*raft.Node, error)
	GetCluster(context.Context, *empty.Empty) (*raft.Cluster, error)
	Snapshot(context.Context, *empty.Empty) (*empty.Empty, error)
	Get(context.Context, *KeyValuePair) (*KeyValuePair, error)
	Put(context.Context, *KeyValuePair) (*empty.Empty, error)
	Delete(context.Context, *KeyValuePair) (*empty.Empty, error)
}

func RegisterKVSServer(s *grpc.Server, srv KVSServer) {
	s.RegisterService(&_KVS_serviceDesc, srv)
}

func _KVS_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Join(ctx, req.(*raft.Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Leave(ctx, req.(*raft.Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).GetNode(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/GetCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).GetCluster(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_Snapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Snapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Snapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Snapshot(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValuePair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Get(ctx, req.(*KeyValuePair))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValuePair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Put(ctx, req.(*KeyValuePair))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValuePair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Delete(ctx, req.(*KeyValuePair))
	}
	return interceptor(ctx, in, info, handler)
}

var _KVS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kvs.KVS",
	HandlerType: (*KVSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _KVS_Join_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _KVS_Leave_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _KVS_GetNode_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _KVS_GetCluster_Handler,
		},
		{
			MethodName: "Snapshot",
			Handler:    _KVS_Snapshot_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _KVS_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _KVS_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _KVS_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/kvs/kvs.proto",
}
