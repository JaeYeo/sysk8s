/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/kubernetes/vendor/k8s.io/api/node/v1/generated.proto

package v1

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	k8s_io_api_core_v1 "k8s.io/api/core/v1"
	v11 "k8s.io/api/core/v1"
	k8s_io_apimachinery_pkg_api_resource "k8s.io/apimachinery/pkg/api/resource"
	resource "k8s.io/apimachinery/pkg/api/resource"

	math "math"
	math_bits "math/bits"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func (m *Overhead) Reset()      { *m = Overhead{} }
func (*Overhead) ProtoMessage() {}
func (*Overhead) Descriptor() ([]byte, []int) {
	return fileDescriptor_f977b0dddc93b4ec, []int{0}
}
func (m *Overhead) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Overhead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Overhead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Overhead.Merge(m, src)
}
func (m *Overhead) XXX_Size() int {
	return m.Size()
}
func (m *Overhead) XXX_DiscardUnknown() {
	xxx_messageInfo_Overhead.DiscardUnknown(m)
}

var xxx_messageInfo_Overhead proto.InternalMessageInfo

func (m *RuntimeClass) Reset()      { *m = RuntimeClass{} }
func (*RuntimeClass) ProtoMessage() {}
func (*RuntimeClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_f977b0dddc93b4ec, []int{1}
}
func (m *RuntimeClass) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RuntimeClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *RuntimeClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeClass.Merge(m, src)
}
func (m *RuntimeClass) XXX_Size() int {
	return m.Size()
}
func (m *RuntimeClass) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeClass.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeClass proto.InternalMessageInfo

func (m *RuntimeClassList) Reset()      { *m = RuntimeClassList{} }
func (*RuntimeClassList) ProtoMessage() {}
func (*RuntimeClassList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f977b0dddc93b4ec, []int{2}
}
func (m *RuntimeClassList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RuntimeClassList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *RuntimeClassList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeClassList.Merge(m, src)
}
func (m *RuntimeClassList) XXX_Size() int {
	return m.Size()
}
func (m *RuntimeClassList) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeClassList.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeClassList proto.InternalMessageInfo

func (m *Scheduling) Reset()      { *m = Scheduling{} }
func (*Scheduling) ProtoMessage() {}
func (*Scheduling) Descriptor() ([]byte, []int) {
	return fileDescriptor_f977b0dddc93b4ec, []int{3}
}
func (m *Scheduling) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Scheduling) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Scheduling) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scheduling.Merge(m, src)
}
func (m *Scheduling) XXX_Size() int {
	return m.Size()
}
func (m *Scheduling) XXX_DiscardUnknown() {
	xxx_messageInfo_Scheduling.DiscardUnknown(m)
}

var xxx_messageInfo_Scheduling proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Overhead)(nil), "k8s.io.api.node.v1.Overhead")
	proto.RegisterMapType((k8s_io_api_core_v1.ResourceList)(nil), "k8s.io.api.node.v1.Overhead.PodFixedEntry")
	proto.RegisterType((*RuntimeClass)(nil), "k8s.io.api.node.v1.RuntimeClass")
	proto.RegisterType((*RuntimeClassList)(nil), "k8s.io.api.node.v1.RuntimeClassList")
	proto.RegisterType((*Scheduling)(nil), "k8s.io.api.node.v1.Scheduling")
	proto.RegisterMapType((map[string]string)(nil), "k8s.io.api.node.v1.Scheduling.NodeSelectorEntry")
}

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/node/v1/generated.proto", fileDescriptor_f977b0dddc93b4ec)
}

