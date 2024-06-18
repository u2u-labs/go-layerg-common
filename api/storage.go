package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Storage objects to delete.
type DeleteStorageObjectId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The collection which stores the object.
	Collection string `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	// The key of the object within the collection.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// The version hash of the object.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *DeleteStorageObjectId) Reset() {
	*x = DeleteStorageObjectId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[34]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStorageObjectId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStorageObjectId) ProtoMessage() {}

func (x *DeleteStorageObjectId) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[34]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStorageObjectId.ProtoReflect.Descriptor instead.
func (*DeleteStorageObjectId) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{34}
}

func (x *DeleteStorageObjectId) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *DeleteStorageObjectId) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *DeleteStorageObjectId) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// Batch delete storage objects.
type DeleteStorageObjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Batch of storage objects.
	ObjectIds []*DeleteStorageObjectId `protobuf:"bytes,1,rep,name=object_ids,json=objectIds,proto3" json:"object_ids,omitempty"`
}

func (x *DeleteStorageObjectsRequest) Reset() {
	*x = DeleteStorageObjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[35]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStorageObjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStorageObjectsRequest) ProtoMessage() {}

func (x *DeleteStorageObjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[35]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStorageObjectsRequest.ProtoReflect.Descriptor instead.
func (*DeleteStorageObjectsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{35}
}

func (x *DeleteStorageObjectsRequest) GetObjectIds() []*DeleteStorageObjectId {
	if x != nil {
		return x.ObjectIds
	}
	return nil
}

// List publicly readable storage objects in a given collection.
type ListStorageObjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the user.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// The collection which stores the object.
	Collection string `protobuf:"bytes,2,opt,name=collection,proto3" json:"collection,omitempty"`
	// The number of storage objects to list. Between 1 and 100.
	Limit *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// The cursor to page through results from.
	Cursor string `protobuf:"bytes,4,opt,name=cursor,proto3" json:"cursor,omitempty"` // value from StorageObjectList.cursor.
}

func (x *ListStorageObjectsRequest) Reset() {
	*x = ListStorageObjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[66]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStorageObjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStorageObjectsRequest) ProtoMessage() {}

func (x *ListStorageObjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[66]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStorageObjectsRequest.ProtoReflect.Descriptor instead.
func (*ListStorageObjectsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{66}
}

func (x *ListStorageObjectsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListStorageObjectsRequest) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *ListStorageObjectsRequest) GetLimit() *wrapperspb.Int32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *ListStorageObjectsRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

// Storage objects to get.
type ReadStorageObjectId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The collection which stores the object.
	Collection string `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	// The key of the object within the collection.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// The user owner of the object.
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ReadStorageObjectId) Reset() {
	*x = ReadStorageObjectId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[78]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadStorageObjectId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadStorageObjectId) ProtoMessage() {}

func (x *ReadStorageObjectId) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[78]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadStorageObjectId.ProtoReflect.Descriptor instead.
func (*ReadStorageObjectId) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{78}
}

func (x *ReadStorageObjectId) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *ReadStorageObjectId) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ReadStorageObjectId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Batch get storage objects.
type ReadStorageObjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Batch of storage objects.
	ObjectIds []*ReadStorageObjectId `protobuf:"bytes,1,rep,name=object_ids,json=objectIds,proto3" json:"object_ids,omitempty"`
}

func (x *ReadStorageObjectsRequest) Reset() {
	*x = ReadStorageObjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[79]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadStorageObjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadStorageObjectsRequest) ProtoMessage() {}

func (x *ReadStorageObjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[79]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadStorageObjectsRequest.ProtoReflect.Descriptor instead.
func (*ReadStorageObjectsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{79}
}

func (x *ReadStorageObjectsRequest) GetObjectIds() []*ReadStorageObjectId {
	if x != nil {
		return x.ObjectIds
	}
	return nil
}

// An object within the storage engine.
type StorageObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The collection which stores the object.
	Collection string `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	// The key of the object within the collection.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// The user owner of the object.
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// The value of the object.
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	// The version hash of the object.
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	// The read access permissions for the object.
	PermissionRead int32 `protobuf:"varint,6,opt,name=permission_read,json=permissionRead,proto3" json:"permission_read,omitempty"`
	// The write access permissions for the object.
	PermissionWrite int32 `protobuf:"varint,7,opt,name=permission_write,json=permissionWrite,proto3" json:"permission_write,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the object was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the object was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *StorageObject) Reset() {
	*x = StorageObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[82]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageObject) ProtoMessage() {}

func (x *StorageObject) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[82]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageObject.ProtoReflect.Descriptor instead.
func (*StorageObject) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{82}
}

func (x *StorageObject) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *StorageObject) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *StorageObject) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *StorageObject) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *StorageObject) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *StorageObject) GetPermissionRead() int32 {
	if x != nil {
		return x.PermissionRead
	}
	return 0
}

