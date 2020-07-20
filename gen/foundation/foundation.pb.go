// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: foundation/v1/foundation.proto

package foundationv1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Connection
type FederateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *FederateRequest) Reset() {
	*x = FederateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foundation_v1_foundation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederateRequest) ProtoMessage() {}

func (x *FederateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foundation_v1_foundation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederateRequest.ProtoReflect.Descriptor instead.
func (*FederateRequest) Descriptor() ([]byte, []int) {
	return file_foundation_v1_foundation_proto_rawDescGZIP(), []int{0}
}

func (x *FederateRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type FederateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Nonce string `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *FederateReply) Reset() {
	*x = FederateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foundation_v1_foundation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederateReply) ProtoMessage() {}

func (x *FederateReply) ProtoReflect() protoreflect.Message {
	mi := &file_foundation_v1_foundation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederateReply.ProtoReflect.Descriptor instead.
func (*FederateReply) Descriptor() ([]byte, []int) {
	return file_foundation_v1_foundation_proto_rawDescGZIP(), []int{1}
}

func (x *FederateReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *FederateReply) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

type KeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KeyRequest) Reset() {
	*x = KeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foundation_v1_foundation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRequest) ProtoMessage() {}

func (x *KeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foundation_v1_foundation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRequest.ProtoReflect.Descriptor instead.
func (*KeyRequest) Descriptor() ([]byte, []int) {
	return file_foundation_v1_foundation_proto_rawDescGZIP(), []int{2}
}

type KeyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *KeyReply) Reset() {
	*x = KeyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foundation_v1_foundation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyReply) ProtoMessage() {}

func (x *KeyReply) ProtoReflect() protoreflect.Message {
	mi := &file_foundation_v1_foundation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyReply.ProtoReflect.Descriptor instead.
func (*KeyReply) Descriptor() ([]byte, []int) {
	return file_foundation_v1_foundation_proto_rawDescGZIP(), []int{3}
}

func (x *KeyReply) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Login:
	//	*LoginRequest_Federated_
	//	*LoginRequest_Local_
	Login isLoginRequest_Login `protobuf_oneof:"login"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foundation_v1_foundation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foundation_v1_foundation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_foundation_v1_foundation_proto_rawDescGZIP(), []int{4}
}

func (m *LoginRequest) GetLogin() isLoginRequest_Login {
	if m != nil {
		return m.Login
	}
	return nil
}

func (x *LoginRequest) GetFederated() *LoginRequest_Federated {
	if x, ok := x.GetLogin().(*LoginRequest_Federated_); ok {
		return x.Federated
	}
	return nil
}

func (x *LoginRequest) GetLocal() *LoginRequest_Local {
	if x, ok := x.GetLogin().(*LoginRequest_Local_); ok {
		return x.Local
	}
	return nil
}

type isLoginRequest_Login interface {
	isLoginRequest_Login()
}

type LoginRequest_Federated_ struct {
	Federated *LoginRequest_Federated `protobuf:"bytes,1,opt,name=federated,proto3,oneof"`
}

type LoginRequest_Local_ struct {
	Local *LoginRequest_Local `protobuf:"bytes,2,opt,name=local,proto3,oneof"`
}

func (*LoginRequest_Federated_) isLoginRequest_Login() {}

func (*LoginRequest_Local_) isLoginRequest_Login() {}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password []byte `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foundation_v1_foundation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foundation_v1_foundation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_foundation_v1_foundation_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterRequest) GetPassword() []byte {
	if x != nil {
		return x.Password
	}
	return nil
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SessionToken string `protobuf:"bytes,2,opt,name=session_token,json=sessionToken,proto3" json:"session_token,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foundation_v1_foundation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_foundation_v1_foundation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_foundation_v1_foundation_proto_rawDescGZIP(), []int{6}
}

func (x *Session) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Session) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

