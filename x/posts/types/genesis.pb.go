// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/posts/v1beta1/genesis.proto

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

// GenesisState contains the data of the genesis state for the posts module
type GenesisState struct {
	Posts               []Post               `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts"`
	UsersPollAnswers    []UserAnswersEntry   `protobuf:"bytes,2,rep,name=users_poll_answers,json=usersPollAnswers,proto3" json:"users_poll_answers"`
	PostsReactions      []PostReactionsEntry `protobuf:"bytes,3,rep,name=posts_reactions,json=postsReactions,proto3" json:"posts_reactions"`
	RegisteredReactions []RegisteredReaction `protobuf:"bytes,4,rep,name=registered_reactions,json=registeredReactions,proto3" json:"registered_reactions"`
	Params              Params               `protobuf:"bytes,5,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e358c996c23f0348, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetPosts() []Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

func (m *GenesisState) GetUsersPollAnswers() []UserAnswersEntry {
	if m != nil {
		return m.UsersPollAnswers
	}
	return nil
}

func (m *GenesisState) GetPostsReactions() []PostReactionsEntry {
	if m != nil {
		return m.PostsReactions
	}
	return nil
}

func (m *GenesisState) GetRegisteredReactions() []RegisteredReaction {
	if m != nil {
		return m.RegisteredReactions
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// UserPollAnswerEntry represents an entry containing all the answers to a poll
type UserAnswersEntry struct {
	PostId      string       `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	UserAnswers []UserAnswer `protobuf:"bytes,2,rep,name=user_answers,json=userAnswers,proto3" json:"user_answers"`
}

func (m *UserAnswersEntry) Reset()         { *m = UserAnswersEntry{} }
func (m *UserAnswersEntry) String() string { return proto.CompactTextString(m) }
func (*UserAnswersEntry) ProtoMessage()    {}
func (*UserAnswersEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e358c996c23f0348, []int{1}
}
func (m *UserAnswersEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserAnswersEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserAnswersEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserAnswersEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAnswersEntry.Merge(m, src)
}
func (m *UserAnswersEntry) XXX_Size() int {
	return m.Size()
}
func (m *UserAnswersEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAnswersEntry.DiscardUnknown(m)
}

var xxx_messageInfo_UserAnswersEntry proto.InternalMessageInfo

func (m *UserAnswersEntry) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

func (m *UserAnswersEntry) GetUserAnswers() []UserAnswer {
	if m != nil {
		return m.UserAnswers
	}
	return nil
}

