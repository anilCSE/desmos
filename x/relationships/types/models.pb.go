// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/relationships/v1beta1/models.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Relationship is the struct of a relationship.
// It represent the concept of "follow" of traditional social networks.
type Relationship struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty" yaml:"recipient"`
	Subspace  string `protobuf:"bytes,3,opt,name=subspace,proto3" json:"subspace,omitempty" yaml:"subspace"`
}

func (m *Relationship) Reset()         { *m = Relationship{} }
func (m *Relationship) String() string { return proto.CompactTextString(m) }
func (*Relationship) ProtoMessage()    {}
func (*Relationship) Descriptor() ([]byte, []int) {
	return fileDescriptor_8511c516358f0efc, []int{0}
}
func (m *Relationship) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Relationship) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Relationship.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Relationship) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Relationship.Merge(m, src)
}
func (m *Relationship) XXX_Size() int {
	return m.Size()
}
func (m *Relationship) XXX_DiscardUnknown() {
	xxx_messageInfo_Relationship.DiscardUnknown(m)
}

var xxx_messageInfo_Relationship proto.InternalMessageInfo

func (m *Relationship) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Relationship) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *Relationship) GetSubspace() string {
	if m != nil {
		return m.Subspace
	}
	return ""
}

// Relationships wraps a list of Relationship objects
type Relationships struct {
	Relationships []Relationship `protobuf:"bytes,1,rep,name=relationships,proto3" json:"relationships"`
}

func (m *Relationships) Reset()         { *m = Relationships{} }
func (m *Relationships) String() string { return proto.CompactTextString(m) }
func (*Relationships) ProtoMessage()    {}
func (*Relationships) Descriptor() ([]byte, []int) {
	return fileDescriptor_8511c516358f0efc, []int{1}
}
func (m *Relationships) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Relationships) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Relationships.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Relationships) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Relationships.Merge(m, src)
}
func (m *Relationships) XXX_Size() int {
	return m.Size()
}
func (m *Relationships) XXX_DiscardUnknown() {
	xxx_messageInfo_Relationships.DiscardUnknown(m)
}

var xxx_messageInfo_Relationships proto.InternalMessageInfo

func (m *Relationships) GetRelationships() []Relationship {
	if m != nil {
		return m.Relationships
	}
	return nil
}

// UserBlock represents the fact that the Blocker has blocked the given Blocked
// user.
type UserBlock struct {
	Blocker  string `protobuf:"bytes,1,opt,name=blocker,proto3" json:"blocker,omitempty" yaml:"blocker"`
	Blocked  string `protobuf:"bytes,2,opt,name=blocked,proto3" json:"blocked,omitempty" yaml:"blocked"`
	Reason   string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty" yaml:"reason"`
	Subspace string `protobuf:"bytes,4,opt,name=subspace,proto3" json:"subspace,omitempty" yaml:"subspace"`
}

func (m *UserBlock) Reset()         { *m = UserBlock{} }
func (m *UserBlock) String() string { return proto.CompactTextString(m) }
func (*UserBlock) ProtoMessage()    {}
func (*UserBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_8511c516358f0efc, []int{2}
}
func (m *UserBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBlock.Merge(m, src)
}
func (m *UserBlock) XXX_Size() int {
	return m.Size()
}
func (m *UserBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBlock.DiscardUnknown(m)
}

var xxx_messageInfo_UserBlock proto.InternalMessageInfo

func (m *UserBlock) GetBlocker() string {
	if m != nil {
		return m.Blocker
	}
	return ""
}

func (m *UserBlock) GetBlocked() string {
	if m != nil {
		return m.Blocked
	}
	return ""
}

func (m *UserBlock) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *UserBlock) GetSubspace() string {
	if m != nil {
		return m.Subspace
	}
	return ""
}

// UserBlocks wraps a list of UserBlock objects
type UserBlocks struct {
	Blocks []UserBlock `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks"`
}

