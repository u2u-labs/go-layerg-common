package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Fetch a subscription by product id.
type GetSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Product id of the subscription
	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *GetSubscriptionRequest) Reset() {
	*x = GetSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[41]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionRequest) ProtoMessage() {}

func (x *GetSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[41]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*GetSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{41}
}

func (x *GetSubscriptionRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

// List user subscriptions.
type ListSubscriptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Max number of results per page
	Limit *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Cursor to retrieve a page of records from
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *ListSubscriptionsRequest) Reset() {
	*x = ListSubscriptionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[67]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubscriptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubscriptionsRequest) ProtoMessage() {}

func (x *ListSubscriptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[67]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubscriptionsRequest.ProtoReflect.Descriptor instead.
func (*ListSubscriptionsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{67}
}

func (x *ListSubscriptionsRequest) GetLimit() *wrapperspb.Int32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *ListSubscriptionsRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

// Apple Subscription validation request
type ValidateSubscriptionAppleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Base64 encoded Apple receipt data payload.
	Receipt string `protobuf:"bytes,1,opt,name=receipt,proto3" json:"receipt,omitempty"`
	// Persist the subscription.
	Persist *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=persist,proto3" json:"persist,omitempty"`
}

func (x *ValidateSubscriptionAppleRequest) Reset() {
	*x = ValidateSubscriptionAppleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[96]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateSubscriptionAppleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateSubscriptionAppleRequest) ProtoMessage() {}

func (x *ValidateSubscriptionAppleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[96]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateSubscriptionAppleRequest.ProtoReflect.Descriptor instead.
func (*ValidateSubscriptionAppleRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{96}
}

func (x *ValidateSubscriptionAppleRequest) GetReceipt() string {
	if x != nil {
		return x.Receipt
	}
	return ""
}

func (x *ValidateSubscriptionAppleRequest) GetPersist() *wrapperspb.BoolValue {
	if x != nil {
		return x.Persist
	}
	return nil
}

// Google Subscription validation request
type ValidateSubscriptionGoogleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// JSON encoded Google purchase payload.
	Receipt string `protobuf:"bytes,1,opt,name=receipt,proto3" json:"receipt,omitempty"`
	// Persist the subscription.
	Persist *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=persist,proto3" json:"persist,omitempty"`
}

func (x *ValidateSubscriptionGoogleRequest) Reset() {
	*x = ValidateSubscriptionGoogleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[98]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateSubscriptionGoogleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateSubscriptionGoogleRequest) ProtoMessage() {}

func (x *ValidateSubscriptionGoogleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[98]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateSubscriptionGoogleRequest.ProtoReflect.Descriptor instead.
func (*ValidateSubscriptionGoogleRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{98}
}

func (x *ValidateSubscriptionGoogleRequest) GetReceipt() string {
	if x != nil {
		return x.Receipt
	}
	return ""
}

func (x *ValidateSubscriptionGoogleRequest) GetPersist() *wrapperspb.BoolValue {
	if x != nil {
		return x.Persist
	}
	return nil
}

// Validate Subscription response.
type ValidateSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidatedSubscription *ValidatedSubscription `protobuf:"bytes,1,opt,name=validated_subscription,json=validatedSubscription,proto3" json:"validated_subscription,omitempty"`
}

func (x *ValidateSubscriptionResponse) Reset() {
	*x = ValidateSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[103]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateSubscriptionResponse) ProtoMessage() {}

func (x *ValidateSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[103]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*ValidateSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{103}
}

func (x *ValidateSubscriptionResponse) GetValidatedSubscription() *ValidatedSubscription {
	if x != nil {
		return x.ValidatedSubscription
	}
	return nil
}

type ValidatedSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Subscription User ID.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Purchase Product ID.
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	// Purchase Original transaction ID (we only keep track of the original subscription, not subsequent renewals).
	OriginalTransactionId string `protobuf:"bytes,3,opt,name=original_transaction_id,json=originalTransactionId,proto3" json:"original_transaction_id,omitempty"`
	// Store identifier
	Store StoreProvider `protobuf:"varint,4,opt,name=store,proto3,enum=nakama.api.StoreProvider" json:"store,omitempty"`
	// UNIX Timestamp when the purchase was done.
	PurchaseTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=purchase_time,json=purchaseTime,proto3" json:"purchase_time,omitempty"`
	// UNIX Timestamp when the receipt validation was stored in DB.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// UNIX Timestamp when the receipt validation was updated in DB.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Whether the purchase was done in production or sandbox environment.
	Environment StoreEnvironment `protobuf:"varint,8,opt,name=environment,proto3,enum=nakama.api.StoreEnvironment" json:"environment,omitempty"`
	// Subscription expiration time. The subscription can still be auto-renewed to extend the expiration time further.
	ExpiryTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=expiry_time,json=expiryTime,proto3" json:"expiry_time,omitempty"`
	// Subscription refund time. If this time is set, the subscription was refunded.
	RefundTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=refund_time,json=refundTime,proto3" json:"refund_time,omitempty"`
	// Raw provider validation response body.
	ProviderResponse string `protobuf:"bytes,11,opt,name=provider_response,json=providerResponse,proto3" json:"provider_response,omitempty"`
	// Raw provider notification body.
	ProviderNotification string `protobuf:"bytes,12,opt,name=provider_notification,json=providerNotification,proto3" json:"provider_notification,omitempty"`
	// Whether the subscription is currently active or not.
	Active bool `protobuf:"varint,13,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *ValidatedSubscription) Reset() {
	*x = ValidatedSubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[104]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatedSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatedSubscription) ProtoMessage() {}

func (x *ValidatedSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[104]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatedSubscription.ProtoReflect.Descriptor instead.
func (*ValidatedSubscription) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{104}
}

func (x *ValidatedSubscription) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ValidatedSubscription) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ValidatedSubscription) GetOriginalTransactionId() string {
	if x != nil {
		return x.OriginalTransactionId
	}
	return ""
}

func (x *ValidatedSubscription) GetStore() StoreProvider {
	if x != nil {
		return x.Store
	}
	return StoreProvider_APPLE_APP_STORE
}

func (x *ValidatedSubscription) GetPurchaseTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PurchaseTime
	}
	return nil
}

func (x *ValidatedSubscription) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ValidatedSubscription) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *ValidatedSubscription) GetEnvironment() StoreEnvironment {
	if x != nil {
		return x.Environment
	}
	return StoreEnvironment_UNKNOWN
}

func (x *ValidatedSubscription) GetExpiryTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiryTime
	}
	return nil
}

func (x *ValidatedSubscription) GetRefundTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RefundTime
	}
	return nil
}

func (x *ValidatedSubscription) GetProviderResponse() string {
	if x != nil {
		return x.ProviderResponse
	}
	return ""
}

func (x *ValidatedSubscription) GetProviderNotification() string {
	if x != nil {
		return x.ProviderNotification
	}
	return ""
}

func (x *ValidatedSubscription) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

// A list of validated subscriptions stored by Layerg.
type SubscriptionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Stored validated subscriptions.
	ValidatedSubscriptions []*ValidatedSubscription `protobuf:"bytes,1,rep,name=validated_subscriptions,json=validatedSubscriptions,proto3" json:"validated_subscriptions,omitempty"`
	// The cursor to send when retrieving the next page, if any.
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	// The cursor to send when retrieving the previous page, if any.
	PrevCursor string `protobuf:"bytes,3,opt,name=prev_cursor,json=prevCursor,proto3" json:"prev_cursor,omitempty"`
}

func (x *SubscriptionList) Reset() {
	*x = SubscriptionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[106]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionList) ProtoMessage() {}

func (x *SubscriptionList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[106]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionList.ProtoReflect.Descriptor instead.
func (*SubscriptionList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{106}
}

func (x *SubscriptionList) GetValidatedSubscriptions() []*ValidatedSubscription {
	if x != nil {
		return x.ValidatedSubscriptions
	}
	return nil
}

func (x *SubscriptionList) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *SubscriptionList) GetPrevCursor() string {
	if x != nil {
		return x.PrevCursor
	}
	return ""
}
