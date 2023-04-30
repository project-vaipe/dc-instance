// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//  **** IMPORTANT NOTE ****
//
//  The proto of BT data has to match exactly the g3 proto, including tag
//  number.

// *** IMPORTANT NOTE ***
// Keep this in sync with the same proto in datacommonsorg/import.
// Eventually we need to move the shared protos out to a separate repo.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: mcf.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Type of MCF sub-graph.
type McfType int32

const (
	McfType_UNKNOWN_MCF_TYPE McfType = 0
	// The sub-graph is made of instances.
	McfType_INSTANCE_MCF McfType = 1
	// The sub-graph is made of templates.
	McfType_TEMPLATE_MCF McfType = 2
)

// Enum value maps for McfType.
var (
	McfType_name = map[int32]string{
		0: "UNKNOWN_MCF_TYPE",
		1: "INSTANCE_MCF",
		2: "TEMPLATE_MCF",
	}
	McfType_value = map[string]int32{
		"UNKNOWN_MCF_TYPE": 0,
		"INSTANCE_MCF":     1,
		"TEMPLATE_MCF":     2,
	}
)

func (x McfType) Enum() *McfType {
	p := new(McfType)
	*p = x
	return p
}

func (x McfType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (McfType) Descriptor() protoreflect.EnumDescriptor {
	return file_mcf_proto_enumTypes[0].Descriptor()
}

func (McfType) Type() protoreflect.EnumType {
	return &file_mcf_proto_enumTypes[0]
}

func (x McfType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *McfType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = McfType(num)
	return nil
}

// Deprecated: Use McfType.Descriptor instead.
func (McfType) EnumDescriptor() ([]byte, []int) {
	return file_mcf_proto_rawDescGZIP(), []int{0}
}

// Represents the type of property value.
type ValueType int32

const (
	ValueType_UNKNOWN_VALUE_TYPE ValueType = 0
	// Any value that has double quotes around it gets considered a text.
	ValueType_TEXT ValueType = 1
	// Any non-text and non-reference value gets interpreted as number (int /
	// uint / double / float).
	ValueType_NUMBER ValueType = 2
	// Represents a reference to a node that has not yet been resolved.  These
	// types should turn into RESOLVED_REF after entity resolution.
	// REQUIRES: the value has a reference prefix (aka, "<id-space>:"), typically
	// "l:".
	ValueType_UNRESOLVED_REF ValueType = 3
	// Represents a resolved datacommons entity. Stores a DCID value.
	ValueType_RESOLVED_REF ValueType = 4
	// Represents a complex value corresponding to Quantity, QuantityRange,
	// LatLng, etc.
	ValueType_COMPLEX_VALUE ValueType = 5
	// Represents a table column for TEMPLATE_MCF.
	ValueType_TABLE_COLUMN ValueType = 6
	// Represents a table entity for TEMPLATE_MCF.
	ValueType_TABLE_ENTITY ValueType = 7
)

// Enum value maps for ValueType.
var (
	ValueType_name = map[int32]string{
		0: "UNKNOWN_VALUE_TYPE",
		1: "TEXT",
		2: "NUMBER",
		3: "UNRESOLVED_REF",
		4: "RESOLVED_REF",
		5: "COMPLEX_VALUE",
		6: "TABLE_COLUMN",
		7: "TABLE_ENTITY",
	}
	ValueType_value = map[string]int32{
		"UNKNOWN_VALUE_TYPE": 0,
		"TEXT":               1,
		"NUMBER":             2,
		"UNRESOLVED_REF":     3,
		"RESOLVED_REF":       4,
		"COMPLEX_VALUE":      5,
		"TABLE_COLUMN":       6,
		"TABLE_ENTITY":       7,
	}
)

func (x ValueType) Enum() *ValueType {
	p := new(ValueType)
	*p = x
	return p
}

func (x ValueType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValueType) Descriptor() protoreflect.EnumDescriptor {
	return file_mcf_proto_enumTypes[1].Descriptor()
}

func (ValueType) Type() protoreflect.EnumType {
	return &file_mcf_proto_enumTypes[1]
}

func (x ValueType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ValueType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ValueType(num)
	return nil
}

// Deprecated: Use ValueType.Descriptor instead.
func (ValueType) EnumDescriptor() ([]byte, []int) {
	return file_mcf_proto_rawDescGZIP(), []int{1}
}

// Proto representation of a Data Commons sub-graph which may be made of
// instance nodes (INSTANCE_MCF) or template nodes (TEMPLATE_MCF).
type McfGraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of the sub-graph.
	Type *McfType `protobuf:"varint,1,opt,name=type,enum=datacommons.McfType,def=1" json:"type,omitempty"`
	// A map from Local Node ID to its property and values.
	Nodes map[string]*McfGraph_PropertyValues `protobuf:"bytes,2,rep,name=nodes" json:"nodes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

// Default values for McfGraph fields.
const (
	Default_McfGraph_Type = McfType_INSTANCE_MCF
)

func (x *McfGraph) Reset() {
	*x = McfGraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *McfGraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*McfGraph) ProtoMessage() {}

func (x *McfGraph) ProtoReflect() protoreflect.Message {
	mi := &file_mcf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use McfGraph.ProtoReflect.Descriptor instead.
func (*McfGraph) Descriptor() ([]byte, []int) {
	return file_mcf_proto_rawDescGZIP(), []int{0}
}

func (x *McfGraph) GetType() McfType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Default_McfGraph_Type
}

func (x *McfGraph) GetNodes() map[string]*McfGraph_PropertyValues {
	if x != nil {
		return x.Nodes
	}
	return nil
}

// Represents a single value.
type McfGraph_TypedValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  *ValueType `protobuf:"varint,1,opt,name=type,enum=datacommons.ValueType" json:"type,omitempty"`
	Value *string    `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (x *McfGraph_TypedValue) Reset() {
	*x = McfGraph_TypedValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *McfGraph_TypedValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*McfGraph_TypedValue) ProtoMessage() {}

func (x *McfGraph_TypedValue) ProtoReflect() protoreflect.Message {
	mi := &file_mcf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use McfGraph_TypedValue.ProtoReflect.Descriptor instead.
func (*McfGraph_TypedValue) Descriptor() ([]byte, []int) {
	return file_mcf_proto_rawDescGZIP(), []int{0, 0}
}

func (x *McfGraph_TypedValue) GetType() ValueType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ValueType_UNKNOWN_VALUE_TYPE
}

func (x *McfGraph_TypedValue) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

// Represents a list of property values.
type McfGraph_Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of typed values.
	TypedValues []*McfGraph_TypedValue `protobuf:"bytes,1,rep,name=typed_values,json=typedValues" json:"typed_values,omitempty"`
}

