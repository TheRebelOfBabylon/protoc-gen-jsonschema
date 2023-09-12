// Custom options for protoc-gen-jsonschema
// Allocated range is 1125-1129
// See https://github.com/protocolbuffers/protobuf/blob/master/docs/options.md

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.22.2
// source: options.proto

package protoc_gen_jsonschema

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

// Custom FieldOptions
type FieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Fields tagged with this will be omitted from generated schemas
	Ignore bool `protobuf:"varint,1,opt,name=ignore,proto3" json:"ignore,omitempty"`
	// Fields tagged with this will be marked as "required" in generated schemas
	Required bool `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
	// Fields tagged with this will constrain strings using the "minLength" keyword in generated schemas
	MinLength int32 `protobuf:"varint,3,opt,name=min_length,json=minLength,proto3" json:"min_length,omitempty"`
	// Fields tagged with this will constrain strings using the "maxLength" keyword in generated schemas
	MaxLength int32 `protobuf:"varint,4,opt,name=max_length,json=maxLength,proto3" json:"max_length,omitempty"`
	// Fields tagged with this will constrain strings using the "pattern" keyword in generated schemas
	Pattern string `protobuf:"bytes,5,opt,name=pattern,proto3" json:"pattern,omitempty"`
	// Fields tagged with this will set the ref field to the given value
	Ref string `protobuf:"bytes,6,opt,name=ref,proto3" json:"ref,omitempty"`
	// Fields tagged with this will constrain arrays using the "minItems" keyword in generated schemas
	MinItems int32 `protobuf:"varint,7,opt,name=min_items,json=minItems,proto3" json:"min_items,omitempty"`
	// Fields tagged with this will constrain strings using the "format" keyword in generated schemas
	Format string `protobuf:"bytes,8,opt,name=format,proto3" json:"format,omitempty"`
}

func (x *FieldOptions) Reset() {
	*x = FieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldOptions) ProtoMessage() {}

func (x *FieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldOptions.ProtoReflect.Descriptor instead.
func (*FieldOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{0}
}

func (x *FieldOptions) GetIgnore() bool {
	if x != nil {
		return x.Ignore
	}
	return false
}

func (x *FieldOptions) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *FieldOptions) GetMinLength() int32 {
	if x != nil {
		return x.MinLength
	}
	return 0
}

func (x *FieldOptions) GetMaxLength() int32 {
	if x != nil {
		return x.MaxLength
	}
	return 0
}

func (x *FieldOptions) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

func (x *FieldOptions) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *FieldOptions) GetMinItems() int32 {
	if x != nil {
		return x.MinItems
	}
	return 0
}

func (x *FieldOptions) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

// Custom MessageOptions
type MessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Messages tagged with this will not be processed
	Ignore bool `protobuf:"varint,1,opt,name=ignore,proto3" json:"ignore,omitempty"`
	// Messages tagged with this will have all fields marked as "required":
	AllFieldsRequired bool `protobuf:"varint,2,opt,name=all_fields_required,json=allFieldsRequired,proto3" json:"all_fields_required,omitempty"`
	// Messages tagged with this will populate the id field with provided value. Default value is filename with json extension
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MessageOptions) Reset() {
	*x = MessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOptions) ProtoMessage() {}

func (x *MessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOptions.ProtoReflect.Descriptor instead.
func (*MessageOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{1}
}

func (x *MessageOptions) GetIgnore() bool {
	if x != nil {
		return x.Ignore
	}
	return false
}

func (x *MessageOptions) GetAllFieldsRequired() bool {
	if x != nil {
		return x.AllFieldsRequired
	}
	return false
}

func (x *MessageOptions) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var file_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldOptions)(nil),
		Field:         1125,
		Name:          "protoc.gen.jsonschema.field_options",
		Tag:           "bytes,1125,opt,name=field_options",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*MessageOptions)(nil),
		Field:         1127,
		Name:          "protoc.gen.jsonschema.message_options",
		Tag:           "bytes,1127,opt,name=message_options",
		Filename:      "options.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional protoc.gen.jsonschema.FieldOptions field_options = 1125;
	E_FieldOptions = &file_options_proto_extTypes[0]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional protoc.gen.jsonschema.MessageOptions message_options = 1127;
	E_MessageOptions = &file_options_proto_extTypes[1]
)

var File_options_proto protoreflect.FileDescriptor

var file_options_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x6a, 0x73, 0x6f, 0x6e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x01, 0x0a, 0x0c, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x67, 0x6e,
	0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x6d, 0x61, 0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x68, 0x0a, 0x0e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x6c, 0x6c, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x11, 0x61, 0x6c, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x68, 0x0a, 0x0d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe5, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x3a, 0x70, 0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe7, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x54, 0x68, 0x65, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x4f, 0x66, 0x42, 0x61, 0x62, 0x79, 0x6c,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6a, 0x73,
	0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_options_proto_rawDescOnce sync.Once
	file_options_proto_rawDescData = file_options_proto_rawDesc
)

func file_options_proto_rawDescGZIP() []byte {
	file_options_proto_rawDescOnce.Do(func() {
		file_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_options_proto_rawDescData)
	})
	return file_options_proto_rawDescData
}

var file_options_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_options_proto_goTypes = []interface{}{
	(*FieldOptions)(nil),                // 0: protoc.gen.jsonschema.FieldOptions
	(*MessageOptions)(nil),              // 1: protoc.gen.jsonschema.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 2: google.protobuf.FieldOptions
	(*descriptorpb.MessageOptions)(nil), // 3: google.protobuf.MessageOptions
}
var file_options_proto_depIdxs = []int32{
	2, // 0: protoc.gen.jsonschema.field_options:extendee -> google.protobuf.FieldOptions
	3, // 1: protoc.gen.jsonschema.message_options:extendee -> google.protobuf.MessageOptions
	0, // 2: protoc.gen.jsonschema.field_options:type_name -> protoc.gen.jsonschema.FieldOptions
	1, // 3: protoc.gen.jsonschema.message_options:type_name -> protoc.gen.jsonschema.MessageOptions
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_options_proto_init() }
func file_options_proto_init() {
	if File_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldOptions); i {
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
		file_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageOptions); i {
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
			RawDescriptor: file_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_options_proto_goTypes,
		DependencyIndexes: file_options_proto_depIdxs,
		MessageInfos:      file_options_proto_msgTypes,
		ExtensionInfos:    file_options_proto_extTypes,
	}.Build()
	File_options_proto = out.File
	file_options_proto_rawDesc = nil
	file_options_proto_goTypes = nil
	file_options_proto_depIdxs = nil
}
