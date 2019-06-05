// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/matcher/value.proto

package matcher

import (
	fmt "fmt"
	io "io"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Specifies the way to match a ProtobufWkt::Value. Primitive values and ListValue are supported.
// StructValue is not supported and is always not matched.
type ValueMatcher struct {
	// Specifies how to match a value.
	//
	// Types that are valid to be assigned to MatchPattern:
	//	*ValueMatcher_NullMatch_
	//	*ValueMatcher_DoubleMatch
	//	*ValueMatcher_StringMatch
	//	*ValueMatcher_BoolMatch
	//	*ValueMatcher_PresentMatch
	//	*ValueMatcher_ListMatch
	MatchPattern         isValueMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ValueMatcher) Reset()         { *m = ValueMatcher{} }
func (m *ValueMatcher) String() string { return proto.CompactTextString(m) }
func (*ValueMatcher) ProtoMessage()    {}
func (*ValueMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_145b36501d266253, []int{0}
}
func (m *ValueMatcher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValueMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValueMatcher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValueMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueMatcher.Merge(m, src)
}
func (m *ValueMatcher) XXX_Size() int {
	return m.Size()
}
func (m *ValueMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_ValueMatcher proto.InternalMessageInfo

type isValueMatcher_MatchPattern interface {
	isValueMatcher_MatchPattern()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ValueMatcher_NullMatch_ struct {
	NullMatch *ValueMatcher_NullMatch `protobuf:"bytes,1,opt,name=null_match,json=nullMatch,proto3,oneof"`
}
type ValueMatcher_DoubleMatch struct {
	DoubleMatch *DoubleMatcher `protobuf:"bytes,2,opt,name=double_match,json=doubleMatch,proto3,oneof"`
}
type ValueMatcher_StringMatch struct {
	StringMatch *StringMatcher `protobuf:"bytes,3,opt,name=string_match,json=stringMatch,proto3,oneof"`
}
type ValueMatcher_BoolMatch struct {
	BoolMatch bool `protobuf:"varint,4,opt,name=bool_match,json=boolMatch,proto3,oneof"`
}
type ValueMatcher_PresentMatch struct {
	PresentMatch bool `protobuf:"varint,5,opt,name=present_match,json=presentMatch,proto3,oneof"`
}
type ValueMatcher_ListMatch struct {
	ListMatch *ListMatcher `protobuf:"bytes,6,opt,name=list_match,json=listMatch,proto3,oneof"`
}

func (*ValueMatcher_NullMatch_) isValueMatcher_MatchPattern()   {}
func (*ValueMatcher_DoubleMatch) isValueMatcher_MatchPattern()  {}
func (*ValueMatcher_StringMatch) isValueMatcher_MatchPattern()  {}
func (*ValueMatcher_BoolMatch) isValueMatcher_MatchPattern()    {}
func (*ValueMatcher_PresentMatch) isValueMatcher_MatchPattern() {}
func (*ValueMatcher_ListMatch) isValueMatcher_MatchPattern()    {}

func (m *ValueMatcher) GetMatchPattern() isValueMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *ValueMatcher) GetNullMatch() *ValueMatcher_NullMatch {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_NullMatch_); ok {
		return x.NullMatch
	}
	return nil
}

func (m *ValueMatcher) GetDoubleMatch() *DoubleMatcher {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_DoubleMatch); ok {
		return x.DoubleMatch
	}
	return nil
}

func (m *ValueMatcher) GetStringMatch() *StringMatcher {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_StringMatch); ok {
		return x.StringMatch
	}
	return nil
}

func (m *ValueMatcher) GetBoolMatch() bool {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_BoolMatch); ok {
		return x.BoolMatch
	}
	return false
}

func (m *ValueMatcher) GetPresentMatch() bool {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_PresentMatch); ok {
		return x.PresentMatch
	}
	return false
}

