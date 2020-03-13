// Code generated by protoc-gen-go. DO NOT EDIT.
// source: place.ext.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HotCitiesByCinemaReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HotCitiesByCinemaReq) Reset()         { *m = HotCitiesByCinemaReq{} }
func (m *HotCitiesByCinemaReq) String() string { return proto.CompactTextString(m) }
func (*HotCitiesByCinemaReq) ProtoMessage()    {}
func (*HotCitiesByCinemaReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c06df2e604fe869, []int{0}
}

func (m *HotCitiesByCinemaReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HotCitiesByCinemaReq.Unmarshal(m, b)
}
func (m *HotCitiesByCinemaReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HotCitiesByCinemaReq.Marshal(b, m, deterministic)
}
func (m *HotCitiesByCinemaReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotCitiesByCinemaReq.Merge(m, src)
}
func (m *HotCitiesByCinemaReq) XXX_Size() int {
	return xxx_messageInfo_HotCitiesByCinemaReq.Size(m)
}
func (m *HotCitiesByCinemaReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HotCitiesByCinemaReq.DiscardUnknown(m)
}

var xxx_messageInfo_HotCitiesByCinemaReq proto.InternalMessageInfo

type HotCitiesByCinemaRep struct {
	P                    []*Place `protobuf:"bytes,1,rep,name=p,proto3" json:"p,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HotCitiesByCinemaRep) Reset()         { *m = HotCitiesByCinemaRep{} }
func (m *HotCitiesByCinemaRep) String() string { return proto.CompactTextString(m) }
func (*HotCitiesByCinemaRep) ProtoMessage()    {}
func (*HotCitiesByCinemaRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c06df2e604fe869, []int{1}
}

func (m *HotCitiesByCinemaRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HotCitiesByCinemaRep.Unmarshal(m, b)
}
func (m *HotCitiesByCinemaRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HotCitiesByCinemaRep.Marshal(b, m, deterministic)
}
func (m *HotCitiesByCinemaRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotCitiesByCinemaRep.Merge(m, src)
}
func (m *HotCitiesByCinemaRep) XXX_Size() int {
	return xxx_messageInfo_HotCitiesByCinemaRep.Size(m)
}
func (m *HotCitiesByCinemaRep) XXX_DiscardUnknown() {
	xxx_messageInfo_HotCitiesByCinemaRep.DiscardUnknown(m)
}

var xxx_messageInfo_HotCitiesByCinemaRep proto.InternalMessageInfo

func (m *HotCitiesByCinemaRep) GetP() []*Place {
	if m != nil {
		return m.P
	}
	return nil
}

type Place struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	N                    string   `protobuf:"bytes,3,opt,name=n,proto3" json:"n,omitempty"`
	PinyinFull           string   `protobuf:"bytes,4,opt,name=pinyinFull,proto3" json:"pinyinFull,omitempty"`
	PinyinShort          string   `protobuf:"bytes,5,opt,name=pinyinShort,proto3" json:"pinyinShort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Place) Reset()         { *m = Place{} }
func (m *Place) String() string { return proto.CompactTextString(m) }
func (*Place) ProtoMessage()    {}
func (*Place) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c06df2e604fe869, []int{2}
}

func (m *Place) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Place.Unmarshal(m, b)
}
func (m *Place) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Place.Marshal(b, m, deterministic)
}
func (m *Place) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Place.Merge(m, src)
}
func (m *Place) XXX_Size() int {
	return xxx_messageInfo_Place.Size(m)
}
func (m *Place) XXX_DiscardUnknown() {
	xxx_messageInfo_Place.DiscardUnknown(m)
}

var xxx_messageInfo_Place proto.InternalMessageInfo

func (m *Place) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Place) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Place) GetN() string {
	if m != nil {
		return m.N
	}
	return ""
}

func (m *Place) GetPinyinFull() string {
	if m != nil {
		return m.PinyinFull
	}
	return ""
}

func (m *Place) GetPinyinShort() string {
	if m != nil {
		return m.PinyinShort
	}
	return ""
}

func init() {
	proto.RegisterType((*HotCitiesByCinemaReq)(nil), "pb.HotCitiesByCinemaReq")
	proto.RegisterType((*HotCitiesByCinemaRep)(nil), "pb.HotCitiesByCinemaRep")
	proto.RegisterType((*Place)(nil), "pb.Place")
}

func init() { proto.RegisterFile("place.ext.proto", fileDescriptor_5c06df2e604fe869) }

var fileDescriptor_5c06df2e604fe869 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x9d, 0xac, 0x2b, 0x74, 0x2a, 0x16, 0x87, 0xa2, 0xc1, 0x83, 0x84, 0x9c, 0xf6, 0x14,
	0xa1, 0xfe, 0x03, 0x8b, 0xa2, 0x37, 0x49, 0xaf, 0x5e, 0xba, 0xdb, 0x80, 0x03, 0x6b, 0x32, 0xae,
	0xa9, 0xb4, 0x07, 0xff, 0xbb, 0x34, 0xbd, 0x14, 0xac, 0xb7, 0x79, 0xdf, 0x0c, 0x33, 0xf3, 0x1e,
	0x4e, 0xa4, 0x5f, 0x76, 0xc1, 0x85, 0x4d, 0x76, 0x32, 0xa4, 0x9c, 0x48, 0x49, 0x6b, 0xaf, 0x70,
	0xfa, 0x9c, 0xf2, 0x9c, 0x33, 0x87, 0xaf, 0x87, 0xed, 0x9c, 0x63, 0xf8, 0x58, 0xfa, 0xf0, 0x69,
	0xef, 0x8e, 0x72, 0xa1, 0x6b, 0x04, 0xd1, 0x60, 0xaa, 0x66, 0x3c, 0x1b, 0x39, 0x69, 0xdd, 0xeb,
	0x6e, 0xa7, 0x07, 0xb1, 0x3f, 0x58, 0x97, 0x9a, 0xa6, 0x58, 0x77, 0x69, 0x1d, 0xb3, 0x06, 0x03,
	0x4d, 0xe5, 0xf7, 0x82, 0x2e, 0x50, 0xf1, 0x4a, 0xab, 0x82, 0x14, 0xaf, 0xe8, 0x1c, 0x21, 0xea,
	0xca, 0x40, 0x33, 0xf2, 0x10, 0xe9, 0x16, 0x51, 0x38, 0x6e, 0x39, 0x3e, 0xad, 0xfb, 0x5e, 0x9f,
	0x16, 0x7c, 0x40, 0xc8, 0xe0, 0x78, 0xaf, 0x16, 0xef, 0x69, 0xc8, 0xba, 0x2e, 0x03, 0x87, 0x68,
	0xf6, 0x86, 0x93, 0x72, 0x7e, 0x11, 0x86, 0x6f, 0xee, 0xc2, 0xe3, 0x26, 0xd3, 0x0b, 0x5e, 0xfe,
	0xb1, 0x40, 0x7a, 0xf7, 0xf4, 0x31, 0xc7, 0x37, 0xff, 0x75, 0xc4, 0x9e, 0xb4, 0x67, 0x25, 0xb0,
	0xfb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0xe3, 0x00, 0x5a, 0x43, 0x01, 0x00, 0x00,
}
