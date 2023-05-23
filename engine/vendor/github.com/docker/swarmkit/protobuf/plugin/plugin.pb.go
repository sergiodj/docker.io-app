// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/docker/swarmkit/protobuf/plugin/plugin.proto

package plugin

import (
	fmt "fmt"
	github_com_docker_swarmkit_api_deepcopy "github.com/docker/swarmkit/api/deepcopy"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
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

type WatchSelectors struct {
	// supported by all object types
	ID           *bool `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	IDPrefix     *bool `protobuf:"varint,2,opt,name=id_prefix,json=idPrefix" json:"id_prefix,omitempty"`
	Name         *bool `protobuf:"varint,3,opt,name=name" json:"name,omitempty"`
	NamePrefix   *bool `protobuf:"varint,4,opt,name=name_prefix,json=namePrefix" json:"name_prefix,omitempty"`
	Custom       *bool `protobuf:"varint,5,opt,name=custom" json:"custom,omitempty"`
	CustomPrefix *bool `protobuf:"varint,6,opt,name=custom_prefix,json=customPrefix" json:"custom_prefix,omitempty"`
	// supported by tasks only
	ServiceID    *bool `protobuf:"varint,7,opt,name=service_id,json=serviceId" json:"service_id,omitempty"`
	NodeID       *bool `protobuf:"varint,8,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Slot         *bool `protobuf:"varint,9,opt,name=slot" json:"slot,omitempty"`
	DesiredState *bool `protobuf:"varint,10,opt,name=desired_state,json=desiredState" json:"desired_state,omitempty"`
	// supported by nodes only
	Role       *bool `protobuf:"varint,11,opt,name=role" json:"role,omitempty"`
	Membership *bool `protobuf:"varint,12,opt,name=membership" json:"membership,omitempty"`
	// supported by: resource
	Kind *bool `protobuf:"varint,13,opt,name=kind" json:"kind,omitempty"`
}

func (m *WatchSelectors) Reset()      { *m = WatchSelectors{} }
func (*WatchSelectors) ProtoMessage() {}
func (*WatchSelectors) Descriptor() ([]byte, []int) {
	return fileDescriptor_3708583e03e1c1e3, []int{0}
}
func (m *WatchSelectors) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WatchSelectors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WatchSelectors.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WatchSelectors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchSelectors.Merge(m, src)
}
func (m *WatchSelectors) XXX_Size() int {
	return m.Size()
}
func (m *WatchSelectors) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchSelectors.DiscardUnknown(m)
}

var xxx_messageInfo_WatchSelectors proto.InternalMessageInfo

type StoreObject struct {
	WatchSelectors *WatchSelectors `protobuf:"bytes,1,req,name=watch_selectors,json=watchSelectors" json:"watch_selectors,omitempty"`
}

func (m *StoreObject) Reset()      { *m = StoreObject{} }
func (*StoreObject) ProtoMessage() {}
func (*StoreObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_3708583e03e1c1e3, []int{1}
}
func (m *StoreObject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoreObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoreObject.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoreObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreObject.Merge(m, src)
}
func (m *StoreObject) XXX_Size() int {
	return m.Size()
}
func (m *StoreObject) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreObject.DiscardUnknown(m)
}

var xxx_messageInfo_StoreObject proto.InternalMessageInfo

