// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chainmain/supply/v1/accounts.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// VestingAccounts stored in keeper
type VestingAccounts struct {
	// addresses defines addresses of all the vesting accounts at genesis
	Addresses []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (m *VestingAccounts) Reset()         { *m = VestingAccounts{} }
func (m *VestingAccounts) String() string { return proto.CompactTextString(m) }
func (*VestingAccounts) ProtoMessage()    {}
func (*VestingAccounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cfb94ff8b239338, []int{0}
}
func (m *VestingAccounts) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VestingAccounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VestingAccounts.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VestingAccounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VestingAccounts.Merge(m, src)
}
func (m *VestingAccounts) XXX_Size() int {
	return m.Size()
}
func (m *VestingAccounts) XXX_DiscardUnknown() {
	xxx_messageInfo_VestingAccounts.DiscardUnknown(m)
}

var xxx_messageInfo_VestingAccounts proto.InternalMessageInfo

func (m *VestingAccounts) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func init() {
	proto.RegisterType((*VestingAccounts)(nil), "chainmain.supply.v1.VestingAccounts")
}

func init() {
	proto.RegisterFile("chainmain/supply/v1/accounts.proto", fileDescriptor_7cfb94ff8b239338)
}

var fileDescriptor_7cfb94ff8b239338 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0xce, 0x48, 0xcc,
	0xcc, 0xcb, 0x4d, 0xcc, 0xcc, 0xd3, 0x2f, 0x2e, 0x2d, 0x28, 0xc8, 0xa9, 0xd4, 0x2f, 0x33, 0xd4,
	0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x86, 0xab, 0xd1, 0x83, 0xa8, 0xd1, 0x2b, 0x33, 0x54, 0xd2, 0xe7, 0xe2, 0x0f, 0x4b, 0x2d, 0x2e,
	0xc9, 0xcc, 0x4b, 0x77, 0x84, 0xaa, 0x16, 0x92, 0xe1, 0xe2, 0x4c, 0x4c, 0x49, 0x29, 0x4a, 0x2d,
	0x2e, 0x4e, 0x2d, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x0c, 0x42, 0x08, 0x38, 0xf9, 0x9f, 0x78,
	0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c,
	0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x69, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92,
	0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x72, 0x51, 0x65, 0x41, 0x49, 0xbe, 0x6e, 0x7e, 0x51, 0xba, 0x2e,
	0xd8, 0x56, 0x7d, 0x30, 0xa9, 0x0b, 0x76, 0x60, 0x05, 0xcc, 0x89, 0x25, 0x95, 0x05, 0xa9, 0xc5,
	0x49, 0x6c, 0x60, 0xd7, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x0b, 0x18, 0xc1, 0xc3,
	0x00, 0x00, 0x00,
}

func (m *VestingAccounts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VestingAccounts) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VestingAccounts) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Addresses[iNdEx])
			copy(dAtA[i:], m.Addresses[iNdEx])
			i = encodeVarintAccounts(dAtA, i, uint64(len(m.Addresses[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccounts(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccounts(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VestingAccounts) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovAccounts(uint64(l))
		}
	}
	return n
}

func sovAccounts(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccounts(x uint64) (n int) {
	return sovAccounts(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VestingAccounts) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccounts
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
			return fmt.Errorf("proto: VestingAccounts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VestingAccounts: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccounts
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
				return ErrInvalidLengthAccounts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccounts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccounts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccounts
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
func skipAccounts(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccounts
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
					return 0, ErrIntOverflowAccounts
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
					return 0, ErrIntOverflowAccounts
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
				return 0, ErrInvalidLengthAccounts
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccounts
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccounts
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccounts        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccounts          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccounts = fmt.Errorf("proto: unexpected end of group")
)
