// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: xia/v4/xia.proto

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

type Role int32

const (
	Role_SYSTEM Role = 0
	Role_USER   Role = 1
	Role_MODEL  Role = 2
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "SYSTEM",
		1: "USER",
		2: "MODEL",
	}
	Role_value = map[string]int32{
		"SYSTEM": 0,
		"USER":   1,
		"MODEL":  2,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_xia_v4_xia_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_xia_v4_xia_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_xia_v4_xia_proto_rawDescGZIP(), []int{0}
}

type GenerationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperature float32 `protobuf:"fixed32,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	TopP        float32 `protobuf:"fixed32,2,opt,name=top_p,json=topP,proto3" json:"top_p,omitempty"`
	TopK        int32   `protobuf:"varint,3,opt,name=top_k,json=topK,proto3" json:"top_k,omitempty"`
}

func (x *GenerationConfig) Reset() {
	*x = GenerationConfig{}
	mi := &file_xia_v4_xia_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerationConfig) ProtoMessage() {}

func (x *GenerationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_xia_v4_xia_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerationConfig.ProtoReflect.Descriptor instead.
func (*GenerationConfig) Descriptor() ([]byte, []int) {
	return file_xia_v4_xia_proto_rawDescGZIP(), []int{0}
}

func (x *GenerationConfig) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *GenerationConfig) GetTopP() float32 {
	if x != nil {
		return x.TopP
	}
	return 0
}

func (x *GenerationConfig) GetTopK() int32 {
	if x != nil {
		return x.TopK
	}
	return 0
}

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//
	//	*Content_Text
	//	*Content_ImageUrl
	Content isContent_Content `protobuf_oneof:"content"`
}

func (x *Content) Reset() {
	*x = Content{}
	mi := &file_xia_v4_xia_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_xia_v4_xia_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_xia_v4_xia_proto_rawDescGZIP(), []int{1}
}

func (m *Content) GetContent() isContent_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *Content) GetText() string {
	if x, ok := x.GetContent().(*Content_Text); ok {
		return x.Text
	}
	return ""
}

func (x *Content) GetImageUrl() string {
	if x, ok := x.GetContent().(*Content_ImageUrl); ok {
		return x.ImageUrl
	}
	return ""
}

type isContent_Content interface {
	isContent_Content()
}

type Content_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

type Content_ImageUrl struct {
	ImageUrl string `protobuf:"bytes,2,opt,name=image_url,json=imageUrl,proto3,oneof"`
}

func (*Content_Text) isContent_Content() {}

func (*Content_ImageUrl) isContent_Content() {}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Role     Role       `protobuf:"varint,2,opt,name=role,proto3,enum=xmux.xia.v4.Role" json:"role,omitempty"`
	Contents []*Content `protobuf:"bytes,3,rep,name=contents,proto3" json:"contents,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_xia_v4_xia_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_xia_v4_xia_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_xia_v4_xia_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_SYSTEM
}

func (x *Message) GetContents() []*Content {
	if x != nil {
		return x.Contents
	}
	return nil
}

type ChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Model    string            `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	Config   *GenerationConfig `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	Messages []*Message        `protobuf:"bytes,4,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *ChatReq) Reset() {
	*x = ChatReq{}
	mi := &file_xia_v4_xia_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatReq) ProtoMessage() {}

func (x *ChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_xia_v4_xia_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatReq.ProtoReflect.Descriptor instead.
func (*ChatReq) Descriptor() ([]byte, []int) {
	return file_xia_v4_xia_proto_rawDescGZIP(), []int{3}
}

func (x *ChatReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChatReq) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *ChatReq) GetConfig() *GenerationConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *ChatReq) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type ChatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message *Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChatResp) Reset() {
	*x = ChatResp{}
	mi := &file_xia_v4_xia_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatResp) ProtoMessage() {}

func (x *ChatResp) ProtoReflect() protoreflect.Message {
	mi := &file_xia_v4_xia_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatResp.ProtoReflect.Descriptor instead.
func (*ChatResp) Descriptor() ([]byte, []int) {
	return file_xia_v4_xia_proto_rawDescGZIP(), []int{4}
}

func (x *ChatResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChatResp) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_xia_v4_xia_proto protoreflect.FileDescriptor

var file_xia_v4_xia_proto_rawDesc = []byte{
	0x0a, 0x10, 0x78, 0x69, 0x61, 0x2f, 0x76, 0x34, 0x2f, 0x78, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x78, 0x6d, 0x75, 0x78, 0x2e, 0x78, 0x69, 0x61, 0x2e, 0x76, 0x34, 0x22,
	0x5e, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x5f, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x50, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x6f,
	0x70, 0x5f, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x4b, 0x22,
	0x49, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x42,
	0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x72, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x78, 0x6d, 0x75, 0x78, 0x2e, 0x78, 0x69, 0x61, 0x2e, 0x76,
	0x34, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x08,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x78, 0x6d, 0x75, 0x78, 0x2e, 0x78, 0x69, 0x61, 0x2e, 0x76, 0x34, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x98,
	0x01, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x12, 0x35, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x78, 0x6d, 0x75, 0x78, 0x2e, 0x78, 0x69, 0x61, 0x2e, 0x76, 0x34, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x78, 0x6d, 0x75, 0x78,
	0x2e, 0x78, 0x69, 0x61, 0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x4a, 0x0a, 0x08, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x78, 0x6d, 0x75, 0x78, 0x2e, 0x78, 0x69,
	0x61, 0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x27, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0a, 0x0a,
	0x06, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45,
	0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x10, 0x02, 0x32, 0x77,
	0x0a, 0x03, 0x58, 0x69, 0x61, 0x12, 0x33, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x14, 0x2e,
	0x78, 0x6d, 0x75, 0x78, 0x2e, 0x78, 0x69, 0x61, 0x2e, 0x76, 0x34, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x78, 0x6d, 0x75, 0x78, 0x2e, 0x78, 0x69, 0x61, 0x2e, 0x76,
	0x34, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x0a, 0x43, 0x68,
	0x61, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x14, 0x2e, 0x78, 0x6d, 0x75, 0x78, 0x2e,
	0x78, 0x69, 0x61, 0x2e, 0x76, 0x34, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15,
	0x2e, 0x78, 0x6d, 0x75, 0x78, 0x2e, 0x78, 0x69, 0x61, 0x2e, 0x76, 0x34, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x30, 0x01, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x2e, 0x78,
	0x64, 0x65, 0x61, 0x2e, 0x78, 0x79, 0x7a, 0x2f, 0x58, 0x4d, 0x55, 0x53, 0x2f, 0x78, 0x69, 0x61,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x78, 0x69, 0x61, 0x2f, 0x76, 0x34, 0x3b, 0x76, 0x34, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xia_v4_xia_proto_rawDescOnce sync.Once
	file_xia_v4_xia_proto_rawDescData = file_xia_v4_xia_proto_rawDesc
)

func file_xia_v4_xia_proto_rawDescGZIP() []byte {
	file_xia_v4_xia_proto_rawDescOnce.Do(func() {
		file_xia_v4_xia_proto_rawDescData = protoimpl.X.CompressGZIP(file_xia_v4_xia_proto_rawDescData)
	})
	return file_xia_v4_xia_proto_rawDescData
}

var file_xia_v4_xia_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_xia_v4_xia_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_xia_v4_xia_proto_goTypes = []any{
	(Role)(0),                // 0: xmux.xia.v4.Role
	(*GenerationConfig)(nil), // 1: xmux.xia.v4.GenerationConfig
	(*Content)(nil),          // 2: xmux.xia.v4.Content
	(*Message)(nil),          // 3: xmux.xia.v4.Message
	(*ChatReq)(nil),          // 4: xmux.xia.v4.ChatReq
	(*ChatResp)(nil),         // 5: xmux.xia.v4.ChatResp
}
var file_xia_v4_xia_proto_depIdxs = []int32{
	0, // 0: xmux.xia.v4.Message.role:type_name -> xmux.xia.v4.Role
	2, // 1: xmux.xia.v4.Message.contents:type_name -> xmux.xia.v4.Content
	1, // 2: xmux.xia.v4.ChatReq.config:type_name -> xmux.xia.v4.GenerationConfig
	3, // 3: xmux.xia.v4.ChatReq.messages:type_name -> xmux.xia.v4.Message
	3, // 4: xmux.xia.v4.ChatResp.message:type_name -> xmux.xia.v4.Message
	4, // 5: xmux.xia.v4.Xia.Chat:input_type -> xmux.xia.v4.ChatReq
	4, // 6: xmux.xia.v4.Xia.ChatStream:input_type -> xmux.xia.v4.ChatReq
	5, // 7: xmux.xia.v4.Xia.Chat:output_type -> xmux.xia.v4.ChatResp
	5, // 8: xmux.xia.v4.Xia.ChatStream:output_type -> xmux.xia.v4.ChatResp
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_xia_v4_xia_proto_init() }
func file_xia_v4_xia_proto_init() {
	if File_xia_v4_xia_proto != nil {
		return
	}
	file_xia_v4_xia_proto_msgTypes[1].OneofWrappers = []any{
		(*Content_Text)(nil),
		(*Content_ImageUrl)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_xia_v4_xia_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_xia_v4_xia_proto_goTypes,
		DependencyIndexes: file_xia_v4_xia_proto_depIdxs,
		EnumInfos:         file_xia_v4_xia_proto_enumTypes,
		MessageInfos:      file_xia_v4_xia_proto_msgTypes,
	}.Build()
	File_xia_v4_xia_proto = out.File
	file_xia_v4_xia_proto_rawDesc = nil
	file_xia_v4_xia_proto_goTypes = nil
	file_xia_v4_xia_proto_depIdxs = nil
}