// PostReactionEntry represents an entry containing all the reactions to a post
type PostReactionsEntry struct {
	PostId    string         `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Reactions []PostReaction `protobuf:"bytes,2,rep,name=reactions,proto3" json:"reactions"`
}

func (m *PostReactionsEntry) Reset()         { *m = PostReactionsEntry{} }
func (m *PostReactionsEntry) String() string { return proto.CompactTextString(m) }
func (*PostReactionsEntry) ProtoMessage()    {}
func (*PostReactionsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e358c996c23f0348, []int{2}
}
func (m *PostReactionsEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostReactionsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PostReactionsEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PostReactionsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostReactionsEntry.Merge(m, src)
}
func (m *PostReactionsEntry) XXX_Size() int {
	return m.Size()
}
func (m *PostReactionsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_PostReactionsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_PostReactionsEntry proto.InternalMessageInfo

func (m *PostReactionsEntry) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

func (m *PostReactionsEntry) GetReactions() []PostReaction {
	if m != nil {
		return m.Reactions
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "desmos.posts.v1beta1.GenesisState")
	proto.RegisterType((*UserAnswersEntry)(nil), "desmos.posts.v1beta1.UserAnswersEntry")
	proto.RegisterType((*PostReactionsEntry)(nil), "desmos.posts.v1beta1.PostReactionsEntry")
}

func init() {
	proto.RegisterFile("desmos/posts/v1beta1/genesis.proto", fileDescriptor_e358c996c23f0348)
}

var fileDescriptor_e358c996c23f0348 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0x86, 0x13, 0xbb, 0x5b, 0xd9, 0xe9, 0xa2, 0xcb, 0x58, 0x30, 0x14, 0x89, 0x31, 0x88, 0x04,
	0xc4, 0x84, 0x5d, 0xc1, 0x83, 0x37, 0x17, 0xaa, 0xf4, 0x56, 0x22, 0x22, 0xf4, 0x12, 0x26, 0xcd,
	0x10, 0x03, 0x69, 0xa6, 0xcc, 0x37, 0xa9, 0xf6, 0x5f, 0x78, 0xf6, 0x17, 0xf5, 0xd8, 0xa3, 0x27,
	0x91, 0xf6, 0x8f, 0x2c, 0x99, 0x99, 0xb4, 0xa5, 0x4d, 0xdb, 0xdb, 0x84, 0x79, 0xde, 0xe7, 0x9b,
	0x79, 0x33, 0xc8, 0x4d, 0x28, 0x4c, 0x18, 0x04, 0x53, 0x06, 0x02, 0x82, 0xd9, 0x6d, 0x4c, 0x05,
	0xb9, 0x0d, 0x52, 0x5a, 0x50, 0xc8, 0xc0, 0x9f, 0x72, 0x26, 0x18, 0xee, 0x2a, 0xc6, 0x97, 0x8c,
	0xaf, 0x99, 0x5e, 0x37, 0x65, 0x29, 0x93, 0x40, 0x50, 0xad, 0x14, 0xdb, 0x73, 0x1a, 0x7d, 0x2a,
	0x79, 0x9a, 0xc8, 0xf3, 0x9a, 0x78, 0xdd, 0x48, 0x70, 0x4a, 0xc6, 0x22, 0x63, 0x45, 0x4d, 0xbd,
	0x6a, 0xf6, 0x10, 0x4e, 0x26, 0x1a, 0x71, 0xff, 0xb4, 0xd0, 0xf5, 0x17, 0x75, 0x95, 0xaf, 0x82,
	0x08, 0x8a, 0x3f, 0xa0, 0x4b, 0x89, 0x5b, 0xa6, 0xd3, 0xf2, 0x3a, 0x77, 0x3d, 0xbf, 0xe9, 0x66,
	0xfe, 0x90, 0x81, 0xb8, 0xbf, 0x58, 0xfc, 0x7b, 0x69, 0x84, 0x0a, 0xc7, 0x23, 0x84, 0x4b, 0xa0,
	0x1c, 0xa2, 0xea, 0x98, 0x11, 0x29, 0xe0, 0x27, 0xe5, 0x60, 0x3d, 0x92, 0x92, 0x37, 0xcd, 0x92,
	0x6f, 0x40, 0xf9, 0x27, 0x05, 0xf6, 0x0b, 0xc1, 0xe7, 0x5a, 0x78, 0x23, 0x3d, 0x43, 0x96, 0xe7,
	0x7a, 0x13, 0x7f, 0x47, 0x4f, 0x65, 0x32, 0xda, 0x5c, 0xd0, 0x6a, 0x49, 0xb1, 0x77, 0xfc, 0x74,
	0x61, 0x8d, 0xee, 0xaa, 0x9f, 0x48, 0x6e, 0xb3, 0x85, 0x09, 0xea, 0x72, 0x9a, 0x66, 0x20, 0x28,
	0xa7, 0xc9, 0x8e, 0xfd, 0xe2, 0x94, 0x3d, 0xdc, 0x24, 0x6a, 0x91, 0xb6, 0x3f, 0xe3, 0x07, 0x3b,
	0x80, 0x3f, 0xa2, 0xb6, 0x2a, 0xdc, 0xba, 0x74, 0x4c, 0xaf, 0x73, 0xf7, 0xe2, 0xc8, 0x91, 0x25,
	0xa3, 0x45, 0x3a, 0xe1, 0xce, 0xd0, 0xcd, 0x7e, 0x47, 0xf8, 0x39, 0x7a, 0x5c, 0x25, 0xa3, 0x2c,
	0xb1, 0x4c, 0xc7, 0xf4, 0xae, 0xc2, 0x76, 0xf5, 0x39, 0x48, 0xf0, 0x00, 0x5d, 0x57, 0xc5, 0xed,
	0x55, 0xef, 0x9c, 0xab, 0x5e, 0x8f, 0xec, 0x94, 0xdb, 0x41, 0x6e, 0x89, 0xf0, 0x61, 0x85, 0xc7,
	0x27, 0x7f, 0x46, 0x57, 0xdb, 0xea, 0xd4, 0x58, 0xf7, 0xfc, 0x8f, 0xd1, 0x83, 0xb7, 0xd1, 0xfb,
	0xfe, 0x62, 0x65, 0x9b, 0xcb, 0x95, 0x6d, 0xfe, 0x5f, 0xd9, 0xe6, 0xef, 0xb5, 0x6d, 0x2c, 0xd7,
	0xb6, 0xf1, 0x77, 0x6d, 0x1b, 0xa3, 0xb7, 0x69, 0x26, 0x7e, 0x94, 0xb1, 0x3f, 0x66, 0x93, 0x40,
	0x89, 0xdf, 0xe5, 0x24, 0x06, 0xbd, 0x0e, 0x7e, 0xe9, 0x17, 0x2e, 0xe6, 0x53, 0x0a, 0x71, 0x5b,
	0xbe, 0xec, 0xf7, 0x0f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x74, 0xdb, 0x14, 0xdc, 0xb8, 0x03, 0x00,
	0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.RegisteredReactions) > 0 {
		for iNdEx := len(m.RegisteredReactions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RegisteredReactions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.PostsReactions) > 0 {
		for iNdEx := len(m.PostsReactions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PostsReactions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.UsersPollAnswers) > 0 {
		for iNdEx := len(m.UsersPollAnswers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UsersPollAnswers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Posts) > 0 {
		for iNdEx := len(m.Posts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Posts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *UserAnswersEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserAnswersEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserAnswersEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UserAnswers) > 0 {
		for iNdEx := len(m.UserAnswers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserAnswers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.PostId) > 0 {
		i -= len(m.PostId)
		copy(dAtA[i:], m.PostId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PostId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PostReactionsEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostReactionsEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PostReactionsEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Reactions) > 0 {
		for iNdEx := len(m.Reactions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Reactions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.PostId) > 0 {
		i -= len(m.PostId)
		copy(dAtA[i:], m.PostId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PostId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Posts) > 0 {
		for _, e := range m.Posts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UsersPollAnswers) > 0 {
		for _, e := range m.UsersPollAnswers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PostsReactions) > 0 {
		for _, e := range m.PostsReactions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RegisteredReactions) > 0 {
		for _, e := range m.RegisteredReactions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *UserAnswersEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PostId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.UserAnswers) > 0 {
		for _, e := range m.UserAnswers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *PostReactionsEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PostId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Reactions) > 0 {
		for _, e := range m.Reactions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Posts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Posts = append(m.Posts, Post{})
			if err := m.Posts[len(m.Posts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsersPollAnswers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UsersPollAnswers = append(m.UsersPollAnswers, UserAnswersEntry{})
			if err := m.UsersPollAnswers[len(m.UsersPollAnswers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostsReactions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostsReactions = append(m.PostsReactions, PostReactionsEntry{})
			if err := m.PostsReactions[len(m.PostsReactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisteredReactions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RegisteredReactions = append(m.RegisteredReactions, RegisteredReaction{})
			if err := m.RegisteredReactions[len(m.RegisteredReactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *UserAnswersEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: UserAnswersEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserAnswersEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserAnswers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserAnswers = append(m.UserAnswers, UserAnswer{})
			if err := m.UserAnswers[len(m.UserAnswers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *PostReactionsEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: PostReactionsEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostReactionsEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reactions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reactions = append(m.Reactions, PostReaction{})
			if err := m.Reactions[len(m.Reactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