func (x *McfGraph_Values) Reset() {
	*x = McfGraph_Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *McfGraph_Values) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*McfGraph_Values) ProtoMessage() {}

func (x *McfGraph_Values) ProtoReflect() protoreflect.Message {
	mi := &file_mcf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use McfGraph_Values.ProtoReflect.Descriptor instead.
func (*McfGraph_Values) Descriptor() ([]byte, []int) {
	return file_mcf_proto_rawDescGZIP(), []int{0, 1}
}

func (x *McfGraph_Values) GetTypedValues() []*McfGraph_TypedValue {
	if x != nil {
		return x.TypedValues
	}
	return nil
}

// Information about a node.
type McfGraph_PropertyValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map of a property name to its values.
	Pvs map[string]*McfGraph_Values `protobuf:"bytes,1,rep,name=pvs" json:"pvs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Information identifying the location of this node in the source.
	// There can be multiple if PVs in this node are merged from different
	// files.
	Location []*Log_Location `protobuf:"bytes,2,rep,name=location" json:"location,omitempty"`
}

func (x *McfGraph_PropertyValues) Reset() {
	*x = McfGraph_PropertyValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *McfGraph_PropertyValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*McfGraph_PropertyValues) ProtoMessage() {}

func (x *McfGraph_PropertyValues) ProtoReflect() protoreflect.Message {
	mi := &file_mcf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use McfGraph_PropertyValues.ProtoReflect.Descriptor instead.
func (*McfGraph_PropertyValues) Descriptor() ([]byte, []int) {
	return file_mcf_proto_rawDescGZIP(), []int{0, 2}
}

func (x *McfGraph_PropertyValues) GetPvs() map[string]*McfGraph_Values {
	if x != nil {
		return x.Pvs
	}
	return nil
}

func (x *McfGraph_PropertyValues) GetLocation() []*Log_Location {
	if x != nil {
		return x.Location
	}
	return nil
}

var File_mcf_proto protoreflect.FileDescriptor

var file_mcf_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x63, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x1a, 0x0b, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x04, 0x0a, 0x08, 0x4d, 0x63, 0x66, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4d,
	0x63, 0x66, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x0c, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45,
	0x5f, 0x4d, 0x43, 0x46, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x63, 0x66, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x1a, 0x4e, 0x0a, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x4d, 0x0a, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x0c,
	0x74, 0x79, 0x70, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x4d, 0x63, 0x66, 0x47, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x1a, 0xde, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x03, 0x70, 0x76, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x4d, 0x63, 0x66, 0x47, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x50, 0x76, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x03, 0x70, 0x76, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x54, 0x0a, 0x08,
	0x50, 0x76, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x63, 0x66, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x5e, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x4d, 0x63, 0x66, 0x47, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x2a, 0x43, 0x0a, 0x07, 0x4d, 0x63, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x10, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4d, 0x43, 0x46, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f,
	0x4d, 0x43, 0x46, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x54,
	0x45, 0x5f, 0x4d, 0x43, 0x46, 0x10, 0x02, 0x2a, 0x96, 0x01, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x55, 0x4d, 0x42, 0x45,
	0x52, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45,
	0x44, 0x5f, 0x52, 0x45, 0x46, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x53, 0x4f, 0x4c,
	0x56, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x46, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x4d,
	0x50, 0x4c, 0x45, 0x58, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c,
	0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x10, 0x06, 0x12, 0x10,
	0x0a, 0x0c, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x10, 0x07,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_mcf_proto_rawDescOnce sync.Once
	file_mcf_proto_rawDescData = file_mcf_proto_rawDesc
)

func file_mcf_proto_rawDescGZIP() []byte {
	file_mcf_proto_rawDescOnce.Do(func() {
		file_mcf_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcf_proto_rawDescData)
	})
	return file_mcf_proto_rawDescData
}

var file_mcf_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_mcf_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mcf_proto_goTypes = []interface{}{
	(McfType)(0),                    // 0: datacommons.McfType
	(ValueType)(0),                  // 1: datacommons.ValueType
	(*McfGraph)(nil),                // 2: datacommons.McfGraph
	(*McfGraph_TypedValue)(nil),     // 3: datacommons.McfGraph.TypedValue
	(*McfGraph_Values)(nil),         // 4: datacommons.McfGraph.Values
	(*McfGraph_PropertyValues)(nil), // 5: datacommons.McfGraph.PropertyValues
	nil,                             // 6: datacommons.McfGraph.NodesEntry
	nil,                             // 7: datacommons.McfGraph.PropertyValues.PvsEntry
	(*Log_Location)(nil),            // 8: datacommons.Log.Location
}
var file_mcf_proto_depIdxs = []int32{
	0, // 0: datacommons.McfGraph.type:type_name -> datacommons.McfType
	6, // 1: datacommons.McfGraph.nodes:type_name -> datacommons.McfGraph.NodesEntry
	1, // 2: datacommons.McfGraph.TypedValue.type:type_name -> datacommons.ValueType
	3, // 3: datacommons.McfGraph.Values.typed_values:type_name -> datacommons.McfGraph.TypedValue
	7, // 4: datacommons.McfGraph.PropertyValues.pvs:type_name -> datacommons.McfGraph.PropertyValues.PvsEntry
	8, // 5: datacommons.McfGraph.PropertyValues.location:type_name -> datacommons.Log.Location
	5, // 6: datacommons.McfGraph.NodesEntry.value:type_name -> datacommons.McfGraph.PropertyValues
	4, // 7: datacommons.McfGraph.PropertyValues.PvsEntry.value:type_name -> datacommons.McfGraph.Values
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_mcf_proto_init() }
func file_mcf_proto_init() {
	if File_mcf_proto != nil {
		return
	}
	file_debug_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mcf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*McfGraph); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mcf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*McfGraph_TypedValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mcf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*McfGraph_Values); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mcf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*McfGraph_PropertyValues); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mcf_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mcf_proto_goTypes,
		DependencyIndexes: file_mcf_proto_depIdxs,
		EnumInfos:         file_mcf_proto_enumTypes,
		MessageInfos:      file_mcf_proto_msgTypes,
	}.Build()
	File_mcf_proto = out.File
	file_mcf_proto_rawDesc = nil
	file_mcf_proto_goTypes = nil
	file_mcf_proto_depIdxs = nil
}