type LoginRequest_Federated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken string `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	Domain    string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *LoginRequest_Federated) Reset() {
	*x = LoginRequest_Federated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foundation_v1_foundation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest_Federated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest_Federated) ProtoMessage() {}

func (x *LoginRequest_Federated) ProtoReflect() protoreflect.Message {
	mi := &file_foundation_v1_foundation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest_Federated.ProtoReflect.Descriptor instead.
func (*LoginRequest_Federated) Descriptor() ([]byte, []int) {
	return file_foundation_v1_foundation_proto_rawDescGZIP(), []int{4, 0}
}

func (x *LoginRequest_Federated) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

func (x *LoginRequest_Federated) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type LoginRequest_Local struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest_Local) Reset() {
	*x = LoginRequest_Local{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foundation_v1_foundation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest_Local) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest_Local) ProtoMessage() {}

func (x *LoginRequest_Local) ProtoReflect() protoreflect.Message {
	mi := &file_foundation_v1_foundation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest_Local.ProtoReflect.Descriptor instead.
func (*LoginRequest_Local) Descriptor() ([]byte, []int) {
	return file_foundation_v1_foundation_proto_rawDescGZIP(), []int{4, 1}
}

func (x *LoginRequest_Local) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest_Local) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_foundation_v1_foundation_proto protoreflect.FileDescriptor

var file_foundation_v1_foundation_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0f, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x22, 0x3b, 0x0a, 0x0d, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x0c,
	0x0a, 0x0a, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1c, 0x0a, 0x08,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xaa, 0x02, 0x0a, 0x0c, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x09, 0x66,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00,
	0x52, 0x09, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x12, 0x42, 0x0a, 0x05, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x1a,
	0x42, 0x0a, 0x09, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x1a, 0x39, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x07,
	0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x5f, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4b, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xe2, 0x02, 0x0a, 0x11, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x08, 0x46,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4b, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x22,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x4e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x24, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x19, 0x5a, 0x17, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foundation_v1_foundation_proto_rawDescOnce sync.Once
	file_foundation_v1_foundation_proto_rawDescData = file_foundation_v1_foundation_proto_rawDesc
)

func file_foundation_v1_foundation_proto_rawDescGZIP() []byte {
	file_foundation_v1_foundation_proto_rawDescOnce.Do(func() {
		file_foundation_v1_foundation_proto_rawDescData = protoimpl.X.CompressGZIP(file_foundation_v1_foundation_proto_rawDescData)
	})
	return file_foundation_v1_foundation_proto_rawDescData
}

var file_foundation_v1_foundation_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_foundation_v1_foundation_proto_goTypes = []interface{}{
	(*FederateRequest)(nil),        // 0: protocol.foundation.v1.FederateRequest
	(*FederateReply)(nil),          // 1: protocol.foundation.v1.FederateReply
	(*KeyRequest)(nil),             // 2: protocol.foundation.v1.KeyRequest
	(*KeyReply)(nil),               // 3: protocol.foundation.v1.KeyReply
	(*LoginRequest)(nil),           // 4: protocol.foundation.v1.LoginRequest
	(*RegisterRequest)(nil),        // 5: protocol.foundation.v1.RegisterRequest
	(*Session)(nil),                // 6: protocol.foundation.v1.Session
	(*LoginRequest_Federated)(nil), // 7: protocol.foundation.v1.LoginRequest.Federated
	(*LoginRequest_Local)(nil),     // 8: protocol.foundation.v1.LoginRequest.Local
}
var file_foundation_v1_foundation_proto_depIdxs = []int32{
	7, // 0: protocol.foundation.v1.LoginRequest.federated:type_name -> protocol.foundation.v1.LoginRequest.Federated
	8, // 1: protocol.foundation.v1.LoginRequest.local:type_name -> protocol.foundation.v1.LoginRequest.Local
	0, // 2: protocol.foundation.v1.FoundationService.Federate:input_type -> protocol.foundation.v1.FederateRequest
	2, // 3: protocol.foundation.v1.FoundationService.Key:input_type -> protocol.foundation.v1.KeyRequest
	4, // 4: protocol.foundation.v1.FoundationService.Login:input_type -> protocol.foundation.v1.LoginRequest
	5, // 5: protocol.foundation.v1.FoundationService.Register:input_type -> protocol.foundation.v1.RegisterRequest
	1, // 6: protocol.foundation.v1.FoundationService.Federate:output_type -> protocol.foundation.v1.FederateReply
	3, // 7: protocol.foundation.v1.FoundationService.Key:output_type -> protocol.foundation.v1.KeyReply
	6, // 8: protocol.foundation.v1.FoundationService.Login:output_type -> protocol.foundation.v1.Session
	6, // 9: protocol.foundation.v1.FoundationService.Register:output_type -> protocol.foundation.v1.Session
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_foundation_v1_foundation_proto_init() }
func file_foundation_v1_foundation_proto_init() {
	if File_foundation_v1_foundation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foundation_v1_foundation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederateRequest); i {
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
		file_foundation_v1_foundation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederateReply); i {
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
		file_foundation_v1_foundation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRequest); i {
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
		file_foundation_v1_foundation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyReply); i {
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
		file_foundation_v1_foundation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_foundation_v1_foundation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_foundation_v1_foundation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_foundation_v1_foundation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest_Federated); i {
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
		file_foundation_v1_foundation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest_Local); i {
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
	file_foundation_v1_foundation_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*LoginRequest_Federated_)(nil),
		(*LoginRequest_Local_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_foundation_v1_foundation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_foundation_v1_foundation_proto_goTypes,
		DependencyIndexes: file_foundation_v1_foundation_proto_depIdxs,
		MessageInfos:      file_foundation_v1_foundation_proto_msgTypes,
	}.Build()
	File_foundation_v1_foundation_proto = out.File
	file_foundation_v1_foundation_proto_rawDesc = nil
	file_foundation_v1_foundation_proto_goTypes = nil
	file_foundation_v1_foundation_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FoundationServiceClient is the client API for FoundationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FoundationServiceClient interface {
	Federate(ctx context.Context, in *FederateRequest, opts ...grpc.CallOption) (*FederateReply, error)
	Key(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*KeyReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Session, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Session, error)
}

type foundationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFoundationServiceClient(cc grpc.ClientConnInterface) FoundationServiceClient {
	return &foundationServiceClient{cc}
}

func (c *foundationServiceClient) Federate(ctx context.Context, in *FederateRequest, opts ...grpc.CallOption) (*FederateReply, error) {
	out := new(FederateReply)
	err := c.cc.Invoke(ctx, "/protocol.foundation.v1.FoundationService/Federate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) Key(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*KeyReply, error) {
	out := new(KeyReply)
	err := c.cc.Invoke(ctx, "/protocol.foundation.v1.FoundationService/Key", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/protocol.foundation.v1.FoundationService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/protocol.foundation.v1.FoundationService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FoundationServiceServer is the server API for FoundationService service.
type FoundationServiceServer interface {
	Federate(context.Context, *FederateRequest) (*FederateReply, error)
	Key(context.Context, *KeyRequest) (*KeyReply, error)
	Login(context.Context, *LoginRequest) (*Session, error)
	Register(context.Context, *RegisterRequest) (*Session, error)
}

// UnimplementedFoundationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFoundationServiceServer struct {
}

func (*UnimplementedFoundationServiceServer) Federate(context.Context, *FederateRequest) (*FederateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Federate not implemented")
}
func (*UnimplementedFoundationServiceServer) Key(context.Context, *KeyRequest) (*KeyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Key not implemented")
}
func (*UnimplementedFoundationServiceServer) Login(context.Context, *LoginRequest) (*Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedFoundationServiceServer) Register(context.Context, *RegisterRequest) (*Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterFoundationServiceServer(s *grpc.Server, srv FoundationServiceServer) {
	s.RegisterService(&_FoundationService_serviceDesc, srv)
}

func _FoundationService_Federate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FederateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).Federate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.foundation.v1.FoundationService/Federate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).Federate(ctx, req.(*FederateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_Key_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).Key(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.foundation.v1.FoundationService/Key",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).Key(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.foundation.v1.FoundationService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.foundation.v1.FoundationService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FoundationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.foundation.v1.FoundationService",
	HandlerType: (*FoundationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Federate",
			Handler:    _FoundationService_Federate_Handler,
		},
		{
			MethodName: "Key",
			Handler:    _FoundationService_Key_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _FoundationService_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _FoundationService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "foundation/v1/foundation.proto",
}
