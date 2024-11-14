// protoc --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: pod/v1/pod.proto

package pod

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

type EciPod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=CustomerId,proto3" json:"CustomerId,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	VNodeId    string `protobuf:"bytes,3,opt,name=VNodeId,proto3" json:"VNodeId,omitempty"`
	Pod        *Pod   `protobuf:"bytes,4,opt,name=Pod,proto3" json:"Pod,omitempty"`
}

func (x *EciPod) Reset() {
	*x = EciPod{}
	mi := &file_pod_v1_pod_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EciPod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EciPod) ProtoMessage() {}

func (x *EciPod) ProtoReflect() protoreflect.Message {
	mi := &file_pod_v1_pod_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EciPod.ProtoReflect.Descriptor instead.
func (*EciPod) Descriptor() ([]byte, []int) {
	return file_pod_v1_pod_proto_rawDescGZIP(), []int{0}
}

func (x *EciPod) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *EciPod) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *EciPod) GetVNodeId() string {
	if x != nil {
		return x.VNodeId
	}
	return ""
}

func (x *EciPod) GetPod() *Pod {
	if x != nil {
		return x.Pod
	}
	return nil
}

type Pod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string       `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Id      string       `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	Status  string       `protobuf:"bytes,3,opt,name=Status,proto3" json:"Status,omitempty"`
	Network *NetworkInfo `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	Cpu     int32        `protobuf:"varint,5,opt,name=Cpu,proto3" json:"Cpu,omitempty"`
	Ram     int32        `protobuf:"varint,6,opt,name=Ram,proto3" json:"Ram,omitempty"`
	Gpu     int32        `protobuf:"varint,7,opt,name=Gpu,proto3" json:"Gpu,omitempty"`
	CpuType string       `protobuf:"bytes,8,opt,name=CpuType,proto3" json:"CpuType,omitempty"`
	GpuType string       `protobuf:"bytes,9,opt,name=GpuType,proto3" json:"GpuType,omitempty"`
	Boot    bool         `protobuf:"varint,10,opt,name=Boot,proto3" json:"Boot,omitempty"`
}

func (x *Pod) Reset() {
	*x = Pod{}
	mi := &file_pod_v1_pod_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pod) ProtoMessage() {}

func (x *Pod) ProtoReflect() protoreflect.Message {
	mi := &file_pod_v1_pod_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pod.ProtoReflect.Descriptor instead.
func (*Pod) Descriptor() ([]byte, []int) {
	return file_pod_v1_pod_proto_rawDescGZIP(), []int{1}
}

func (x *Pod) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pod) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pod) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Pod) GetNetwork() *NetworkInfo {
	if x != nil {
		return x.Network
	}
	return nil
}

func (x *Pod) GetCpu() int32 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *Pod) GetRam() int32 {
	if x != nil {
		return x.Ram
	}
	return 0
}

func (x *Pod) GetGpu() int32 {
	if x != nil {
		return x.Gpu
	}
	return 0
}

func (x *Pod) GetCpuType() string {
	if x != nil {
		return x.CpuType
	}
	return ""
}

func (x *Pod) GetGpuType() string {
	if x != nil {
		return x.GpuType
	}
	return ""
}

func (x *Pod) GetBoot() bool {
	if x != nil {
		return x.Boot
	}
	return false
}

type NetworkInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip      string `protobuf:"bytes,1,opt,name=Ip,proto3" json:"Ip,omitempty"`
	Vlan    string `protobuf:"bytes,2,opt,name=Vlan,proto3" json:"Vlan,omitempty"`
	Gateway string `protobuf:"bytes,3,opt,name=Gateway,proto3" json:"Gateway,omitempty"`
}

func (x *NetworkInfo) Reset() {
	*x = NetworkInfo{}
	mi := &file_pod_v1_pod_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInfo) ProtoMessage() {}

func (x *NetworkInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pod_v1_pod_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInfo.ProtoReflect.Descriptor instead.
func (*NetworkInfo) Descriptor() ([]byte, []int) {
	return file_pod_v1_pod_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *NetworkInfo) GetVlan() string {
	if x != nil {
		return x.Vlan
	}
	return ""
}

func (x *NetworkInfo) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

type SavePodResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"` // v: required
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`   // v: required
}

func (x *SavePodResp) Reset() {
	*x = SavePodResp{}
	mi := &file_pod_v1_pod_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SavePodResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePodResp) ProtoMessage() {}

func (x *SavePodResp) ProtoReflect() protoreflect.Message {
	mi := &file_pod_v1_pod_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePodResp.ProtoReflect.Descriptor instead.
func (*SavePodResp) Descriptor() ([]byte, []int) {
	return file_pod_v1_pod_proto_rawDescGZIP(), []int{3}
}

func (x *SavePodResp) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SavePodResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_pod_v1_pod_proto protoreflect.FileDescriptor

var file_pod_v1_pod_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x6f, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x70, 0x6f, 0x64, 0x22, 0x76, 0x0a, 0x06, 0x45, 0x63, 0x69, 0x50, 0x6f,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x03, 0x50, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x50, 0x6f, 0x64, 0x52, 0x03, 0x50, 0x6f, 0x64, 0x22,
	0xeb, 0x01, 0x0a, 0x03, 0x50, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x10, 0x0a, 0x03, 0x43, 0x70, 0x75, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x43, 0x70,
	0x75, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x61, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x52, 0x61, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x47, 0x70, 0x75, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x47, 0x70, 0x75, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x70, 0x75, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x70, 0x75, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x47, 0x70, 0x75, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x47, 0x70, 0x75, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x6f,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x42, 0x6f, 0x6f, 0x74, 0x22, 0x4b, 0x0a,
	0x0b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x56, 0x6c, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x56, 0x6c, 0x61, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0x33, 0x0a, 0x0b, 0x53, 0x61,
	0x76, 0x65, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x32,
	0x38, 0x0a, 0x0a, 0x45, 0x63, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a,
	0x07, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6f, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x45,
	0x63, 0x69, 0x50, 0x6f, 0x64, 0x1a, 0x10, 0x2e, 0x70, 0x6f, 0x64, 0x2e, 0x53, 0x61, 0x76, 0x65,
	0x50, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x70,
	0x6f, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pod_v1_pod_proto_rawDescOnce sync.Once
	file_pod_v1_pod_proto_rawDescData = file_pod_v1_pod_proto_rawDesc
)

func file_pod_v1_pod_proto_rawDescGZIP() []byte {
	file_pod_v1_pod_proto_rawDescOnce.Do(func() {
		file_pod_v1_pod_proto_rawDescData = protoimpl.X.CompressGZIP(file_pod_v1_pod_proto_rawDescData)
	})
	return file_pod_v1_pod_proto_rawDescData
}

var file_pod_v1_pod_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pod_v1_pod_proto_goTypes = []any{
	(*EciPod)(nil),      // 0: pod.EciPod
	(*Pod)(nil),         // 1: pod.Pod
	(*NetworkInfo)(nil), // 2: pod.NetworkInfo
	(*SavePodResp)(nil), // 3: pod.SavePodResp
}
var file_pod_v1_pod_proto_depIdxs = []int32{
	1, // 0: pod.EciPod.Pod:type_name -> pod.Pod
	2, // 1: pod.Pod.network:type_name -> pod.NetworkInfo
	0, // 2: pod.EciService.SavePod:input_type -> pod.EciPod
	3, // 3: pod.EciService.SavePod:output_type -> pod.SavePodResp
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pod_v1_pod_proto_init() }
func file_pod_v1_pod_proto_init() {
	if File_pod_v1_pod_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pod_v1_pod_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pod_v1_pod_proto_goTypes,
		DependencyIndexes: file_pod_v1_pod_proto_depIdxs,
		MessageInfos:      file_pod_v1_pod_proto_msgTypes,
	}.Build()
	File_pod_v1_pod_proto = out.File
	file_pod_v1_pod_proto_rawDesc = nil
	file_pod_v1_pod_proto_goTypes = nil
	file_pod_v1_pod_proto_depIdxs = nil
}
