// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: peasant/v1/account_service.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peasant_v1_account_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_peasant_v1_account_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_peasant_v1_account_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateAccountRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateAccountRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peasant_v1_account_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_peasant_v1_account_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_peasant_v1_account_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountResponse) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type GetAccountByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *GetAccountByIdRequest) Reset() {
	*x = GetAccountByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peasant_v1_account_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountByIdRequest) ProtoMessage() {}

func (x *GetAccountByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_peasant_v1_account_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountByIdRequest.ProtoReflect.Descriptor instead.
func (*GetAccountByIdRequest) Descriptor() ([]byte, []int) {
	return file_peasant_v1_account_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAccountByIdRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type GetAccountByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *GetAccountByIdResponse) Reset() {
	*x = GetAccountByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peasant_v1_account_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountByIdResponse) ProtoMessage() {}

func (x *GetAccountByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_peasant_v1_account_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountByIdResponse.ProtoReflect.Descriptor instead.
func (*GetAccountByIdResponse) Descriptor() ([]byte, []int) {
	return file_peasant_v1_account_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccountByIdResponse) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type DeleteAccountByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *DeleteAccountByIdRequest) Reset() {
	*x = DeleteAccountByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peasant_v1_account_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccountByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountByIdRequest) ProtoMessage() {}

func (x *DeleteAccountByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_peasant_v1_account_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountByIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteAccountByIdRequest) Descriptor() ([]byte, []int) {
	return file_peasant_v1_account_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteAccountByIdRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type DeleteAccountByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAccountByIdResponse) Reset() {
	*x = DeleteAccountByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peasant_v1_account_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccountByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountByIdResponse) ProtoMessage() {}

func (x *DeleteAccountByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_peasant_v1_account_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountByIdResponse.ProtoReflect.Descriptor instead.
func (*DeleteAccountByIdResponse) Descriptor() ([]byte, []int) {
	return file_peasant_v1_account_service_proto_rawDescGZIP(), []int{5}
}

var File_peasant_v1_account_service_proto protoreflect.FileDescriptor

var file_peasant_v1_account_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x70, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9b, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05,
	0x18, 0xc0, 0x02, 0x60, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x3a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e,
	0xfa, 0x42, 0x1b, 0x72, 0x19, 0x32, 0x17, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30,
	0x2d, 0x39, 0x5d, 0x5b, 0x5c, 0x77, 0x5d, 0x7b, 0x33, 0x2c, 0x32, 0x34, 0x7d, 0x24, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72,
	0x04, 0x10, 0x08, 0x18, 0x40, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x46, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x65, 0x61, 0x73,
	0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x40, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x27, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x43, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27,
	0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa1, 0x02, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x65, 0x61, 0x73, 0x61,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x65, 0x61,
	0x73, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x21, 0x2e, 0x70, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x24, 0x2e, 0x70, 0x65,
	0x61, 0x73, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x70, 0x65, 0x61, 0x73, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x73, 0x6f, 0x6d, 0x61, 0x64, 0x2f, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x73, 0x75, 0x63, 0x6b, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x65, 0x61,
	0x73, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peasant_v1_account_service_proto_rawDescOnce sync.Once
	file_peasant_v1_account_service_proto_rawDescData = file_peasant_v1_account_service_proto_rawDesc
)

func file_peasant_v1_account_service_proto_rawDescGZIP() []byte {
	file_peasant_v1_account_service_proto_rawDescOnce.Do(func() {
		file_peasant_v1_account_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_peasant_v1_account_service_proto_rawDescData)
	})
	return file_peasant_v1_account_service_proto_rawDescData
}

var file_peasant_v1_account_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_peasant_v1_account_service_proto_goTypes = []interface{}{
	(*CreateAccountRequest)(nil),      // 0: peasant.v1.CreateAccountRequest
	(*CreateAccountResponse)(nil),     // 1: peasant.v1.CreateAccountResponse
	(*GetAccountByIdRequest)(nil),     // 2: peasant.v1.GetAccountByIdRequest
	(*GetAccountByIdResponse)(nil),    // 3: peasant.v1.GetAccountByIdResponse
	(*DeleteAccountByIdRequest)(nil),  // 4: peasant.v1.DeleteAccountByIdRequest
	(*DeleteAccountByIdResponse)(nil), // 5: peasant.v1.DeleteAccountByIdResponse
	(*Account)(nil),                   // 6: peasant.v1.Account
}
var file_peasant_v1_account_service_proto_depIdxs = []int32{
	6, // 0: peasant.v1.CreateAccountResponse.account:type_name -> peasant.v1.Account
	6, // 1: peasant.v1.GetAccountByIdResponse.account:type_name -> peasant.v1.Account
	0, // 2: peasant.v1.AccountService.CreateAccount:input_type -> peasant.v1.CreateAccountRequest
	2, // 3: peasant.v1.AccountService.GetAccountById:input_type -> peasant.v1.GetAccountByIdRequest
	4, // 4: peasant.v1.AccountService.DeleteAccountById:input_type -> peasant.v1.DeleteAccountByIdRequest
	1, // 5: peasant.v1.AccountService.CreateAccount:output_type -> peasant.v1.CreateAccountResponse
	3, // 6: peasant.v1.AccountService.GetAccountById:output_type -> peasant.v1.GetAccountByIdResponse
	5, // 7: peasant.v1.AccountService.DeleteAccountById:output_type -> peasant.v1.DeleteAccountByIdResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_peasant_v1_account_service_proto_init() }
func file_peasant_v1_account_service_proto_init() {
	if File_peasant_v1_account_service_proto != nil {
		return
	}
	file_peasant_v1_account_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_peasant_v1_account_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
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
		file_peasant_v1_account_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountResponse); i {
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
		file_peasant_v1_account_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountByIdRequest); i {
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
		file_peasant_v1_account_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountByIdResponse); i {
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
		file_peasant_v1_account_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccountByIdRequest); i {
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
		file_peasant_v1_account_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccountByIdResponse); i {
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
			RawDescriptor: file_peasant_v1_account_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_peasant_v1_account_service_proto_goTypes,
		DependencyIndexes: file_peasant_v1_account_service_proto_depIdxs,
		MessageInfos:      file_peasant_v1_account_service_proto_msgTypes,
	}.Build()
	File_peasant_v1_account_service_proto = out.File
	file_peasant_v1_account_service_proto_rawDesc = nil
	file_peasant_v1_account_service_proto_goTypes = nil
	file_peasant_v1_account_service_proto_depIdxs = nil
}