func (m *ValueMatcher) GetListMatch() *ListMatcher {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_ListMatch); ok {
		return x.ListMatch
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ValueMatcher) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ValueMatcher_OneofMarshaler, _ValueMatcher_OneofUnmarshaler, _ValueMatcher_OneofSizer, []interface{}{
		(*ValueMatcher_NullMatch_)(nil),
		(*ValueMatcher_DoubleMatch)(nil),
		(*ValueMatcher_StringMatch)(nil),
		(*ValueMatcher_BoolMatch)(nil),
		(*ValueMatcher_PresentMatch)(nil),
		(*ValueMatcher_ListMatch)(nil),
	}
}

func _ValueMatcher_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ValueMatcher)
	// match_pattern
	switch x := m.MatchPattern.(type) {
	case *ValueMatcher_NullMatch_:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NullMatch); err != nil {
			return err
		}
	case *ValueMatcher_DoubleMatch:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DoubleMatch); err != nil {
			return err
		}
	case *ValueMatcher_StringMatch:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StringMatch); err != nil {
			return err
		}
	case *ValueMatcher_BoolMatch:
		t := uint64(0)
		if x.BoolMatch {
			t = 1
		}
		_ = b.EncodeVarint(4<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case *ValueMatcher_PresentMatch:
		t := uint64(0)
		if x.PresentMatch {
			t = 1
		}
		_ = b.EncodeVarint(5<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case *ValueMatcher_ListMatch:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ListMatch); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ValueMatcher.MatchPattern has unexpected type %T", x)
	}
	return nil
}

func _ValueMatcher_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ValueMatcher)
	switch tag {
	case 1: // match_pattern.null_match
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ValueMatcher_NullMatch)
		err := b.DecodeMessage(msg)
		m.MatchPattern = &ValueMatcher_NullMatch_{msg}
		return true, err
	case 2: // match_pattern.double_match
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DoubleMatcher)
		err := b.DecodeMessage(msg)
		m.MatchPattern = &ValueMatcher_DoubleMatch{msg}
		return true, err
	case 3: // match_pattern.string_match
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StringMatcher)
		err := b.DecodeMessage(msg)
		m.MatchPattern = &ValueMatcher_StringMatch{msg}
		return true, err
	case 4: // match_pattern.bool_match
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.MatchPattern = &ValueMatcher_BoolMatch{x != 0}
		return true, err
	case 5: // match_pattern.present_match
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.MatchPattern = &ValueMatcher_PresentMatch{x != 0}
		return true, err
	case 6: // match_pattern.list_match
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ListMatcher)
		err := b.DecodeMessage(msg)
		m.MatchPattern = &ValueMatcher_ListMatch{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ValueMatcher_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ValueMatcher)
	// match_pattern
	switch x := m.MatchPattern.(type) {
	case *ValueMatcher_NullMatch_:
		s := proto.Size(x.NullMatch)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ValueMatcher_DoubleMatch:
		s := proto.Size(x.DoubleMatch)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ValueMatcher_StringMatch:
		s := proto.Size(x.StringMatch)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ValueMatcher_BoolMatch:
		n += 1 // tag and wire
		n += 1
	case *ValueMatcher_PresentMatch:
		n += 1 // tag and wire
		n += 1
	case *ValueMatcher_ListMatch:
		s := proto.Size(x.ListMatch)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// NullMatch is an empty message to specify a null value.
type ValueMatcher_NullMatch struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValueMatcher_NullMatch) Reset()         { *m = ValueMatcher_NullMatch{} }
func (m *ValueMatcher_NullMatch) String() string { return proto.CompactTextString(m) }
func (*ValueMatcher_NullMatch) ProtoMessage()    {}
func (*ValueMatcher_NullMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_145b36501d266253, []int{0, 0}
}
func (m *ValueMatcher_NullMatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValueMatcher_NullMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValueMatcher_NullMatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValueMatcher_NullMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueMatcher_NullMatch.Merge(m, src)
}
func (m *ValueMatcher_NullMatch) XXX_Size() int {
	return m.Size()
}
func (m *ValueMatcher_NullMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueMatcher_NullMatch.DiscardUnknown(m)
}

