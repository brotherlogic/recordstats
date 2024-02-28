// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: recordstats.proto

package proto

import (
	proto "github.com/brotherlogic/recordcollection/proto"
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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastListenTime int64                                       `protobuf:"varint,1,opt,name=last_listen_time,json=lastListenTime,proto3" json:"last_listen_time,omitempty"`
	CompleteSales  []*CompleteSale                             `protobuf:"bytes,2,rep,name=complete_sales,json=completeSales,proto3" json:"complete_sales,omitempty"`
	Auditions      []*Auditioned                               `protobuf:"bytes,3,rep,name=auditions,proto3" json:"auditions,omitempty"`
	Filed          map[int32]proto.ReleaseMetadata_FileSize    `protobuf:"bytes,4,rep,name=filed,proto3" json:"filed,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=recordcollection.ReleaseMetadata_FileSize"`
	FiledTime      map[int32]int64                             `protobuf:"bytes,5,rep,name=filed_time,json=filedTime,proto3" json:"filed_time,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	LbLastTime     map[int32]int64                             `protobuf:"bytes,6,rep,name=lb_last_time,json=lbLastTime,proto3" json:"lb_last_time,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	LbLastTimeHigh map[int32]int64                             `protobuf:"bytes,7,rep,name=lb_last_time_high,json=lbLastTimeHigh,proto3" json:"lb_last_time_high,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	LastListen     map[int32]int64                             `protobuf:"bytes,8,rep,name=last_listen,json=lastListen,proto3" json:"last_listen,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Folder         map[int32]int32                             `protobuf:"bytes,11,rep,name=folder,proto3" json:"folder,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Keeps          map[int32]proto.ReleaseMetadata_KeepState   `protobuf:"bytes,10,rep,name=keeps,proto3" json:"keeps,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=recordcollection.ReleaseMetadata_KeepState"`
	Values         []*Values                                   `protobuf:"bytes,9,rep,name=values,proto3" json:"values,omitempty"`
	Categories     map[int32]proto.ReleaseMetadata_Category    `protobuf:"bytes,12,rep,name=categories,proto3" json:"categories,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=recordcollection.ReleaseMetadata_Category"`
	Weights        map[int32]int32                             `protobuf:"bytes,13,rep,name=weights,proto3" json:"weights,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Sleeves        map[int32]proto.ReleaseMetadata_SleeveState `protobuf:"bytes,14,rep,name=sleeves,proto3" json:"sleeves,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=recordcollection.ReleaseMetadata_SleeveState"`
	IsDirty        map[int32]bool                              `protobuf:"bytes,15,rep,name=is_dirty,json=isDirty,proto3" json:"is_dirty,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recordstats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_recordstats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_recordstats_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetLastListenTime() int64 {
	if x != nil {
		return x.LastListenTime
	}
	return 0
}

func (x *Config) GetCompleteSales() []*CompleteSale {
	if x != nil {
		return x.CompleteSales
	}
	return nil
}

func (x *Config) GetAuditions() []*Auditioned {
	if x != nil {
		return x.Auditions
	}
	return nil
}

func (x *Config) GetFiled() map[int32]proto.ReleaseMetadata_FileSize {
	if x != nil {
		return x.Filed
	}
	return nil
}

func (x *Config) GetFiledTime() map[int32]int64 {
	if x != nil {
		return x.FiledTime
	}
	return nil
}

func (x *Config) GetLbLastTime() map[int32]int64 {
	if x != nil {
		return x.LbLastTime
	}
	return nil
}

func (x *Config) GetLbLastTimeHigh() map[int32]int64 {
	if x != nil {
		return x.LbLastTimeHigh
	}
	return nil
}

func (x *Config) GetLastListen() map[int32]int64 {
	if x != nil {
		return x.LastListen
	}
	return nil
}

func (x *Config) GetFolder() map[int32]int32 {
	if x != nil {
		return x.Folder
	}
	return nil
}

func (x *Config) GetKeeps() map[int32]proto.ReleaseMetadata_KeepState {
	if x != nil {
		return x.Keeps
	}
	return nil
}

