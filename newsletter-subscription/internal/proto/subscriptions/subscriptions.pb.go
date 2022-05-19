// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: subscriptions.proto

package newsletter

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Genre struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Genre) Reset() {
	*x = Genre{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Genre) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genre) ProtoMessage() {}

func (x *Genre) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genre.ProtoReflect.Descriptor instead.
func (*Genre) Descriptor() ([]byte, []int) {
	return file_subscriptions_proto_rawDescGZIP(), []int{0}
}

func (x *Genre) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//A Representation of subscription request from user
type SubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//User email from which request has been raised
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SubscriptionRequest) Reset() {
	*x = SubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionRequest) ProtoMessage() {}

func (x *SubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionRequest.ProtoReflect.Descriptor instead.
func (*SubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_subscriptions_proto_rawDescGZIP(), []int{1}
}

func (x *SubscriptionRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ListPlansRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPlansRequest) Reset() {
	*x = ListPlansRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPlansRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlansRequest) ProtoMessage() {}

func (x *ListPlansRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlansRequest.ProtoReflect.Descriptor instead.
func (*ListPlansRequest) Descriptor() ([]byte, []int) {
	return file_subscriptions_proto_rawDescGZIP(), []int{2}
}

type Plans struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subs []*Plan `protobuf:"bytes,1,rep,name=subs,proto3" json:"subs,omitempty"`
}

func (x *Plans) Reset() {
	*x = Plans{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plans) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plans) ProtoMessage() {}

func (x *Plans) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plans.ProtoReflect.Descriptor instead.
func (*Plans) Descriptor() ([]byte, []int) {
	return file_subscriptions_proto_rawDescGZIP(), []int{3}
}

func (x *Plans) GetSubs() []*Plan {
	if x != nil {
		return x.Subs
	}
	return nil
}

type Plan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Validity int32    `protobuf:"varint,3,opt,name=validity,proto3" json:"validity,omitempty"`
	Genres   []*Genre `protobuf:"bytes,4,rep,name=genres,proto3" json:"genres,omitempty"`
}

func (x *Plan) Reset() {
	*x = Plan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plan) ProtoMessage() {}

func (x *Plan) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plan.ProtoReflect.Descriptor instead.
func (*Plan) Descriptor() ([]byte, []int) {
	return file_subscriptions_proto_rawDescGZIP(), []int{4}
}

func (x *Plan) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Plan) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plan) GetValidity() int32 {
	if x != nil {
		return x.Validity
	}
	return 0
}

func (x *Plan) GetGenres() []*Genre {
	if x != nil {
		return x.Genres
	}
	return nil
}

type CreateSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email  string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Subsid int32  `protobuf:"varint,2,opt,name=subsid,proto3" json:"subsid,omitempty"`
}

func (x *CreateSubscriptionRequest) Reset() {
	*x = CreateSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubscriptionRequest) ProtoMessage() {}

func (x *CreateSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*CreateSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_subscriptions_proto_rawDescGZIP(), []int{5}
}

func (x *CreateSubscriptionRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateSubscriptionRequest) GetSubsid() int32 {
	if x != nil {
		return x.Subsid
	}
	return 0
}

type SubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Active    bool                   `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	Starttime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=starttime,proto3" json:"starttime,omitempty"`
	Validity  int32                  `protobuf:"varint,4,opt,name=validity,proto3" json:"validity,omitempty"`
	Endtime   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=endtime,proto3" json:"endtime,omitempty"`
}

func (x *SubscriptionResponse) Reset() {
	*x = SubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptions_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionResponse) ProtoMessage() {}

func (x *SubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptions_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionResponse.ProtoReflect.Descriptor instead.
func (*SubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_subscriptions_proto_rawDescGZIP(), []int{6}
}

func (x *SubscriptionResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SubscriptionResponse) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *SubscriptionResponse) GetStarttime() *timestamppb.Timestamp {
	if x != nil {
		return x.Starttime
	}
	return nil
}

func (x *SubscriptionResponse) GetValidity() int32 {
	if x != nil {
		return x.Validity
	}
	return 0
}

func (x *SubscriptionResponse) GetEndtime() *timestamppb.Timestamp {
	if x != nil {
		return x.Endtime
	}
	return nil
}

