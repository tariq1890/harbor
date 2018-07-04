// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hapi/release/hook.proto

/*
Package release is a generated protocol buffer package.

It is generated from these files:
	hapi/release/hook.proto
	hapi/release/info.proto
	hapi/release/release.proto
	hapi/release/status.proto
	hapi/release/test_run.proto
	hapi/release/test_suite.proto

It has these top-level messages:
	Hook
	Info
	Release
	Status
	TestRun
	TestSuite
*/
package release

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Hook_Event int32

const (
	Hook_UNKNOWN              Hook_Event = 0
	Hook_PRE_INSTALL          Hook_Event = 1
	Hook_POST_INSTALL         Hook_Event = 2
	Hook_PRE_DELETE           Hook_Event = 3
	Hook_POST_DELETE          Hook_Event = 4
	Hook_PRE_UPGRADE          Hook_Event = 5
	Hook_POST_UPGRADE         Hook_Event = 6
	Hook_PRE_ROLLBACK         Hook_Event = 7
	Hook_POST_ROLLBACK        Hook_Event = 8
	Hook_RELEASE_TEST_SUCCESS Hook_Event = 9
	Hook_RELEASE_TEST_FAILURE Hook_Event = 10
)

var Hook_Event_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "PRE_INSTALL",
	2:  "POST_INSTALL",
	3:  "PRE_DELETE",
	4:  "POST_DELETE",
	5:  "PRE_UPGRADE",
	6:  "POST_UPGRADE",
	7:  "PRE_ROLLBACK",
	8:  "POST_ROLLBACK",
	9:  "RELEASE_TEST_SUCCESS",
	10: "RELEASE_TEST_FAILURE",
}
var Hook_Event_value = map[string]int32{
	"UNKNOWN":              0,
	"PRE_INSTALL":          1,
	"POST_INSTALL":         2,
	"PRE_DELETE":           3,
	"POST_DELETE":          4,
	"PRE_UPGRADE":          5,
	"POST_UPGRADE":         6,
	"PRE_ROLLBACK":         7,
	"POST_ROLLBACK":        8,
	"RELEASE_TEST_SUCCESS": 9,
	"RELEASE_TEST_FAILURE": 10,
}

func (x Hook_Event) String() string {
	return proto.EnumName(Hook_Event_name, int32(x))
}
func (Hook_Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Hook_DeletePolicy int32

const (
	Hook_SUCCEEDED            Hook_DeletePolicy = 0
	Hook_FAILED               Hook_DeletePolicy = 1
	Hook_BEFORE_HOOK_CREATION Hook_DeletePolicy = 2
)

var Hook_DeletePolicy_name = map[int32]string{
	0: "SUCCEEDED",
	1: "FAILED",
	2: "BEFORE_HOOK_CREATION",
}
var Hook_DeletePolicy_value = map[string]int32{
	"SUCCEEDED":            0,
	"FAILED":               1,
	"BEFORE_HOOK_CREATION": 2,
}

func (x Hook_DeletePolicy) String() string {
	return proto.EnumName(Hook_DeletePolicy_name, int32(x))
}
func (Hook_DeletePolicy) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

// Hook defines a hook object.
type Hook struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Kind is the Kubernetes kind.
	Kind string `protobuf:"bytes,2,opt,name=kind" json:"kind,omitempty"`
	// Path is the chart-relative path to the template.
	Path string `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	// Manifest is the manifest contents.
	Manifest string `protobuf:"bytes,4,opt,name=manifest" json:"manifest,omitempty"`
	// Events are the events that this hook fires on.
	Events []Hook_Event `protobuf:"varint,5,rep,packed,name=events,enum=hapi.release.Hook_Event" json:"events,omitempty"`
	// LastRun indicates the date/time this was last run.
	LastRun *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=last_run,json=lastRun" json:"last_run,omitempty"`
	// Weight indicates the sort order for execution among similar Hook type
	Weight int32 `protobuf:"varint,7,opt,name=weight" json:"weight,omitempty"`
	// DeletePolicies are the policies that indicate when to delete the hook
	DeletePolicies []Hook_DeletePolicy `protobuf:"varint,8,rep,packed,name=delete_policies,json=deletePolicies,enum=hapi.release.Hook_DeletePolicy" json:"delete_policies,omitempty"`
}

func (m *Hook) Reset()                    { *m = Hook{} }
func (m *Hook) String() string            { return proto.CompactTextString(m) }
func (*Hook) ProtoMessage()               {}
func (*Hook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Hook) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Hook) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Hook) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Hook) GetManifest() string {
	if m != nil {
		return m.Manifest
	}
	return ""
}

func (m *Hook) GetEvents() []Hook_Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Hook) GetLastRun() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastRun
	}
	return nil
}

func (m *Hook) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Hook) GetDeletePolicies() []Hook_DeletePolicy {
	if m != nil {
		return m.DeletePolicies
	}
	return nil
}

func init() {
	proto.RegisterType((*Hook)(nil), "hapi.release.Hook")
	proto.RegisterEnum("hapi.release.Hook_Event", Hook_Event_name, Hook_Event_value)
	proto.RegisterEnum("hapi.release.Hook_DeletePolicy", Hook_DeletePolicy_name, Hook_DeletePolicy_value)
}

