// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: interpret/operation.proto

package interpret

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

type OpUnary_Operator int32

const (
	OpUnary_UNSPECIFIED OpUnary_Operator = 0
	OpUnary_LEN         OpUnary_Operator = 1
	OpUnary_NOT         OpUnary_Operator = 2
	OpUnary_FLAT        OpUnary_Operator = 3
	OpUnary_FLOOR       OpUnary_Operator = 4
	OpUnary_CEIL        OpUnary_Operator = 5
	OpUnary_ABORT       OpUnary_Operator = 6
)

// Enum value maps for OpUnary_Operator.
var (
	OpUnary_Operator_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "LEN",
		2: "NOT",
		3: "FLAT",
		4: "FLOOR",
		5: "CEIL",
		6: "ABORT",
	}
	OpUnary_Operator_value = map[string]int32{
		"UNSPECIFIED": 0,
		"LEN":         1,
		"NOT":         2,
		"FLAT":        3,
		"FLOOR":       4,
		"CEIL":        5,
		"ABORT":       6,
	}
)

func (x OpUnary_Operator) Enum() *OpUnary_Operator {
	p := new(OpUnary_Operator)
	*p = x
	return p
}

func (x OpUnary_Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpUnary_Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_interpret_operation_proto_enumTypes[0].Descriptor()
}

func (OpUnary_Operator) Type() protoreflect.EnumType {
	return &file_interpret_operation_proto_enumTypes[0]
}

func (x OpUnary_Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpUnary_Operator.Descriptor instead.
func (OpUnary_Operator) EnumDescriptor() ([]byte, []int) {
	return file_interpret_operation_proto_rawDescGZIP(), []int{0, 0}
}

type OpBinary_Operator int32

const (
	OpBinary_UNSPECIFIED OpBinary_Operator = 0
	OpBinary_SUB         OpBinary_Operator = 1
	OpBinary_DIV         OpBinary_Operator = 2
	OpBinary_MOD         OpBinary_Operator = 3
	OpBinary_EQ          OpBinary_Operator = 4
	OpBinary_NEQ         OpBinary_Operator = 5
	OpBinary_LT          OpBinary_Operator = 6
	OpBinary_LTE         OpBinary_Operator = 7
	OpBinary_GT          OpBinary_Operator = 8
	OpBinary_GTE         OpBinary_Operator = 9
)

// Enum value maps for OpBinary_Operator.
var (
	OpBinary_Operator_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "SUB",
		2: "DIV",
		3: "MOD",
		4: "EQ",
		5: "NEQ",
		6: "LT",
		7: "LTE",
		8: "GT",
		9: "GTE",
	}
	OpBinary_Operator_value = map[string]int32{
		"UNSPECIFIED": 0,
		"SUB":         1,
		"DIV":         2,
		"MOD":         3,
		"EQ":          4,
		"NEQ":         5,
		"LT":          6,
		"LTE":         7,
		"GT":          8,
		"GTE":         9,
	}
)

func (x OpBinary_Operator) Enum() *OpBinary_Operator {
	p := new(OpBinary_Operator)
	*p = x
	return p
}

func (x OpBinary_Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpBinary_Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_interpret_operation_proto_enumTypes[1].Descriptor()
}

func (OpBinary_Operator) Type() protoreflect.EnumType {
	return &file_interpret_operation_proto_enumTypes[1]
}

func (x OpBinary_Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpBinary_Operator.Descriptor instead.
func (OpBinary_Operator) EnumDescriptor() ([]byte, []int) {
	return file_interpret_operation_proto_rawDescGZIP(), []int{1, 0}
}

type OpVariadic_Operator int32

const (
	OpVariadic_UNSPECIFIED OpVariadic_Operator = 0
	OpVariadic_ADD         OpVariadic_Operator = 1
	OpVariadic_MUL         OpVariadic_Operator = 2
	OpVariadic_AND         OpVariadic_Operator = 3
	OpVariadic_OR          OpVariadic_Operator = 4
	OpVariadic_CAT         OpVariadic_Operator = 5
	OpVariadic_MIN         OpVariadic_Operator = 6
	OpVariadic_MAX         OpVariadic_Operator = 7
	OpVariadic_MERGE       OpVariadic_Operator = 8
)

// Enum value maps for OpVariadic_Operator.
var (
	OpVariadic_Operator_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "ADD",
		2: "MUL",
		3: "AND",
		4: "OR",
		5: "CAT",
		6: "MIN",
		7: "MAX",
		8: "MERGE",
	}
	OpVariadic_Operator_value = map[string]int32{
		"UNSPECIFIED": 0,
		"ADD":         1,
		"MUL":         2,
		"AND":         3,
		"OR":          4,
		"CAT":         5,
		"MIN":         6,
		"MAX":         7,
		"MERGE":       8,
	}
)

func (x OpVariadic_Operator) Enum() *OpVariadic_Operator {
	p := new(OpVariadic_Operator)
	*p = x
	return p
}

func (x OpVariadic_Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpVariadic_Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_interpret_operation_proto_enumTypes[2].Descriptor()
}

func (OpVariadic_Operator) Type() protoreflect.EnumType {
	return &file_interpret_operation_proto_enumTypes[2]
}

func (x OpVariadic_Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpVariadic_Operator.Descriptor instead.
func (OpVariadic_Operator) EnumDescriptor() ([]byte, []int) {
	return file_interpret_operation_proto_rawDescGZIP(), []int{2, 0}
}

type OpUnary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OpUnary) Reset() {
	*x = OpUnary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interpret_operation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpUnary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpUnary) ProtoMessage() {}

