// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: host.proto

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

type SetHostnameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *SetHostnameRequest) Reset() {
	*x = SetHostnameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetHostnameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetHostnameRequest) ProtoMessage() {}

func (x *SetHostnameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_host_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetHostnameRequest.ProtoReflect.Descriptor instead.
func (*SetHostnameRequest) Descriptor() ([]byte, []int) {
	return file_host_proto_rawDescGZIP(), []int{0}
}

func (x *SetHostnameRequest) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

type SetHostnameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetHostnameResponse) Reset() {
	*x = SetHostnameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetHostnameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetHostnameResponse) ProtoMessage() {}

func (x *SetHostnameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_host_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetHostnameResponse.ProtoReflect.Descriptor instead.
func (*SetHostnameResponse) Descriptor() ([]byte, []int) {
	return file_host_proto_rawDescGZIP(), []int{1}
}

func (x *SetHostnameResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetDnsServersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDnsServersRequest) Reset() {
	*x = GetDnsServersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDnsServersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDnsServersRequest) ProtoMessage() {}

func (x *GetDnsServersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_host_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDnsServersRequest.ProtoReflect.Descriptor instead.
func (*GetDnsServersRequest) Descriptor() ([]byte, []int) {
	return file_host_proto_rawDescGZIP(), []int{2}
}

type GetDnsServersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DnsServers []string `protobuf:"bytes,1,rep,name=dns_servers,json=dnsServers,proto3" json:"dns_servers,omitempty"`
}

func (x *GetDnsServersResponse) Reset() {
	*x = GetDnsServersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDnsServersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDnsServersResponse) ProtoMessage() {}

func (x *GetDnsServersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_host_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDnsServersResponse.ProtoReflect.Descriptor instead.
func (*GetDnsServersResponse) Descriptor() ([]byte, []int) {
	return file_host_proto_rawDescGZIP(), []int{3}
}

func (x *GetDnsServersResponse) GetDnsServers() []string {
	if x != nil {
		return x.DnsServers
	}
	return nil
}

type AddDnsServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DnsServer string `protobuf:"bytes,1,opt,name=dns_server,json=dnsServer,proto3" json:"dns_server,omitempty"`
}

func (x *AddDnsServerRequest) Reset() {
	*x = AddDnsServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDnsServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDnsServerRequest) ProtoMessage() {}

func (x *AddDnsServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_host_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDnsServerRequest.ProtoReflect.Descriptor instead.
func (*AddDnsServerRequest) Descriptor() ([]byte, []int) {
	return file_host_proto_rawDescGZIP(), []int{4}
}

func (x *AddDnsServerRequest) GetDnsServer() string {
	if x != nil {
		return x.DnsServer
	}
	return ""
}

type AddDnsServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AddDnsServerResponse) Reset() {
	*x = AddDnsServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDnsServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDnsServerResponse) ProtoMessage() {}

func (x *AddDnsServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_host_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDnsServerResponse.ProtoReflect.Descriptor instead.
func (*AddDnsServerResponse) Descriptor() ([]byte, []int) {
	return file_host_proto_rawDescGZIP(), []int{5}
}

func (x *AddDnsServerResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type RemoveDnsServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DnsServer string `protobuf:"bytes,1,opt,name=dns_server,json=dnsServer,proto3" json:"dns_server,omitempty"`
}

func (x *RemoveDnsServerRequest) Reset() {
	*x = RemoveDnsServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveDnsServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveDnsServerRequest) ProtoMessage() {}

func (x *RemoveDnsServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_host_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveDnsServerRequest.ProtoReflect.Descriptor instead.
func (*RemoveDnsServerRequest) Descriptor() ([]byte, []int) {
	return file_host_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveDnsServerRequest) GetDnsServer() string {
	if x != nil {
		return x.DnsServer
	}
	return ""
}

type RemoveDnsServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RemoveDnsServerResponse) Reset() {
	*x = RemoveDnsServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveDnsServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveDnsServerResponse) ProtoMessage() {}

func (x *RemoveDnsServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_host_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveDnsServerResponse.ProtoReflect.Descriptor instead.
func (*RemoveDnsServerResponse) Descriptor() ([]byte, []int) {
	return file_host_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveDnsServerResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_host_proto protoreflect.FileDescriptor

var file_host_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x22, 0x30, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6e, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x22, 0x34, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x44, 0x6e, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x2e, 0x0a, 0x14, 0x41,
	0x64, 0x64, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x37, 0x0a, 0x16, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x6e, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x6e,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xb5, 0x02, 0x0a, 0x0e, 0x48, 0x6f, 0x73, 0x74,
	0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x53, 0x65,
	0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x2e, 0x68, 0x6f, 0x73, 0x74,
	0x2e, 0x53, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x48, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12,
	0x1a, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x6f,
	0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x44,
	0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e,
	0x41, 0x64, 0x64, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x6e,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4e, 0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x1c, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x6e,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0d, 0x5a, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x6f, 0x76, 0x6f, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_host_proto_rawDescOnce sync.Once
	file_host_proto_rawDescData = file_host_proto_rawDesc
)

func file_host_proto_rawDescGZIP() []byte {
	file_host_proto_rawDescOnce.Do(func() {
		file_host_proto_rawDescData = protoimpl.X.CompressGZIP(file_host_proto_rawDescData)
	})
	return file_host_proto_rawDescData
}

var file_host_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_host_proto_goTypes = []any{
	(*SetHostnameRequest)(nil),      // 0: host.SetHostnameRequest
	(*SetHostnameResponse)(nil),     // 1: host.SetHostnameResponse
	(*GetDnsServersRequest)(nil),    // 2: host.GetDnsServersRequest
	(*GetDnsServersResponse)(nil),   // 3: host.GetDnsServersResponse
	(*AddDnsServerRequest)(nil),     // 4: host.AddDnsServerRequest
	(*AddDnsServerResponse)(nil),    // 5: host.AddDnsServerResponse
	(*RemoveDnsServerRequest)(nil),  // 6: host.RemoveDnsServerRequest
	(*RemoveDnsServerResponse)(nil), // 7: host.RemoveDnsServerResponse
}
var file_host_proto_depIdxs = []int32{
	0, // 0: host.HostDnsService.SetHostname:input_type -> host.SetHostnameRequest
	2, // 1: host.HostDnsService.GetDnsServers:input_type -> host.GetDnsServersRequest
	4, // 2: host.HostDnsService.AddDnsServer:input_type -> host.AddDnsServerRequest
	6, // 3: host.HostDnsService.RemoveDnsServer:input_type -> host.RemoveDnsServerRequest
	1, // 4: host.HostDnsService.SetHostname:output_type -> host.SetHostnameResponse
	3, // 5: host.HostDnsService.GetDnsServers:output_type -> host.GetDnsServersResponse
	5, // 6: host.HostDnsService.AddDnsServer:output_type -> host.AddDnsServerResponse
	7, // 7: host.HostDnsService.RemoveDnsServer:output_type -> host.RemoveDnsServerResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_host_proto_init() }
func file_host_proto_init() {
	if File_host_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_host_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SetHostnameRequest); i {
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
		file_host_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SetHostnameResponse); i {
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
		file_host_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetDnsServersRequest); i {
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
		file_host_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetDnsServersResponse); i {
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
		file_host_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AddDnsServerRequest); i {
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
		file_host_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AddDnsServerResponse); i {
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
		file_host_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveDnsServerRequest); i {
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
		file_host_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveDnsServerResponse); i {
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
			RawDescriptor: file_host_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_host_proto_goTypes,
		DependencyIndexes: file_host_proto_depIdxs,
		MessageInfos:      file_host_proto_msgTypes,
	}.Build()
	File_host_proto = out.File
	file_host_proto_rawDesc = nil
	file_host_proto_goTypes = nil
	file_host_proto_depIdxs = nil
}
