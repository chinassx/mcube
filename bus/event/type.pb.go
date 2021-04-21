//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pb/event/type.proto

package event

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Type int32

const (
	// 全局配置, 所有服务可以读取
	Type_OperateEvent Type = 0
	// 组内配置, 组里面的服务可以读取
	Type_StatusEvent Type = 1
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "OperateEvent",
		1: "StatusEvent",
	}
	Type_value = map[string]int32{
		"OperateEvent": 0,
		"StatusEvent":  1,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_event_type_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_pb_event_type_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_pb_event_type_proto_rawDescGZIP(), []int{0}
}

var File_pb_event_type_proto protoreflect.FileDescriptor

var file_pb_event_type_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2a, 0x29, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x62, 0x75, 0x73,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_event_type_proto_rawDescOnce sync.Once
	file_pb_event_type_proto_rawDescData = file_pb_event_type_proto_rawDesc
)

func file_pb_event_type_proto_rawDescGZIP() []byte {
	file_pb_event_type_proto_rawDescOnce.Do(func() {
		file_pb_event_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_event_type_proto_rawDescData)
	})
	return file_pb_event_type_proto_rawDescData
}

var file_pb_event_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_event_type_proto_goTypes = []interface{}{
	(Type)(0), // 0: mcube.event.Type
}
var file_pb_event_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_event_type_proto_init() }
func file_pb_event_type_proto_init() {
	if File_pb_event_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_event_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_event_type_proto_goTypes,
		DependencyIndexes: file_pb_event_type_proto_depIdxs,
		EnumInfos:         file_pb_event_type_proto_enumTypes,
	}.Build()
	File_pb_event_type_proto = out.File
	file_pb_event_type_proto_rawDesc = nil
	file_pb_event_type_proto_goTypes = nil
	file_pb_event_type_proto_depIdxs = nil
}