var File_subscriptions_proto protoreflect.FileDescriptor

var file_subscriptions_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x05, 0x47,
	0x65, 0x6e, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x05, 0x50, 0x6c, 0x61,
	0x6e, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x73, 0x75, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x73, 0x75, 0x62,
	0x73, 0x22, 0x7a, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x06, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x22, 0x49, 0x0a,
	0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x75, 0x62, 0x73, 0x69, 0x64, 0x22, 0xd0, 0x01, 0x0a, 0x14, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x74, 0x69, 0x6d, 0x65, 0x32, 0xbe, 0x02, 0x0a, 0x1d,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x22, 0x00, 0x12, 0x71, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x09, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c,
	0x2e, 0x3b, 0x6e, 0x65, 0x77, 0x73, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subscriptions_proto_rawDescOnce sync.Once
	file_subscriptions_proto_rawDescData = file_subscriptions_proto_rawDesc
)

func file_subscriptions_proto_rawDescGZIP() []byte {
	file_subscriptions_proto_rawDescOnce.Do(func() {
		file_subscriptions_proto_rawDescData = protoimpl.X.CompressGZIP(file_subscriptions_proto_rawDescData)
	})
	return file_subscriptions_proto_rawDescData
}

var file_subscriptions_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_subscriptions_proto_goTypes = []interface{}{
	(*Genre)(nil),                     // 0: proto.subscriptions.Genre
	(*SubscriptionRequest)(nil),       // 1: proto.subscriptions.SubscriptionRequest
	(*ListPlansRequest)(nil),          // 2: proto.subscriptions.ListPlansRequest
	(*Plans)(nil),                     // 3: proto.subscriptions.Plans
	(*Plan)(nil),                      // 4: proto.subscriptions.Plan
	(*CreateSubscriptionRequest)(nil), // 5: proto.subscriptions.CreateSubscriptionRequest
	(*SubscriptionResponse)(nil),      // 6: proto.subscriptions.SubscriptionResponse
	(*timestamppb.Timestamp)(nil),     // 7: google.protobuf.Timestamp
}
var file_subscriptions_proto_depIdxs = []int32{
	4, // 0: proto.subscriptions.Plans.subs:type_name -> proto.subscriptions.Plan
	0, // 1: proto.subscriptions.Plan.genres:type_name -> proto.subscriptions.Genre
	7, // 2: proto.subscriptions.SubscriptionResponse.starttime:type_name -> google.protobuf.Timestamp
	7, // 3: proto.subscriptions.SubscriptionResponse.endtime:type_name -> google.protobuf.Timestamp
	1, // 4: proto.subscriptions.SubscriptionManagementService.GetSubscription:input_type -> proto.subscriptions.SubscriptionRequest
	5, // 5: proto.subscriptions.SubscriptionManagementService.CreateSubscription:input_type -> proto.subscriptions.CreateSubscriptionRequest
	2, // 6: proto.subscriptions.SubscriptionManagementService.ListPlans:input_type -> proto.subscriptions.ListPlansRequest
	4, // 7: proto.subscriptions.SubscriptionManagementService.GetSubscription:output_type -> proto.subscriptions.Plan
	6, // 8: proto.subscriptions.SubscriptionManagementService.CreateSubscription:output_type -> proto.subscriptions.SubscriptionResponse
	3, // 9: proto.subscriptions.SubscriptionManagementService.ListPlans:output_type -> proto.subscriptions.Plans
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_subscriptions_proto_init() }
func file_subscriptions_proto_init() {
	if File_subscriptions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_subscriptions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Genre); i {
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
		file_subscriptions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionRequest); i {
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
		file_subscriptions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPlansRequest); i {
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
		file_subscriptions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plans); i {
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
		file_subscriptions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plan); i {
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
		file_subscriptions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSubscriptionRequest); i {
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
		file_subscriptions_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionResponse); i {
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
			RawDescriptor: file_subscriptions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_subscriptions_proto_goTypes,
		DependencyIndexes: file_subscriptions_proto_depIdxs,
		MessageInfos:      file_subscriptions_proto_msgTypes,
	}.Build()
	File_subscriptions_proto = out.File
	file_subscriptions_proto_rawDesc = nil
	file_subscriptions_proto_goTypes = nil
	file_subscriptions_proto_depIdxs = nil
}