func (x *Config) GetValues() []*Values {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *Config) GetCategories() map[int32]proto.ReleaseMetadata_Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *Config) GetWeights() map[int32]int32 {
	if x != nil {
		return x.Weights
	}
	return nil
}

func (x *Config) GetSleeves() map[int32]proto.ReleaseMetadata_SleeveState {
	if x != nil {
		return x.Sleeves
	}
	return nil
}

func (x *Config) GetIsDirty() map[int32]bool {
	if x != nil {
		return x.IsDirty
	}
	return nil
}

type Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category string            `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Filling  string            `protobuf:"bytes,2,opt,name=filling,proto3" json:"filling,omitempty"`
	Value    map[int32]float32 `protobuf:"bytes,3,rep,name=value,proto3" json:"value,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
}

func (x *Values) Reset() {
	*x = Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recordstats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Values) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Values) ProtoMessage() {}

func (x *Values) ProtoReflect() protoreflect.Message {
	mi := &file_recordstats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Values.ProtoReflect.Descriptor instead.
func (*Values) Descriptor() ([]byte, []int) {
	return file_recordstats_proto_rawDescGZIP(), []int{1}
}

func (x *Values) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Values) GetFilling() string {
	if x != nil {
		return x.Filling
	}
	return ""
}

func (x *Values) GetValue() map[int32]float32 {
	if x != nil {
		return x.Value
	}
	return nil
}

type Auditioned struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId   int32 `protobuf:"varint,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Valid        bool  `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	LastAudition int64 `protobuf:"varint,3,opt,name=last_audition,json=lastAudition,proto3" json:"last_audition,omitempty"`
	AudScore     int32 `protobuf:"varint,4,opt,name=audScore,proto3" json:"audScore,omitempty"`
}

func (x *Auditioned) Reset() {
	*x = Auditioned{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recordstats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auditioned) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auditioned) ProtoMessage() {}

func (x *Auditioned) ProtoReflect() protoreflect.Message {
	mi := &file_recordstats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auditioned.ProtoReflect.Descriptor instead.
func (*Auditioned) Descriptor() ([]byte, []int) {
	return file_recordstats_proto_rawDescGZIP(), []int{2}
}

func (x *Auditioned) GetInstanceId() int32 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

func (x *Auditioned) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *Auditioned) GetLastAudition() int64 {
	if x != nil {
		return x.LastAudition
	}
	return 0
}

func (x *Auditioned) GetAudScore() int32 {
	if x != nil {
		return x.AudScore
	}
	return 0
}

type CompleteSale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId int32 `protobuf:"varint,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	HasCost    bool  `protobuf:"varint,2,opt,name=has_cost,json=hasCost,proto3" json:"has_cost,omitempty"`
}

func (x *CompleteSale) Reset() {
	*x = CompleteSale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recordstats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteSale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteSale) ProtoMessage() {}

func (x *CompleteSale) ProtoReflect() protoreflect.Message {
	mi := &file_recordstats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteSale.ProtoReflect.Descriptor instead.
func (*CompleteSale) Descriptor() ([]byte, []int) {
	return file_recordstats_proto_rawDescGZIP(), []int{3}
}

func (x *CompleteSale) GetInstanceId() int32 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

func (x *CompleteSale) GetHasCost() bool {
	if x != nil {
		return x.HasCost
	}
	return false
}

type GetStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId int32 `protobuf:"varint,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
}

func (x *GetStatsRequest) Reset() {
	*x = GetStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recordstats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsRequest) ProtoMessage() {}

func (x *GetStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recordstats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsRequest.ProtoReflect.Descriptor instead.
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return file_recordstats_proto_rawDescGZIP(), []int{4}
}

func (x *GetStatsRequest) GetInstanceId() int32 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

type GetStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStatsResponse) Reset() {
	*x = GetStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recordstats_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsResponse) ProtoMessage() {}

func (x *GetStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recordstats_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsResponse.ProtoReflect.Descriptor instead.
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return file_recordstats_proto_rawDescGZIP(), []int{5}
}

