// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: due.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type S2CKickoffPlayerNotify struct {
	ErrorCode            int32    `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Uid                  uint64   `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Seq                  int32    `protobuf:"varint,4,opt,name=seq,proto3" json:"seq,omitempty"`
	Route                int32    `protobuf:"varint,5,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CKickoffPlayerNotify) Reset()         { *m = S2CKickoffPlayerNotify{} }
func (m *S2CKickoffPlayerNotify) String() string { return proto.CompactTextString(m) }
func (*S2CKickoffPlayerNotify) ProtoMessage()    {}
func (*S2CKickoffPlayerNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_23444ade4e9c50f9, []int{0}
}
func (m *S2CKickoffPlayerNotify) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *S2CKickoffPlayerNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_S2CKickoffPlayerNotify.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *S2CKickoffPlayerNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CKickoffPlayerNotify.Merge(m, src)
}
func (m *S2CKickoffPlayerNotify) XXX_Size() int {
	return m.Size()
}
func (m *S2CKickoffPlayerNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CKickoffPlayerNotify.DiscardUnknown(m)
}

var xxx_messageInfo_S2CKickoffPlayerNotify proto.InternalMessageInfo

func (m *S2CKickoffPlayerNotify) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *S2CKickoffPlayerNotify) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *S2CKickoffPlayerNotify) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *S2CKickoffPlayerNotify) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *S2CKickoffPlayerNotify) GetRoute() int32 {
	if m != nil {
		return m.Route
	}
	return 0
}

func init() {
	proto.RegisterType((*S2CKickoffPlayerNotify)(nil), "pb.s2c_kickoff_player_notify")
}

func init() { proto.RegisterFile("due.proto", fileDescriptor_23444ade4e9c50f9) }

var fileDescriptor_23444ade4e9c50f9 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x29, 0x4d, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xea, 0x62, 0xe4, 0x92, 0x2c, 0x36,
	0x4a, 0x8e, 0xcf, 0xce, 0x4c, 0xce, 0xce, 0x4f, 0x4b, 0x8b, 0x2f, 0xc8, 0x49, 0xac, 0x4c, 0x2d,
	0x8a, 0xcf, 0xcb, 0x2f, 0xc9, 0x4c, 0xab, 0x14, 0x92, 0xe5, 0xe2, 0x4a, 0x2d, 0x2a, 0xca, 0x2f,
	0x8a, 0x4f, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0xe2, 0x04, 0x8b, 0x38,
	0xe7, 0xa7, 0xa4, 0x0a, 0x89, 0x70, 0xb1, 0x82, 0x39, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x10, 0x8e, 0x90, 0x00, 0x17, 0x73, 0x69, 0x66, 0x8a, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x4b, 0x10,
	0x88, 0x09, 0x12, 0x29, 0x4e, 0x2d, 0x94, 0x60, 0x01, 0xeb, 0x07, 0x31, 0x41, 0x3a, 0x8b, 0xf2,
	0x4b, 0x4b, 0x52, 0x25, 0x58, 0xc1, 0x62, 0x10, 0x8e, 0x93, 0xd8, 0x89, 0x47, 0x72, 0x8c, 0x17,
	0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe3, 0xb1, 0x1c, 0x43, 0x14, 0x8b, 0x9e, 0x7e,
	0x41, 0x52, 0x12, 0x1b, 0xd8, 0xbd, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xda, 0x7d,
	0xca, 0xbc, 0x00, 0x00, 0x00,
}

func (m *S2CKickoffPlayerNotify) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *S2CKickoffPlayerNotify) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *S2CKickoffPlayerNotify) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Route != 0 {
		i = encodeVarintDue(dAtA, i, uint64(m.Route))
		i--
		dAtA[i] = 0x28
	}
	if m.Seq != 0 {
		i = encodeVarintDue(dAtA, i, uint64(m.Seq))
		i--
		dAtA[i] = 0x20
	}
	if m.Uid != 0 {
		i = encodeVarintDue(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintDue(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x12
	}
	if m.ErrorCode != 0 {
		i = encodeVarintDue(dAtA, i, uint64(m.ErrorCode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDue(dAtA []byte, offset int, v uint64) int {
	offset -= sovDue(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *S2CKickoffPlayerNotify) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ErrorCode != 0 {
		n += 1 + sovDue(uint64(m.ErrorCode))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovDue(uint64(l))
	}
	if m.Uid != 0 {
		n += 1 + sovDue(uint64(m.Uid))
	}
	if m.Seq != 0 {
		n += 1 + sovDue(uint64(m.Seq))
	}
	if m.Route != 0 {
		n += 1 + sovDue(uint64(m.Route))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDue(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDue(x uint64) (n int) {
	return sovDue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *S2CKickoffPlayerNotify) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDue
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
			return fmt.Errorf("proto: s2c_kickoff_player_notify: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: s2c_kickoff_player_notify: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorCode", wireType)
			}
			m.ErrorCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrorCode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDue
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
				return ErrInvalidLengthDue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seq", wireType)
			}
			m.Seq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seq |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Route", wireType)
			}
			m.Route = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Route |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDue
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
func skipDue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDue
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
					return 0, ErrIntOverflowDue
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
					return 0, ErrIntOverflowDue
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
				return 0, ErrInvalidLengthDue
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDue
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDue
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDue        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDue          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDue = fmt.Errorf("proto: unexpected end of group")
)