var xxx_messageInfo_ValueMatcher_NullMatch proto.InternalMessageInfo

// Specifies the way to match a list value.
type ListMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*ListMatcher_OneOf
	MatchPattern         isListMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ListMatcher) Reset()         { *m = ListMatcher{} }
func (m *ListMatcher) String() string { return proto.CompactTextString(m) }
func (*ListMatcher) ProtoMessage()    {}
func (*ListMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_145b36501d266253, []int{1}
}
func (m *ListMatcher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListMatcher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMatcher.Merge(m, src)
}
func (m *ListMatcher) XXX_Size() int {
	return m.Size()
}
func (m *ListMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_ListMatcher proto.InternalMessageInfo

type isListMatcher_MatchPattern interface {
	isListMatcher_MatchPattern()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ListMatcher_OneOf struct {
	OneOf *ValueMatcher `protobuf:"bytes,1,opt,name=one_of,json=oneOf,proto3,oneof"`
}

func (*ListMatcher_OneOf) isListMatcher_MatchPattern() {}

func (m *ListMatcher) GetMatchPattern() isListMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *ListMatcher) GetOneOf() *ValueMatcher {
	if x, ok := m.GetMatchPattern().(*ListMatcher_OneOf); ok {
		return x.OneOf
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ListMatcher) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ListMatcher_OneofMarshaler, _ListMatcher_OneofUnmarshaler, _ListMatcher_OneofSizer, []interface{}{
		(*ListMatcher_OneOf)(nil),
	}
}

func _ListMatcher_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ListMatcher)
	// match_pattern
	switch x := m.MatchPattern.(type) {
	case *ListMatcher_OneOf:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OneOf); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ListMatcher.MatchPattern has unexpected type %T", x)
	}
	return nil
}

func _ListMatcher_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ListMatcher)
	switch tag {
	case 1: // match_pattern.one_of
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ValueMatcher)
		err := b.DecodeMessage(msg)
		m.MatchPattern = &ListMatcher_OneOf{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ListMatcher_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ListMatcher)
	// match_pattern
	switch x := m.MatchPattern.(type) {
	case *ListMatcher_OneOf:
		s := proto.Size(x.OneOf)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ValueMatcher)(nil), "envoy.type.matcher.ValueMatcher")
	proto.RegisterType((*ValueMatcher_NullMatch)(nil), "envoy.type.matcher.ValueMatcher.NullMatch")
	proto.RegisterType((*ListMatcher)(nil), "envoy.type.matcher.ListMatcher")
}

func init() { proto.RegisterFile("envoy/type/matcher/value.proto", fileDescriptor_145b36501d266253) }

