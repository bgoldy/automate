// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/cfgmgmt/response/eventstrings.proto

package response // import "github.com/chef/automate/api/interservice/cfgmgmt/response"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EventStrings struct {
	Strings              []*EventString `protobuf:"bytes,1,rep,name=strings,proto3" json:"strings,omitempty" toml:"strings,omitempty" mapstructure:"strings,omitempty"`
	Start                string         `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty" toml:"start,omitempty" mapstructure:"start,omitempty"`
	End                  string         `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty" toml:"end,omitempty" mapstructure:"end,omitempty"`
	HoursBetween         int32          `protobuf:"varint,4,opt,name=hours_between,json=hoursBetween,proto3" json:"hours_between,omitempty" toml:"hours_between,omitempty" mapstructure:"hours_between,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte         `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32          `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *EventStrings) Reset()         { *m = EventStrings{} }
func (m *EventStrings) String() string { return proto.CompactTextString(m) }
func (*EventStrings) ProtoMessage()    {}
func (*EventStrings) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventstrings_4865921dbc7605c7, []int{0}
}
func (m *EventStrings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventStrings.Unmarshal(m, b)
}
func (m *EventStrings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventStrings.Marshal(b, m, deterministic)
}
func (dst *EventStrings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventStrings.Merge(dst, src)
}
func (m *EventStrings) XXX_Size() int {
	return xxx_messageInfo_EventStrings.Size(m)
}
func (m *EventStrings) XXX_DiscardUnknown() {
	xxx_messageInfo_EventStrings.DiscardUnknown(m)
}

var xxx_messageInfo_EventStrings proto.InternalMessageInfo

func (m *EventStrings) GetStrings() []*EventString {
	if m != nil {
		return m.Strings
	}
	return nil
}

func (m *EventStrings) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *EventStrings) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func (m *EventStrings) GetHoursBetween() int32 {
	if m != nil {
		return m.HoursBetween
	}
	return 0
}

type EventString struct {
	Collection           []*EventCollection `protobuf:"bytes,1,rep,name=collection,proto3" json:"collection,omitempty" toml:"collection,omitempty" mapstructure:"collection,omitempty"`
	EventAction          string             `protobuf:"bytes,2,opt,name=event_action,json=eventAction,proto3" json:"event_action,omitempty" toml:"event_action,omitempty" mapstructure:"event_action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32              `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *EventString) Reset()         { *m = EventString{} }
func (m *EventString) String() string { return proto.CompactTextString(m) }
func (*EventString) ProtoMessage()    {}
func (*EventString) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventstrings_4865921dbc7605c7, []int{1}
}
func (m *EventString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventString.Unmarshal(m, b)
}
func (m *EventString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventString.Marshal(b, m, deterministic)
}
func (dst *EventString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventString.Merge(dst, src)
}
func (m *EventString) XXX_Size() int {
	return xxx_messageInfo_EventString.Size(m)
}
func (m *EventString) XXX_DiscardUnknown() {
	xxx_messageInfo_EventString.DiscardUnknown(m)
}

var xxx_messageInfo_EventString proto.InternalMessageInfo

func (m *EventString) GetCollection() []*EventCollection {
	if m != nil {
		return m.Collection
	}
	return nil
}

func (m *EventString) GetEventAction() string {
	if m != nil {
		return m.EventAction
	}
	return ""
}

