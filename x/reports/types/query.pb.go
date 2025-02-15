// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/reports/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryPostReportsRequest is the request type for the Query/PostReports RPC
// method.
type QueryPostReportsRequest struct {
	// ID of the post to which query the reports for
	PostId string `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (m *QueryPostReportsRequest) Reset()         { *m = QueryPostReportsRequest{} }
func (m *QueryPostReportsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryPostReportsRequest) ProtoMessage()    {}
func (*QueryPostReportsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e7d62955adee61e, []int{0}
}
func (m *QueryPostReportsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPostReportsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPostReportsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPostReportsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPostReportsRequest.Merge(m, src)
}
func (m *QueryPostReportsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryPostReportsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPostReportsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPostReportsRequest proto.InternalMessageInfo

// QueryPostReportsResponse is the response type for the Query/PostReports RPC
// method.
type QueryPostReportsResponse struct {
	// relationships represent the list of all the relationships for the queried
	// user
	Reports []Report `protobuf:"bytes,1,rep,name=reports,proto3" json:"reports"`
}

func (m *QueryPostReportsResponse) Reset()         { *m = QueryPostReportsResponse{} }
func (m *QueryPostReportsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryPostReportsResponse) ProtoMessage()    {}
func (*QueryPostReportsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e7d62955adee61e, []int{1}
}
func (m *QueryPostReportsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPostReportsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPostReportsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPostReportsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPostReportsResponse.Merge(m, src)
}
func (m *QueryPostReportsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryPostReportsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPostReportsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPostReportsResponse proto.InternalMessageInfo

func (m *QueryPostReportsResponse) GetReports() []Report {
	if m != nil {
		return m.Reports
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryPostReportsRequest)(nil), "desmos.reports.v1beta1.QueryPostReportsRequest")
	proto.RegisterType((*QueryPostReportsResponse)(nil), "desmos.reports.v1beta1.QueryPostReportsResponse")
}

func init() {
	proto.RegisterFile("desmos/reports/v1beta1/query.proto", fileDescriptor_3e7d62955adee61e)
}

var fileDescriptor_3e7d62955adee61e = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0x49, 0x2d, 0xce,
	0xcd, 0x2f, 0xd6, 0x2f, 0x4a, 0x2d, 0xc8, 0x2f, 0x2a, 0x29, 0xd6, 0x2f, 0x33, 0x4c, 0x4a, 0x2d,
	0x49, 0x34, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x83, 0xa8, 0xd1, 0x83, 0xaa, 0xd1, 0x83, 0xaa, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b,
	0xd1, 0x07, 0xb1, 0x20, 0xaa, 0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b,
	0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0xa1, 0xb2, 0xca,
	0x38, 0xec, 0xcb, 0xcd, 0x4f, 0x49, 0xcd, 0x81, 0x2a, 0x52, 0xb2, 0xe1, 0x12, 0x0f, 0x04, 0xd9,
	0x1f, 0x90, 0x5f, 0x5c, 0x12, 0x04, 0x51, 0x18, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24,
	0xce, 0xc5, 0x5e, 0x90, 0x5f, 0x5c, 0x12, 0x9f, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19,
	0xc4, 0x06, 0xe2, 0x7a, 0xa6, 0x58, 0x71, 0x74, 0x2c, 0x90, 0x67, 0x78, 0xb1, 0x40, 0x9e, 0x41,
	0x29, 0x8a, 0x4b, 0x02, 0x53, 0x77, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x1d, 0x17, 0x3b,
	0xd4, 0x66, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x39, 0x3d, 0xec, 0x9e, 0xd3, 0x83, 0xe8,
	0x74, 0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21, 0x08, 0xa6, 0xc9, 0x68, 0x2d, 0x23, 0x17, 0x2b, 0xd8,
	0x70, 0xa1, 0xc5, 0x8c, 0x5c, 0xdc, 0x48, 0x36, 0x08, 0xe9, 0xe3, 0x32, 0x08, 0x87, 0x4f, 0xa4,
	0x0c, 0x88, 0xd7, 0x00, 0x71, 0xbc, 0x92, 0x61, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0xb4, 0x85, 0x34,
	0xf5, 0x71, 0x04, 0x22, 0x8c, 0x5f, 0x0d, 0x0d, 0xa2, 0x5a, 0x27, 0xf7, 0x13, 0x8f, 0xe4, 0x18,
	0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5,
	0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce,
	0xcf, 0x85, 0x1a, 0xa7, 0x9b, 0x93, 0x98, 0x54, 0x0c, 0x33, 0xba, 0x02, 0x6e, 0x58, 0x49, 0x65,
	0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0x66, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xb7,
	0x45, 0x10, 0x30, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// PostReports queries the reports for the post having the given id
	PostReports(ctx context.Context, in *QueryPostReportsRequest, opts ...grpc.CallOption) (*QueryPostReportsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) PostReports(ctx context.Context, in *QueryPostReportsRequest, opts ...grpc.CallOption) (*QueryPostReportsResponse, error) {
	out := new(QueryPostReportsResponse)
	err := c.cc.Invoke(ctx, "/desmos.reports.v1beta1.Query/PostReports", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// PostReports queries the reports for the post having the given id
	PostReports(context.Context, *QueryPostReportsRequest) (*QueryPostReportsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) PostReports(ctx context.Context, req *QueryPostReportsRequest) (*QueryPostReportsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostReports not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_PostReports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPostReportsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PostReports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.reports.v1beta1.Query/PostReports",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PostReports(ctx, req.(*QueryPostReportsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "desmos.reports.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostReports",
			Handler:    _Query_PostReports_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "desmos/reports/v1beta1/query.proto",
}

func (m *QueryPostReportsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPostReportsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPostReportsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PostId) > 0 {
		i -= len(m.PostId)
		copy(dAtA[i:], m.PostId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PostId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryPostReportsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPostReportsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPostReportsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Reports) > 0 {
		for iNdEx := len(m.Reports) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Reports[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryPostReportsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PostId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryPostReportsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Reports) > 0 {
		for _, e := range m.Reports {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryPostReportsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryPostReportsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPostReportsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryPostReportsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryPostReportsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPostReportsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reports", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reports = append(m.Reports, Report{})
			if err := m.Reports[len(m.Reports)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
