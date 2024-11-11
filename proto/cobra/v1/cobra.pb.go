// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: proto/cobra/v1/cobra.proto

package cobrapb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Define the structure of the Cobra CLI options
type CommandOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Alias       string   `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	Summary     string   `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Aliases     []string `protobuf:"bytes,5,rep,name=aliases,proto3" json:"aliases,omitempty"`
}

func (x *CommandOptions) Reset() {
	*x = CommandOptions{}
	mi := &file_proto_cobra_v1_cobra_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommandOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandOptions) ProtoMessage() {}

func (x *CommandOptions) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cobra_v1_cobra_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandOptions.ProtoReflect.Descriptor instead.
func (*CommandOptions) Descriptor() ([]byte, []int) {
	return file_proto_cobra_v1_cobra_proto_rawDescGZIP(), []int{0}
}

func (x *CommandOptions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommandOptions) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *CommandOptions) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *CommandOptions) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CommandOptions) GetAliases() []string {
	if x != nil {
		return x.Aliases
	}
	return nil
}

// Define the structure of a flag
type FlagOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Alias       string `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	Default     string `protobuf:"bytes,4,opt,name=default,proto3" json:"default,omitempty"`
}

func (x *FlagOptions) Reset() {
	*x = FlagOptions{}
	mi := &file_proto_cobra_v1_cobra_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlagOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagOptions) ProtoMessage() {}

func (x *FlagOptions) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cobra_v1_cobra_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagOptions.ProtoReflect.Descriptor instead.
func (*FlagOptions) Descriptor() ([]byte, []int) {
	return file_proto_cobra_v1_cobra_proto_rawDescGZIP(), []int{1}
}

func (x *FlagOptions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FlagOptions) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FlagOptions) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *FlagOptions) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

var file_proto_cobra_v1_cobra_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*CommandOptions)(nil),
		Field:         50001,
		Name:          "cobra.command",
		Tag:           "bytes,50001,opt,name=command",
		Filename:      "proto/cobra/v1/cobra.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*CommandOptions)(nil),
		Field:         50002,
		Name:          "cobra.subcommand",
		Tag:           "bytes,50002,opt,name=subcommand",
		Filename:      "proto/cobra/v1/cobra.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FlagOptions)(nil),
		Field:         50003,
		Name:          "cobra.flag",
		Tag:           "bytes,50003,opt,name=flag",
		Filename:      "proto/cobra/v1/cobra.proto",
	},
}

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional cobra.CommandOptions command = 50001;
	E_Command = &file_proto_cobra_v1_cobra_proto_extTypes[0]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional cobra.CommandOptions subcommand = 50002;
	E_Subcommand = &file_proto_cobra_v1_cobra_proto_extTypes[1]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional cobra.FlagOptions flag = 50003;
	E_Flag = &file_proto_cobra_v1_cobra_proto_extTypes[2]
)

var File_proto_cobra_v1_cobra_proto protoreflect.FileDescriptor

var file_proto_cobra_v1_cobra_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x62, 0x72, 0x61, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x62, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x6f,
	0x62, 0x72, 0x61, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x22, 0x73, 0x0a, 0x0b, 0x46, 0x6c, 0x61, 0x67,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x3a, 0x52, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd1, 0x86, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x62, 0x72, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x3a, 0x57, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xd2, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x62, 0x72, 0x61, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a,
	0x73, 0x75, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x3a, 0x47, 0x0a, 0x04, 0x66, 0x6c,
	0x61, 0x67, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xd3, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x62, 0x72,
	0x61, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x04, 0x66,
	0x6c, 0x61, 0x67, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x64, 0x73, 0x6f, 0x6e, 0x6d, 0x69, 0x63, 0x68, 0x61, 0x71, 0x75, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x63, 0x6f, 0x62, 0x72, 0x61,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x62, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x3b,
	0x63, 0x6f, 0x62, 0x72, 0x61, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cobra_v1_cobra_proto_rawDescOnce sync.Once
	file_proto_cobra_v1_cobra_proto_rawDescData = file_proto_cobra_v1_cobra_proto_rawDesc
)

func file_proto_cobra_v1_cobra_proto_rawDescGZIP() []byte {
	file_proto_cobra_v1_cobra_proto_rawDescOnce.Do(func() {
		file_proto_cobra_v1_cobra_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cobra_v1_cobra_proto_rawDescData)
	})
	return file_proto_cobra_v1_cobra_proto_rawDescData
}

var file_proto_cobra_v1_cobra_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_cobra_v1_cobra_proto_goTypes = []any{
	(*CommandOptions)(nil),              // 0: cobra.CommandOptions
	(*FlagOptions)(nil),                 // 1: cobra.FlagOptions
	(*descriptorpb.ServiceOptions)(nil), // 2: google.protobuf.ServiceOptions
	(*descriptorpb.MethodOptions)(nil),  // 3: google.protobuf.MethodOptions
	(*descriptorpb.FieldOptions)(nil),   // 4: google.protobuf.FieldOptions
}
var file_proto_cobra_v1_cobra_proto_depIdxs = []int32{
	2, // 0: cobra.command:extendee -> google.protobuf.ServiceOptions
	3, // 1: cobra.subcommand:extendee -> google.protobuf.MethodOptions
	4, // 2: cobra.flag:extendee -> google.protobuf.FieldOptions
	0, // 3: cobra.command:type_name -> cobra.CommandOptions
	0, // 4: cobra.subcommand:type_name -> cobra.CommandOptions
	1, // 5: cobra.flag:type_name -> cobra.FlagOptions
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	3, // [3:6] is the sub-list for extension type_name
	0, // [0:3] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_cobra_v1_cobra_proto_init() }
func file_proto_cobra_v1_cobra_proto_init() {
	if File_proto_cobra_v1_cobra_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_cobra_v1_cobra_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_proto_cobra_v1_cobra_proto_goTypes,
		DependencyIndexes: file_proto_cobra_v1_cobra_proto_depIdxs,
		MessageInfos:      file_proto_cobra_v1_cobra_proto_msgTypes,
		ExtensionInfos:    file_proto_cobra_v1_cobra_proto_extTypes,
	}.Build()
	File_proto_cobra_v1_cobra_proto = out.File
	file_proto_cobra_v1_cobra_proto_rawDesc = nil
	file_proto_cobra_v1_cobra_proto_goTypes = nil
	file_proto_cobra_v1_cobra_proto_depIdxs = nil
}