type TLSAuthorization struct {
	// Roles contains the acceptable TLS OU roles for the handler.
	Roles []string `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
	// Insecure is set to true if this method does not require
	// authorization. NOTE: Specifying both "insecure" and a nonempty
	// list of roles is invalid. This would fail at codegen time.
	Insecure *bool `protobuf:"varint,2,opt,name=insecure" json:"insecure,omitempty"`
}

func (m *TLSAuthorization) Reset()      { *m = TLSAuthorization{} }
func (*TLSAuthorization) ProtoMessage() {}
func (*TLSAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_3708583e03e1c1e3, []int{2}
}
func (m *TLSAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLSAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLSAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TLSAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLSAuthorization.Merge(m, src)
}
func (m *TLSAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *TLSAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_TLSAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_TLSAuthorization proto.InternalMessageInfo

var E_Deepcopy = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         70000,
	Name:          "docker.protobuf.plugin.deepcopy",
	Tag:           "varint,70000,opt,name=deepcopy,def=1",
	Filename:      "github.com/docker/swarmkit/protobuf/plugin/plugin.proto",
}

var E_StoreObject = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*StoreObject)(nil),
	Field:         70001,
	Name:          "docker.protobuf.plugin.store_object",
	Tag:           "bytes,70001,opt,name=store_object",
	Filename:      "github.com/docker/swarmkit/protobuf/plugin/plugin.proto",
}

var E_TlsAuthorization = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*TLSAuthorization)(nil),
	Field:         73626345,
	Name:          "docker.protobuf.plugin.tls_authorization",
	Tag:           "bytes,73626345,opt,name=tls_authorization",
	Filename:      "github.com/docker/swarmkit/protobuf/plugin/plugin.proto",
}

func init() {
	proto.RegisterType((*WatchSelectors)(nil), "docker.protobuf.plugin.WatchSelectors")
	proto.RegisterType((*StoreObject)(nil), "docker.protobuf.plugin.StoreObject")
	proto.RegisterType((*TLSAuthorization)(nil), "docker.protobuf.plugin.TLSAuthorization")
	proto.RegisterExtension(E_Deepcopy)
	proto.RegisterExtension(E_StoreObject)
	proto.RegisterExtension(E_TlsAuthorization)
}

func init() {
	proto.RegisterFile("github.com/docker/swarmkit/protobuf/plugin/plugin.proto", fileDescriptor_3708583e03e1c1e3)
}

var fileDescriptor_3708583e03e1c1e3 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0x8d, 0xd3, 0x36, 0x4d, 0x26, 0x69, 0xbf, 0x7e, 0x2b, 0x54, 0x59, 0x3d, 0x38, 0x55, 0x83,
	0x50, 0x90, 0x90, 0x23, 0xf5, 0x82, 0x94, 0x1b, 0x25, 0x97, 0x48, 0x40, 0x91, 0x83, 0xc4, 0x8d,
	0xc8, 0xf1, 0x4e, 0x93, 0xa5, 0x8e, 0xd7, 0xda, 0x5d, 0xd3, 0xc2, 0x89, 0x3f, 0x80, 0xc4, 0x95,
	0x2b, 0xbf, 0xa6, 0xc7, 0x1e, 0x7b, 0x8a, 0xa8, 0x23, 0x0e, 0xdc, 0xe0, 0x1f, 0xa0, 0xdd, 0x75,
	0x1a, 0x82, 0x5a, 0x71, 0xf2, 0xcc, 0x9b, 0xf7, 0x66, 0xe6, 0xed, 0x18, 0x1e, 0x8f, 0x99, 0x9a,
	0x64, 0x23, 0x3f, 0xe2, 0xd3, 0x0e, 0xe5, 0xd1, 0x29, 0x8a, 0x8e, 0x3c, 0x0b, 0xc5, 0xf4, 0x94,
	0xa9, 0x4e, 0x2a, 0xb8, 0xe2, 0xa3, 0xec, 0xa4, 0x93, 0xc6, 0xd9, 0x98, 0x25, 0xc5, 0xc7, 0x37,
	0x30, 0xd9, 0xb5, 0x6c, 0x7f, 0x41, 0xf2, 0x6d, 0x75, 0x6f, 0x7f, 0xcc, 0xf9, 0x38, 0xc6, 0xa5,
	0x98, 0xa2, 0x8c, 0x04, 0x4b, 0x15, 0x2f, 0xb8, 0x07, 0x5f, 0xd6, 0x60, 0xfb, 0x75, 0xa8, 0xa2,
	0xc9, 0x00, 0x63, 0x8c, 0x14, 0x17, 0x92, 0xec, 0x42, 0x99, 0x51, 0xd7, 0xd9, 0x77, 0xda, 0xd5,
	0xa3, 0x4a, 0x3e, 0x6b, 0x96, 0xfb, 0xbd, 0xa0, 0xcc, 0x28, 0x79, 0x08, 0x35, 0x46, 0x87, 0xa9,
	0xc0, 0x13, 0x76, 0xee, 0x96, 0x4d, 0xb9, 0x91, 0xcf, 0x9a, 0xd5, 0x7e, 0xef, 0xa5, 0xc1, 0x82,
	0x2a, 0xa3, 0x36, 0x22, 0x04, 0xd6, 0x93, 0x70, 0x8a, 0xee, 0x9a, 0x66, 0x05, 0x26, 0x26, 0x4d,
	0xa8, 0xeb, 0xef, 0xa2, 0xc1, 0xba, 0x29, 0x81, 0x86, 0x0a, 0xd1, 0x2e, 0x54, 0xa2, 0x4c, 0x2a,
	0x3e, 0x75, 0x37, 0x4c, 0xad, 0xc8, 0x48, 0x0b, 0xb6, 0x6c, 0xb4, 0x90, 0x56, 0x4c, 0xb9, 0x61,
	0xc1, 0x42, 0xfc, 0x08, 0x40, 0xa2, 0x78, 0xc7, 0x22, 0x1c, 0x32, 0xea, 0x6e, 0x9a, 0xed, 0xb6,
	0xf2, 0x59, 0xb3, 0x36, 0xb0, 0x68, 0xbf, 0x17, 0xd4, 0x0a, 0x42, 0x9f, 0x92, 0x16, 0x6c, 0x26,
	0x9c, 0x1a, 0x6a, 0xd5, 0x50, 0x21, 0x9f, 0x35, 0x2b, 0x2f, 0x38, 0xd5, 0xbc, 0x8a, 0x2e, 0xf5,
	0xa9, 0x36, 0x21, 0x63, 0xae, 0xdc, 0x9a, 0x35, 0xa1, 0x63, 0xbd, 0x0b, 0x45, 0xc9, 0x04, 0xd2,
	0xa1, 0x54, 0xa1, 0x42, 0x17, 0xec, 0x2e, 0x05, 0x38, 0xd0, 0x98, 0x16, 0x0a, 0x1e, 0xa3, 0x5b,
	0xb7, 0x42, 0x1d, 0x13, 0x0f, 0x60, 0x8a, 0xd3, 0x11, 0x0a, 0x39, 0x61, 0xa9, 0xdb, 0xb0, 0xe6,
	0x97, 0x88, 0xd6, 0x9c, 0xb2, 0x84, 0xba, 0x5b, 0x56, 0xa3, 0xe3, 0x83, 0x37, 0x50, 0x1f, 0x28,
	0x2e, 0xf0, 0x78, 0xf4, 0x16, 0x23, 0x45, 0x8e, 0xe1, 0xbf, 0x33, 0x7d, 0xa9, 0xa1, 0x5c, 0x9c,
	0xca, 0x75, 0xf6, 0xcb, 0xed, 0xfa, 0xe1, 0x03, 0xff, 0xf6, 0xf3, 0xfb, 0xab, 0x87, 0x0d, 0xb6,
	0xcf, 0x56, 0xf2, 0x83, 0x1e, 0xec, 0xbc, 0x7a, 0x36, 0x78, 0x92, 0xa9, 0x09, 0x17, 0xec, 0x43,
	0xa8, 0x18, 0x4f, 0xc8, 0x3d, 0xd8, 0xd0, 0xfb, 0xea, 0xd6, 0x6b, 0xed, 0x5a, 0x60, 0x13, 0xb2,
	0x07, 0x55, 0x96, 0x48, 0x8c, 0x32, 0x81, 0xf6, 0xf2, 0xc1, 0x4d, 0xde, 0x7d, 0x0a, 0x55, 0x8a,
	0x98, 0x46, 0x3c, 0x7d, 0x4f, 0x9a, 0xbe, 0xfd, 0xe1, 0x96, 0x9b, 0x3c, 0x47, 0x29, 0xc3, 0x31,
	0x1e, 0xa7, 0xba, 0xbb, 0x74, 0x7f, 0x7e, 0x35, 0x77, 0xef, 0xae, 0x2b, 0x91, 0x61, 0x70, 0x23,
	0xec, 0x32, 0x68, 0x48, 0x6d, 0x75, 0xc8, 0xad, 0xd7, 0x7f, 0x36, 0xfa, 0x65, 0x1a, 0xd5, 0x0f,
	0x5b, 0x77, 0x79, 0xff, 0xe3, 0xe5, 0x82, 0xba, 0x5c, 0x26, 0xdd, 0x73, 0xf8, 0x5f, 0xc5, 0x72,
	0x18, 0xae, 0xd8, 0xf6, 0x6e, 0x99, 0xa7, 0x26, 0x9c, 0x2e, 0xc6, 0xfd, 0xf8, 0xfe, 0xa9, 0x65,
	0xe6, 0xb5, 0xef, 0x9a, 0xf7, 0xf7, 0x4b, 0x06, 0x3b, 0x2a, 0x96, 0x2b, 0xc8, 0xd1, 0xfd, 0x8b,
	0x6b, 0xaf, 0x74, 0x75, 0xed, 0x95, 0x3e, 0xe6, 0x9e, 0x73, 0x91, 0x7b, 0xce, 0x65, 0xee, 0x39,
	0xdf, 0x72, 0xcf, 0xf9, 0x3c, 0xf7, 0x4a, 0x97, 0x73, 0xaf, 0x74, 0x35, 0xf7, 0x4a, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x0f, 0x50, 0xb9, 0xa3, 0x05, 0x04, 0x00, 0x00,
}

func (m *WatchSelectors) Copy() *WatchSelectors {
	if m == nil {
		return nil
	}
	o := &WatchSelectors{}
	o.CopyFrom(m)
	return o
}

func (m *WatchSelectors) CopyFrom(src interface{}) {

	o := src.(*WatchSelectors)
	*m = *o
}

func (m *StoreObject) Copy() *StoreObject {
	if m == nil {
		return nil
	}
	o := &StoreObject{}
	o.CopyFrom(m)
	return o
}

func (m *StoreObject) CopyFrom(src interface{}) {

	o := src.(*StoreObject)
	*m = *o
	if o.WatchSelectors != nil {
		m.WatchSelectors = &WatchSelectors{}
		github_com_docker_swarmkit_api_deepcopy.Copy(m.WatchSelectors, o.WatchSelectors)
	}
}

func (m *TLSAuthorization) Copy() *TLSAuthorization {
	if m == nil {
		return nil
	}
	o := &TLSAuthorization{}
	o.CopyFrom(m)
	return o
}

func (m *TLSAuthorization) CopyFrom(src interface{}) {

	o := src.(*TLSAuthorization)
	*m = *o
	if o.Roles != nil {
		m.Roles = make([]string, len(o.Roles))
		copy(m.Roles, o.Roles)
	}

}

func (m *WatchSelectors) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WatchSelectors) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != nil {
		dAtA[i] = 0x8
		i++
		if *m.ID {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.IDPrefix != nil {
		dAtA[i] = 0x10
		i++
		if *m.IDPrefix {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Name != nil {
		dAtA[i] = 0x18
		i++
		if *m.Name {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.NamePrefix != nil {
		dAtA[i] = 0x20
		i++
		if *m.NamePrefix {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Custom != nil {
		dAtA[i] = 0x28
		i++
		if *m.Custom {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.CustomPrefix != nil {
		dAtA[i] = 0x30
		i++
		if *m.CustomPrefix {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.ServiceID != nil {
		dAtA[i] = 0x38
		i++
		if *m.ServiceID {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.NodeID != nil {
		dAtA[i] = 0x40
		i++
		if *m.NodeID {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Slot != nil {
		dAtA[i] = 0x48
		i++
		if *m.Slot {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.DesiredState != nil {
		dAtA[i] = 0x50
		i++
		if *m.DesiredState {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Role != nil {
		dAtA[i] = 0x58
		i++
		if *m.Role {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Membership != nil {
		dAtA[i] = 0x60
		i++
		if *m.Membership {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Kind != nil {
		dAtA[i] = 0x68
		i++
		if *m.Kind {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *StoreObject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreObject) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.WatchSelectors == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("watch_selectors")
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPlugin(dAtA, i, uint64(m.WatchSelectors.Size()))
		n1, err := m.WatchSelectors.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *TLSAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLSAuthorization) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.Insecure != nil {
		dAtA[i] = 0x10
		i++
		if *m.Insecure {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintPlugin(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *WatchSelectors) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != nil {
		n += 2
	}
	if m.IDPrefix != nil {
		n += 2
	}
	if m.Name != nil {
		n += 2
	}
	if m.NamePrefix != nil {
		n += 2
	}
	if m.Custom != nil {
		n += 2
	}
	if m.CustomPrefix != nil {
		n += 2
	}
	if m.ServiceID != nil {
		n += 2
	}
	if m.NodeID != nil {
		n += 2
	}
	if m.Slot != nil {
		n += 2
	}
	if m.DesiredState != nil {
		n += 2
	}
	if m.Role != nil {
		n += 2
	}
	if m.Membership != nil {
		n += 2
	}
	if m.Kind != nil {
		n += 2
	}
	return n
}

func (m *StoreObject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.WatchSelectors != nil {
		l = m.WatchSelectors.Size()
		n += 1 + l + sovPlugin(uint64(l))
	}
	return n
}

func (m *TLSAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			l = len(s)
			n += 1 + l + sovPlugin(uint64(l))
		}
	}
	if m.Insecure != nil {
		n += 2
	}
	return n
}

func sovPlugin(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPlugin(x uint64) (n int) {
	return sovPlugin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *WatchSelectors) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WatchSelectors{`,
		`ID:` + valueToStringPlugin(this.ID) + `,`,
		`IDPrefix:` + valueToStringPlugin(this.IDPrefix) + `,`,
		`Name:` + valueToStringPlugin(this.Name) + `,`,
		`NamePrefix:` + valueToStringPlugin(this.NamePrefix) + `,`,
		`Custom:` + valueToStringPlugin(this.Custom) + `,`,
		`CustomPrefix:` + valueToStringPlugin(this.CustomPrefix) + `,`,
		`ServiceID:` + valueToStringPlugin(this.ServiceID) + `,`,
		`NodeID:` + valueToStringPlugin(this.NodeID) + `,`,
		`Slot:` + valueToStringPlugin(this.Slot) + `,`,
		`DesiredState:` + valueToStringPlugin(this.DesiredState) + `,`,
		`Role:` + valueToStringPlugin(this.Role) + `,`,
		`Membership:` + valueToStringPlugin(this.Membership) + `,`,
		`Kind:` + valueToStringPlugin(this.Kind) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StoreObject) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StoreObject{`,
		`WatchSelectors:` + strings.Replace(fmt.Sprintf("%v", this.WatchSelectors), "WatchSelectors", "WatchSelectors", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TLSAuthorization) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TLSAuthorization{`,
		`Roles:` + fmt.Sprintf("%v", this.Roles) + `,`,
		`Insecure:` + valueToStringPlugin(this.Insecure) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPlugin(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *WatchSelectors) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlugin
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
			return fmt.Errorf("proto: WatchSelectors: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchSelectors: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
			m.ID = &b
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IDPrefix", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
			m.IDPrefix = &b
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
			m.Name = &b
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NamePrefix", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
			m.NamePrefix = &b
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Custom", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
			m.Custom = &b
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomPrefix", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
			m.CustomPrefix = &b
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceID", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
			m.ServiceID = &b
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
			m.NodeID = &b
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slot", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
			m.Slot = &b
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DesiredState", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
			m.DesiredState = &b
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
			m.Role = &b
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Membership", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
			m.Membership = &b
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
			m.Kind = &b
		default:
			iNdEx = preIndex
			skippy, err := skipPlugin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlugin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPlugin
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
func (m *StoreObject) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlugin
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
			return fmt.Errorf("proto: StoreObject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreObject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WatchSelectors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlugin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.WatchSelectors == nil {
				m.WatchSelectors = &WatchSelectors{}
			}
			if err := m.WatchSelectors.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipPlugin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlugin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPlugin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("watch_selectors")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLSAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlugin
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
			return fmt.Errorf("proto: TLSAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TLSAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlugin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Insecure", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
			m.Insecure = &b
		default:
			iNdEx = preIndex
			skippy, err := skipPlugin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlugin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPlugin
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
func skipPlugin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlugin
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
					return 0, ErrIntOverflowPlugin
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
					return 0, ErrIntOverflowPlugin
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
				return 0, ErrInvalidLengthPlugin
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthPlugin
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPlugin
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
				next, err := skipPlugin(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthPlugin
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
	ErrInvalidLengthPlugin = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlugin   = fmt.Errorf("proto: integer overflow")
)