func (x *StorageObject) GetPermissionWrite() int32 {
	if x != nil {
		return x.PermissionWrite
	}
	return 0
}

func (x *StorageObject) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *StorageObject) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// A storage acknowledgement.
type StorageObjectAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The collection which stores the object.
	Collection string `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	// The key of the object within the collection.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// The version hash of the object.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// The owner of the object.
	UserId string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the object was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the object was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *StorageObjectAck) Reset() {
	*x = StorageObjectAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[83]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageObjectAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageObjectAck) ProtoMessage() {}

func (x *StorageObjectAck) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[83]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageObjectAck.ProtoReflect.Descriptor instead.
func (*StorageObjectAck) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{83}
}

func (x *StorageObjectAck) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *StorageObjectAck) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *StorageObjectAck) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *StorageObjectAck) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *StorageObjectAck) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *StorageObjectAck) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// Batch of acknowledgements for the storage object write.
type StorageObjectAcks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Batch of storage write acknowledgements.
	Acks []*StorageObjectAck `protobuf:"bytes,1,rep,name=acks,proto3" json:"acks,omitempty"`
}

func (x *StorageObjectAcks) Reset() {
	*x = StorageObjectAcks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[84]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageObjectAcks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageObjectAcks) ProtoMessage() {}

func (x *StorageObjectAcks) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[84]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageObjectAcks.ProtoReflect.Descriptor instead.
func (*StorageObjectAcks) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{84}
}

func (x *StorageObjectAcks) GetAcks() []*StorageObjectAck {
	if x != nil {
		return x.Acks
	}
	return nil
}

// Batch of storage objects.
type StorageObjects struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The batch of storage objects.
	Objects []*StorageObject `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *StorageObjects) Reset() {
	*x = StorageObjects{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[85]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageObjects) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageObjects) ProtoMessage() {}

func (x *StorageObjects) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[85]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageObjects.ProtoReflect.Descriptor instead.
func (*StorageObjects) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{85}
}

func (x *StorageObjects) GetObjects() []*StorageObject {
	if x != nil {
		return x.Objects
	}
	return nil
}

// List of storage objects.
type StorageObjectList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of storage objects.
	Objects []*StorageObject `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	// The cursor for the next page of results, if any.
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *StorageObjectList) Reset() {
	*x = StorageObjectList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[86]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageObjectList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageObjectList) ProtoMessage() {}

func (x *StorageObjectList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[86]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageObjectList.ProtoReflect.Descriptor instead.
func (*StorageObjectList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{86}
}

func (x *StorageObjectList) GetObjects() []*StorageObject {
	if x != nil {
		return x.Objects
	}
	return nil
}

func (x *StorageObjectList) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

// The object to store.
type WriteStorageObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The collection to store the object.
	Collection string `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	// The key for the object within the collection.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// The value of the object.
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// The version hash of the object to check. Possible values are: ["", "*", "#hash#"].
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"` // if-match and if-none-match
	// The read access permissions for the object.
	PermissionRead *wrapperspb.Int32Value `protobuf:"bytes,5,opt,name=permission_read,json=permissionRead,proto3" json:"permission_read,omitempty"`
	// The write access permissions for the object.
	PermissionWrite *wrapperspb.Int32Value `protobuf:"bytes,6,opt,name=permission_write,json=permissionWrite,proto3" json:"permission_write,omitempty"`
}

func (x *WriteStorageObject) Reset() {
	*x = WriteStorageObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[108]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteStorageObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteStorageObject) ProtoMessage() {}

func (x *WriteStorageObject) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[108]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteStorageObject.ProtoReflect.Descriptor instead.
func (*WriteStorageObject) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{108}
}

func (x *WriteStorageObject) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *WriteStorageObject) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *WriteStorageObject) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *WriteStorageObject) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *WriteStorageObject) GetPermissionRead() *wrapperspb.Int32Value {
	if x != nil {
		return x.PermissionRead
	}
	return nil
}

func (x *WriteStorageObject) GetPermissionWrite() *wrapperspb.Int32Value {
	if x != nil {
		return x.PermissionWrite
	}
	return nil
}

// Write objects to the storage engine.
type WriteStorageObjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The objects to store on the server.
	Objects []*WriteStorageObject `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *WriteStorageObjectsRequest) Reset() {
	*x = WriteStorageObjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[109]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteStorageObjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteStorageObjectsRequest) ProtoMessage() {}

func (x *WriteStorageObjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[109]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteStorageObjectsRequest.ProtoReflect.Descriptor instead.
func (*WriteStorageObjectsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{109}
}

func (x *WriteStorageObjectsRequest) GetObjects() []*WriteStorageObject {
	if x != nil {
		return x.Objects
	}
	return nil
}
