// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: api/app/v1alpha1/message.proto

package v1alpha1

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

type MetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string            `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Labels    map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MetaData) Reset() {
	*x = MetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_v1alpha1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaData) ProtoMessage() {}

func (x *MetaData) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_v1alpha1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaData.ProtoReflect.Descriptor instead.
func (*MetaData) Descriptor() ([]byte, []int) {
	return file_api_app_v1alpha1_message_proto_rawDescGZIP(), []int{0}
}

func (x *MetaData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetaData) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *MetaData) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type AppV1Alpha1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion string    `protobuf:"bytes,1,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Kind       string    `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	MetaData   *MetaData `protobuf:"bytes,3,opt,name=metaData,proto3" json:"metaData,omitempty"`
	Spec       *App      `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *AppV1Alpha1) Reset() {
	*x = AppV1Alpha1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_v1alpha1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppV1Alpha1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppV1Alpha1) ProtoMessage() {}

func (x *AppV1Alpha1) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_v1alpha1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppV1Alpha1.ProtoReflect.Descriptor instead.
func (*AppV1Alpha1) Descriptor() ([]byte, []int) {
	return file_api_app_v1alpha1_message_proto_rawDescGZIP(), []int{1}
}

func (x *AppV1Alpha1) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *AppV1Alpha1) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *AppV1Alpha1) GetMetaData() *MetaData {
	if x != nil {
		return x.MetaData
	}
	return nil
}

func (x *AppV1Alpha1) GetSpec() *App {
	if x != nil {
		return x.Spec
	}
	return nil
}

type Apps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Apps []*App `protobuf:"bytes,1,rep,name=apps,proto3" json:"apps,omitempty"`
}

func (x *Apps) Reset() {
	*x = Apps{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_v1alpha1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Apps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Apps) ProtoMessage() {}

func (x *Apps) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_v1alpha1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Apps.ProtoReflect.Descriptor instead.
func (*Apps) Descriptor() ([]byte, []int) {
	return file_api_app_v1alpha1_message_proto_rawDescGZIP(), []int{2}
}

func (x *Apps) GetApps() []*App {
	if x != nil {
		return x.Apps
	}
	return nil
}

type ClusterID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterID int32 `protobuf:"varint,1,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
}

func (x *ClusterID) Reset() {
	*x = ClusterID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_v1alpha1_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterID) ProtoMessage() {}

func (x *ClusterID) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_v1alpha1_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterID.ProtoReflect.Descriptor instead.
func (*ClusterID) Descriptor() ([]byte, []int) {
	return file_api_app_v1alpha1_message_proto_rawDescGZIP(), []int{3}
}

func (x *ClusterID) GetClusterID() int32 {
	if x != nil {
		return x.ClusterID
	}
	return 0
}

type AppID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterID int32 `protobuf:"varint,1,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
	AppID     int32 `protobuf:"varint,2,opt,name=appID,proto3" json:"appID,omitempty"`
}

func (x *AppID) Reset() {
	*x = AppID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_v1alpha1_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppID) ProtoMessage() {}

func (x *AppID) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_v1alpha1_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppID.ProtoReflect.Descriptor instead.
func (*AppID) Descriptor() ([]byte, []int) {
	return file_api_app_v1alpha1_message_proto_rawDescGZIP(), []int{4}
}

func (x *AppID) GetClusterID() int32 {
	if x != nil {
		return x.ClusterID
	}
	return 0
}

func (x *AppID) GetAppID() int32 {
	if x != nil {
		return x.AppID
	}
	return 0
}

type App struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RepoName  string `protobuf:"bytes,3,opt,name=repoName,proto3" json:"repoName,omitempty"`
	RepoUrl   string `protobuf:"bytes,4,opt,name=repoUrl,proto3" json:"repoUrl,omitempty"`
	ChartName string `protobuf:"bytes,5,opt,name=chartName,proto3" json:"chartName,omitempty"`
	Version   string `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	Config    string `protobuf:"bytes,7,opt,name=config,proto3" json:"config,omitempty"`
	Secret    string `protobuf:"bytes,8,opt,name=secret,proto3" json:"secret,omitempty"`
	Namespace string `protobuf:"bytes,9,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ClusterID int32  `protobuf:"varint,10,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
}

func (x *App) Reset() {
	*x = App{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_app_v1alpha1_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *App) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App) ProtoMessage() {}

