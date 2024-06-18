package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Apple IAP Purchases validation request
type ValidatePurchaseAppleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Base64 encoded Apple receipt data payload.
	Receipt string `protobuf:"bytes,1,opt,name=receipt,proto3" json:"receipt,omitempty"`
	// Persist the purchase
	Persist *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=persist,proto3" json:"persist,omitempty"`
}

func (x *ValidatePurchaseAppleRequest) Reset() {
	*x = ValidatePurchaseAppleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[95]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatePurchaseAppleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatePurchaseAppleRequest) ProtoMessage() {}

func (x *ValidatePurchaseAppleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[95]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatePurchaseAppleRequest.ProtoReflect.Descriptor instead.
func (*ValidatePurchaseAppleRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{95}
}

func (x *ValidatePurchaseAppleRequest) GetReceipt() string {
	if x != nil {
		return x.Receipt
	}
	return ""
}

func (x *ValidatePurchaseAppleRequest) GetPersist() *wrapperspb.BoolValue {
	if x != nil {
		return x.Persist
	}
	return nil
}

// Google IAP Purchase validation request
type ValidatePurchaseGoogleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// JSON encoded Google purchase payload.
	Purchase string `protobuf:"bytes,1,opt,name=purchase,proto3" json:"purchase,omitempty"`
	// Persist the purchase
	Persist *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=persist,proto3" json:"persist,omitempty"`
}

func (x *ValidatePurchaseGoogleRequest) Reset() {
	*x = ValidatePurchaseGoogleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[97]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatePurchaseGoogleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatePurchaseGoogleRequest) ProtoMessage() {}

func (x *ValidatePurchaseGoogleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[97]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatePurchaseGoogleRequest.ProtoReflect.Descriptor instead.
func (*ValidatePurchaseGoogleRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{97}
}

func (x *ValidatePurchaseGoogleRequest) GetPurchase() string {
	if x != nil {
		return x.Purchase
	}
	return ""
}

func (x *ValidatePurchaseGoogleRequest) GetPersist() *wrapperspb.BoolValue {
	if x != nil {
		return x.Persist
	}
	return nil
}

// Huawei IAP Purchase validation request
type ValidatePurchaseHuaweiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// JSON encoded Huawei InAppPurchaseData.
	Purchase string `protobuf:"bytes,1,opt,name=purchase,proto3" json:"purchase,omitempty"`
	// InAppPurchaseData signature.
	Signature string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	// Persist the purchase
	Persist *wrapperspb.BoolValue `protobuf:"bytes,3,opt,name=persist,proto3" json:"persist,omitempty"`
}

func (x *ValidatePurchaseHuaweiRequest) Reset() {
	*x = ValidatePurchaseHuaweiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[99]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatePurchaseHuaweiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatePurchaseHuaweiRequest) ProtoMessage() {}

func (x *ValidatePurchaseHuaweiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[99]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatePurchaseHuaweiRequest.ProtoReflect.Descriptor instead.
func (*ValidatePurchaseHuaweiRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{99}
}

func (x *ValidatePurchaseHuaweiRequest) GetPurchase() string {
	if x != nil {
		return x.Purchase
	}
	return ""
}

func (x *ValidatePurchaseHuaweiRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *ValidatePurchaseHuaweiRequest) GetPersist() *wrapperspb.BoolValue {
	if x != nil {
		return x.Persist
	}
	return nil
}

// Facebook Instant IAP Purchase validation request
type ValidatePurchaseFacebookInstantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Base64 encoded Facebook Instant signedRequest receipt data payload.
	SignedRequest string `protobuf:"bytes,1,opt,name=signed_request,json=signedRequest,proto3" json:"signed_request,omitempty"`
	// Persist the purchase
	Persist *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=persist,proto3" json:"persist,omitempty"`
}

func (x *ValidatePurchaseFacebookInstantRequest) Reset() {
	*x = ValidatePurchaseFacebookInstantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[100]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatePurchaseFacebookInstantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatePurchaseFacebookInstantRequest) ProtoMessage() {}

func (x *ValidatePurchaseFacebookInstantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[100]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatePurchaseFacebookInstantRequest.ProtoReflect.Descriptor instead.
func (*ValidatePurchaseFacebookInstantRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{100}
}

func (x *ValidatePurchaseFacebookInstantRequest) GetSignedRequest() string {
	if x != nil {
		return x.SignedRequest
	}
	return ""
}

func (x *ValidatePurchaseFacebookInstantRequest) GetPersist() *wrapperspb.BoolValue {
	if x != nil {
		return x.Persist
	}
	return nil
}

// Validated Purchase stored by Nakama.
type ValidatedPurchase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Purchase User ID.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Purchase Product ID.
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	// Purchase Transaction ID.
	TransactionId string `protobuf:"bytes,3,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	// Store identifier
	Store StoreProvider `protobuf:"varint,4,opt,name=store,proto3,enum=nakama.api.StoreProvider" json:"store,omitempty"`
	// Timestamp when the purchase was done.
	PurchaseTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=purchase_time,json=purchaseTime,proto3" json:"purchase_time,omitempty"`
	// Timestamp when the receipt validation was stored in DB.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Timestamp when the receipt validation was updated in DB.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Timestamp when the purchase was refunded. Set to UNIX
	RefundTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=refund_time,json=refundTime,proto3" json:"refund_time,omitempty"`
	// Raw provider validation response.
	ProviderResponse string `protobuf:"bytes,9,opt,name=provider_response,json=providerResponse,proto3" json:"provider_response,omitempty"`
	// Whether the purchase was done in production or sandbox environment.
	Environment StoreEnvironment `protobuf:"varint,10,opt,name=environment,proto3,enum=nakama.api.StoreEnvironment" json:"environment,omitempty"`
	// Whether the purchase had already been validated by Nakama before.
	SeenBefore bool `protobuf:"varint,11,opt,name=seen_before,json=seenBefore,proto3" json:"seen_before,omitempty"`
}

func (x *ValidatedPurchase) Reset() {
	*x = ValidatedPurchase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[101]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatedPurchase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatedPurchase) ProtoMessage() {}

func (x *ValidatedPurchase) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[101]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatedPurchase.ProtoReflect.Descriptor instead.
func (*ValidatedPurchase) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{101}
}

func (x *ValidatedPurchase) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ValidatedPurchase) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ValidatedPurchase) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *ValidatedPurchase) GetStore() StoreProvider {
	if x != nil {
		return x.Store
	}
	return StoreProvider_APPLE_APP_STORE
}

func (x *ValidatedPurchase) GetPurchaseTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PurchaseTime
	}
	return nil
}

func (x *ValidatedPurchase) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ValidatedPurchase) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *ValidatedPurchase) GetRefundTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RefundTime
	}
	return nil
}

func (x *ValidatedPurchase) GetProviderResponse() string {
	if x != nil {
		return x.ProviderResponse
	}
	return ""
}

func (x *ValidatedPurchase) GetEnvironment() StoreEnvironment {
	if x != nil {
		return x.Environment
	}
	return StoreEnvironment_UNKNOWN
}

func (x *ValidatedPurchase) GetSeenBefore() bool {
	if x != nil {
		return x.SeenBefore
	}
	return false
}

// Validate IAP response.
type ValidatePurchaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Newly seen validated purchases.
	ValidatedPurchases []*ValidatedPurchase `protobuf:"bytes,1,rep,name=validated_purchases,json=validatedPurchases,proto3" json:"validated_purchases,omitempty"`
}

func (x *ValidatePurchaseResponse) Reset() {
	*x = ValidatePurchaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[102]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatePurchaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatePurchaseResponse) ProtoMessage() {}

func (x *ValidatePurchaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[102]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatePurchaseResponse.ProtoReflect.Descriptor instead.
func (*ValidatePurchaseResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{102}
}

func (x *ValidatePurchaseResponse) GetValidatedPurchases() []*ValidatedPurchase {
	if x != nil {
		return x.ValidatedPurchases
	}
	return nil
}

// A list of validated purchases stored by Nakama.
type PurchaseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Stored validated purchases.
	ValidatedPurchases []*ValidatedPurchase `protobuf:"bytes,1,rep,name=validated_purchases,json=validatedPurchases,proto3" json:"validated_purchases,omitempty"`
	// The cursor to send when retrieving the next page, if any.
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	// The cursor to send when retrieving the previous page, if any.
	PrevCursor string `protobuf:"bytes,3,opt,name=prev_cursor,json=prevCursor,proto3" json:"prev_cursor,omitempty"`
}

func (x *PurchaseList) Reset() {
	*x = PurchaseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[105]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseList) ProtoMessage() {}

func (x *PurchaseList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[105]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseList.ProtoReflect.Descriptor instead.
func (*PurchaseList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{105}
}

func (x *PurchaseList) GetValidatedPurchases() []*ValidatedPurchase {
	if x != nil {
		return x.ValidatedPurchases
	}
	return nil
}

func (x *PurchaseList) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *PurchaseList) GetPrevCursor() string {
	if x != nil {
		return x.PrevCursor
	}
	return ""
}
