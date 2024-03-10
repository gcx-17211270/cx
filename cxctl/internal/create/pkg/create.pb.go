// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.3
// source: create.proto

package pkg

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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId int32     `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Mdata     *Metadata `protobuf:"bytes,2,opt,name=mdata,proto3" json:"mdata,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_create_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetRequestId() int32 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *CreateRequest) GetMdata() *Metadata {
	if x != nil {
		return x.Mdata
	}
	return nil
}

type CreateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *CreateReply) Reset() {
	*x = CreateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReply) ProtoMessage() {}

func (x *CreateReply) ProtoReflect() protoreflect.Message {
	mi := &file_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReply.ProtoReflect.Descriptor instead.
func (*CreateReply) Descriptor() ([]byte, []int) {
	return file_create_proto_rawDescGZIP(), []int{1}
}

func (x *CreateReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_create_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_create_proto_rawDescGZIP(), []int{2}
}

func (x *Metadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metadata) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Replicas int32     `protobuf:"varint,1,opt,name=replicas,proto3" json:"replicas,omitempty"`
	Template *Template `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *Spec) Reset() {
	*x = Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec) ProtoMessage() {}

func (x *Spec) ProtoReflect() protoreflect.Message {
	mi := &file_create_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spec.ProtoReflect.Descriptor instead.
func (*Spec) Descriptor() ([]byte, []int) {
	return file_create_proto_rawDescGZIP(), []int{3}
}

func (x *Spec) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *Spec) GetTemplate() *Template {
	if x != nil {
		return x.Template
	}
	return nil
}

type Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec     *TemplateSpec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Affinity string        `protobuf:"bytes,2,opt,name=affinity,proto3" json:"affinity,omitempty"` // 先只实现根据hostname亲和
}

func (x *Template) Reset() {
	*x = Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template) ProtoMessage() {}

func (x *Template) ProtoReflect() protoreflect.Message {
	mi := &file_create_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template.ProtoReflect.Descriptor instead.
func (*Template) Descriptor() ([]byte, []int) {
	return file_create_proto_rawDescGZIP(), []int{4}
}

func (x *Template) GetSpec() *TemplateSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Template) GetAffinity() string {
	if x != nil {
		return x.Affinity
	}
	return ""
}

type TemplateSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Containers *Containers `protobuf:"bytes,1,opt,name=containers,proto3" json:"containers,omitempty"`
}

func (x *TemplateSpec) Reset() {
	*x = TemplateSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateSpec) ProtoMessage() {}

func (x *TemplateSpec) ProtoReflect() protoreflect.Message {
	mi := &file_create_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateSpec.ProtoReflect.Descriptor instead.
func (*TemplateSpec) Descriptor() ([]byte, []int) {
	return file_create_proto_rawDescGZIP(), []int{5}
}

func (x *TemplateSpec) GetContainers() *Containers {
	if x != nil {
		return x.Containers
	}
	return nil
}

type Containers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C []*Container `protobuf:"bytes,1,rep,name=c,proto3" json:"c,omitempty"`
}

func (x *Containers) Reset() {
	*x = Containers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Containers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Containers) ProtoMessage() {}

func (x *Containers) ProtoReflect() protoreflect.Message {
	mi := &file_create_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Containers.ProtoReflect.Descriptor instead.
func (*Containers) Descriptor() ([]byte, []int) {
	return file_create_proto_rawDescGZIP(), []int{6}
}

func (x *Containers) GetC() []*Container {
	if x != nil {
		return x.C
	}
	return nil
}

type Container struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Image        string        `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	VolumeMounts *VolumeMounts `protobuf:"bytes,3,opt,name=volumeMounts,proto3" json:"volumeMounts,omitempty"`
	Resource     *Resource     `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *Container) Reset() {
	*x = Container{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_create_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container.ProtoReflect.Descriptor instead.
func (*Container) Descriptor() ([]byte, []int) {
	return file_create_proto_rawDescGZIP(), []int{7}
}

func (x *Container) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Container) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Container) GetVolumeMounts() *VolumeMounts {
	if x != nil {
		return x.VolumeMounts
	}
	return nil
}

func (x *Container) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

type VolumeMounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VolumeMount []*VolumeMount `protobuf:"bytes,1,rep,name=volumeMount,proto3" json:"volumeMount,omitempty"`
}

func (x *VolumeMounts) Reset() {
	*x = VolumeMounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeMounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeMounts) ProtoMessage() {}

func (x *VolumeMounts) ProtoReflect() protoreflect.Message {
	mi := &file_create_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeMounts.ProtoReflect.Descriptor instead.
func (*VolumeMounts) Descriptor() ([]byte, []int) {
	return file_create_proto_rawDescGZIP(), []int{8}
}

func (x *VolumeMounts) GetVolumeMount() []*VolumeMount {
	if x != nil {
		return x.VolumeMount
	}
	return nil
}

type VolumeMount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PathOutSide string `protobuf:"bytes,1,opt,name=pathOutSide,proto3" json:"pathOutSide,omitempty"`
	PathIn      string `protobuf:"bytes,2,opt,name=pathIn,proto3" json:"pathIn,omitempty"`
	Type        int32  `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"` // 1只读 2读写
}