var fileDescriptor_f977b0dddc93b4ec = []byte{
	// 666 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xbd, 0x6f, 0xd3, 0x40,
	0x14, 0xcf, 0xa5, 0x54, 0x4d, 0x2f, 0x29, 0x14, 0x53, 0x89, 0x28, 0x83, 0x53, 0x82, 0x90, 0xca,
	0xd0, 0x33, 0xad, 0x00, 0x55, 0x2c, 0x20, 0xf3, 0x21, 0x3e, 0x5b, 0x70, 0x61, 0x41, 0x0c, 0x5c,
	0xec, 0x87, 0x63, 0x12, 0xfb, 0xa2, 0xf3, 0x39, 0x22, 0x1b, 0x62, 0x41, 0x62, 0x62, 0xe1, 0xbf,
	0x81, 0xbd, 0x1b, 0x5d, 0x90, 0x3a, 0xb5, 0x34, 0xfc, 0x17, 0x4c, 0xe8, 0xec, 0x73, 0x72, 0x6d,
	0x9a, 0xb6, 0x6c, 0xbe, 0xf3, 0xef, 0xe3, 0xbd, 0xdf, 0xbb, 0x87, 0xef, 0xb4, 0xd7, 0x62, 0x12,
	0x30, 0xab, 0x9d, 0x34, 0x81, 0x47, 0x20, 0x20, 0xb6, 0x7a, 0x10, 0x79, 0x8c, 0x5b, 0xea, 0x07,
	0xed, 0x06, 0x56, 0xc4, 0x3c, 0xb0, 0x7a, 0x2b, 0x4d, 0x10, 0x74, 0xc5, 0xf2, 0x21, 0x02, 0x4e,
	0x05, 0x78, 0xa4, 0xcb, 0x99, 0x60, 0xc6, 0xc5, 0x0c, 0x48, 0x68, 0x37, 0x20, 0x12, 0x48, 0x14,
	0xb0, 0xb6, 0xec, 0x07, 0xa2, 0x95, 0x34, 0x89, 0xcb, 0x42, 0xcb, 0x67, 0x3e, 0xb3, 0x52, 0x7c,
	0x33, 0x79, 0x97, 0x9e, 0xd2, 0x43, 0xfa, 0x95, 0xe9, 0xd4, 0x1a, 0x9a, 0xa1, 0xcb, 0xb8, 0x34,
	0x3c, 0xec, 0x55, 0xbb, 0x3e, 0xc2, 0x84, 0xd4, 0x6d, 0x05, 0x11, 0xf0, 0xbe, 0xd5, 0x6d, 0xfb,
	0x29, 0x89, 0x43, 0xcc, 0x12, 0xee, 0xc2, 0x7f, 0xb1, 0x62, 0x2b, 0x04, 0x41, 0x8f, 0xf2, 0xb2,
	0x26, 0xb1, 0x78, 0x12, 0x89, 0x20, 0x1c, 0xb7, 0xb9, 0x79, 0x12, 0x21, 0x76, 0x5b, 0x10, 0xd2,
	0xc3, 0xbc, 0xc6, 0xcf, 0x22, 0x2e, 0x6d, 0xf4, 0x80, 0xb7, 0x80, 0x7a, 0xc6, 0x2f, 0x84, 0x4b,
	0x5d, 0xe6, 0x3d, 0x08, 0x3e, 0x80, 0x57, 0x45, 0x8b, 0x53, 0x4b, 0xe5, 0x55, 0x8b, 0x4c, 0x48,
	0x98, 0xe4, 0x2c, 0xf2, 0x5c, 0x31, 0xee, 0x47, 0x82, 0xf7, 0xed, 0xcf, 0x68, 0x6b, 0xb7, 0x5e,
	0x18, 0xec, 0xd6, 0x4b, 0xf9, 0xfd, 0xdf, 0xdd, 0x7a, 0x7d, 0x3c, 0x5e, 0xe2, 0xa8, 0xc4, 0x9e,
	0x06, 0xb1, 0xf8, 0xb4, 0x77, 0x2c, 0x64, 0x9d, 0x86, 0xf0, 0x65, 0xaf, 0xbe, 0x7c, 0x9a, 0x01,
	0x90, 0x17, 0x09, 0x8d, 0x44, 0x20, 0xfa, 0xce, 0xb0, 0x95, 0x5a, 0x1b, 0xcf, 0x1d, 0x28, 0xd2,
	0x98, 0xc7, 0x53, 0x6d, 0xe8, 0x57, 0xd1, 0x22, 0x5a, 0x9a, 0x75, 0xe4, 0xa7, 0x71, 0x0f, 0x4f,
	0xf7, 0x68, 0x27, 0x81, 0x6a, 0x71, 0x11, 0x2d, 0x95, 0x57, 0x89, 0xd6, 0xf6, 0xd0, 0x8b, 0x74,
	0xdb, 0x7e, 0x9a, 0xc3, 0xb8, 0x57, 0x46, 0xbe, 0x55, 0x5c, 0x43, 0x8d, 0x1f, 0x45, 0x5c, 0x71,
	0xb2, 0xd0, 0xef, 0x76, 0x68, 0x1c, 0x1b, 0x6f, 0x71, 0x49, 0x8e, 0xd9, 0xa3, 0x82, 0xa6, 0x8e,
	0xe5, 0xd5, 0x6b, 0xc7, 0xa9, 0xc7, 0x44, 0xa2, 0x49, 0x6f, 0x85, 0x6c, 0x34, 0xdf, 0x83, 0x2b,
	0x9e, 0x81, 0xa0, 0xb6, 0xa1, 0x42, 0xc5, 0xa3, 0x3b, 0x67, 0xa8, 0x6a, 0x5c, 0xc5, 0x33, 0x2d,
	0x1a, 0x79, 0x1d, 0xe0, 0x69, 0xf9, 0xb3, 0xf6, 0x39, 0x05, 0x9f, 0x79, 0x98, 0x5d, 0x3b, 0xf9,
	0x7f, 0xe3, 0x09, 0x2e, 0x31, 0x35, 0xb8, 0xea, 0x54, 0x5a, 0xcc, 0xa5, 0x13, 0x27, 0x6c, 0x57,
	0xe4, 0x38, 0xf3, 0x93, 0x33, 0x14, 0x30, 0x36, 0x31, 0x96, 0xcf, 0xca, 0x4b, 0x3a, 0x41, 0xe4,
	0x57, 0xcf, 0xa4, 0x72, 0x97, 0x27, 0xca, 0x6d, 0x0e, 0xa1, 0xf6, 0x59, 0xd9, 0xca, 0xe8, 0xec,
	0x68, 0x32, 0x8d, 0xef, 0x08, 0xcf, 0xeb, 0xf9, 0xc9, 0xf7, 0x61, 0xbc, 0x19, 0xcb, 0x90, 0x9c,
	0x2e, 0x43, 0xc9, 0x4e, 0x13, 0x9c, 0xcf, 0x9f, 0x65, 0x7e, 0xa3, 0xe5, 0xf7, 0x18, 0x4f, 0x07,
	0x02, 0xc2, 0xb8, 0x5a, 0x4c, 0xdf, 0xfc, 0x95, 0x89, 0x2d, 0xe8, 0x75, 0xd9, 0x73, 0x4a, 0x71,
	0xfa, 0x91, 0xe4, 0x3a, 0x99, 0x44, 0xe3, 0x5b, 0x11, 0x6b, 0x9d, 0x19, 0x0c, 0x57, 0xa4, 0xc2,
	0x26, 0x74, 0xc0, 0x15, 0x8c, 0xab, 0xad, 0xba, 0x71, 0x8a, 0x90, 0xc8, 0xba, 0xc6, 0xcb, 0x76,
	0x6b, 0x41, 0x39, 0x56, 0xf4, 0x5f, 0xce, 0x01, 0x03, 0xe3, 0x15, 0x2e, 0x0b, 0xd6, 0x91, 0x3b,
	0x1e, 0xb0, 0x28, 0xef, 0xc8, 0xd4, 0xfd, 0xe4, 0x76, 0xc9, 0x68, 0x5e, 0x0e, 0x61, 0xf6, 0x05,
	0x25, 0x5c, 0x1e, 0xdd, 0xc5, 0x8e, 0xae, 0x53, 0xbb, 0x8d, 0xcf, 0x8f, 0xd5, 0x73, 0xc4, 0x1a,
	0x2d, 0xe8, 0x6b, 0x34, 0xab, 0xad, 0x85, 0xbd, 0xbc, 0xb5, 0x6f, 0x16, 0xb6, 0xf7, 0xcd, 0xc2,
	0xce, 0xbe, 0x59, 0xf8, 0x38, 0x30, 0xd1, 0xd6, 0xc0, 0x44, 0xdb, 0x03, 0x13, 0xed, 0x0c, 0x4c,
	0xf4, 0x7b, 0x60, 0xa2, 0xaf, 0x7f, 0xcc, 0xc2, 0xeb, 0x19, 0x95, 0xc3, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x5b, 0xcf, 0x13, 0x0c, 0x1b, 0x06, 0x00, 0x00,
}

