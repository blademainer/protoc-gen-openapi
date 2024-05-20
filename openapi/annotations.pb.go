// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: openapi/annotations.proto

package openapi

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

type Parameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headers   []*Header `protobuf:"bytes,1,rep,name=headers" json:"headers,omitempty"`
	BuildTags []string  `protobuf:"bytes,2,rep,name=build_tags,json=buildTags" json:"build_tags,omitempty"`
}

func (x *Parameters) Reset() {
	*x = Parameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openapi_annotations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Parameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parameters) ProtoMessage() {}

func (x *Parameters) ProtoReflect() protoreflect.Message {
	mi := &file_openapi_annotations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parameters.ProtoReflect.Descriptor instead.
func (*Parameters) Descriptor() ([]byte, []int) {
	return file_openapi_annotations_proto_rawDescGZIP(), []int{0}
}

func (x *Parameters) GetHeaders() []*Header {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Parameters) GetBuildTags() []string {
	if x != nil {
		return x.BuildTags
	}
	return nil
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Pattern     *string `protobuf:"bytes,2,opt,name=pattern" json:"pattern,omitempty"`
	Description *string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Required    *bool   `protobuf:"varint,4,opt,name=required" json:"required,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openapi_annotations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_openapi_annotations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_openapi_annotations_proto_rawDescGZIP(), []int{1}
}

func (x *Header) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Header) GetPattern() string {
	if x != nil && x.Pattern != nil {
		return *x.Pattern
	}
	return ""
}

func (x *Header) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Header) GetRequired() bool {
	if x != nil && x.Required != nil {
		return *x.Required
	}
	return false
}

var file_openapi_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Parameters)(nil),
		Field:         66700,
		Name:          "openapi.method_params",
		Tag:           "bytes,66700,opt,name=method_params",
		Filename:      "openapi/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*Parameters)(nil),
		Field:         66701,
		Name:          "openapi.service_params",
		Tag:           "bytes,66701,opt,name=service_params",
		Filename:      "openapi/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*Parameters)(nil),
		Field:         66702,
		Name:          "openapi.file_params",
		Tag:           "bytes,66702,opt,name=file_params",
		Filename:      "openapi/annotations.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional openapi.Parameters method_params = 66700;
	E_MethodParams = &file_openapi_annotations_proto_extTypes[0]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional openapi.Parameters service_params = 66701;
	E_ServiceParams = &file_openapi_annotations_proto_extTypes[1]
)

// Extension fields to descriptorpb.FileOptions.
var (
	// optional openapi.Parameters file_params = 66702;
	E_FileParams = &file_openapi_annotations_proto_extTypes[2]
)

var File_openapi_annotations_proto protoreflect.FileDescriptor

var file_openapi_annotations_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x73, 0x22, 0x74,
	0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x3a, 0x5a, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8c, 0x89, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x3a, 0x5d, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x8d, 0x89, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x3a,
	0x54, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8e, 0x89, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6c, 0x6c, 0x61, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x3b, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
}

var (
	file_openapi_annotations_proto_rawDescOnce sync.Once
	file_openapi_annotations_proto_rawDescData = file_openapi_annotations_proto_rawDesc
)

func file_openapi_annotations_proto_rawDescGZIP() []byte {
	file_openapi_annotations_proto_rawDescOnce.Do(func() {
		file_openapi_annotations_proto_rawDescData = protoimpl.X.CompressGZIP(file_openapi_annotations_proto_rawDescData)
	})
	return file_openapi_annotations_proto_rawDescData
}

var file_openapi_annotations_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_openapi_annotations_proto_goTypes = []interface{}{
	(*Parameters)(nil),                  // 0: openapi.Parameters
	(*Header)(nil),                      // 1: openapi.Header
	(*descriptorpb.MethodOptions)(nil),  // 2: google.protobuf.MethodOptions
	(*descriptorpb.ServiceOptions)(nil), // 3: google.protobuf.ServiceOptions
	(*descriptorpb.FileOptions)(nil),    // 4: google.protobuf.FileOptions
}
var file_openapi_annotations_proto_depIdxs = []int32{
	1, // 0: openapi.Parameters.headers:type_name -> openapi.Header
	2, // 1: openapi.method_params:extendee -> google.protobuf.MethodOptions
	3, // 2: openapi.service_params:extendee -> google.protobuf.ServiceOptions
	4, // 3: openapi.file_params:extendee -> google.protobuf.FileOptions
	0, // 4: openapi.method_params:type_name -> openapi.Parameters
	0, // 5: openapi.service_params:type_name -> openapi.Parameters
	0, // 6: openapi.file_params:type_name -> openapi.Parameters
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	4, // [4:7] is the sub-list for extension type_name
	1, // [1:4] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_openapi_annotations_proto_init() }
func file_openapi_annotations_proto_init() {
	if File_openapi_annotations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_openapi_annotations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Parameters); i {
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
		file_openapi_annotations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
			RawDescriptor: file_openapi_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_openapi_annotations_proto_goTypes,
		DependencyIndexes: file_openapi_annotations_proto_depIdxs,
		MessageInfos:      file_openapi_annotations_proto_msgTypes,
		ExtensionInfos:    file_openapi_annotations_proto_extTypes,
	}.Build()
	File_openapi_annotations_proto = out.File
	file_openapi_annotations_proto_rawDesc = nil
	file_openapi_annotations_proto_goTypes = nil
	file_openapi_annotations_proto_depIdxs = nil
}
