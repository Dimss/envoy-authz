// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: envoy/extensions/filters/network/meta_protocol_proxy/v3/route.proto

package meta_protocol_proxyv3

import (
	_ "github.com/Dimss/envoy-authz/gen/proto/go/udpa/annotations"
	_ "github.com/Dimss/envoy-authz/gen/proto/go/validate"
	_ "github.com/Dimss/envoy-authz/gen/proto/go/xds/annotations/v3"
	v3 "github.com/Dimss/envoy-authz/gen/proto/go/xds/type/matcher/v3"
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

// [#protodoc-title: Meta Protocol Proxy Route Configuration]
// The meta protocol proxy makes use of the `xds matching API` for routing configurations.
//
// In the below example, we combine a top level tree matcher with a linear matcher to match
// the incoming requests, and send the matching requests to v1 of the upstream service.
//
// name: demo-v1
// route:
//   matcher_tree:
//     input:
//       name: request-service
//       typed_config:
//         "@type": type.googleapis.com/envoy.extensions.filters.network.meta_protocol_proxy.matcher.v3.ServiceMatchInput
//     exact_match_map:
//       map:
//         org.apache.dubbo.samples.basic.api.DemoService:
//           matcher:
//             matcher_list:
//               matchers:
//               - predicate:
//                   and_matcher:
//                     predicate:
//                     - single_predicate:
//                         input:
//                           name: request-properties
//                           typed_config:
//                             "@type": type.googleapis.com/envoy.extensions.filters.network.meta_protocol_proxy.matcher.v3.PropertyMatchInput
//                             property_name: version
//                         value_match:
//                           exact: v1
//                     - single_predicate:
//                         input:
//                           name: request-properties
//                           typed_config:
//                             "@type": type.googleapis.com/envoy.extensions.filters.network.meta_protocol_proxy.matcher.v3.PropertyMatchInput
//                             property_name: user
//                         value_match:
//                           exact: john
//                 on_match:
//                   action:
//                     name: route
//                     typed_config:
//                       "@type": type.googleapis.com/envoy.extensions.filters.network.meta_protocol_proxy.matcher.action.v3.routeAction
//                       cluster: outbound|20880|v1|org.apache.dubbo.samples.basic.api.demoservice
// [#not-implemented-hide:]
type RouteConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the route configuration. For example, it might match route_config_name in
	// envoy.extensions.filters.network.meta_protocol_proxy.v3.Rds.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The match tree to use when resolving route actions for incoming requests.
	Route *v3.Matcher `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
}

func (x *RouteConfiguration) Reset() {
	*x = RouteConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteConfiguration) ProtoMessage() {}

func (x *RouteConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteConfiguration.ProtoReflect.Descriptor instead.
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_rawDescGZIP(), []int{0}
}

func (x *RouteConfiguration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RouteConfiguration) GetRoute() *v3.Matcher {
	if x != nil {
		return x.Route
	}
	return nil
}

var File_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_rawDesc = []byte{
	0x0a, 0x43, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x37, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x1a, 0x1f,
	0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x78, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x12, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a,
	0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x78,
	0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x76, 0x33, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x42, 0xcf, 0x03, 0x0a, 0x3b,
	0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x42, 0x0a, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a, 0x77, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x69, 0x6d, 0x73, 0x73, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x3b, 0x6d,
	0x65, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x76, 0x33, 0xa2, 0x02, 0x05, 0x45, 0x45, 0x46, 0x4e, 0x4d, 0xaa, 0x02, 0x35, 0x45,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x56, 0x33, 0xca, 0x02, 0x35, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x5c, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x5c,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5c, 0x4d, 0x65, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x5c, 0x56, 0x33, 0xe2, 0x02, 0x41, 0x45,
	0x6e, 0x76, 0x6f, 0x79, 0x5c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5c,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x5c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5c,
	0x4d, 0x65, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x5c, 0x56, 0x33, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x3a, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x3a, 0x3a, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x3a, 0x3a, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x3a, 0x3a, 0x4d, 0x65, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x3a, 0x3a, 0x56, 0x33, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_rawDescData = file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_rawDesc
)

func file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_rawDescData)
	})
	return file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_rawDescData
}

var file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_goTypes = []interface{}{
	(*RouteConfiguration)(nil), // 0: envoy.extensions.filters.network.meta_protocol_proxy.v3.RouteConfiguration
	(*v3.Matcher)(nil),         // 1: xds.type.matcher.v3.Matcher
}
var file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.network.meta_protocol_proxy.v3.RouteConfiguration.route:type_name -> xds.type.matcher.v3.Matcher
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_init() }
func file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_init() {
	if File_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteConfiguration); i {
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
			RawDescriptor: file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto = out.File
	file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_rawDesc = nil
	file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_goTypes = nil
	file_envoy_extensions_filters_network_meta_protocol_proxy_v3_route_proto_depIdxs = nil
}