func init() { proto.RegisterFile("hapi/release/hook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x51, 0x8f, 0x9a, 0x40,
	0x10, 0x80, 0x8f, 0x13, 0x41, 0x47, 0xcf, 0xdb, 0x6e, 0x9a, 0x76, 0xe3, 0xcb, 0x19, 0x9f, 0x7c,
	0xc2, 0xe6, 0x9a, 0xfe, 0x00, 0x84, 0xb9, 0x6a, 0x24, 0x60, 0x16, 0x4c, 0x93, 0xbe, 0x10, 0xae,
	0xee, 0x29, 0x11, 0x81, 0x08, 0xb6, 0xe9, 0x0f, 0xec, 0x3f, 0xe8, 0x0f, 0x6a, 0x76, 0x45, 0x7b,
	0x49, 0xfb, 0x36, 0xf3, 0xcd, 0x37, 0xc3, 0x0c, 0x0b, 0xef, 0x77, 0x49, 0x99, 0x4e, 0x8f, 0x22,
	0x13, 0x49, 0x25, 0xa6, 0xbb, 0xa2, 0xd8, 0x5b, 0xe5, 0xb1, 0xa8, 0x0b, 0xda, 0x97, 0x05, 0xab,
	0x29, 0x0c, 0x1f, 0xb6, 0x45, 0xb1, 0xcd, 0xc4, 0x54, 0xd5, 0x9e, 0x4f, 0x2f, 0xd3, 0x3a, 0x3d,
	0x88, 0xaa, 0x4e, 0x0e, 0xe5, 0x59, 0x1f, 0xff, 0xd2, 0x41, 0x9f, 0x17, 0xc5, 0x9e, 0x52, 0xd0,
	0xf3, 0xe4, 0x20, 0x98, 0x36, 0xd2, 0x26, 0x5d, 0xae, 0x62, 0xc9, 0xf6, 0x69, 0xbe, 0x61, 0xb7,
	0x67, 0x26, 0x63, 0xc9, 0xca, 0xa4, 0xde, 0xb1, 0xd6, 0x99, 0xc9, 0x98, 0x0e, 0xa1, 0x73, 0x48,
	0xf2, 0xf4, 0x45, 0x54, 0x35, 0xd3, 0x15, 0xbf, 0xe6, 0xf4, 0x03, 0x18, 0xe2, 0xbb, 0xc8, 0xeb,
	0x8a, 0xb5, 0x47, 0xad, 0xc9, 0xe0, 0x91, 0x59, 0xaf, 0x17, 0xb4, 0xe4, 0xb7, 0x2d, 0x94, 0x02,
	0x6f, 0x3c, 0xfa, 0x09, 0x3a, 0x59, 0x52, 0xd5, 0xf1, 0xf1, 0x94, 0x33, 0x63, 0xa4, 0x4d, 0x7a,
	0x8f, 0x43, 0xeb, 0x7c, 0x86, 0x75, 0x39, 0xc3, 0x8a, 0x2e, 0x67, 0x70, 0x53, 0xba, 0xfc, 0x94,
	0xd3, 0x77, 0x60, 0xfc, 0x10, 0xe9, 0x76, 0x57, 0x33, 0x73, 0xa4, 0x4d, 0xda, 0xbc, 0xc9, 0xe8,
	0x1c, 0xee, 0x37, 0x22, 0x13, 0xb5, 0x88, 0xcb, 0x22, 0x4b, 0xbf, 0xa5, 0xa2, 0x62, 0x1d, 0xb5,
	0xc9, 0xc3, 0x7f, 0x36, 0x71, 0x95, 0xb9, 0x92, 0xe2, 0x4f, 0x3e, 0xd8, 0xfc, 0xcd, 0x52, 0x51,
	0x8d, 0x7f, 0x6b, 0xd0, 0x56, 0xab, 0xd2, 0x1e, 0x98, 0x6b, 0x7f, 0xe9, 0x07, 0x5f, 0x7c, 0x72,
	0x43, 0xef, 0xa1, 0xb7, 0xe2, 0x18, 0x2f, 0xfc, 0x30, 0xb2, 0x3d, 0x8f, 0x68, 0x94, 0x40, 0x7f,
	0x15, 0x84, 0xd1, 0x95, 0xdc, 0xd2, 0x01, 0x80, 0x54, 0x5c, 0xf4, 0x30, 0x42, 0xd2, 0x52, 0x2d,
	0xd2, 0x68, 0x80, 0x7e, 0x99, 0xb1, 0x5e, 0x7d, 0xe6, 0xb6, 0x8b, 0xa4, 0x7d, 0x9d, 0x71, 0x21,
	0x86, 0x22, 0x1c, 0x63, 0x1e, 0x78, 0xde, 0xcc, 0x76, 0x96, 0xc4, 0xa4, 0x6f, 0xe0, 0x4e, 0x39,
	0x57, 0xd4, 0xa1, 0x0c, 0xde, 0x72, 0xf4, 0xd0, 0x0e, 0x31, 0x8e, 0x30, 0x8c, 0xe2, 0x70, 0xed,
	0x38, 0x18, 0x86, 0xa4, 0xfb, 0x4f, 0xe5, 0xc9, 0x5e, 0x78, 0x6b, 0x8e, 0x04, 0xc6, 0x0e, 0xf4,
	0x5f, 0x9f, 0x4d, 0xef, 0xa0, 0xab, 0xda, 0xd0, 0x45, 0x97, 0xdc, 0x50, 0x00, 0x43, 0xba, 0xe8,
	0x12, 0x4d, 0x0e, 0x99, 0xe1, 0x53, 0xc0, 0x31, 0x9e, 0x07, 0xc1, 0x32, 0x76, 0x38, 0xda, 0xd1,
	0x22, 0xf0, 0xc9, 0xed, 0xac, 0xfb, 0xd5, 0x6c, 0x7e, 0xe4, 0xb3, 0xa1, 0x5e, 0xe9, 0xe3, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x13, 0x64, 0x75, 0x6c, 0xa3, 0x02, 0x00, 0x00,
}