var fileDescriptor_145b36501d266253 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbd, 0x4e, 0xf3, 0x30,
	0x18, 0x85, 0xeb, 0xaf, 0x5f, 0x0b, 0x7d, 0xd3, 0x2e, 0x1e, 0x00, 0x75, 0x48, 0x43, 0x25, 0xa4,
	0x8a, 0x21, 0x91, 0x60, 0x62, 0x43, 0x15, 0xa0, 0x4a, 0xfc, 0x55, 0x45, 0x62, 0x60, 0x29, 0x09,
	0x75, 0x21, 0x92, 0x6b, 0x47, 0x8e, 0x53, 0xd1, 0x5b, 0xe1, 0x4a, 0x18, 0x19, 0x19, 0xb9, 0x04,
	0xd4, 0x8d, 0xbb, 0x40, 0xfe, 0x09, 0x8d, 0x44, 0x2a, 0xb6, 0xf8, 0xbc, 0xe7, 0x3c, 0x3e, 0x8e,
	0x0d, 0x2e, 0x61, 0x73, 0xbe, 0x08, 0xe4, 0x22, 0x21, 0xc1, 0x2c, 0x94, 0x0f, 0x4f, 0x44, 0x04,
	0xf3, 0x90, 0x66, 0xc4, 0x4f, 0x04, 0x97, 0x1c, 0x63, 0x3d, 0xf7, 0xd5, 0xdc, 0xb7, 0xf3, 0x76,
	0xa7, 0x24, 0xc3, 0xb2, 0x59, 0x44, 0x84, 0x09, 0x95, 0x1a, 0x52, 0x29, 0x62, 0xf6, 0x68, 0x0d,
	0xdb, 0xf3, 0x90, 0xc6, 0x93, 0x50, 0x92, 0x20, 0xff, 0x30, 0x83, 0xee, 0x4b, 0x15, 0x9a, 0xb7,
	0x6a, 0xfb, 0x4b, 0x13, 0xc3, 0xe7, 0x00, 0x2c, 0xa3, 0x74, 0xac, 0x31, 0x3b, 0xc8, 0x43, 0x3d,
	0xe7, 0x60, 0xdf, 0xff, 0x5d, 0xca, 0x2f, 0xa6, 0xfc, 0xab, 0x8c, 0x52, 0xfd, 0x3d, 0xa8, 0x8c,
	0x1a, 0x2c, 0x5f, 0xe0, 0x33, 0x68, 0x4e, 0x78, 0x16, 0x51, 0x62, 0x71, 0xff, 0x34, 0x6e, 0xb7,
	0x0c, 0x77, 0xa2, 0x7d, 0x96, 0x37, 0xa8, 0x8c, 0x9c, 0xc9, 0x4a, 0x50, 0x1c, 0x73, 0x1c, 0xcb,
	0xa9, 0xae, 0xe7, 0xdc, 0x68, 0x5f, 0x81, 0x93, 0xae, 0x04, 0xdc, 0x01, 0x88, 0x38, 0xcf, 0x0f,
	0xf7, 0xdf, 0x43, 0xbd, 0x4d, 0x55, 0x58, 0x69, 0xc6, 0xb0, 0x07, 0xad, 0x44, 0x90, 0x94, 0x30,
	0x69, 0x3d, 0x35, 0xeb, 0x69, 0x5a, 0xd9, 0xd8, 0x8e, 0x01, 0x68, 0x9c, 0xe6, 0x9e, 0xba, 0x6e,
	0xd3, 0x29, 0x6b, 0x73, 0x11, 0xa7, 0x72, 0xd5, 0xa5, 0x41, 0xf3, 0x65, 0xdb, 0x81, 0xc6, 0xcf,
	0x3f, 0xeb, 0x6f, 0x41, 0x4b, 0x07, 0xc6, 0x49, 0x28, 0x25, 0x11, 0x0c, 0xd7, 0x5e, 0xbf, 0xde,
	0xaa, 0xa8, 0x7b, 0x0f, 0x4e, 0x01, 0x80, 0x8f, 0xa0, 0xce, 0x19, 0x19, 0xf3, 0xa9, 0xbd, 0x16,
	0xef, 0xaf, 0x6b, 0x19, 0x54, 0x46, 0x35, 0xce, 0xc8, 0xf5, 0x74, 0xdd, 0x0e, 0xfd, 0xd3, 0xf7,
	0xa5, 0x8b, 0x3e, 0x96, 0x2e, 0xfa, 0x5c, 0xba, 0x08, 0xbc, 0x98, 0x1b, 0x64, 0x22, 0xf8, 0xf3,
	0xa2, 0x84, 0xde, 0x07, 0x8d, 0x1f, 0xaa, 0xa7, 0x33, 0x44, 0x77, 0x1b, 0x56, 0x8e, 0xea, 0xfa,
	0x31, 0x1d, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x0f, 0xb6, 0x44, 0xdd, 0x02, 0x00, 0x00,
}

func (m *ValueMatcher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValueMatcher) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MatchPattern != nil {
		nn1, err := m.MatchPattern.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ValueMatcher_NullMatch_) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.NullMatch != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintValue(dAtA, i, uint64(m.NullMatch.Size()))
		n2, err := m.NullMatch.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *ValueMatcher_DoubleMatch) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.DoubleMatch != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintValue(dAtA, i, uint64(m.DoubleMatch.Size()))
		n3, err := m.DoubleMatch.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *ValueMatcher_StringMatch) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.StringMatch != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintValue(dAtA, i, uint64(m.StringMatch.Size()))
		n4, err := m.StringMatch.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *ValueMatcher_BoolMatch) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x20
	i++
	if m.BoolMatch {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}