func (x *App) ProtoReflect() protoreflect.Message {
	mi := &file_api_app_v1alpha1_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App.ProtoReflect.Descriptor instead.
func (*App) Descriptor() ([]byte, []int) {
	return file_api_app_v1alpha1_message_proto_rawDescGZIP(), []int{5}
}

func (x *App) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *App) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *App) GetRepoName() string {
	if x != nil {
		return x.RepoName
	}
	return ""
}

func (x *App) GetRepoUrl() string {
	if x != nil {
		return x.RepoUrl
	}
	return ""
}

func (x *App) GetChartName() string {
	if x != nil {
		return x.ChartName
	}
	return ""
}

func (x *App) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *App) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *App) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *App) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *App) GetClusterID() int32 {
	if x != nil {
		return x.ClusterID
	}
	return 0
}

var File_api_app_v1alpha1_message_proto protoreflect.FileDescriptor

var file_api_app_v1alpha1_message_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0xb3,
	0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x3a, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x9c, 0x01, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x22, 0x2d, 0x0a, 0x04, 0x41, 0x70, 0x70, 0x73, 0x12, 0x25, 0x0a, 0x04, 0x61,
	0x70, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x04, 0x61, 0x70,
	0x70, 0x73, 0x22, 0x29, 0x0a, 0x09, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x22, 0x3b, 0x0a,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x22, 0x83, 0x02, 0x0a, 0x03, 0x41,
	0x70, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x68, 0x61, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x68, 0x61, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44,
	0x42, 0x1b, 0x5a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x3b, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_app_v1alpha1_message_proto_rawDescOnce sync.Once
	file_api_app_v1alpha1_message_proto_rawDescData = file_api_app_v1alpha1_message_proto_rawDesc
)

func file_api_app_v1alpha1_message_proto_rawDescGZIP() []byte {
	file_api_app_v1alpha1_message_proto_rawDescOnce.Do(func() {
		file_api_app_v1alpha1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_app_v1alpha1_message_proto_rawDescData)
	})
	return file_api_app_v1alpha1_message_proto_rawDescData
}

var file_api_app_v1alpha1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_app_v1alpha1_message_proto_goTypes = []interface{}{
	(*MetaData)(nil),    // 0: app.v1alpha1.MetaData
	(*AppV1Alpha1)(nil), // 1: app.v1alpha1.AppV1alpha1
	(*Apps)(nil),        // 2: app.v1alpha1.Apps
	(*ClusterID)(nil),   // 3: app.v1alpha1.ClusterID
	(*AppID)(nil),       // 4: app.v1alpha1.AppID
	(*App)(nil),         // 5: app.v1alpha1.App
	nil,                 // 6: app.v1alpha1.MetaData.LabelsEntry
}
var file_api_app_v1alpha1_message_proto_depIdxs = []int32{
	6, // 0: app.v1alpha1.MetaData.labels:type_name -> app.v1alpha1.MetaData.LabelsEntry
	0, // 1: app.v1alpha1.AppV1alpha1.metaData:type_name -> app.v1alpha1.MetaData
	5, // 2: app.v1alpha1.AppV1alpha1.spec:type_name -> app.v1alpha1.App
	5, // 3: app.v1alpha1.Apps.apps:type_name -> app.v1alpha1.App
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_app_v1alpha1_message_proto_init() }
func file_api_app_v1alpha1_message_proto_init() {
	if File_api_app_v1alpha1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_app_v1alpha1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaData); i {
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
		file_api_app_v1alpha1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppV1Alpha1); i {
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
		file_api_app_v1alpha1_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Apps); i {
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
		file_api_app_v1alpha1_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterID); i {
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
		file_api_app_v1alpha1_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppID); i {
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
		file_api_app_v1alpha1_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*App); i {
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
			RawDescriptor: file_api_app_v1alpha1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_app_v1alpha1_message_proto_goTypes,
		DependencyIndexes: file_api_app_v1alpha1_message_proto_depIdxs,
		MessageInfos:      file_api_app_v1alpha1_message_proto_msgTypes,
	}.Build()
	File_api_app_v1alpha1_message_proto = out.File
	file_api_app_v1alpha1_message_proto_rawDesc = nil
	file_api_app_v1alpha1_message_proto_goTypes = nil
	file_api_app_v1alpha1_message_proto_depIdxs = nil
}
