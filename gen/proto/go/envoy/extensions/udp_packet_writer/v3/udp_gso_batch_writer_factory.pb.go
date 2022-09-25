// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: envoy/extensions/udp_packet_writer/v3/udp_gso_batch_writer_factory.proto

package udp_packet_writerv3

import (
	_ "github.com/Dimss/envoy-authz/gen/proto/go/udpa/annotations"
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

// Configuration for the UDP GSO batch packet writer factory.
type UdpGsoBatchWriterFactory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UdpGsoBatchWriterFactory) Reset() {
	*x = UdpGsoBatchWriterFactory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UdpGsoBatchWriterFactory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UdpGsoBatchWriterFactory) ProtoMessage() {}

func (x *UdpGsoBatchWriterFactory) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UdpGsoBatchWriterFactory.ProtoReflect.Descriptor instead.
func (*UdpGsoBatchWriterFactory) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_rawDescGZIP(), []int{0}
}

var File_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto protoreflect.FileDescriptor

var file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_rawDesc = []byte{
	0x0a, 0x48, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x75, 0x64, 0x70, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x64, 0x70, 0x5f, 0x67, 0x73, 0x6f, 0x5f,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x5f, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x75, 0x64, 0x70,
	0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x33, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1a, 0x0a, 0x18, 0x55, 0x64, 0x70, 0x47, 0x73, 0x6f, 0x42, 0x61, 0x74, 0x63, 0x68, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x42, 0xe8, 0x02, 0x0a,
	0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x75, 0x64, 0x70, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x1d, 0x55, 0x64, 0x70, 0x47,
	0x73, 0x6f, 0x42, 0x61, 0x74, 0x63, 0x68, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x46, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a, 0x63, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x69, 0x6d, 0x73, 0x73, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x75, 0x64, 0x70, 0x5f, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x3b, 0x75,
	0x64, 0x70, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x76, 0x33, 0xa2, 0x02, 0x03, 0x45, 0x45, 0x55, 0xaa, 0x02, 0x23, 0x45, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x64, 0x70, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x33, 0xca, 0x02,
	0x23, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x5c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x5c, 0x55, 0x64, 0x70, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x5c, 0x56, 0x33, 0xe2, 0x02, 0x2f, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x5c, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x55, 0x64, 0x70, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x33, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x26, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x3a, 0x3a,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x55, 0x64, 0x70, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x33, 0xba,
	0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_rawDescOnce sync.Once
	file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_rawDescData = file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_rawDesc
)

func file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_rawDescGZIP() []byte {
	file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_rawDescData)
	})
	return file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_rawDescData
}

var file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_goTypes = []interface{}{
	(*UdpGsoBatchWriterFactory)(nil), // 0: envoy.extensions.udp_packet_writer.v3.UdpGsoBatchWriterFactory
}
var file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_init() }
func file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_init() {
	if File_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UdpGsoBatchWriterFactory); i {
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
			RawDescriptor: file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_msgTypes,
	}.Build()
	File_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto = out.File
	file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_rawDesc = nil
	file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_goTypes = nil
	file_envoy_extensions_udp_packet_writer_v3_udp_gso_batch_writer_factory_proto_depIdxs = nil
}
