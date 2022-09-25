// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: envoy/service/secret/v3/sds.proto

package secretv3

import (
	_ "github.com/Dimss/envoy-authz/gen/proto/go/envoy/annotations"
	v3 "github.com/Dimss/envoy-authz/gen/proto/go/envoy/service/discovery/v3"
	_ "github.com/Dimss/envoy-authz/gen/proto/go/google/api"
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

// [#not-implemented-hide:] Not configuration. Workaround c++ protobuf issue with importing
// services: https://github.com/google/protobuf/issues/4221
type SdsDummy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SdsDummy) Reset() {
	*x = SdsDummy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_secret_v3_sds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SdsDummy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SdsDummy) ProtoMessage() {}

func (x *SdsDummy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_secret_v3_sds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SdsDummy.ProtoReflect.Descriptor instead.
func (*SdsDummy) Descriptor() ([]byte, []int) {
	return file_envoy_service_secret_v3_sds_proto_rawDescGZIP(), []int{0}
}

var File_envoy_service_secret_v3_sds_proto protoreflect.FileDescriptor

var file_envoy_service_secret_v3_sds_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x17, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x76, 0x33, 0x1a, 0x2a, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x08, 0x53, 0x64,
	0x73, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x3a, 0x2a, 0x9a, 0xc5, 0x88, 0x1e, 0x25, 0x0a, 0x23, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x64, 0x73, 0x44, 0x75, 0x6d,
	0x6d, 0x79, 0x32, 0xd3, 0x03, 0x0a, 0x16, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7b, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x31, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x32, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65,
	0x6c, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x72, 0x0a, 0x0d, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x2c, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x8d,
	0x01, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12,
	0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x3a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x1a, 0x38,
	0x8a, 0xa4, 0x96, 0xf3, 0x07, 0x32, 0x0a, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x76,
	0x33, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x42, 0xff, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x76, 0x33, 0x42, 0x08, 0x53, 0x64, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x44, 0x69, 0x6d, 0x73, 0x73, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2d, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2f, 0x76, 0x33, 0x3b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x76, 0x33, 0x88, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x45, 0x53, 0x53, 0xaa, 0x02, 0x17, 0x45, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x2e, 0x56, 0x33, 0xca, 0x02, 0x17, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x5c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5c, 0x56, 0x33, 0xe2,
	0x02, 0x23, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5c, 0x56, 0x33, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x3a, 0x3a, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x3a, 0x3a,
	0x56, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_envoy_service_secret_v3_sds_proto_rawDescOnce sync.Once
	file_envoy_service_secret_v3_sds_proto_rawDescData = file_envoy_service_secret_v3_sds_proto_rawDesc
)

func file_envoy_service_secret_v3_sds_proto_rawDescGZIP() []byte {
	file_envoy_service_secret_v3_sds_proto_rawDescOnce.Do(func() {
		file_envoy_service_secret_v3_sds_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_service_secret_v3_sds_proto_rawDescData)
	})
	return file_envoy_service_secret_v3_sds_proto_rawDescData
}

var file_envoy_service_secret_v3_sds_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_service_secret_v3_sds_proto_goTypes = []interface{}{
	(*SdsDummy)(nil),                  // 0: envoy.service.secret.v3.SdsDummy
	(*v3.DeltaDiscoveryRequest)(nil),  // 1: envoy.service.discovery.v3.DeltaDiscoveryRequest
	(*v3.DiscoveryRequest)(nil),       // 2: envoy.service.discovery.v3.DiscoveryRequest
	(*v3.DeltaDiscoveryResponse)(nil), // 3: envoy.service.discovery.v3.DeltaDiscoveryResponse
	(*v3.DiscoveryResponse)(nil),      // 4: envoy.service.discovery.v3.DiscoveryResponse
}
var file_envoy_service_secret_v3_sds_proto_depIdxs = []int32{
	1, // 0: envoy.service.secret.v3.SecretDiscoveryService.DeltaSecrets:input_type -> envoy.service.discovery.v3.DeltaDiscoveryRequest
	2, // 1: envoy.service.secret.v3.SecretDiscoveryService.StreamSecrets:input_type -> envoy.service.discovery.v3.DiscoveryRequest
	2, // 2: envoy.service.secret.v3.SecretDiscoveryService.FetchSecrets:input_type -> envoy.service.discovery.v3.DiscoveryRequest
	3, // 3: envoy.service.secret.v3.SecretDiscoveryService.DeltaSecrets:output_type -> envoy.service.discovery.v3.DeltaDiscoveryResponse
	4, // 4: envoy.service.secret.v3.SecretDiscoveryService.StreamSecrets:output_type -> envoy.service.discovery.v3.DiscoveryResponse
	4, // 5: envoy.service.secret.v3.SecretDiscoveryService.FetchSecrets:output_type -> envoy.service.discovery.v3.DiscoveryResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_service_secret_v3_sds_proto_init() }
func file_envoy_service_secret_v3_sds_proto_init() {
	if File_envoy_service_secret_v3_sds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_service_secret_v3_sds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SdsDummy); i {
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
			RawDescriptor: file_envoy_service_secret_v3_sds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_envoy_service_secret_v3_sds_proto_goTypes,
		DependencyIndexes: file_envoy_service_secret_v3_sds_proto_depIdxs,
		MessageInfos:      file_envoy_service_secret_v3_sds_proto_msgTypes,
	}.Build()
	File_envoy_service_secret_v3_sds_proto = out.File
	file_envoy_service_secret_v3_sds_proto_rawDesc = nil
	file_envoy_service_secret_v3_sds_proto_goTypes = nil
	file_envoy_service_secret_v3_sds_proto_depIdxs = nil
}
