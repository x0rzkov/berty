// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/p2p/push.proto

package p2p // import "berty.tech/core/api/p2p"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DevicePushType int32

const (
	DevicePushType_UnknownDevicePushType DevicePushType = 0
	DevicePushType_APNS                  DevicePushType = 1
	DevicePushType_FCM                   DevicePushType = 2
)

var DevicePushType_name = map[int32]string{
	0: "UnknownDevicePushType",
	1: "APNS",
	2: "FCM",
}
var DevicePushType_value = map[string]int32{
	"UnknownDevicePushType": 0,
	"APNS":                  1,
	"FCM":                   2,
}

func (x DevicePushType) String() string {
	return proto.EnumName(DevicePushType_name, int32(x))
}
func (DevicePushType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_push_14e365c2f17f4685, []int{0}
}

type DevicePushToDecrypted struct {
	Nonce                []byte         `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	PushType             DevicePushType `protobuf:"varint,2,opt,name=push_type,json=pushType,proto3,enum=berty.p2p.DevicePushType" json:"push_type,omitempty"`
	PushId               string         `protobuf:"bytes,3,opt,name=push_id,json=pushId,proto3" json:"push_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DevicePushToDecrypted) Reset()         { *m = DevicePushToDecrypted{} }
func (m *DevicePushToDecrypted) String() string { return proto.CompactTextString(m) }
func (*DevicePushToDecrypted) ProtoMessage()    {}
func (*DevicePushToDecrypted) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_14e365c2f17f4685, []int{0}
}
func (m *DevicePushToDecrypted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DevicePushToDecrypted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DevicePushToDecrypted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *DevicePushToDecrypted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DevicePushToDecrypted.Merge(dst, src)
}
func (m *DevicePushToDecrypted) XXX_Size() int {
	return m.Size()
}
func (m *DevicePushToDecrypted) XXX_DiscardUnknown() {
	xxx_messageInfo_DevicePushToDecrypted.DiscardUnknown(m)
}

var xxx_messageInfo_DevicePushToDecrypted proto.InternalMessageInfo

func (m *DevicePushToDecrypted) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *DevicePushToDecrypted) GetPushType() DevicePushType {
	if m != nil {
		return m.PushType
	}
	return DevicePushType_UnknownDevicePushType
}

func (m *DevicePushToDecrypted) GetPushId() string {
	if m != nil {
		return m.PushId
	}
	return ""
}

func init() {
	proto.RegisterType((*DevicePushToDecrypted)(nil), "berty.p2p.DevicePushToDecrypted")
	proto.RegisterEnum("berty.p2p.DevicePushType", DevicePushType_name, DevicePushType_value)
}
func (m *DevicePushToDecrypted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DevicePushToDecrypted) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Nonce) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPush(dAtA, i, uint64(len(m.Nonce)))
		i += copy(dAtA[i:], m.Nonce)
	}
	if m.PushType != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPush(dAtA, i, uint64(m.PushType))
	}
	if len(m.PushId) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPush(dAtA, i, uint64(len(m.PushId)))
		i += copy(dAtA[i:], m.PushId)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintPush(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DevicePushToDecrypted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovPush(uint64(l))
	}
	if m.PushType != 0 {
		n += 1 + sovPush(uint64(m.PushType))
	}
	l = len(m.PushId)
	if l > 0 {
		n += 1 + l + sovPush(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPush(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPush(x uint64) (n int) {
	return sovPush(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DevicePushToDecrypted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPush
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DevicePushToDecrypted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DevicePushToDecrypted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPush
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = append(m.Nonce[:0], dAtA[iNdEx:postIndex]...)
			if m.Nonce == nil {
				m.Nonce = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PushType", wireType)
			}
			m.PushType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PushType |= (DevicePushType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PushId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPush
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPush
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PushId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPush(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPush
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
func skipPush(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPush
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
					return 0, ErrIntOverflowPush
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPush
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPush
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPush
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPush(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPush = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPush   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api/p2p/push.proto", fileDescriptor_push_14e365c2f17f4685) }

var fileDescriptor_push_14e365c2f17f4685 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x2c, 0xc8, 0xd4,
	0x2f, 0x30, 0x2a, 0xd0, 0x2f, 0x28, 0x2d, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x4c, 0x4a, 0x2d, 0x2a, 0xa9, 0xd4, 0x2b, 0x30, 0x2a, 0x50, 0xaa, 0xe3, 0x12, 0x75, 0x49, 0x2d,
	0xcb, 0x4c, 0x4e, 0x0d, 0x28, 0x2d, 0xce, 0x08, 0xc9, 0x77, 0x49, 0x4d, 0x2e, 0xaa, 0x2c, 0x28,
	0x49, 0x4d, 0x11, 0x12, 0xe1, 0x62, 0xcd, 0xcb, 0xcf, 0x4b, 0x4e, 0x95, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x09, 0x82, 0x70, 0x84, 0xcc, 0xb8, 0x38, 0x41, 0xe6, 0xc4, 0x97, 0x54, 0x16, 0xa4, 0x4a,
	0x30, 0x29, 0x30, 0x6a, 0xf0, 0x19, 0x49, 0xea, 0xc1, 0x4d, 0xd3, 0x43, 0x32, 0xaa, 0xb2, 0x20,
	0x35, 0x88, 0xa3, 0x00, 0xca, 0x12, 0x12, 0xe7, 0x62, 0x07, 0xeb, 0xcb, 0x4c, 0x91, 0x60, 0x56,
	0x60, 0xd4, 0xe0, 0x0c, 0x62, 0x03, 0x71, 0x3d, 0x53, 0xb4, 0xec, 0xb8, 0xf8, 0x50, 0x35, 0x09,
	0x49, 0x72, 0x89, 0x86, 0xe6, 0x65, 0xe7, 0xe5, 0x97, 0xe7, 0xa1, 0x4a, 0x08, 0x30, 0x08, 0x71,
	0x70, 0xb1, 0x38, 0x06, 0xf8, 0x05, 0x0b, 0x30, 0x0a, 0xb1, 0x73, 0x31, 0xbb, 0x39, 0xfb, 0x0a,
	0x30, 0x39, 0x69, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c,
	0x33, 0x1e, 0xcb, 0x31, 0x44, 0x89, 0x43, 0x9c, 0x53, 0x92, 0x9a, 0x9c, 0xa1, 0x9f, 0x9c, 0x5f,
	0x94, 0xaa, 0x0f, 0x0d, 0x80, 0x24, 0x36, 0xb0, 0xe7, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x8f, 0xf1, 0x92, 0xae, 0x12, 0x01, 0x00, 0x00,
}
