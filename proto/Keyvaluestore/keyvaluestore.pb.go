// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.3
// source: keyvaluestore.proto

package keyvaluestore

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The request message containing the key
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keyvaluestore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_keyvaluestore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_keyvaluestore_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// The response message containing the value associated with the key
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keyvaluestore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_keyvaluestore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_keyvaluestore_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_keyvaluestore_proto protoreflect.FileDescriptor

var file_keyvaluestore_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6b, 0x65, 0x79, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x22, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x32, 0x33, 0x0a, 0x0d, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x6b,
	0x65, 0x79, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keyvaluestore_proto_rawDescOnce sync.Once
	file_keyvaluestore_proto_rawDescData = file_keyvaluestore_proto_rawDesc
)

func file_keyvaluestore_proto_rawDescGZIP() []byte {
	file_keyvaluestore_proto_rawDescOnce.Do(func() {
		file_keyvaluestore_proto_rawDescData = protoimpl.X.CompressGZIP(file_keyvaluestore_proto_rawDescData)
	})
	return file_keyvaluestore_proto_rawDescData
}

var file_keyvaluestore_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_keyvaluestore_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: Request
	(*Response)(nil), // 1: Response
}
var file_keyvaluestore_proto_depIdxs = []int32{
	0, // 0: KeyValueStore.GetValues:input_type -> Request
	1, // 1: KeyValueStore.GetValues:output_type -> Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_keyvaluestore_proto_init() }
func file_keyvaluestore_proto_init() {
	if File_keyvaluestore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_keyvaluestore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_keyvaluestore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_keyvaluestore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_keyvaluestore_proto_goTypes,
		DependencyIndexes: file_keyvaluestore_proto_depIdxs,
		MessageInfos:      file_keyvaluestore_proto_msgTypes,
	}.Build()
	File_keyvaluestore_proto = out.File
	file_keyvaluestore_proto_rawDesc = nil
	file_keyvaluestore_proto_goTypes = nil
	file_keyvaluestore_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeyValueStoreClient is the client API for KeyValueStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyValueStoreClient interface {
	// Provides a value for each key request
	GetValues(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type keyValueStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyValueStoreClient(cc grpc.ClientConnInterface) KeyValueStoreClient {
	return &keyValueStoreClient{cc}
}

func (c *keyValueStoreClient) GetValues(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/KeyValueStore/GetValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyValueStoreServer is the server API for KeyValueStore service.
type KeyValueStoreServer interface {
	// Provides a value for each key request
	GetValues(context.Context, *Request) (*Response, error)
}

// UnimplementedKeyValueStoreServer can be embedded to have forward compatible implementations.
type UnimplementedKeyValueStoreServer struct {
}

func (*UnimplementedKeyValueStoreServer) GetValues(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValues not implemented")
}

func RegisterKeyValueStoreServer(s *grpc.Server, srv KeyValueStoreServer) {
	s.RegisterService(&_KeyValueStore_serviceDesc, srv)
}

func _KeyValueStore_GetValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueStoreServer).GetValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KeyValueStore/GetValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueStoreServer).GetValues(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyValueStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "KeyValueStore",
	HandlerType: (*KeyValueStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetValues",
			Handler:    _KeyValueStore_GetValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keyvaluestore.proto",
}
