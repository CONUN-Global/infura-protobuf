// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: v1/pool.proto

package v1

import (
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

type Pool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoolAddress string `protobuf:"bytes,1,opt,name=pool_address,json=poolAddress,proto3" json:"pool_address,omitempty"`
}

func (x *Pool) Reset() {
	*x = Pool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pool_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pool) ProtoMessage() {}

func (x *Pool) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pool_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pool.ProtoReflect.Descriptor instead.
func (*Pool) Descriptor() ([]byte, []int) {
	return file_v1_pool_proto_rawDescGZIP(), []int{0}
}

func (x *Pool) GetPoolAddress() string {
	if x != nil {
		return x.PoolAddress
	}
	return ""
}

type CreatePoolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pool *Pool `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool,omitempty"`
}

func (x *CreatePoolRequest) Reset() {
	*x = CreatePoolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pool_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePoolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePoolRequest) ProtoMessage() {}

func (x *CreatePoolRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pool_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePoolRequest.ProtoReflect.Descriptor instead.
func (*CreatePoolRequest) Descriptor() ([]byte, []int) {
	return file_v1_pool_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePoolRequest) GetPool() *Pool {
	if x != nil {
		return x.Pool
	}
	return nil
}

type CreatePoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreatePoolResponse) Reset() {
	*x = CreatePoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pool_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePoolResponse) ProtoMessage() {}

func (x *CreatePoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pool_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePoolResponse.ProtoReflect.Descriptor instead.
func (*CreatePoolResponse) Descriptor() ([]byte, []int) {
	return file_v1_pool_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePoolResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreatePoolResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_v1_pool_proto protoreflect.FileDescriptor

var file_v1_pool_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x31, 0x22, 0x29, 0x0a, 0x04, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x6f, 0x6f, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x6f, 0x6f, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x31,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x04, 0x70, 0x6f, 0x6f,
	0x6c, 0x22, 0x46, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x4c, 0x0a, 0x0b, 0x50, 0x6f, 0x6f,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_pool_proto_rawDescOnce sync.Once
	file_v1_pool_proto_rawDescData = file_v1_pool_proto_rawDesc
)

func file_v1_pool_proto_rawDescGZIP() []byte {
	file_v1_pool_proto_rawDescOnce.Do(func() {
		file_v1_pool_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_pool_proto_rawDescData)
	})
	return file_v1_pool_proto_rawDescData
}

var file_v1_pool_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_pool_proto_goTypes = []interface{}{
	(*Pool)(nil),               // 0: v1.Pool
	(*CreatePoolRequest)(nil),  // 1: v1.CreatePoolRequest
	(*CreatePoolResponse)(nil), // 2: v1.CreatePoolResponse
}
var file_v1_pool_proto_depIdxs = []int32{
	0, // 0: v1.CreatePoolRequest.pool:type_name -> v1.Pool
	1, // 1: v1.PoolService.CreatePool:input_type -> v1.CreatePoolRequest
	2, // 2: v1.PoolService.CreatePool:output_type -> v1.CreatePoolResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_pool_proto_init() }
func file_v1_pool_proto_init() {
	if File_v1_pool_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_pool_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pool); i {
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
		file_v1_pool_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePoolRequest); i {
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
		file_v1_pool_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePoolResponse); i {
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
			RawDescriptor: file_v1_pool_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_pool_proto_goTypes,
		DependencyIndexes: file_v1_pool_proto_depIdxs,
		MessageInfos:      file_v1_pool_proto_msgTypes,
	}.Build()
	File_v1_pool_proto = out.File
	file_v1_pool_proto_rawDesc = nil
	file_v1_pool_proto_goTypes = nil
	file_v1_pool_proto_depIdxs = nil
}
