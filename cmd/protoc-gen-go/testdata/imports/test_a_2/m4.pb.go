// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imports/test_a_2/m4.proto

package test_a_2

import (
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoregistry "github.com/golang/protobuf/v2/reflect/protoregistry"
	protoiface "github.com/golang/protobuf/v2/runtime/protoiface"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

type M4 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (x *M4) Reset() {
	*x = M4{}
}

func (x *M4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M4) ProtoMessage() {}

func (x *M4) ProtoReflect() protoreflect.Message {
	return xxx_File_imports_test_a_2_m4_proto_messageTypes[0].MessageOf(x)
}

func (m *M4) XXX_Methods() *protoiface.Methods {
	return xxx_File_imports_test_a_2_m4_proto_messageTypes[0].Methods()
}

// Deprecated: Use M4.ProtoReflect.Type instead.
func (*M4) Descriptor() ([]byte, []int) {
	return xxx_File_imports_test_a_2_m4_proto_rawDescGZIP(), []int{0}
}

var File_imports_test_a_2_m4_proto protoreflect.FileDescriptor

var xxx_File_imports_test_a_2_m4_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61,
	0x5f, 0x32, 0x2f, 0x6d, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x61, 0x22, 0x04, 0x0a, 0x02, 0x4d, 0x34, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x61, 0x5f, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	xxx_File_imports_test_a_2_m4_proto_rawDesc_once sync.Once
	xxx_File_imports_test_a_2_m4_proto_rawDesc_data = xxx_File_imports_test_a_2_m4_proto_rawDesc
)

func xxx_File_imports_test_a_2_m4_proto_rawDescGZIP() []byte {
	xxx_File_imports_test_a_2_m4_proto_rawDesc_once.Do(func() {
		xxx_File_imports_test_a_2_m4_proto_rawDesc_data = protoimpl.X.CompressGZIP(xxx_File_imports_test_a_2_m4_proto_rawDesc_data)
	})
	return xxx_File_imports_test_a_2_m4_proto_rawDesc_data
}

var xxx_File_imports_test_a_2_m4_proto_messageTypes = make([]protoimpl.MessageType, 1)
var xxx_File_imports_test_a_2_m4_proto_goTypes = []interface{}{
	(*M4)(nil), // 0: test.a.M4
}
var xxx_File_imports_test_a_2_m4_proto_depIdxs = []int32{}

func init() { xxx_File_imports_test_a_2_m4_proto_init() }
func xxx_File_imports_test_a_2_m4_proto_init() {
	if File_imports_test_a_2_m4_proto != nil {
		return
	}
	File_imports_test_a_2_m4_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_imports_test_a_2_m4_proto_rawDesc,
		GoTypes:            xxx_File_imports_test_a_2_m4_proto_goTypes,
		DependencyIndexes:  xxx_File_imports_test_a_2_m4_proto_depIdxs,
		MessageOutputTypes: xxx_File_imports_test_a_2_m4_proto_messageTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	xxx_File_imports_test_a_2_m4_proto_rawDesc = nil
	xxx_File_imports_test_a_2_m4_proto_goTypes = nil
	xxx_File_imports_test_a_2_m4_proto_depIdxs = nil
}