func (m *ValueMatcher_PresentMatch) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x28
	i++
	if m.PresentMatch {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}
func (m *ValueMatcher_ListMatch) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.ListMatch != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintValue(dAtA, i, uint64(m.ListMatch.Size()))
		n5, err := m.ListMatch.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}
func (m *ValueMatcher_NullMatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValueMatcher_NullMatch) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ListMatcher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListMatcher) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MatchPattern != nil {
		nn6, err := m.MatchPattern.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn6
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ListMatcher_OneOf) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.OneOf != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintValue(dAtA, i, uint64(m.OneOf.Size()))
		n7, err := m.OneOf.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}
func encodeVarintValue(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ValueMatcher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MatchPattern != nil {
		n += m.MatchPattern.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ValueMatcher_NullMatch_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NullMatch != nil {
		l = m.NullMatch.Size()
		n += 1 + l + sovValue(uint64(l))
	}
	return n
}
func (m *ValueMatcher_DoubleMatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DoubleMatch != nil {
		l = m.DoubleMatch.Size()
		n += 1 + l + sovValue(uint64(l))
	}
	return n
}
func (m *ValueMatcher_StringMatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StringMatch != nil {
		l = m.StringMatch.Size()
		n += 1 + l + sovValue(uint64(l))
	}
	return n
}
func (m *ValueMatcher_BoolMatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	return n
}
func (m *ValueMatcher_PresentMatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	return n
}
func (m *ValueMatcher_ListMatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ListMatch != nil {
		l = m.ListMatch.Size()
		n += 1 + l + sovValue(uint64(l))
	}
	return n
}
func (m *ValueMatcher_NullMatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ListMatcher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MatchPattern != nil {
		n += m.MatchPattern.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ListMatcher_OneOf) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OneOf != nil {
		l = m.OneOf.Size()
		n += 1 + l + sovValue(uint64(l))
	}
	return n
}

func sovValue(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozValue(x uint64) (n int) {
	return sovValue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValueMatcher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValue
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
			return fmt.Errorf("proto: ValueMatcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValueMatcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NullMatch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValue
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
				return ErrInvalidLengthValue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ValueMatcher_NullMatch{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.MatchPattern = &ValueMatcher_NullMatch_{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DoubleMatch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValue
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
				return ErrInvalidLengthValue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DoubleMatcher{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.MatchPattern = &ValueMatcher_DoubleMatch{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringMatch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValue
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
				return ErrInvalidLengthValue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &StringMatcher{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.MatchPattern = &ValueMatcher_StringMatch{v}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoolMatch", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.MatchPattern = &ValueMatcher_BoolMatch{b}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PresentMatch", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.MatchPattern = &ValueMatcher_PresentMatch{b}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMatch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValue
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
				return ErrInvalidLengthValue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ListMatcher{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.MatchPattern = &ValueMatcher_ListMatch{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthValue
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthValue
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
func (m *ValueMatcher_NullMatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValue
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
			return fmt.Errorf("proto: NullMatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NullMatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthValue
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthValue
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
func (m *ListMatcher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValue
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
			return fmt.Errorf("proto: ListMatcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListMatcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OneOf", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValue
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
				return ErrInvalidLengthValue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ValueMatcher{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.MatchPattern = &ListMatcher_OneOf{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthValue
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthValue
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
func skipValue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValue
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
					return 0, ErrIntOverflowValue
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
					return 0, ErrIntOverflowValue
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
				return 0, ErrInvalidLengthValue
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthValue
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowValue
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
				next, err := skipValue(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthValue
				}
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
	ErrInvalidLengthValue = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValue   = fmt.Errorf("proto: integer overflow")
)