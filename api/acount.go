package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// A user with additional account details. Always the current user.
type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The user object.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// The user's wallet data.
	Wallet string `protobuf:"bytes,2,opt,name=wallet,proto3" json:"wallet,omitempty"`
	// The email address of the user.
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	// The devices which belong to the user's account.
	Devices []*AccountDevice `protobuf:"bytes,4,rep,name=devices,proto3" json:"devices,omitempty"`
	// The custom id in the user's account.
	CustomId string `protobuf:"bytes,5,opt,name=custom_id,json=customId,proto3" json:"custom_id,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the user's email was verified.
	VerifyTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=verify_time,json=verifyTime,proto3" json:"verify_time,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the user's account was disabled/banned.
	DisableTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=disable_time,json=disableTime,proto3" json:"disable_time,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Account) GetWallet() string {
	if x != nil {
		return x.Wallet
	}
	return ""
}

func (x *Account) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Account) GetDevices() []*AccountDevice {
	if x != nil {
		return x.Devices
	}
	return nil
}

func (x *Account) GetCustomId() string {
	if x != nil {
		return x.CustomId
	}
	return ""
}

func (x *Account) GetVerifyTime() *timestamppb.Timestamp {
	if x != nil {
		return x.VerifyTime
	}
	return nil
}

func (x *Account) GetDisableTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DisableTime
	}
	return nil
}

// Update a user's account details.
type UpdateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The username of the user's account.
	Username *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// The display name of the user.
	DisplayName *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// A URL for an avatar image.
	AvatarUrl *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	// The language expected to be a tag which follows the BCP-47 spec.
	LangTag *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=lang_tag,json=langTag,proto3" json:"lang_tag,omitempty"`
	// The location set by the user.
	Location *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	// The timezone set by the user.
	Timezone *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=timezone,proto3" json:"timezone,omitempty"`
}

func (x *UpdateAccountRequest) Reset() {
	*x = UpdateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[90]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountRequest) ProtoMessage() {}

func (x *UpdateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[90]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountRequest.ProtoReflect.Descriptor instead.
func (*UpdateAccountRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{90}
}

func (x *UpdateAccountRequest) GetUsername() *wrapperspb.StringValue {
	if x != nil {
		return x.Username
	}
	return nil
}

func (x *UpdateAccountRequest) GetDisplayName() *wrapperspb.StringValue {
	if x != nil {
		return x.DisplayName
	}
	return nil
}

func (x *UpdateAccountRequest) GetAvatarUrl() *wrapperspb.StringValue {
	if x != nil {
		return x.AvatarUrl
	}
	return nil
}

func (x *UpdateAccountRequest) GetLangTag() *wrapperspb.StringValue {
	if x != nil {
		return x.LangTag
	}
	return nil
}

func (x *UpdateAccountRequest) GetLocation() *wrapperspb.StringValue {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *UpdateAccountRequest) GetTimezone() *wrapperspb.StringValue {
	if x != nil {
		return x.Timezone
	}
	return nil
}
