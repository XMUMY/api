// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetTimetableResp struct {
	Timetable            []*TimetableClass `protobuf:"bytes,1,rep,name=timetable,proto3" json:"timetable,omitempty"`
	RecentUpdateS        string            `protobuf:"bytes,2,opt,name=recentUpdateS,proto3" json:"recentUpdateS,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetTimetableResp) Reset()         { *m = GetTimetableResp{} }
func (m *GetTimetableResp) String() string { return proto.CompactTextString(m) }
func (*GetTimetableResp) ProtoMessage()    {}
func (*GetTimetableResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}
func (m *GetTimetableResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTimetableResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTimetableResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTimetableResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTimetableResp.Merge(m, src)
}
func (m *GetTimetableResp) XXX_Size() int {
	return m.Size()
}
func (m *GetTimetableResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTimetableResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetTimetableResp proto.InternalMessageInfo

type TimetableClass struct {
	Cid                  string   `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lecturer             string   `protobuf:"bytes,3,opt,name=lecturer,proto3" json:"lecturer,omitempty"`
	Room                 string   `protobuf:"bytes,4,opt,name=room,proto3" json:"room,omitempty"`
	Day                  int32    `protobuf:"varint,5,opt,name=day,proto3" json:"day,omitempty"`
	Start                string   `protobuf:"bytes,6,opt,name=start,proto3" json:"start,omitempty"`
	End                  string   `protobuf:"bytes,7,opt,name=end,proto3" json:"end,omitempty"`
	StartDay             string   `protobuf:"bytes,8,opt,name=startDay,proto3" json:"startDay,omitempty"`
	EndDay               string   `protobuf:"bytes,9,opt,name=endDay,proto3" json:"endDay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimetableClass) Reset()         { *m = TimetableClass{} }
func (m *TimetableClass) String() string { return proto.CompactTextString(m) }
func (*TimetableClass) ProtoMessage()    {}
func (*TimetableClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}
func (m *TimetableClass) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimetableClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimetableClass.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimetableClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimetableClass.Merge(m, src)
}
func (m *TimetableClass) XXX_Size() int {
	return m.Size()
}
func (m *TimetableClass) XXX_DiscardUnknown() {
	xxx_messageInfo_TimetableClass.DiscardUnknown(m)
}

var xxx_messageInfo_TimetableClass proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetTimetableResp)(nil), "xmux.base.bridge.v3.GetTimetableResp")
	proto.RegisterType((*TimetableClass)(nil), "xmux.base.bridge.v3.TimetableClass")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0xae, 0xd3, 0x30,
	0x10, 0xc6, 0x9f, 0x5f, 0xda, 0xf0, 0x62, 0xfe, 0xa8, 0x32, 0xe8, 0xc9, 0x04, 0x14, 0x55, 0x05,
	0xa4, 0x6e, 0x70, 0xa4, 0x57, 0x89, 0x3d, 0x05, 0xc4, 0x16, 0x05, 0xd8, 0xb0, 0x73, 0x92, 0xc1,
	0x58, 0x4a, 0xec, 0x28, 0x71, 0xaa, 0x56, 0xec, 0xb8, 0x02, 0x1b, 0xae, 0xc0, 0x4d, 0xba, 0x44,
	0x82, 0x03, 0x40, 0xe1, 0x20, 0xc8, 0x4e, 0xda, 0x52, 0xd4, 0xee, 0xbe, 0x6f, 0xbe, 0x9f, 0x67,
	0xac, 0x19, 0x1c, 0xf0, 0x4a, 0xb2, 0xaa, 0xd6, 0x46, 0x93, 0xdb, 0xcb, 0xb2, 0x5d, 0xb2, 0x94,
	0x37, 0xc0, 0xd2, 0x5a, 0xe6, 0x02, 0xd8, 0x62, 0x16, 0x3e, 0x16, 0xd2, 0x7c, 0x68, 0x53, 0x96,
	0xe9, 0x32, 0x16, 0x5a, 0xe8, 0xd8, 0xb1, 0x69, 0xfb, 0xde, 0x39, 0x67, 0x9c, 0xea, 0x7a, 0x84,
	0xf7, 0x84, 0xd6, 0xa2, 0x80, 0x3d, 0x05, 0x65, 0x65, 0x56, 0x7d, 0x78, 0xbf, 0x0f, 0x79, 0x25,
	0x63, 0xae, 0x94, 0x36, 0xdc, 0x48, 0xad, 0x9a, 0x2e, 0x9d, 0x7c, 0xc4, 0xa3, 0x97, 0x60, 0xde,
	0xc8, 0x12, 0x0c, 0x4f, 0x0b, 0x48, 0xa0, 0xa9, 0xc8, 0x53, 0x1c, 0x98, 0x6d, 0x81, 0xa2, 0xb1,
	0x37, 0xbd, 0x7e, 0xf5, 0x80, 0x1d, 0xf9, 0x26, 0xdb, 0x3d, 0x7b, 0x56, 0xf0, 0xa6, 0x49, 0xf6,
	0xaf, 0xc8, 0x43, 0x7c, 0xb3, 0x86, 0x0c, 0x94, 0x79, 0x5b, 0xe5, 0xdc, 0xc0, 0x6b, 0x7a, 0x3e,
	0x46, 0xd3, 0x20, 0x39, 0x2c, 0x4e, 0x7e, 0x20, 0x7c, 0xeb, 0xb0, 0x07, 0x19, 0x61, 0x2f, 0x93,
	0x39, 0x45, 0x0e, 0xb7, 0x92, 0x10, 0x3c, 0x50, 0xbc, 0x84, 0xbe, 0x83, 0xd3, 0x24, 0xc4, 0x17,
	0x05, 0x64, 0xa6, 0xad, 0xa1, 0xa6, 0x9e, 0xab, 0xef, 0xbc, 0xe5, 0x6b, 0xad, 0x4b, 0x3a, 0xe8,
	0x78, 0xab, 0x6d, 0xd7, 0x9c, 0xaf, 0xe8, 0x70, 0x8c, 0xa6, 0xc3, 0xc4, 0x4a, 0x72, 0x07, 0x0f,
	0x1b, 0xc3, 0x6b, 0x43, 0x7d, 0x87, 0x75, 0xc6, 0x72, 0xa0, 0x72, 0x7a, 0xad, 0x9b, 0x0e, 0x2a,
	0xb7, 0x93, 0x5c, 0xf4, 0x9c, 0xaf, 0xe8, 0x45, 0x37, 0x69, 0xeb, 0xc9, 0x25, 0xf6, 0x41, 0xe5,
	0x36, 0x09, 0x5c, 0xd2, 0xbb, 0xab, 0xaf, 0x08, 0xfb, 0x73, 0xb7, 0x24, 0x22, 0xf0, 0x8d, 0x7f,
	0xd7, 0x4b, 0x2e, 0x59, 0x77, 0x0d, 0xb6, 0x3d, 0x15, 0x7b, 0x61, 0x4f, 0x15, 0x3e, 0x3a, 0xba,
	0xdf, 0xff, 0x2f, 0x33, 0xa1, 0x9f, 0xbe, 0xff, 0xf9, 0x7c, 0x4e, 0xc8, 0x28, 0x5e, 0xcc, 0x62,
	0x9e, 0xc5, 0xfb, 0x85, 0x3f, 0xc1, 0x83, 0x57, 0x52, 0x89, 0x93, 0x03, 0x4e, 0xd4, 0xe7, 0x77,
	0xd7, 0xbf, 0xa2, 0xb3, 0xf5, 0x26, 0x42, 0xdf, 0x36, 0x11, 0xfa, 0xb9, 0x89, 0xd0, 0x97, 0xdf,
	0xd1, 0xd9, 0x3b, 0x8f, 0x57, 0x32, 0xf5, 0x1d, 0x3a, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xfb,
	0xb5, 0x53, 0xfc, 0xad, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BridgeClient is the client API for Bridge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BridgeClient interface {
	GetTimetable(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetTimetableResp, error)
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type bridgeClient struct {
	cc *grpc.ClientConn
}

func NewBridgeClient(cc *grpc.ClientConn) BridgeClient {
	return &bridgeClient{cc}
}

func (c *bridgeClient) GetTimetable(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetTimetableResp, error) {
	out := new(GetTimetableResp)
	err := c.cc.Invoke(ctx, "/xmux.base.bridge.v3.Bridge/GetTimetable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/xmux.base.bridge.v3.Bridge/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BridgeServer is the server API for Bridge service.
type BridgeServer interface {
	GetTimetable(context.Context, *empty.Empty) (*GetTimetableResp, error)
	Ping(context.Context, *empty.Empty) (*empty.Empty, error)
}

// UnimplementedBridgeServer can be embedded to have forward compatible implementations.
type UnimplementedBridgeServer struct {
}

func (*UnimplementedBridgeServer) GetTimetable(ctx context.Context, req *empty.Empty) (*GetTimetableResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimetable not implemented")
}
func (*UnimplementedBridgeServer) Ping(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterBridgeServer(s *grpc.Server, srv BridgeServer) {
	s.RegisterService(&_Bridge_serviceDesc, srv)
}

func _Bridge_GetTimetable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServer).GetTimetable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xmux.base.bridge.v3.Bridge/GetTimetable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServer).GetTimetable(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bridge_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xmux.base.bridge.v3.Bridge/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bridge_serviceDesc = grpc.ServiceDesc{
	ServiceName: "xmux.base.bridge.v3.Bridge",
	HandlerType: (*BridgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTimetable",
			Handler:    _Bridge_GetTimetable_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Bridge_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func (m *GetTimetableResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTimetableResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTimetableResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.RecentUpdateS) > 0 {
		i -= len(m.RecentUpdateS)
		copy(dAtA[i:], m.RecentUpdateS)
		i = encodeVarintApi(dAtA, i, uint64(len(m.RecentUpdateS)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Timetable) > 0 {
		for iNdEx := len(m.Timetable) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Timetable[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintApi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *TimetableClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimetableClass) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimetableClass) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.EndDay) > 0 {
		i -= len(m.EndDay)
		copy(dAtA[i:], m.EndDay)
		i = encodeVarintApi(dAtA, i, uint64(len(m.EndDay)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.StartDay) > 0 {
		i -= len(m.StartDay)
		copy(dAtA[i:], m.StartDay)
		i = encodeVarintApi(dAtA, i, uint64(len(m.StartDay)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.End) > 0 {
		i -= len(m.End)
		copy(dAtA[i:], m.End)
		i = encodeVarintApi(dAtA, i, uint64(len(m.End)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Start) > 0 {
		i -= len(m.Start)
		copy(dAtA[i:], m.Start)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Start)))
		i--
		dAtA[i] = 0x32
	}
	if m.Day != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Day))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Room) > 0 {
		i -= len(m.Room)
		copy(dAtA[i:], m.Room)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Room)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Lecturer) > 0 {
		i -= len(m.Lecturer)
		copy(dAtA[i:], m.Lecturer)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Lecturer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	offset -= sovApi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetTimetableResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Timetable) > 0 {
		for _, e := range m.Timetable {
			l = e.Size()
			n += 1 + l + sovApi(uint64(l))
		}
	}
	l = len(m.RecentUpdateS)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TimetableClass) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Lecturer)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Room)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.Day != 0 {
		n += 1 + sovApi(uint64(m.Day))
	}
	l = len(m.Start)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.End)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.StartDay)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.EndDay)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetTimetableResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetTimetableResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTimetableResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timetable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timetable = append(m.Timetable, &TimetableClass{})
			if err := m.Timetable[len(m.Timetable)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecentUpdateS", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecentUpdateS = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TimetableClass) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TimetableClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimetableClass: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lecturer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lecturer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Room", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Room = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Day", wireType)
			}
			m.Day = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Day |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Start = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.End = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartDay", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartDay = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDay", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndDay = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthApi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApi = fmt.Errorf("proto: unexpected end of group")
)
