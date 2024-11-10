// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: oneof.proto

package oneof

import (
	common "github.com/edsonmichaque/protoc-gen-cobra/testdata/oneof/common"
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

type OneOfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paging *common.PaginationRequest `protobuf:"bytes,1,opt,name=paging,proto3" json:"paging,omitempty"`
}

func (x *OneOfRequest) Reset() {
	*x = OneOfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oneof_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneOfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneOfRequest) ProtoMessage() {}

func (x *OneOfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oneof_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneOfRequest.ProtoReflect.Descriptor instead.
func (*OneOfRequest) Descriptor() ([]byte, []int) {
	return file_oneof_proto_rawDescGZIP(), []int{0}
}

func (x *OneOfRequest) GetPaging() *common.PaginationRequest {
	if x != nil {
		return x.Paging
	}
	return nil
}

type OneOfResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OneOfResponse) Reset() {
	*x = OneOfResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oneof_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneOfResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneOfResponse) ProtoMessage() {}

func (x *OneOfResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oneof_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneOfResponse.ProtoReflect.Descriptor instead.
func (*OneOfResponse) Descriptor() ([]byte, []int) {
	return file_oneof_proto_rawDescGZIP(), []int{1}
}

var File_oneof_proto protoreflect.FileDescriptor

var file_oneof_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0c, 0x4f, 0x6e, 0x65,
	0x4f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x0f, 0x0a, 0x0d,
	0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x44, 0x0a,
	0x0c, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a,
	0x05, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x12, 0x13, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x4f,
	0x6e, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6f, 0x6e,
	0x65, 0x6f, 0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oneof_proto_rawDescOnce sync.Once
	file_oneof_proto_rawDescData = file_oneof_proto_rawDesc
)

func file_oneof_proto_rawDescGZIP() []byte {
	file_oneof_proto_rawDescOnce.Do(func() {
		file_oneof_proto_rawDescData = protoimpl.X.CompressGZIP(file_oneof_proto_rawDescData)
	})
	return file_oneof_proto_rawDescData
}

var file_oneof_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_oneof_proto_goTypes = []interface{}{
	(*OneOfRequest)(nil),             // 0: oneof.OneOfRequest
	(*OneOfResponse)(nil),            // 1: oneof.OneOfResponse
	(*common.PaginationRequest)(nil), // 2: common.PaginationRequest
}
var file_oneof_proto_depIdxs = []int32{
	2, // 0: oneof.OneOfRequest.paging:type_name -> common.PaginationRequest
	0, // 1: oneof.OneOfService.OneOf:input_type -> oneof.OneOfRequest
	1, // 2: oneof.OneOfService.OneOf:output_type -> oneof.OneOfResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_oneof_proto_init() }
func file_oneof_proto_init() {
	if File_oneof_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oneof_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneOfRequest); i {
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
		file_oneof_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneOfResponse); i {
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
			RawDescriptor: file_oneof_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oneof_proto_goTypes,
		DependencyIndexes: file_oneof_proto_depIdxs,
		MessageInfos:      file_oneof_proto_msgTypes,
	}.Build()
	File_oneof_proto = out.File
	file_oneof_proto_rawDesc = nil
	file_oneof_proto_goTypes = nil
	file_oneof_proto_depIdxs = nil
}
