// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.2
// source: pb/bank.proto

package pb

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

type DepositRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent                string                                  `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Tenant                string                                  `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Environment           string                                  `protobuf:"bytes,3,opt,name=environment,proto3" json:"environment,omitempty"`
	Clusters              []*DepositRequest_ClusterWithNamespaces `protobuf:"bytes,4,rep,name=clusters,proto3" json:"clusters,omitempty"`
	ClusterWithNamespaces *DepositRequest_ClusterWithNamespaces   `protobuf:"bytes,5,opt,name=cluster_with_namespaces,json=clusterWithNamespaces,proto3" json:"cluster_with_namespaces,omitempty"`
}

func (x *DepositRequest) Reset() {
	*x = DepositRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_bank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositRequest) ProtoMessage() {}

func (x *DepositRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_bank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositRequest.ProtoReflect.Descriptor instead.
func (*DepositRequest) Descriptor() ([]byte, []int) {
	return file_pb_bank_proto_rawDescGZIP(), []int{0}
}

func (x *DepositRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *DepositRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *DepositRequest) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *DepositRequest) GetClusters() []*DepositRequest_ClusterWithNamespaces {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *DepositRequest) GetClusterWithNamespaces() *DepositRequest_ClusterWithNamespaces {
	if x != nil {
		return x.ClusterWithNamespaces
	}
	return nil
}

type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_bank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_pb_bank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_pb_bank_proto_rawDescGZIP(), []int{1}
}

type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_bank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_pb_bank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_pb_bank_proto_rawDescGZIP(), []int{2}
}

type Namespace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Namespace) Reset() {
	*x = Namespace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_bank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Namespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Namespace) ProtoMessage() {}

func (x *Namespace) ProtoReflect() protoreflect.Message {
	mi := &file_pb_bank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Namespace.ProtoReflect.Descriptor instead.
func (*Namespace) Descriptor() ([]byte, []int) {
	return file_pb_bank_proto_rawDescGZIP(), []int{3}
}

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_bank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_pb_bank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_pb_bank_proto_rawDescGZIP(), []int{4}
}

type DepositReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DepositReply) Reset() {
	*x = DepositReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_bank_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositReply) ProtoMessage() {}