var File_recordstats_proto protoreflect.FileDescriptor

var file_recordstats_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x1a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x0d, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6c, 0x61,
	0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0e,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x52,
	0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x35,
	0x0a, 0x09, 0x61, 0x75, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x64, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x41, 0x0a, 0x0a, 0x66,
	0x69, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x45,
	0x0a, 0x0c, 0x6c, 0x62, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x62, 0x4c, 0x61, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x6c, 0x62, 0x4c, 0x61, 0x73,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x52, 0x0a, 0x11, 0x6c, 0x62, 0x5f, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x68, 0x69, 0x67, 0x68, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x62, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x48, 0x69, 0x67, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x6c, 0x62, 0x4c, 0x61, 0x73,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x48, 0x69, 0x67, 0x68, 0x12, 0x44, 0x0a, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x12,
	0x37, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x05, 0x6b, 0x65, 0x65, 0x70,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x65, 0x65,
	0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6b, 0x65, 0x65, 0x70, 0x73, 0x12, 0x2b,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x3a, 0x0a, 0x07, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x07,
	0x73, 0x6c, 0x65, 0x65, 0x76, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x53, 0x6c, 0x65, 0x65, 0x76, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x73, 0x6c, 0x65, 0x65, 0x76, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x64,
	0x69, 0x72, 0x74, 0x79, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x49, 0x73, 0x44, 0x69, 0x72, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x69, 0x73,
	0x44, 0x69, 0x72, 0x74, 0x79, 0x1a, 0x64, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x64, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x40, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x46,
	0x69, 0x6c, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3d, 0x0a, 0x0f, 0x4c, 0x62, 0x4c,
	0x61, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x41, 0x0a, 0x13, 0x4c, 0x62, 0x4c, 0x61,
	0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x48, 0x69, 0x67, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3d, 0x0a, 0x0f, 0x4c,
	0x61, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x65, 0x0a, 0x0a, 0x4b, 0x65, 0x65, 0x70, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x41, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4b, 0x65, 0x65, 0x70, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x69, 0x0a, 0x0f,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x40, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2a, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x57, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x69, 0x0a, 0x0c, 0x53, 0x6c, 0x65, 0x65, 0x76, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x43, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x6c, 0x65, 0x65, 0x76, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a,
	0x0a, 0x0c, 0x49, 0x73, 0x44, 0x69, 0x72, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xae, 0x01, 0x0a, 0x06, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x34, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x1a, 0x38, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x84, 0x01, 0x0a, 0x0a,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x75, 0x64, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x22, 0x4a, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61,
	0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x43, 0x6f, 0x73, 0x74, 0x22, 0x32,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x5f, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x6c, 0x6f, 0x67,
	0x69, 0x63, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recordstats_proto_rawDescOnce sync.Once
	file_recordstats_proto_rawDescData = file_recordstats_proto_rawDesc
)

func file_recordstats_proto_rawDescGZIP() []byte {
	file_recordstats_proto_rawDescOnce.Do(func() {
		file_recordstats_proto_rawDescData = protoimpl.X.CompressGZIP(file_recordstats_proto_rawDescData)
	})
	return file_recordstats_proto_rawDescData
}

