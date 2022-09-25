// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: envoy/data/cluster/v3/outlier_detection_event.proto

package clusterv3

import (
	_ "github.com/Dimss/envoy-authz/gen/proto/go/udpa/annotations"
	_ "github.com/Dimss/envoy-authz/gen/proto/go/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Type of ejection that took place
type OutlierEjectionType int32

const (
	// In case upstream host returns certain number of consecutive 5xx.
	// If
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_v3_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is ``false``, all type of errors are treated as HTTP 5xx errors.
	// See :ref:`Cluster outlier detection <arch_overview_outlier_detection>` documentation for
	// details.
	OutlierEjectionType_CONSECUTIVE_5XX OutlierEjectionType = 0
	// In case upstream host returns certain number of consecutive gateway errors
	OutlierEjectionType_CONSECUTIVE_GATEWAY_FAILURE OutlierEjectionType = 1
	// Runs over aggregated success rate statistics from every host in cluster
	// and selects hosts for which ratio of successful replies deviates from other hosts
	// in the cluster.
	// If
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_v3_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is ``false``, all errors (externally and locally generated) are used to calculate success rate
	// statistics. See :ref:`Cluster outlier detection <arch_overview_outlier_detection>`
	// documentation for details.
	OutlierEjectionType_SUCCESS_RATE OutlierEjectionType = 2
	// Consecutive local origin failures: Connection failures, resets, timeouts, etc
	// This type of ejection happens only when
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_v3_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is set to ``true``.
	// See :ref:`Cluster outlier detection <arch_overview_outlier_detection>` documentation for
	OutlierEjectionType_CONSECUTIVE_LOCAL_ORIGIN_FAILURE OutlierEjectionType = 3
	// Runs over aggregated success rate statistics for local origin failures
	// for all hosts in the cluster and selects hosts for which success rate deviates from other
	// hosts in the cluster. This type of ejection happens only when
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_v3_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is set to ``true``.
	// See :ref:`Cluster outlier detection <arch_overview_outlier_detection>` documentation for
	OutlierEjectionType_SUCCESS_RATE_LOCAL_ORIGIN OutlierEjectionType = 4
	// Runs over aggregated success rate statistics from every host in cluster and selects hosts for
	// which ratio of failed replies is above configured value.
	OutlierEjectionType_FAILURE_PERCENTAGE OutlierEjectionType = 5
	// Runs over aggregated success rate statistics for local origin failures from every host in
	// cluster and selects hosts for which ratio of failed replies is above configured value.
	OutlierEjectionType_FAILURE_PERCENTAGE_LOCAL_ORIGIN OutlierEjectionType = 6
)

// Enum value maps for OutlierEjectionType.
var (
	OutlierEjectionType_name = map[int32]string{
		0: "CONSECUTIVE_5XX",
		1: "CONSECUTIVE_GATEWAY_FAILURE",
		2: "SUCCESS_RATE",
		3: "CONSECUTIVE_LOCAL_ORIGIN_FAILURE",
		4: "SUCCESS_RATE_LOCAL_ORIGIN",
		5: "FAILURE_PERCENTAGE",
		6: "FAILURE_PERCENTAGE_LOCAL_ORIGIN",
	}
	OutlierEjectionType_value = map[string]int32{
		"CONSECUTIVE_5XX":                  0,
		"CONSECUTIVE_GATEWAY_FAILURE":      1,
		"SUCCESS_RATE":                     2,
		"CONSECUTIVE_LOCAL_ORIGIN_FAILURE": 3,
		"SUCCESS_RATE_LOCAL_ORIGIN":        4,
		"FAILURE_PERCENTAGE":               5,
		"FAILURE_PERCENTAGE_LOCAL_ORIGIN":  6,
	}
)

func (x OutlierEjectionType) Enum() *OutlierEjectionType {
	p := new(OutlierEjectionType)
	*p = x
	return p
}

func (x OutlierEjectionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OutlierEjectionType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_data_cluster_v3_outlier_detection_event_proto_enumTypes[0].Descriptor()
}

func (OutlierEjectionType) Type() protoreflect.EnumType {
	return &file_envoy_data_cluster_v3_outlier_detection_event_proto_enumTypes[0]
}

func (x OutlierEjectionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OutlierEjectionType.Descriptor instead.
func (OutlierEjectionType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_data_cluster_v3_outlier_detection_event_proto_rawDescGZIP(), []int{0}
}

// Represents possible action applied to upstream host
type Action int32

const (
	// In case host was excluded from service
	Action_EJECT Action = 0
	// In case host was brought back into service
	Action_UNEJECT Action = 1
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "EJECT",
		1: "UNEJECT",
	}
	Action_value = map[string]int32{
		"EJECT":   0,
		"UNEJECT": 1,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_data_cluster_v3_outlier_detection_event_proto_enumTypes[1].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_envoy_data_cluster_v3_outlier_detection_event_proto_enumTypes[1]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_envoy_data_cluster_v3_outlier_detection_event_proto_rawDescGZIP(), []int{1}
}

// [#next-free-field: 12]
type OutlierDetectionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// In case of eject represents type of ejection that took place.
	Type OutlierEjectionType `protobuf:"varint,1,opt,name=type,proto3,enum=envoy.data.cluster.v3.OutlierEjectionType" json:"type,omitempty"`
	// Timestamp for event.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The time in seconds since the last action (either an ejection or unejection) took place.
	SecsSinceLastAction *wrapperspb.UInt64Value `protobuf:"bytes,3,opt,name=secs_since_last_action,json=secsSinceLastAction,proto3" json:"secs_since_last_action,omitempty"`
	// The :ref:`cluster <envoy_v3_api_msg_config.cluster.v3.Cluster>` that owns the ejected host.
	ClusterName string `protobuf:"bytes,4,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// The URL of the ejected host. E.g., ``tcp://1.2.3.4:80``.
	UpstreamUrl string `protobuf:"bytes,5,opt,name=upstream_url,json=upstreamUrl,proto3" json:"upstream_url,omitempty"`
	// The action that took place.
	Action Action `protobuf:"varint,6,opt,name=action,proto3,enum=envoy.data.cluster.v3.Action" json:"action,omitempty"`
	// If ``action`` is ``eject``, specifies the number of times the host has been ejected (local to
	// that Envoy and gets reset if the host gets removed from the upstream cluster for any reason and
	// then re-added).
	NumEjections uint32 `protobuf:"varint,7,opt,name=num_ejections,json=numEjections,proto3" json:"num_ejections,omitempty"`
	// If ``action`` is ``eject``, specifies if the ejection was enforced. ``true`` means the host was
	// ejected. ``false`` means the event was logged but the host was not actually ejected.
	Enforced bool `protobuf:"varint,8,opt,name=enforced,proto3" json:"enforced,omitempty"`
	// Types that are assignable to Event:
	//	*OutlierDetectionEvent_EjectSuccessRateEvent
	//	*OutlierDetectionEvent_EjectConsecutiveEvent
	//	*OutlierDetectionEvent_EjectFailurePercentageEvent
	Event isOutlierDetectionEvent_Event `protobuf_oneof:"event"`
}

func (x *OutlierDetectionEvent) Reset() {
	*x = OutlierDetectionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_cluster_v3_outlier_detection_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutlierDetectionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutlierDetectionEvent) ProtoMessage() {}

func (x *OutlierDetectionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_cluster_v3_outlier_detection_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutlierDetectionEvent.ProtoReflect.Descriptor instead.
func (*OutlierDetectionEvent) Descriptor() ([]byte, []int) {
	return file_envoy_data_cluster_v3_outlier_detection_event_proto_rawDescGZIP(), []int{0}
}

func (x *OutlierDetectionEvent) GetType() OutlierEjectionType {
	if x != nil {
		return x.Type
	}
	return OutlierEjectionType_CONSECUTIVE_5XX
}

func (x *OutlierDetectionEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *OutlierDetectionEvent) GetSecsSinceLastAction() *wrapperspb.UInt64Value {
	if x != nil {
		return x.SecsSinceLastAction
	}
	return nil
}

func (x *OutlierDetectionEvent) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *OutlierDetectionEvent) GetUpstreamUrl() string {
	if x != nil {
		return x.UpstreamUrl
	}
	return ""
}

func (x *OutlierDetectionEvent) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_EJECT
}

func (x *OutlierDetectionEvent) GetNumEjections() uint32 {
	if x != nil {
		return x.NumEjections
	}
	return 0
}

func (x *OutlierDetectionEvent) GetEnforced() bool {
	if x != nil {
		return x.Enforced
	}
	return false
}

func (m *OutlierDetectionEvent) GetEvent() isOutlierDetectionEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *OutlierDetectionEvent) GetEjectSuccessRateEvent() *OutlierEjectSuccessRate {
	if x, ok := x.GetEvent().(*OutlierDetectionEvent_EjectSuccessRateEvent); ok {
		return x.EjectSuccessRateEvent
	}
	return nil
}

func (x *OutlierDetectionEvent) GetEjectConsecutiveEvent() *OutlierEjectConsecutive {
	if x, ok := x.GetEvent().(*OutlierDetectionEvent_EjectConsecutiveEvent); ok {
		return x.EjectConsecutiveEvent
	}
	return nil
}

func (x *OutlierDetectionEvent) GetEjectFailurePercentageEvent() *OutlierEjectFailurePercentage {
	if x, ok := x.GetEvent().(*OutlierDetectionEvent_EjectFailurePercentageEvent); ok {
		return x.EjectFailurePercentageEvent
	}
	return nil
}

type isOutlierDetectionEvent_Event interface {
	isOutlierDetectionEvent_Event()
}

type OutlierDetectionEvent_EjectSuccessRateEvent struct {
	EjectSuccessRateEvent *OutlierEjectSuccessRate `protobuf:"bytes,9,opt,name=eject_success_rate_event,json=ejectSuccessRateEvent,proto3,oneof"`
}

type OutlierDetectionEvent_EjectConsecutiveEvent struct {
	EjectConsecutiveEvent *OutlierEjectConsecutive `protobuf:"bytes,10,opt,name=eject_consecutive_event,json=ejectConsecutiveEvent,proto3,oneof"`
}

type OutlierDetectionEvent_EjectFailurePercentageEvent struct {
	EjectFailurePercentageEvent *OutlierEjectFailurePercentage `protobuf:"bytes,11,opt,name=eject_failure_percentage_event,json=ejectFailurePercentageEvent,proto3,oneof"`
}

func (*OutlierDetectionEvent_EjectSuccessRateEvent) isOutlierDetectionEvent_Event() {}

func (*OutlierDetectionEvent_EjectConsecutiveEvent) isOutlierDetectionEvent_Event() {}

func (*OutlierDetectionEvent_EjectFailurePercentageEvent) isOutlierDetectionEvent_Event() {}

type OutlierEjectSuccessRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Host’s success rate at the time of the ejection event on a 0-100 range.
	HostSuccessRate uint32 `protobuf:"varint,1,opt,name=host_success_rate,json=hostSuccessRate,proto3" json:"host_success_rate,omitempty"`
	// Average success rate of the hosts in the cluster at the time of the ejection event on a 0-100
	// range.
	ClusterAverageSuccessRate uint32 `protobuf:"varint,2,opt,name=cluster_average_success_rate,json=clusterAverageSuccessRate,proto3" json:"cluster_average_success_rate,omitempty"`
	// Success rate ejection threshold at the time of the ejection event.
	ClusterSuccessRateEjectionThreshold uint32 `protobuf:"varint,3,opt,name=cluster_success_rate_ejection_threshold,json=clusterSuccessRateEjectionThreshold,proto3" json:"cluster_success_rate_ejection_threshold,omitempty"`
}

func (x *OutlierEjectSuccessRate) Reset() {
	*x = OutlierEjectSuccessRate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_cluster_v3_outlier_detection_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutlierEjectSuccessRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutlierEjectSuccessRate) ProtoMessage() {}

func (x *OutlierEjectSuccessRate) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_cluster_v3_outlier_detection_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutlierEjectSuccessRate.ProtoReflect.Descriptor instead.
func (*OutlierEjectSuccessRate) Descriptor() ([]byte, []int) {
	return file_envoy_data_cluster_v3_outlier_detection_event_proto_rawDescGZIP(), []int{1}
}

func (x *OutlierEjectSuccessRate) GetHostSuccessRate() uint32 {
	if x != nil {
		return x.HostSuccessRate
	}
	return 0
}

func (x *OutlierEjectSuccessRate) GetClusterAverageSuccessRate() uint32 {
	if x != nil {
		return x.ClusterAverageSuccessRate
	}
	return 0
}

func (x *OutlierEjectSuccessRate) GetClusterSuccessRateEjectionThreshold() uint32 {
	if x != nil {
		return x.ClusterSuccessRateEjectionThreshold
	}
	return 0
}

type OutlierEjectConsecutive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OutlierEjectConsecutive) Reset() {
	*x = OutlierEjectConsecutive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_cluster_v3_outlier_detection_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutlierEjectConsecutive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutlierEjectConsecutive) ProtoMessage() {}

func (x *OutlierEjectConsecutive) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_cluster_v3_outlier_detection_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutlierEjectConsecutive.ProtoReflect.Descriptor instead.
func (*OutlierEjectConsecutive) Descriptor() ([]byte, []int) {
	return file_envoy_data_cluster_v3_outlier_detection_event_proto_rawDescGZIP(), []int{2}
}

type OutlierEjectFailurePercentage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Host's success rate at the time of the ejection event on a 0-100 range.
	HostSuccessRate uint32 `protobuf:"varint,1,opt,name=host_success_rate,json=hostSuccessRate,proto3" json:"host_success_rate,omitempty"`
}

func (x *OutlierEjectFailurePercentage) Reset() {
	*x = OutlierEjectFailurePercentage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_cluster_v3_outlier_detection_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutlierEjectFailurePercentage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutlierEjectFailurePercentage) ProtoMessage() {}

func (x *OutlierEjectFailurePercentage) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_cluster_v3_outlier_detection_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutlierEjectFailurePercentage.ProtoReflect.Descriptor instead.
func (*OutlierEjectFailurePercentage) Descriptor() ([]byte, []int) {
	return file_envoy_data_cluster_v3_outlier_detection_event_proto_rawDescGZIP(), []int{3}
}

func (x *OutlierEjectFailurePercentage) GetHostSuccessRate() uint32 {
	if x != nil {
		return x.HostSuccessRate
	}
	return 0
}

var File_envoy_data_cluster_v3_outlier_detection_event_proto protoreflect.FileDescriptor

var file_envoy_data_cluster_v3_outlier_detection_event_proto_rawDesc = []byte{
	0x0a, 0x33, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x6f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x5f,
	0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x06, 0x0a, 0x15, 0x4f, 0x75, 0x74,
	0x6c, 0x69, 0x65, 0x72, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x48, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72,
	0x45, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x51, 0x0a, 0x16, 0x73, 0x65, 0x63, 0x73, 0x5f, 0x73,
	0x69, 0x6e, 0x63, 0x65, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x13, 0x73, 0x65, 0x63, 0x73, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x4c,
	0x61, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x0c, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0c, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55, 0x72,
	0x6c, 0x12, 0x3f, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x5f, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x45, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x66, 0x6f, 0x72,
	0x63, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x6e, 0x66, 0x6f, 0x72,
	0x63, 0x65, 0x64, 0x12, 0x69, 0x0a, 0x18, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x75,
	0x74, 0x6c, 0x69, 0x65, 0x72, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x15, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x68,
	0x0a, 0x17, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x45,
	0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x48,
	0x00, 0x52, 0x15, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x7b, 0x0a, 0x1e, 0x65, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x34, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72,
	0x45, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x1b, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x46,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x3a, 0x37, 0x9a, 0xc5, 0x88, 0x1e, 0x32, 0x0a, 0x30, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x0c,
	0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0xb2, 0x02, 0x0a,
	0x17, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x11, 0x68, 0x6f, 0x73, 0x74,
	0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x64, 0x52, 0x0f, 0x68, 0x6f,
	0x73, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x12, 0x48, 0x0a,
	0x1c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x64, 0x52, 0x19, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x12, 0x5d, 0x0a, 0x27, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f,
	0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18,
	0x64, 0x52, 0x23, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x61, 0x74, 0x65, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x3a, 0x39, 0x9a, 0xc5, 0x88, 0x1e, 0x34, 0x0a, 0x32, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65,
	0x72, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74,
	0x65, 0x22, 0x54, 0x0a, 0x17, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x45, 0x6a, 0x65, 0x63,
	0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x3a, 0x39, 0x9a, 0xc5,
	0x88, 0x1e, 0x34, 0x0a, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x73,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x1d, 0x4f, 0x75, 0x74, 0x6c,
	0x69, 0x65, 0x72, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x11, 0x68, 0x6f, 0x73,
	0x74, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x64, 0x52, 0x0f, 0x68,
	0x6f, 0x73, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x3a, 0x3f,
	0x9a, 0xc5, 0x88, 0x1e, 0x3a, 0x0a, 0x38, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x2a,
	0xdf, 0x01, 0x0a, 0x13, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x45, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x53, 0x45,
	0x43, 0x55, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x35, 0x58, 0x58, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b,
	0x43, 0x4f, 0x4e, 0x53, 0x45, 0x43, 0x55, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x47, 0x41, 0x54, 0x45,
	0x57, 0x41, 0x59, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x01, 0x12, 0x10, 0x0a,
	0x0c, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12,
	0x24, 0x0a, 0x20, 0x43, 0x4f, 0x4e, 0x53, 0x45, 0x43, 0x55, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x4c,
	0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x55, 0x52, 0x45, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x5f, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x4f, 0x52, 0x49, 0x47,
	0x49, 0x4e, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x5f,
	0x50, 0x45, 0x52, 0x43, 0x45, 0x4e, 0x54, 0x41, 0x47, 0x45, 0x10, 0x05, 0x12, 0x23, 0x0a, 0x1f,
	0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x43, 0x45, 0x4e, 0x54, 0x41,
	0x47, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x10,
	0x06, 0x2a, 0x20, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x0a, 0x05, 0x45,
	0x4a, 0x45, 0x43, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x45, 0x4a, 0x45, 0x43,
	0x54, 0x10, 0x01, 0x42, 0x83, 0x02, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x33, 0x42, 0x1a, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50,
	0x01, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x69,
	0x6d, 0x73, 0x73, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f,
	0x76, 0x33, 0x3b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x76, 0x33, 0xa2, 0x02, 0x03, 0x45,
	0x44, 0x43, 0xaa, 0x02, 0x15, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x33, 0xca, 0x02, 0x15, 0x45, 0x6e, 0x76,
	0x6f, 0x79, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x5c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5c,
	0x56, 0x33, 0xe2, 0x02, 0x21, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x5c,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x33, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x3a, 0x3a,
	0x44, 0x61, 0x74, 0x61, 0x3a, 0x3a, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56,
	0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envoy_data_cluster_v3_outlier_detection_event_proto_rawDescOnce sync.Once
	file_envoy_data_cluster_v3_outlier_detection_event_proto_rawDescData = file_envoy_data_cluster_v3_outlier_detection_event_proto_rawDesc
)

func file_envoy_data_cluster_v3_outlier_detection_event_proto_rawDescGZIP() []byte {
	file_envoy_data_cluster_v3_outlier_detection_event_proto_rawDescOnce.Do(func() {
		file_envoy_data_cluster_v3_outlier_detection_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_data_cluster_v3_outlier_detection_event_proto_rawDescData)
	})
	return file_envoy_data_cluster_v3_outlier_detection_event_proto_rawDescData
}

var file_envoy_data_cluster_v3_outlier_detection_event_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_envoy_data_cluster_v3_outlier_detection_event_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_data_cluster_v3_outlier_detection_event_proto_goTypes = []interface{}{
	(OutlierEjectionType)(0),              // 0: envoy.data.cluster.v3.OutlierEjectionType
	(Action)(0),                           // 1: envoy.data.cluster.v3.Action
	(*OutlierDetectionEvent)(nil),         // 2: envoy.data.cluster.v3.OutlierDetectionEvent
	(*OutlierEjectSuccessRate)(nil),       // 3: envoy.data.cluster.v3.OutlierEjectSuccessRate
	(*OutlierEjectConsecutive)(nil),       // 4: envoy.data.cluster.v3.OutlierEjectConsecutive
	(*OutlierEjectFailurePercentage)(nil), // 5: envoy.data.cluster.v3.OutlierEjectFailurePercentage
	(*timestamppb.Timestamp)(nil),         // 6: google.protobuf.Timestamp
	(*wrapperspb.UInt64Value)(nil),        // 7: google.protobuf.UInt64Value
}
var file_envoy_data_cluster_v3_outlier_detection_event_proto_depIdxs = []int32{
	0, // 0: envoy.data.cluster.v3.OutlierDetectionEvent.type:type_name -> envoy.data.cluster.v3.OutlierEjectionType
	6, // 1: envoy.data.cluster.v3.OutlierDetectionEvent.timestamp:type_name -> google.protobuf.Timestamp
	7, // 2: envoy.data.cluster.v3.OutlierDetectionEvent.secs_since_last_action:type_name -> google.protobuf.UInt64Value
	1, // 3: envoy.data.cluster.v3.OutlierDetectionEvent.action:type_name -> envoy.data.cluster.v3.Action
	3, // 4: envoy.data.cluster.v3.OutlierDetectionEvent.eject_success_rate_event:type_name -> envoy.data.cluster.v3.OutlierEjectSuccessRate
	4, // 5: envoy.data.cluster.v3.OutlierDetectionEvent.eject_consecutive_event:type_name -> envoy.data.cluster.v3.OutlierEjectConsecutive
	5, // 6: envoy.data.cluster.v3.OutlierDetectionEvent.eject_failure_percentage_event:type_name -> envoy.data.cluster.v3.OutlierEjectFailurePercentage
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_envoy_data_cluster_v3_outlier_detection_event_proto_init() }
func file_envoy_data_cluster_v3_outlier_detection_event_proto_init() {
	if File_envoy_data_cluster_v3_outlier_detection_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_data_cluster_v3_outlier_detection_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutlierDetectionEvent); i {
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
		file_envoy_data_cluster_v3_outlier_detection_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutlierEjectSuccessRate); i {
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
		file_envoy_data_cluster_v3_outlier_detection_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutlierEjectConsecutive); i {
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
		file_envoy_data_cluster_v3_outlier_detection_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutlierEjectFailurePercentage); i {
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
	file_envoy_data_cluster_v3_outlier_detection_event_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OutlierDetectionEvent_EjectSuccessRateEvent)(nil),
		(*OutlierDetectionEvent_EjectConsecutiveEvent)(nil),
		(*OutlierDetectionEvent_EjectFailurePercentageEvent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_data_cluster_v3_outlier_detection_event_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_data_cluster_v3_outlier_detection_event_proto_goTypes,
		DependencyIndexes: file_envoy_data_cluster_v3_outlier_detection_event_proto_depIdxs,
		EnumInfos:         file_envoy_data_cluster_v3_outlier_detection_event_proto_enumTypes,
		MessageInfos:      file_envoy_data_cluster_v3_outlier_detection_event_proto_msgTypes,
	}.Build()
	File_envoy_data_cluster_v3_outlier_detection_event_proto = out.File
	file_envoy_data_cluster_v3_outlier_detection_event_proto_rawDesc = nil
	file_envoy_data_cluster_v3_outlier_detection_event_proto_goTypes = nil
	file_envoy_data_cluster_v3_outlier_detection_event_proto_depIdxs = nil
}