func (m *UserBlocks) Reset()         { *m = UserBlocks{} }
func (m *UserBlocks) String() string { return proto.CompactTextString(m) }
func (*UserBlocks) ProtoMessage()    {}
func (*UserBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_8511c516358f0efc, []int{3}
}
func (m *UserBlocks) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserBlocks.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBlocks.Merge(m, src)
}
func (m *UserBlocks) XXX_Size() int {
	return m.Size()
}
func (m *UserBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_UserBlocks proto.InternalMessageInfo

func (m *UserBlocks) GetBlocks() []UserBlock {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func init() {
	proto.RegisterType((*Relationship)(nil), "desmos.relationships.v1beta1.Relationship")
	proto.RegisterType((*Relationships)(nil), "desmos.relationships.v1beta1.Relationships")
	proto.RegisterType((*UserBlock)(nil), "desmos.relationships.v1beta1.UserBlock")
	proto.RegisterType((*UserBlocks)(nil), "desmos.relationships.v1beta1.UserBlocks")
}

func init() {
	proto.RegisterFile("desmos/relationships/v1beta1/models.proto", fileDescriptor_8511c516358f0efc)
}

var fileDescriptor_8511c516358f0efc = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbd, 0x4e, 0xeb, 0x30,
	0x18, 0x86, 0xe3, 0xd3, 0xaa, 0xa7, 0xf5, 0x39, 0x3d, 0x3f, 0xa6, 0x43, 0x84, 0x50, 0x52, 0x79,
	0xa1, 0x45, 0x90, 0xa8, 0xed, 0xd6, 0x31, 0x12, 0x23, 0x4b, 0x10, 0x0c, 0x6c, 0xf9, 0xb1, 0xd2,
	0x88, 0xa4, 0x8e, 0xe2, 0x14, 0xd1, 0xbb, 0x60, 0x64, 0xec, 0xce, 0x8d, 0x74, 0x41, 0xea, 0xc8,
	0x14, 0xa1, 0x76, 0x61, 0xce, 0x15, 0xa0, 0x24, 0x4e, 0xda, 0x00, 0xaa, 0xc4, 0xf6, 0x29, 0xef,
	0xf3, 0x7e, 0xb1, 0x1f, 0x19, 0xf6, 0x6d, 0xc2, 0x7c, 0xca, 0xd4, 0x90, 0x78, 0x46, 0xe4, 0xd2,
	0x29, 0x9b, 0xb8, 0x01, 0x53, 0xef, 0x06, 0x26, 0x89, 0x8c, 0x81, 0xea, 0x53, 0x9b, 0x78, 0x4c,
	0x09, 0x42, 0x1a, 0x51, 0x74, 0x94, 0xa3, 0x4a, 0x05, 0x55, 0x38, 0x7a, 0xd8, 0x71, 0xa8, 0x43,
	0x33, 0x50, 0x4d, 0xa7, 0xbc, 0x83, 0x9f, 0x00, 0xfc, 0xad, 0xef, 0xf0, 0xe8, 0x14, 0xfe, 0xb4,
	0x42, 0x62, 0x44, 0x34, 0x14, 0x41, 0x17, 0xf4, 0x5a, 0x1a, 0x4a, 0x62, 0xf9, 0xcf, 0xdc, 0xf0,
	0xbd, 0x31, 0xe6, 0x01, 0xd6, 0x0b, 0x04, 0x0d, 0x61, 0x2b, 0x24, 0x96, 0x1b, 0xb8, 0x64, 0x1a,
	0x89, 0x3f, 0x32, 0xbe, 0x93, 0xc4, 0xf2, 0xbf, 0x9c, 0x2f, 0x23, 0xac, 0x6f, 0x31, 0xa4, 0xc2,
	0x26, 0x9b, 0x99, 0x2c, 0x30, 0x2c, 0x22, 0xd6, 0xb2, 0xca, 0x41, 0x12, 0xcb, 0x7f, 0xf3, 0x4a,
	0x91, 0x60, 0xbd, 0x84, 0xc6, 0xcd, 0xc7, 0x85, 0x0c, 0xde, 0x16, 0x32, 0xc0, 0x0e, 0x6c, 0xef,
	0x1e, 0x96, 0xa1, 0x6b, 0xd8, 0xae, 0xdc, 0x56, 0x04, 0xdd, 0x5a, 0xef, 0xd7, 0xf0, 0x44, 0xd9,
	0xa7, 0x42, 0xd9, 0xdd, 0xa1, 0xd5, 0x97, 0xb1, 0x2c, 0xe8, 0xd5, 0x35, 0xf8, 0x19, 0xc0, 0xd6,
	0x15, 0x23, 0xa1, 0xe6, 0x51, 0xeb, 0x36, 0x75, 0x62, 0xa6, 0x03, 0xf9, 0xc2, 0x09, 0x0f, 0xb0,
	0x5e, 0x20, 0x5b, 0xda, 0xe6, 0x46, 0x3e, 0xd1, 0x76, 0x49, 0xdb, 0xa8, 0x0f, 0x1b, 0x21, 0x31,
	0x18, 0x9d, 0x72, 0x17, 0xff, 0x93, 0x58, 0x6e, 0x17, 0xfa, 0xd2, 0xef, 0x58, 0xe7, 0x40, 0x45,
	0x5c, 0xfd, 0x7b, 0xe2, 0x2e, 0x21, 0x2c, 0xaf, 0xc3, 0xd0, 0x39, 0x6c, 0x64, 0xbf, 0x2f, 0x74,
	0x1d, 0xef, 0xd7, 0x55, 0x36, 0xb9, 0x2b, 0x5e, 0xd6, 0x2e, 0x96, 0x6b, 0x09, 0xac, 0xd6, 0x12,
	0x78, 0x5d, 0x4b, 0xe0, 0x61, 0x23, 0x09, 0xab, 0x8d, 0x24, 0xbc, 0x6c, 0x24, 0xe1, 0x66, 0xe4,
	0xb8, 0xd1, 0x64, 0x66, 0x2a, 0x16, 0xf5, 0xd5, 0x7c, 0xf5, 0x99, 0x67, 0x98, 0x8c, 0xcf, 0xea,
	0xfd, 0x87, 0xd7, 0x1c, 0xcd, 0x03, 0xc2, 0xcc, 0x46, 0xf6, 0x22, 0x47, 0xef, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x41, 0xbc, 0xa0, 0x19, 0xf2, 0x02, 0x00, 0x00,
}

func (this *Relationship) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Relationship)
	if !ok {
		that2, ok := that.(Relationship)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Creator != that1.Creator {
		return false
	}
	if this.Recipient != that1.Recipient {
		return false
	}
	if this.Subspace != that1.Subspace {
		return false
	}
	return true
}
func (this *UserBlock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UserBlock)
	if !ok {
		that2, ok := that.(UserBlock)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Blocker != that1.Blocker {
		return false
	}
	if this.Blocked != that1.Blocked {
		return false
	}
	if this.Reason != that1.Reason {
		return false
	}
	if this.Subspace != that1.Subspace {
		return false
	}
	return true
}
func (m *Relationship) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Relationship) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Relationship) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Subspace) > 0 {
		i -= len(m.Subspace)
		copy(dAtA[i:], m.Subspace)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Subspace)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Relationships) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Relationships) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Relationships) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Relationships) > 0 {
		for iNdEx := len(m.Relationships) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Relationships[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintModels(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *UserBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Subspace) > 0 {
		i -= len(m.Subspace)
		copy(dAtA[i:], m.Subspace)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Subspace)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Reason) > 0 {
		i -= len(m.Reason)
		copy(dAtA[i:], m.Reason)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Reason)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Blocked) > 0 {
		i -= len(m.Blocked)
		copy(dAtA[i:], m.Blocked)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Blocked)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Blocker) > 0 {
		i -= len(m.Blocker)
		copy(dAtA[i:], m.Blocker)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Blocker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserBlocks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserBlocks) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserBlocks) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Blocks) > 0 {
		for iNdEx := len(m.Blocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Blocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintModels(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintModels(dAtA []byte, offset int, v uint64) int {
	offset -= sovModels(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Relationship) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Subspace)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	return n
}

func (m *Relationships) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Relationships) > 0 {
		for _, e := range m.Relationships {
			l = e.Size()
			n += 1 + l + sovModels(uint64(l))
		}
	}
	return n
}