func (x *VolumeMount) Reset() {
	*x = VolumeMount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeMount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeMount) ProtoMessage() {}

func (x *VolumeMount) ProtoReflect() protoreflect.Message {
	mi := &file_create_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeMount.ProtoReflect.Descriptor instead.
func (*VolumeMount) Descriptor() ([]byte, []int) {
	return file_create_proto_rawDescGZIP(), []int{9}
}

func (x *VolumeMount) GetPathOutSide() string {
	if x != nil {
		return x.PathOutSide
	}
	return ""
}

func (x *VolumeMount) GetPathIn() string {
	if x != nil {
		return x.PathIn
	}
	return ""
}

func (x *VolumeMount) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *Request `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Limit   *Limit   `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_create_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_create_proto_rawDescGZIP(), []int{10}
}

func (x *Resource) GetRequest() *Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Resource) GetLimit() *Limit {
	if x != nil {
		return x.Limit
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cm *CpuAndMem `protobuf:"bytes,1,opt,name=cm,proto3" json:"cm,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_create_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_create_proto_rawDescGZIP(), []int{11}
}

func (x *Request) GetCm() *CpuAndMem {
	if x != nil {
		return x.Cm
	}
	return nil
}

type Limit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cm *CpuAndMem `protobuf:"bytes,1,opt,name=cm,proto3" json:"cm,omitempty"`
}

func (x *Limit) Reset() {
	*x = Limit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Limit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Limit) ProtoMessage() {}

func (x *Limit) ProtoReflect() protoreflect.Message {
	mi := &file_create_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Limit.ProtoReflect.Descriptor instead.
func (*Limit) Descriptor() ([]byte, []int) {
	return file_create_proto_rawDescGZIP(), []int{12}
}

func (x *Limit) GetCm() *CpuAndMem {
	if x != nil {
		return x.Cm
	}
	return nil
}

type CpuAndMem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpu    uint32 `protobuf:"varint,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory uint32 `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty"`
}

func (x *CpuAndMem) Reset() {
	*x = CpuAndMem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuAndMem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuAndMem) ProtoMessage() {}

func (x *CpuAndMem) ProtoReflect() protoreflect.Message {
	mi := &file_create_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuAndMem.ProtoReflect.Descriptor instead.
func (*CpuAndMem) Descriptor() ([]byte, []int) {
	return file_create_proto_rawDescGZIP(), []int{13}
}

func (x *CpuAndMem) GetCpu() uint32 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *CpuAndMem) GetMemory() uint32 {
	if x != nil {
		return x.Memory
	}
	return 0
}

var File_create_proto protoreflect.FileDescriptor

var file_create_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x70, 0x6b, 0x67, 0x22, 0x52, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x05, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x05, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1d, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x3c, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x22, 0x4d, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x29, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x6b, 0x67,
	0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x22, 0x4d, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12,
	0x25, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x66, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x66, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x79, 0x22, 0x3f, 0x0a, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x2f, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x22, 0x2a, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x73, 0x12, 0x1c, 0x0a, 0x01, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x01, 0x63, 0x22,
	0x97, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x52, 0x0c, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x29,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x42, 0x0a, 0x0c, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x0b, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4d, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x0b, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5b, 0x0a,
	0x0b, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x61, 0x74, 0x68, 0x4f, 0x75, 0x74, 0x53, 0x69, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x61, 0x74, 0x68, 0x4f, 0x75, 0x74, 0x53, 0x69, 0x64, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x61, 0x74, 0x68, 0x49, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x61, 0x74, 0x68, 0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x54, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x29, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x02, 0x63,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x43, 0x70,
	0x75, 0x41, 0x6e, 0x64, 0x4d, 0x65, 0x6d, 0x52, 0x02, 0x63, 0x6d, 0x22, 0x27, 0x0a, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x1e, 0x0a, 0x02, 0x63, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x43, 0x70, 0x75, 0x41, 0x6e, 0x64, 0x4d, 0x65, 0x6d,
	0x52, 0x02, 0x63, 0x6d, 0x22, 0x35, 0x0a, 0x09, 0x43, 0x70, 0x75, 0x41, 0x6e, 0x64, 0x4d, 0x65,
	0x6d, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x63, 0x70, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x32, 0x3a, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x12, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x54, 0x0a, 0x1a, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x67, 0x63, 0x78, 0x2d, 0x31, 0x37, 0x32, 0x31, 0x31, 0x32, 0x37, 0x30, 0x2e, 0x63, 0x78, 0x2e,
	0x63, 0x78, 0x63, 0x74, 0x6c, 0x42, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x63, 0x78, 0x2d, 0x31, 0x37, 0x32, 0x31, 0x31, 0x32, 0x37, 0x30, 0x2f, 0x63, 0x78,
	0x2f, 0x63, 0x78, 0x63, 0x74, 0x6c, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_create_proto_rawDescOnce sync.Once
	file_create_proto_rawDescData = file_create_proto_rawDesc
)

func file_create_proto_rawDescGZIP() []byte {
	file_create_proto_rawDescOnce.Do(func() {
		file_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_create_proto_rawDescData)
	})
	return file_create_proto_rawDescData
}

var file_create_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_create_proto_goTypes = []interface{}{
	(*CreateRequest)(nil), // 0: pkg.CreateRequest
	(*CreateReply)(nil),   // 1: pkg.CreateReply
	(*Metadata)(nil),      // 2: pkg.Metadata
	(*Spec)(nil),          // 3: pkg.Spec
	(*Template)(nil),      // 4: pkg.Template
	(*TemplateSpec)(nil),  // 5: pkg.TemplateSpec
	(*Containers)(nil),    // 6: pkg.Containers
	(*Container)(nil),     // 7: pkg.Container
	(*VolumeMounts)(nil),  // 8: pkg.VolumeMounts
	(*VolumeMount)(nil),   // 9: pkg.VolumeMount
	(*Resource)(nil),      // 10: pkg.Resource
	(*Request)(nil),       // 11: pkg.Request
	(*Limit)(nil),         // 12: pkg.Limit
	(*CpuAndMem)(nil),     // 13: pkg.CpuAndMem
}
var file_create_proto_depIdxs = []int32{
	2,  // 0: pkg.CreateRequest.mdata:type_name -> pkg.Metadata
	4,  // 1: pkg.Spec.template:type_name -> pkg.Template
	5,  // 2: pkg.Template.spec:type_name -> pkg.TemplateSpec
	6,  // 3: pkg.TemplateSpec.containers:type_name -> pkg.Containers
	7,  // 4: pkg.Containers.c:type_name -> pkg.Container
	8,  // 5: pkg.Container.volumeMounts:type_name -> pkg.VolumeMounts
	10, // 6: pkg.Container.resource:type_name -> pkg.Resource
	9,  // 7: pkg.VolumeMounts.volumeMount:type_name -> pkg.VolumeMount
	11, // 8: pkg.Resource.request:type_name -> pkg.Request
	12, // 9: pkg.Resource.limit:type_name -> pkg.Limit
	13, // 10: pkg.Request.cm:type_name -> pkg.CpuAndMem
	13, // 11: pkg.Limit.cm:type_name -> pkg.CpuAndMem
	0,  // 12: pkg.Create.Create:input_type -> pkg.CreateRequest
	1,  // 13: pkg.Create.Create:output_type -> pkg.CreateReply
	13, // [13:14] is the sub-list for method output_type
	12, // [12:13] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_create_proto_init() }
func file_create_proto_init() {
	if File_create_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReply); i {
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
		file_create_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_create_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spec); i {
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
		file_create_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Template); i {
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
		file_create_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateSpec); i {
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
		file_create_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Containers); i {
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
		file_create_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Container); i {
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
		file_create_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeMounts); i {
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
		file_create_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeMount); i {
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
		file_create_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_create_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_create_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Limit); i {
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
		file_create_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpuAndMem); i {
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
			RawDescriptor: file_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_create_proto_goTypes,
		DependencyIndexes: file_create_proto_depIdxs,
		MessageInfos:      file_create_proto_msgTypes,
	}.Build()
	File_create_proto = out.File
	file_create_proto_rawDesc = nil
	file_create_proto_goTypes = nil
	file_create_proto_depIdxs = nil
}