var file_recordstats_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file_recordstats_proto_goTypes = []interface{}{
	(*Config)(nil),                         // 0: recordstats.Config
	(*Values)(nil),                         // 1: recordstats.Values
	(*Auditioned)(nil),                     // 2: recordstats.Auditioned
	(*CompleteSale)(nil),                   // 3: recordstats.CompleteSale
	(*GetStatsRequest)(nil),                // 4: recordstats.GetStatsRequest
	(*GetStatsResponse)(nil),               // 5: recordstats.GetStatsResponse
	nil,                                    // 6: recordstats.Config.FiledEntry
	nil,                                    // 7: recordstats.Config.FiledTimeEntry
	nil,                                    // 8: recordstats.Config.LbLastTimeEntry
	nil,                                    // 9: recordstats.Config.LbLastTimeHighEntry
	nil,                                    // 10: recordstats.Config.LastListenEntry
	nil,                                    // 11: recordstats.Config.FolderEntry
	nil,                                    // 12: recordstats.Config.KeepsEntry
	nil,                                    // 13: recordstats.Config.CategoriesEntry
	nil,                                    // 14: recordstats.Config.WeightsEntry
	nil,                                    // 15: recordstats.Config.SleevesEntry
	nil,                                    // 16: recordstats.Config.IsDirtyEntry
	nil,                                    // 17: recordstats.Values.ValueEntry
	(proto.ReleaseMetadata_FileSize)(0),    // 18: recordcollection.ReleaseMetadata.FileSize
	(proto.ReleaseMetadata_KeepState)(0),   // 19: recordcollection.ReleaseMetadata.KeepState
	(proto.ReleaseMetadata_Category)(0),    // 20: recordcollection.ReleaseMetadata.Category
	(proto.ReleaseMetadata_SleeveState)(0), // 21: recordcollection.ReleaseMetadata.SleeveState
}
var file_recordstats_proto_depIdxs = []int32{
	3,  // 0: recordstats.Config.complete_sales:type_name -> recordstats.CompleteSale
	2,  // 1: recordstats.Config.auditions:type_name -> recordstats.Auditioned
	6,  // 2: recordstats.Config.filed:type_name -> recordstats.Config.FiledEntry
	7,  // 3: recordstats.Config.filed_time:type_name -> recordstats.Config.FiledTimeEntry
	8,  // 4: recordstats.Config.lb_last_time:type_name -> recordstats.Config.LbLastTimeEntry
	9,  // 5: recordstats.Config.lb_last_time_high:type_name -> recordstats.Config.LbLastTimeHighEntry
	10, // 6: recordstats.Config.last_listen:type_name -> recordstats.Config.LastListenEntry
	11, // 7: recordstats.Config.folder:type_name -> recordstats.Config.FolderEntry
	12, // 8: recordstats.Config.keeps:type_name -> recordstats.Config.KeepsEntry
	1,  // 9: recordstats.Config.values:type_name -> recordstats.Values
	13, // 10: recordstats.Config.categories:type_name -> recordstats.Config.CategoriesEntry
	14, // 11: recordstats.Config.weights:type_name -> recordstats.Config.WeightsEntry
	15, // 12: recordstats.Config.sleeves:type_name -> recordstats.Config.SleevesEntry
	16, // 13: recordstats.Config.is_dirty:type_name -> recordstats.Config.IsDirtyEntry
	17, // 14: recordstats.Values.value:type_name -> recordstats.Values.ValueEntry
	18, // 15: recordstats.Config.FiledEntry.value:type_name -> recordcollection.ReleaseMetadata.FileSize
	19, // 16: recordstats.Config.KeepsEntry.value:type_name -> recordcollection.ReleaseMetadata.KeepState
	20, // 17: recordstats.Config.CategoriesEntry.value:type_name -> recordcollection.ReleaseMetadata.Category
	21, // 18: recordstats.Config.SleevesEntry.value:type_name -> recordcollection.ReleaseMetadata.SleeveState
	4,  // 19: recordstats.RecordStatsService.GetStats:input_type -> recordstats.GetStatsRequest
	5,  // 20: recordstats.RecordStatsService.GetStats:output_type -> recordstats.GetStatsResponse
	20, // [20:21] is the sub-list for method output_type
	19, // [19:20] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_recordstats_proto_init() }
func file_recordstats_proto_init() {
	if File_recordstats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recordstats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_recordstats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Values); i {
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
		file_recordstats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auditioned); i {
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
		file_recordstats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteSale); i {
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
		file_recordstats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatsRequest); i {
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
		file_recordstats_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatsResponse); i {
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
			RawDescriptor: file_recordstats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recordstats_proto_goTypes,
		DependencyIndexes: file_recordstats_proto_depIdxs,
		MessageInfos:      file_recordstats_proto_msgTypes,
	}.Build()
	File_recordstats_proto = out.File
	file_recordstats_proto_rawDesc = nil
	file_recordstats_proto_goTypes = nil
	file_recordstats_proto_depIdxs = nil
}