func (m *UserBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Blocker)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Blocked)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Subspace)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	return n
}

func (m *UserBlocks) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Blocks) > 0 {
		for _, e := range m.Blocks {
			l = e.Size()
			n += 1 + l + sovModels(uint64(l))
		}
	}
	return n
}

func sovModels(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModels(x uint64) (n int) {
	return sovModels(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Relationship) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModels
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
			return fmt.Errorf("proto: Relationship: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Relationship: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subspace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subspace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthModels
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthModels
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Relationships) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModels
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
			return fmt.Errorf("proto: Relationships: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Relationships: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relationships", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Relationships = append(m.Relationships, Relationship{})
			if err := m.Relationships[len(m.Relationships)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthModels
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthModels
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UserBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModels
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
			return fmt.Errorf("proto: UserBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocked", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocked = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subspace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subspace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthModels
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthModels
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UserBlocks) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModels
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
			return fmt.Errorf("proto: UserBlocks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserBlocks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocks = append(m.Blocks, UserBlock{})
			if err := m.Blocks[len(m.Blocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthModels
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthModels
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipModels(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModels
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
					return 0, ErrIntOverflowModels
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
					return 0, ErrIntOverflowModels
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
				return 0, ErrInvalidLengthModels
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModels
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModels
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModels        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModels          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModels = fmt.Errorf("proto: unexpected end of group")
)