func (x *DepositReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_bank_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositReply.ProtoReflect.Descriptor instead.
func (*DepositReply) Descriptor() ([]byte, []int) {
	return file_pb_bank_proto_rawDescGZIP(), []int{5}
}

func (x *DepositReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DepositRequest_ClusterWithNamespaces struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster    *Cluster                                   `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespaces []*DepositRequest_NamespaceWithDeployments `protobuf:"bytes,2,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
}

func (x *DepositRequest_ClusterWithNamespaces) Reset() {
	*x = DepositRequest_ClusterWithNamespaces{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_bank_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositRequest_ClusterWithNamespaces) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositRequest_ClusterWithNamespaces) ProtoMessage() {}

func (x *DepositRequest_ClusterWithNamespaces) ProtoReflect() protoreflect.Message {
	mi := &file_pb_bank_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositRequest_ClusterWithNamespaces.ProtoReflect.Descriptor instead.
func (*DepositRequest_ClusterWithNamespaces) Descriptor() ([]byte, []int) {
	return file_pb_bank_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DepositRequest_ClusterWithNamespaces) GetCluster() *Cluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *DepositRequest_ClusterWithNamespaces) GetNamespaces() []*DepositRequest_NamespaceWithDeployments {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

type DepositRequest_NamespaceWithDeployments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace   *Namespace                                `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Deployments []*DepositRequest_DeploymentWithEndpoints `protobuf:"bytes,2,rep,name=deployments,proto3" json:"deployments,omitempty"`
}

func (x *DepositRequest_NamespaceWithDeployments) Reset() {
	*x = DepositRequest_NamespaceWithDeployments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_bank_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositRequest_NamespaceWithDeployments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositRequest_NamespaceWithDeployments) ProtoMessage() {}

func (x *DepositRequest_NamespaceWithDeployments) ProtoReflect() protoreflect.Message {
	mi := &file_pb_bank_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositRequest_NamespaceWithDeployments.ProtoReflect.Descriptor instead.
func (*DepositRequest_NamespaceWithDeployments) Descriptor() ([]byte, []int) {
	return file_pb_bank_proto_rawDescGZIP(), []int{0, 1}
}

func (x *DepositRequest_NamespaceWithDeployments) GetNamespace() *Namespace {
	if x != nil {
		return x.Namespace
	}
	return nil
}

func (x *DepositRequest_NamespaceWithDeployments) GetDeployments() []*DepositRequest_DeploymentWithEndpoints {
	if x != nil {
		return x.Deployments
	}
	return nil
}

type DepositRequest_DeploymentWithEndpoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deployment *Deployment `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	Endpoints  []*Endpoint `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *DepositRequest_DeploymentWithEndpoints) Reset() {
	*x = DepositRequest_DeploymentWithEndpoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_bank_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositRequest_DeploymentWithEndpoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositRequest_DeploymentWithEndpoints) ProtoMessage() {}

func (x *DepositRequest_DeploymentWithEndpoints) ProtoReflect() protoreflect.Message {
	mi := &file_pb_bank_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositRequest_DeploymentWithEndpoints.ProtoReflect.Descriptor instead.
func (*DepositRequest_DeploymentWithEndpoints) Descriptor() ([]byte, []int) {
	return file_pb_bank_proto_rawDescGZIP(), []int{0, 2}
}

func (x *DepositRequest_DeploymentWithEndpoints) GetDeployment() *Deployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

func (x *DepositRequest_DeploymentWithEndpoints) GetEndpoints() []*Endpoint {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

var File_pb_bank_proto protoreflect.FileDescriptor

var file_pb_bank_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0xa7, 0x05, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x62, 0x2e,
	0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x60,
	0x0a, 0x17, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x15, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x1a, 0x8b, 0x01, 0x0a, 0x15, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x4b, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x1a, 0x95,
	0x01, 0x0a, 0x18, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x57, 0x69, 0x74, 0x68,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x75, 0x0a, 0x17, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x12, 0x2e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x2a, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x0c, 0x0a,
	0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x0a, 0x0a, 0x08, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x0b, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x22, 0x09, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22,
	0x1e, 0x0a, 0x0c, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32,
	0x37, 0x0a, 0x04, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x2f, 0x0a, 0x07, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_bank_proto_rawDescOnce sync.Once
	file_pb_bank_proto_rawDescData = file_pb_bank_proto_rawDesc
)

func file_pb_bank_proto_rawDescGZIP() []byte {
	file_pb_bank_proto_rawDescOnce.Do(func() {
		file_pb_bank_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_bank_proto_rawDescData)
	})
	return file_pb_bank_proto_rawDescData
}

var file_pb_bank_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pb_bank_proto_goTypes = []interface{}{
	(*DepositRequest)(nil),                          // 0: pb.DepositRequest
	(*Deployment)(nil),                              // 1: pb.Deployment
	(*Endpoint)(nil),                                // 2: pb.Endpoint
	(*Namespace)(nil),                               // 3: pb.Namespace
	(*Cluster)(nil),                                 // 4: pb.Cluster
	(*DepositReply)(nil),                            // 5: pb.DepositReply
	(*DepositRequest_ClusterWithNamespaces)(nil),    // 6: pb.DepositRequest.ClusterWithNamespaces
	(*DepositRequest_NamespaceWithDeployments)(nil), // 7: pb.DepositRequest.NamespaceWithDeployments
	(*DepositRequest_DeploymentWithEndpoints)(nil),  // 8: pb.DepositRequest.DeploymentWithEndpoints
}
var file_pb_bank_proto_depIdxs = []int32{
	6, // 0: pb.DepositRequest.clusters:type_name -> pb.DepositRequest.ClusterWithNamespaces
	6, // 1: pb.DepositRequest.cluster_with_namespaces:type_name -> pb.DepositRequest.ClusterWithNamespaces
	4, // 2: pb.DepositRequest.ClusterWithNamespaces.cluster:type_name -> pb.Cluster
	7, // 3: pb.DepositRequest.ClusterWithNamespaces.namespaces:type_name -> pb.DepositRequest.NamespaceWithDeployments
	3, // 4: pb.DepositRequest.NamespaceWithDeployments.namespace:type_name -> pb.Namespace
	8, // 5: pb.DepositRequest.NamespaceWithDeployments.deployments:type_name -> pb.DepositRequest.DeploymentWithEndpoints
	1, // 6: pb.DepositRequest.DeploymentWithEndpoints.deployment:type_name -> pb.Deployment
	2, // 7: pb.DepositRequest.DeploymentWithEndpoints.endpoints:type_name -> pb.Endpoint
	0, // 8: pb.Bank.Deposit:input_type -> pb.DepositRequest
	5, // 9: pb.Bank.Deposit:output_type -> pb.DepositReply
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_pb_bank_proto_init() }
func file_pb_bank_proto_init() {
	if File_pb_bank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_bank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositRequest); i {
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
		file_pb_bank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment); i {
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
		file_pb_bank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint); i {
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
		file_pb_bank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Namespace); i {
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
		file_pb_bank_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
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
		file_pb_bank_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositReply); i {
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
		file_pb_bank_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositRequest_ClusterWithNamespaces); i {
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
		file_pb_bank_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositRequest_NamespaceWithDeployments); i {
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
		file_pb_bank_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositRequest_DeploymentWithEndpoints); i {
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
			RawDescriptor: file_pb_bank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_bank_proto_goTypes,
		DependencyIndexes: file_pb_bank_proto_depIdxs,
		MessageInfos:      file_pb_bank_proto_msgTypes,
	}.Build()
	File_pb_bank_proto = out.File
	file_pb_bank_proto_rawDesc = nil
	file_pb_bank_proto_goTypes = nil
	file_pb_bank_proto_depIdxs = nil
}