type EventCollection struct {
	EventsCount          []*EventCount `protobuf:"bytes,1,rep,name=events_count,json=eventsCount,proto3" json:"events_count,omitempty" toml:"events_count,omitempty" mapstructure:"events_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte        `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32         `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *EventCollection) Reset()         { *m = EventCollection{} }
func (m *EventCollection) String() string { return proto.CompactTextString(m) }
func (*EventCollection) ProtoMessage()    {}
func (*EventCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventstrings_4865921dbc7605c7, []int{2}
}
func (m *EventCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventCollection.Unmarshal(m, b)
}
func (m *EventCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventCollection.Marshal(b, m, deterministic)
}
func (dst *EventCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCollection.Merge(dst, src)
}
func (m *EventCollection) XXX_Size() int {
	return xxx_messageInfo_EventCollection.Size(m)
}
func (m *EventCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCollection.DiscardUnknown(m)
}

var xxx_messageInfo_EventCollection proto.InternalMessageInfo

func (m *EventCollection) GetEventsCount() []*EventCount {
	if m != nil {
		return m.EventsCount
	}
	return nil
}

func init() {
	proto.RegisterType((*EventStrings)(nil), "chef.automate.domain.cfgmgmt.response.EventStrings")
	proto.RegisterType((*EventString)(nil), "chef.automate.domain.cfgmgmt.response.EventString")
	proto.RegisterType((*EventCollection)(nil), "chef.automate.domain.cfgmgmt.response.EventCollection")
}

func init() {
	proto.RegisterFile("api/interservice/cfgmgmt/response/eventstrings.proto", fileDescriptor_eventstrings_4865921dbc7605c7)
}

var fileDescriptor_eventstrings_4865921dbc7605c7 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4a, 0x33, 0x31,
	0x14, 0x85, 0x99, 0xbf, 0x7f, 0x15, 0xd3, 0x8a, 0x12, 0x5c, 0x0c, 0xae, 0x6a, 0x45, 0xe8, 0xc6,
	0x04, 0xab, 0xb8, 0x10, 0x37, 0x56, 0xdc, 0xb9, 0x1a, 0xc5, 0x85, 0x9b, 0x92, 0xa6, 0xb7, 0x33,
	0x81, 0x4e, 0x32, 0x24, 0x77, 0xea, 0x2b, 0xf8, 0x2a, 0xbe, 0xa5, 0x4c, 0x92, 0x91, 0xc1, 0x8d,
	0xed, 0x2e, 0xe7, 0xc0, 0xf9, 0xce, 0x21, 0x97, 0xdc, 0x88, 0x4a, 0x71, 0xa5, 0x11, 0xac, 0x03,
	0xbb, 0x51, 0x12, 0xb8, 0x5c, 0xe5, 0x65, 0x5e, 0x22, 0xb7, 0xe0, 0x2a, 0xa3, 0x1d, 0x70, 0xd8,
	0x80, 0x46, 0x87, 0x56, 0xe9, 0xdc, 0xb1, 0xca, 0x1a, 0x34, 0xf4, 0x42, 0x16, 0xb0, 0x62, 0xa2,
	0x46, 0x53, 0x0a, 0x04, 0xb6, 0x34, 0xa5, 0x50, 0x9a, 0xc5, 0x24, 0x6b, 0x93, 0xa7, 0x97, 0x5b,
	0xc2, 0x03, 0x75, 0xfc, 0x95, 0x90, 0xe1, 0x53, 0xa3, 0x5f, 0x42, 0x19, 0x7d, 0x26, 0xfb, 0xb1,
	0x37, 0x4d, 0x46, 0xbd, 0xc9, 0x60, 0x3a, 0x65, 0x5b, 0x15, 0xb3, 0x0e, 0x25, 0x6b, 0x11, 0xf4,
	0x84, 0xf4, 0x1d, 0x0a, 0x8b, 0xe9, 0xbf, 0x51, 0x32, 0x39, 0xc8, 0x82, 0xa0, 0xc7, 0xa4, 0x07,
	0x7a, 0x99, 0xf6, 0xbc, 0xd7, 0x3c, 0xe9, 0x39, 0x39, 0x2c, 0x4c, 0x6d, 0xdd, 0x7c, 0x01, 0xf8,
	0x01, 0xa0, 0xd3, 0xff, 0xa3, 0x64, 0xd2, 0xcf, 0x86, 0xde, 0x9c, 0x05, 0x6f, 0xfc, 0x99, 0x90,
	0x41, 0xa7, 0x85, 0xbe, 0x11, 0x22, 0xcd, 0x7a, 0x0d, 0x12, 0x95, 0xd1, 0x71, 0xed, 0xed, 0x2e,
	0x6b, 0x1f, 0x7f, 0xd2, 0x59, 0x87, 0x44, 0xcf, 0xc8, 0xd0, 0x7f, 0xd1, 0x5c, 0x04, 0x72, 0xd8,
	0x3e, 0xf0, 0xde, 0x83, 0xb7, 0xc6, 0x39, 0x39, 0xfa, 0x45, 0xa0, 0xaf, 0x31, 0xe5, 0xe6, 0xd2,
	0xd4, 0x1a, 0xe3, 0x9e, 0xab, 0xdd, 0xf6, 0xd4, 0x1a, 0x63, 0x91, 0xf3, 0x62, 0x76, 0xff, 0x7e,
	0x97, 0x2b, 0x2c, 0xea, 0x05, 0x93, 0xa6, 0xe4, 0x0d, 0x8b, 0xb7, 0x2c, 0xfe, 0xe7, 0xa5, 0x17,
	0x7b, 0xfe, 0xc8, 0xd7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x50, 0xa7, 0x61, 0x3f, 0x72, 0x02,
	0x00, 0x00,
}