func (m *Overhead) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Overhead) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Overhead) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PodFixed) > 0 {
		keysForPodFixed := make([]string, 0, len(m.PodFixed))
		for k := range m.PodFixed {
			keysForPodFixed = append(keysForPodFixed, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForPodFixed)
		for iNdEx := len(keysForPodFixed) - 1; iNdEx >= 0; iNdEx-- {
			v := m.PodFixed[k8s_io_api_core_v1.ResourceName(keysForPodFixed[iNdEx])]
			baseI := i
			{
				size, err := ((*resource.Quantity)(&v)).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(keysForPodFixed[iNdEx])
			copy(dAtA[i:], keysForPodFixed[iNdEx])
			i = encodeVarintGenerated(dAtA, i, uint64(len(keysForPodFixed[iNdEx])))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenerated(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RuntimeClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RuntimeClass) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RuntimeClass) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Scheduling != nil {
		{
			size, err := m.Scheduling.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Overhead != nil {
		{
			size, err := m.Overhead.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	i -= len(m.Handler)
	copy(dAtA[i:], m.Handler)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Handler)))
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *RuntimeClassList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RuntimeClassList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RuntimeClassList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ListMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Scheduling) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Scheduling) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Scheduling) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tolerations) > 0 {
		for iNdEx := len(m.Tolerations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tolerations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.NodeSelector) > 0 {
		keysForNodeSelector := make([]string, 0, len(m.NodeSelector))
		for k := range m.NodeSelector {
			keysForNodeSelector = append(keysForNodeSelector, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForNodeSelector)
		for iNdEx := len(keysForNodeSelector) - 1; iNdEx >= 0; iNdEx-- {
			v := m.NodeSelector[string(keysForNodeSelector[iNdEx])]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintGenerated(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(keysForNodeSelector[iNdEx])
			copy(dAtA[i:], keysForNodeSelector[iNdEx])
			i = encodeVarintGenerated(dAtA, i, uint64(len(keysForNodeSelector[iNdEx])))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenerated(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Overhead) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PodFixed) > 0 {
		for k, v := range m.PodFixed {
			_ = k
			_ = v
			l = ((*resource.Quantity)(&v)).Size()
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + l + sovGenerated(uint64(l))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *RuntimeClass) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Handler)
	n += 1 + l + sovGenerated(uint64(l))
	if m.Overhead != nil {
		l = m.Overhead.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Scheduling != nil {
		l = m.Scheduling.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *RuntimeClassList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *Scheduling) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.NodeSelector) > 0 {
		for k, v := range m.NodeSelector {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	if len(m.Tolerations) > 0 {
		for _, e := range m.Tolerations {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Overhead) String() string {
	if this == nil {
		return "nil"
	}
	keysForPodFixed := make([]string, 0, len(this.PodFixed))
	for k := range this.PodFixed {
		keysForPodFixed = append(keysForPodFixed, string(k))
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForPodFixed)
	mapStringForPodFixed := "k8s_io_api_core_v1.ResourceList{"
	for _, k := range keysForPodFixed {
		mapStringForPodFixed += fmt.Sprintf("%v: %v,", k, this.PodFixed[k8s_io_api_core_v1.ResourceName(k)])
	}
	mapStringForPodFixed += "}"
	s := strings.Join([]string{`&Overhead{`,
		`PodFixed:` + mapStringForPodFixed + `,`,
		`}`,
	}, "")
	return s
}
func (this *RuntimeClass) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RuntimeClass{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Handler:` + fmt.Sprintf("%v", this.Handler) + `,`,
		`Overhead:` + strings.Replace(this.Overhead.String(), "Overhead", "Overhead", 1) + `,`,
		`Scheduling:` + strings.Replace(this.Scheduling.String(), "Scheduling", "Scheduling", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RuntimeClassList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]RuntimeClass{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "RuntimeClass", "RuntimeClass", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&RuntimeClassList{`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func (this *Scheduling) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForTolerations := "[]Toleration{"
	for _, f := range this.Tolerations {
		repeatedStringForTolerations += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForTolerations += "}"
	keysForNodeSelector := make([]string, 0, len(this.NodeSelector))
	for k := range this.NodeSelector {
		keysForNodeSelector = append(keysForNodeSelector, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForNodeSelector)
	mapStringForNodeSelector := "map[string]string{"
	for _, k := range keysForNodeSelector {
		mapStringForNodeSelector += fmt.Sprintf("%v: %v,", k, this.NodeSelector[k])
	}
	mapStringForNodeSelector += "}"
	s := strings.Join([]string{`&Scheduling{`,
		`NodeSelector:` + mapStringForNodeSelector + `,`,
		`Tolerations:` + repeatedStringForTolerations + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Overhead) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: Overhead: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Overhead: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodFixed", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PodFixed == nil {
				m.PodFixed = make(k8s_io_api_core_v1.ResourceList)
			}
			var mapkey k8s_io_api_core_v1.ResourceName
			mapvalue := &resource.Quantity{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = k8s_io_api_core_v1.ResourceName(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthGenerated
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthGenerated
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &resource.Quantity{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenerated(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthGenerated
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.PodFixed[k8s_io_api_core_v1.ResourceName(mapkey)] = ((k8s_io_apimachinery_pkg_api_resource.Quantity)(*mapvalue))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *RuntimeClass) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: RuntimeClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RuntimeClass: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handler", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Handler = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Overhead", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Overhead == nil {
				m.Overhead = &Overhead{}
			}
			if err := m.Overhead.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scheduling", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Scheduling == nil {
				m.Scheduling = &Scheduling{}
			}
			if err := m.Scheduling.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *RuntimeClassList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: RuntimeClassList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RuntimeClassList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, RuntimeClass{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *Scheduling) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: Scheduling: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Scheduling: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeSelector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NodeSelector == nil {
				m.NodeSelector = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthGenerated
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenerated(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthGenerated
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.NodeSelector[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tolerations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tolerations = append(m.Tolerations, v11.Toleration{})
			if err := m.Tolerations[len(m.Tolerations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
