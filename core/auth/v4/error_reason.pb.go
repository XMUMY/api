// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: api/auth/v4/error_reason.proto

package v4

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

type ErrorReason int32

const (
	ErrorReason_UNKNOWN                 ErrorReason = 0
	ErrorReason_INVALID_CREDENTIAL      ErrorReason = 1
	ErrorReason_INVALID_CAMPUS_ID       ErrorReason = 2
	ErrorReason_INVALID_PASSWORD        ErrorReason = 3
	ErrorReason_WRONG_PASSWORD          ErrorReason = 4
	ErrorReason_LDAP_SERVER_DOWN        ErrorReason = 5
	ErrorReason_WRONG_FIREBASE_ID_TOKEN ErrorReason = 6
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "UNKNOWN",
		1: "INVALID_CREDENTIAL",
		2: "INVALID_CAMPUS_ID",
		3: "INVALID_PASSWORD",
		4: "WRONG_PASSWORD",
		5: "LDAP_SERVER_DOWN",
		6: "WRONG_FIREBASE_ID_TOKEN",
	}
	ErrorReason_value = map[string]int32{
		"UNKNOWN":                 0,
		"INVALID_CREDENTIAL":      1,
		"INVALID_CAMPUS_ID":       2,
		"INVALID_PASSWORD":        3,
		"WRONG_PASSWORD":          4,
		"LDAP_SERVER_DOWN":        5,
		"WRONG_FIREBASE_ID_TOKEN": 6,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_auth_v4_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_api_auth_v4_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_api_auth_v4_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_api_auth_v4_error_reason_proto protoreflect.FileDescriptor

var file_api_auth_v4_error_reason_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x34, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x78, 0x6d, 0x75, 0x78, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x34, 0x2a, 0xa6, 0x01, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x52, 0x45, 0x44,
	0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x55, 0x53, 0x5f, 0x49, 0x44, 0x10, 0x02, 0x12,
	0x14, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57,
	0x4f, 0x52, 0x44, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x57, 0x52, 0x4f, 0x4e, 0x47, 0x5f, 0x50,
	0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x44, 0x41,
	0x50, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x05, 0x12,
	0x1b, 0x0a, 0x17, 0x57, 0x52, 0x4f, 0x4e, 0x47, 0x5f, 0x46, 0x49, 0x52, 0x45, 0x42, 0x41, 0x53,
	0x45, 0x5f, 0x49, 0x44, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x06, 0x42, 0x23, 0x5a, 0x21,
	0x67, 0x69, 0x74, 0x2e, 0x78, 0x64, 0x65, 0x61, 0x2e, 0x78, 0x79, 0x7a, 0x2f, 0x58, 0x4d, 0x55,
	0x53, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x34, 0x3b, 0x76,
	0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_auth_v4_error_reason_proto_rawDescOnce sync.Once
	file_api_auth_v4_error_reason_proto_rawDescData = file_api_auth_v4_error_reason_proto_rawDesc
)

func file_api_auth_v4_error_reason_proto_rawDescGZIP() []byte {
	file_api_auth_v4_error_reason_proto_rawDescOnce.Do(func() {
		file_api_auth_v4_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_auth_v4_error_reason_proto_rawDescData)
	})
	return file_api_auth_v4_error_reason_proto_rawDescData
}

var file_api_auth_v4_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_auth_v4_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: xmux.core.auth.v4.ErrorReason
}
var file_api_auth_v4_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_auth_v4_error_reason_proto_init() }
func file_api_auth_v4_error_reason_proto_init() {
	if File_api_auth_v4_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_auth_v4_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_auth_v4_error_reason_proto_goTypes,
		DependencyIndexes: file_api_auth_v4_error_reason_proto_depIdxs,
		EnumInfos:         file_api_auth_v4_error_reason_proto_enumTypes,
	}.Build()
	File_api_auth_v4_error_reason_proto = out.File
	file_api_auth_v4_error_reason_proto_rawDesc = nil
	file_api_auth_v4_error_reason_proto_goTypes = nil
	file_api_auth_v4_error_reason_proto_depIdxs = nil
}
