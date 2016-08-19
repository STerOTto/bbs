// Code generated by protoc-gen-gogo.
// source: modification_tag.proto
// DO NOT EDIT!

package models

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ModificationTag struct {
	Epoch string `protobuf:"bytes,1,opt,name=epoch" json:"epoch"`
	Index uint32 `protobuf:"varint,2,opt,name=index" json:"index"`
}

func (m *ModificationTag) Reset()                    { *m = ModificationTag{} }
func (*ModificationTag) ProtoMessage()               {}
func (*ModificationTag) Descriptor() ([]byte, []int) { return fileDescriptorModificationTag, []int{0} }

func (m *ModificationTag) GetEpoch() string {
	if m != nil {
		return m.Epoch
	}
	return ""
}

func (m *ModificationTag) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func init() {
	proto.RegisterType((*ModificationTag)(nil), "models.ModificationTag")
}
func (this *ModificationTag) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ModificationTag)
	if !ok {
		that2, ok := that.(ModificationTag)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Epoch != that1.Epoch {
		return false
	}
	if this.Index != that1.Index {
		return false
	}
	return true
}
func (this *ModificationTag) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&models.ModificationTag{")
	s = append(s, "Epoch: "+fmt.Sprintf("%#v", this.Epoch)+",\n")
	s = append(s, "Index: "+fmt.Sprintf("%#v", this.Index)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringModificationTag(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringModificationTag(m github_com_gogo_protobuf_proto.Message) string {
	e := github_com_gogo_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}
func (m *ModificationTag) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ModificationTag) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintModificationTag(data, i, uint64(len(m.Epoch)))
	i += copy(data[i:], m.Epoch)
	data[i] = 0x10
	i++
	i = encodeVarintModificationTag(data, i, uint64(m.Index))
	return i, nil
}

func encodeFixed64ModificationTag(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32ModificationTag(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintModificationTag(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ModificationTag) Size() (n int) {
	var l int
	_ = l
	l = len(m.Epoch)
	n += 1 + l + sovModificationTag(uint64(l))
	n += 1 + sovModificationTag(uint64(m.Index))
	return n
}

func sovModificationTag(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozModificationTag(x uint64) (n int) {
	return sovModificationTag(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ModificationTag) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ModificationTag{`,
		`Epoch:` + fmt.Sprintf("%v", this.Epoch) + `,`,
		`Index:` + fmt.Sprintf("%v", this.Index) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringModificationTag(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ModificationTag) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModificationTag
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ModificationTag: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModificationTag: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModificationTag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthModificationTag
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Epoch = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModificationTag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Index |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipModificationTag(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthModificationTag
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
func skipModificationTag(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModificationTag
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowModificationTag
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowModificationTag
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthModificationTag
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowModificationTag
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipModificationTag(data[start:])
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
	ErrInvalidLengthModificationTag = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModificationTag   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("modification_tag.proto", fileDescriptorModificationTag) }

var fileDescriptorModificationTag = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0xcd, 0x4f, 0xc9,
	0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x8b, 0x2f, 0x49, 0x4c, 0xd7, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x03, 0x8a, 0xa7, 0xe6, 0x14, 0x4b, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94,
	0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0xa5, 0x93, 0x4a, 0xd3,
	0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x68, 0x53, 0xf2, 0xe4, 0xe2, 0xf7, 0x45, 0x32, 0x30, 0x24,
	0x31, 0x5d, 0x48, 0x8a, 0x8b, 0x35, 0xb5, 0x20, 0x3f, 0x39, 0x43, 0x82, 0x51, 0x81, 0x51, 0x83,
	0xd3, 0x89, 0xe5, 0xc4, 0x3d, 0x79, 0x86, 0x20, 0x88, 0x10, 0x48, 0x2e, 0x33, 0x2f, 0x25, 0xb5,
	0x42, 0x82, 0x09, 0x28, 0xc7, 0x0b, 0x93, 0x03, 0x0b, 0x39, 0xe9, 0x5c, 0x78, 0x28, 0xc7, 0x70,
	0x03, 0x88, 0x3f, 0x3c, 0x94, 0x63, 0x6c, 0x78, 0x24, 0xc7, 0xb8, 0x02, 0x88, 0x4f, 0x00, 0xf1,
	0x05, 0x20, 0x7e, 0x00, 0xc4, 0x2f, 0x1e, 0x01, 0xe5, 0x80, 0xf4, 0x84, 0xc7, 0x72, 0x0c, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x08, 0x48, 0xaf, 0xc8, 0x00, 0x00, 0x00,
}