func (x *OpUnary) ProtoReflect() protoreflect.Message {
	mi := &file_interpret_operation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpUnary.ProtoReflect.Descriptor instead.
func (*OpUnary) Descriptor() ([]byte, []int) {
	return file_interpret_operation_proto_rawDescGZIP(), []int{0}
}

type OpBinary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OpBinary) Reset() {
	*x = OpBinary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interpret_operation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpBinary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpBinary) ProtoMessage() {}

func (x *OpBinary) ProtoReflect() protoreflect.Message {
	mi := &file_interpret_operation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpBinary.ProtoReflect.Descriptor instead.
func (*OpBinary) Descriptor() ([]byte, []int) {
	return file_interpret_operation_proto_rawDescGZIP(), []int{1}
}

type OpVariadic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OpVariadic) Reset() {
	*x = OpVariadic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interpret_operation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpVariadic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpVariadic) ProtoMessage() {}

func (x *OpVariadic) ProtoReflect() protoreflect.Message {
	mi := &file_interpret_operation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpVariadic.ProtoReflect.Descriptor instead.
func (*OpVariadic) Descriptor() ([]byte, []int) {
	return file_interpret_operation_proto_rawDescGZIP(), []int{2}
}

var File_interpret_operation_proto protoreflect.FileDescriptor

var file_interpret_operation_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x2f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x22, 0x62, 0x0a, 0x07, 0x4f, 0x70, 0x55, 0x6e, 0x61, 0x72,
	0x79, 0x22, 0x57, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x4c, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x4f, 0x54, 0x10, 0x02,
	0x12, 0x08, 0x0a, 0x04, 0x46, 0x4c, 0x41, 0x54, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x4c,
	0x4f, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x45, 0x49, 0x4c, 0x10, 0x05, 0x12,
	0x09, 0x0a, 0x05, 0x41, 0x42, 0x4f, 0x52, 0x54, 0x10, 0x06, 0x22, 0x75, 0x0a, 0x08, 0x4f, 0x70,
	0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x22, 0x69, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x55, 0x42, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03,
	0x44, 0x49, 0x56, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x4f, 0x44, 0x10, 0x03, 0x12, 0x06,
	0x0a, 0x02, 0x45, 0x51, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x45, 0x51, 0x10, 0x05, 0x12,
	0x06, 0x0a, 0x02, 0x4c, 0x54, 0x10, 0x06, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x54, 0x45, 0x10, 0x07,
	0x12, 0x06, 0x0a, 0x02, 0x47, 0x54, 0x10, 0x08, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x54, 0x45, 0x10,
	0x09, 0x22, 0x72, 0x0a, 0x0a, 0x4f, 0x70, 0x56, 0x61, 0x72, 0x69, 0x61, 0x64, 0x69, 0x63, 0x22,
	0x64, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x41, 0x44, 0x44, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x55, 0x4c, 0x10, 0x02, 0x12, 0x07,
	0x0a, 0x03, 0x41, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x52, 0x10, 0x04, 0x12,
	0x07, 0x0a, 0x03, 0x43, 0x41, 0x54, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x49, 0x4e, 0x10,
	0x06, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x58, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x45,
	0x52, 0x47, 0x45, 0x10, 0x08, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x75, 0x6d, 0x70, 0x61, 0x6b, 0x75, 0x2f, 0x65, 0x79, 0x61, 0x6d,
	0x6c, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interpret_operation_proto_rawDescOnce sync.Once
	file_interpret_operation_proto_rawDescData = file_interpret_operation_proto_rawDesc
)

func file_interpret_operation_proto_rawDescGZIP() []byte {
	file_interpret_operation_proto_rawDescOnce.Do(func() {
		file_interpret_operation_proto_rawDescData = protoimpl.X.CompressGZIP(file_interpret_operation_proto_rawDescData)
	})
	return file_interpret_operation_proto_rawDescData
}

var file_interpret_operation_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_interpret_operation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_interpret_operation_proto_goTypes = []any{
	(OpUnary_Operator)(0),    // 0: interpret.OpUnary.Operator
	(OpBinary_Operator)(0),   // 1: interpret.OpBinary.Operator
	(OpVariadic_Operator)(0), // 2: interpret.OpVariadic.Operator
	(*OpUnary)(nil),          // 3: interpret.OpUnary
	(*OpBinary)(nil),         // 4: interpret.OpBinary
	(*OpVariadic)(nil),       // 5: interpret.OpVariadic
}
var file_interpret_operation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_interpret_operation_proto_init() }
func file_interpret_operation_proto_init() {
	if File_interpret_operation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interpret_operation_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*OpUnary); i {
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
		file_interpret_operation_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*OpBinary); i {
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
		file_interpret_operation_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*OpVariadic); i {
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
			RawDescriptor: file_interpret_operation_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_interpret_operation_proto_goTypes,
		DependencyIndexes: file_interpret_operation_proto_depIdxs,
		EnumInfos:         file_interpret_operation_proto_enumTypes,
		MessageInfos:      file_interpret_operation_proto_msgTypes,
	}.Build()
	File_interpret_operation_proto = out.File
	file_interpret_operation_proto_rawDesc = nil
	file_interpret_operation_proto_goTypes = nil
	file_interpret_operation_proto_depIdxs = nil
}
