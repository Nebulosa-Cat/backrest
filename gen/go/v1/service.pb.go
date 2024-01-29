// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: v1/service.proto

package v1

import (
	types "github.com/garethgeorge/backrest/gen/go/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[0]
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
	return file_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

type ClearHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoId     string `protobuf:"bytes,1,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty"`
	PlanId     string `protobuf:"bytes,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	OnlyFailed bool   `protobuf:"varint,3,opt,name=only_failed,json=onlyFailed,proto3" json:"only_failed,omitempty"`
}

func (x *ClearHistoryRequest) Reset() {
	*x = ClearHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearHistoryRequest) ProtoMessage() {}

func (x *ClearHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearHistoryRequest.ProtoReflect.Descriptor instead.
func (*ClearHistoryRequest) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *ClearHistoryRequest) GetRepoId() string {
	if x != nil {
		return x.RepoId
	}
	return ""
}

func (x *ClearHistoryRequest) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

func (x *ClearHistoryRequest) GetOnlyFailed() bool {
	if x != nil {
		return x.OnlyFailed
	}
	return false
}

type ListSnapshotsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoId string `protobuf:"bytes,1,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty"`
	PlanId string `protobuf:"bytes,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
}

func (x *ListSnapshotsRequest) Reset() {
	*x = ListSnapshotsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSnapshotsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSnapshotsRequest) ProtoMessage() {}

func (x *ListSnapshotsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSnapshotsRequest.ProtoReflect.Descriptor instead.
func (*ListSnapshotsRequest) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListSnapshotsRequest) GetRepoId() string {
	if x != nil {
		return x.RepoId
	}
	return ""
}

func (x *ListSnapshotsRequest) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

type GetOperationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoId     string  `protobuf:"bytes,1,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty"`
	PlanId     string  `protobuf:"bytes,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	SnapshotId string  `protobuf:"bytes,4,opt,name=snapshot_id,json=snapshotId,proto3" json:"snapshot_id,omitempty"`
	Ids        []int64 `protobuf:"varint,5,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	LastN      int64   `protobuf:"varint,3,opt,name=last_n,json=lastN,proto3" json:"last_n,omitempty"` // limit to the last n operations
}

func (x *GetOperationsRequest) Reset() {
	*x = GetOperationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOperationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOperationsRequest) ProtoMessage() {}

func (x *GetOperationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOperationsRequest.ProtoReflect.Descriptor instead.
func (*GetOperationsRequest) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetOperationsRequest) GetRepoId() string {
	if x != nil {
		return x.RepoId
	}
	return ""
}

func (x *GetOperationsRequest) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

func (x *GetOperationsRequest) GetSnapshotId() string {
	if x != nil {
		return x.SnapshotId
	}
	return ""
}

func (x *GetOperationsRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GetOperationsRequest) GetLastN() int64 {
	if x != nil {
		return x.LastN
	}
	return 0
}

type RestoreSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlanId     string `protobuf:"bytes,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	SnapshotId string `protobuf:"bytes,2,opt,name=snapshot_id,json=snapshotId,proto3" json:"snapshot_id,omitempty"`
	Path       string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Target     string `protobuf:"bytes,4,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *RestoreSnapshotRequest) Reset() {
	*x = RestoreSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestoreSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreSnapshotRequest) ProtoMessage() {}

func (x *RestoreSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreSnapshotRequest.ProtoReflect.Descriptor instead.
func (*RestoreSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *RestoreSnapshotRequest) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

func (x *RestoreSnapshotRequest) GetSnapshotId() string {
	if x != nil {
		return x.SnapshotId
	}
	return ""
}

func (x *RestoreSnapshotRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *RestoreSnapshotRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type ListSnapshotFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoId     string `protobuf:"bytes,1,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty"`
	SnapshotId string `protobuf:"bytes,2,opt,name=snapshot_id,json=snapshotId,proto3" json:"snapshot_id,omitempty"`
	Path       string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ListSnapshotFilesRequest) Reset() {
	*x = ListSnapshotFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSnapshotFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSnapshotFilesRequest) ProtoMessage() {}

func (x *ListSnapshotFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSnapshotFilesRequest.ProtoReflect.Descriptor instead.
func (*ListSnapshotFilesRequest) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListSnapshotFilesRequest) GetRepoId() string {
	if x != nil {
		return x.RepoId
	}
	return ""
}

func (x *ListSnapshotFilesRequest) GetSnapshotId() string {
	if x != nil {
		return x.SnapshotId
	}
	return ""
}

func (x *ListSnapshotFilesRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ListSnapshotFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path    string     `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Entries []*LsEntry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *ListSnapshotFilesResponse) Reset() {
	*x = ListSnapshotFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSnapshotFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSnapshotFilesResponse) ProtoMessage() {}

func (x *ListSnapshotFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSnapshotFilesResponse.ProtoReflect.Descriptor instead.
func (*ListSnapshotFilesResponse) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListSnapshotFilesResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ListSnapshotFilesResponse) GetEntries() []*LsEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type OperationDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`  // operation id
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"` // data key
}

func (x *OperationDataRequest) Reset() {
	*x = OperationDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationDataRequest) ProtoMessage() {}

func (x *OperationDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationDataRequest.ProtoReflect.Descriptor instead.
func (*OperationDataRequest) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{8}
}

func (x *OperationDataRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OperationDataRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type LsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Path  string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Uid   int64  `protobuf:"varint,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid   int64  `protobuf:"varint,5,opt,name=gid,proto3" json:"gid,omitempty"`
	Size  int64  `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	Mode  int64  `protobuf:"varint,7,opt,name=mode,proto3" json:"mode,omitempty"`
	Mtime string `protobuf:"bytes,8,opt,name=mtime,proto3" json:"mtime,omitempty"`
	Atime string `protobuf:"bytes,9,opt,name=atime,proto3" json:"atime,omitempty"`
	Ctime string `protobuf:"bytes,10,opt,name=ctime,proto3" json:"ctime,omitempty"`
}

func (x *LsEntry) Reset() {
	*x = LsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LsEntry) ProtoMessage() {}

func (x *LsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LsEntry.ProtoReflect.Descriptor instead.
func (*LsEntry) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{9}
}

func (x *LsEntry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LsEntry) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *LsEntry) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *LsEntry) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *LsEntry) GetGid() int64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *LsEntry) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *LsEntry) GetMode() int64 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *LsEntry) GetMtime() string {
	if x != nil {
		return x.Mtime
	}
	return ""
}

func (x *LsEntry) GetAtime() string {
	if x != nil {
		return x.Atime
	}
	return ""
}

func (x *LsEntry) GetCtime() string {
	if x != nil {
		return x.Ctime
	}
	return ""
}

var File_v1_service_proto protoreflect.FileDescriptor

var file_v1_service_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x0f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x74,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0c, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x21, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x22, 0x68, 0x0a, 0x13, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x70, 0x6f, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x22, 0x48, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x22, 0x92, 0x01, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x22,
	0x7e, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22,
	0x68, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72,
	0x65, 0x70, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x70, 0x6f, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x56, 0x0a, 0x19, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x07, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x22, 0x38, 0x0a, 0x14, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xd3, 0x01, 0x0a, 0x07,
	0x4c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d,
	0x65, 0x32, 0xb7, 0x08, 0x0a, 0x08, 0x42, 0x61, 0x63, 0x6b, 0x72, 0x65, 0x73, 0x74, 0x12, 0x31,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x00, 0x12, 0x25, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0a,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x0a, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x70, 0x6f, 0x12, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x1a, 0x0a, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x73, 0x12, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x69, 0x63, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0e, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x12, 0x12, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x12, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x05, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x12, 0x12, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x46, 0x6f, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x3f, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x12, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x35, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x11, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0c, 0x43,
	0x6c, 0x65, 0x61, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x17, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3b,
	0x0a, 0x10, 0x50, 0x61, 0x74, 0x68, 0x41, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x72, 0x65, 0x74, 0x68,
	0x67, 0x65, 0x6f, 0x72, 0x67, 0x65, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x72, 0x65, 0x73, 0x74, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_v1_service_proto_rawDescOnce sync.Once
	file_v1_service_proto_rawDescData = file_v1_service_proto_rawDesc
)

func file_v1_service_proto_rawDescGZIP() []byte {
	file_v1_service_proto_rawDescOnce.Do(func() {
		file_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_service_proto_rawDescData)
	})
	return file_v1_service_proto_rawDescData
}

var file_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_v1_service_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),              // 0: v1.LoginRequest
	(*LoginResponse)(nil),             // 1: v1.LoginResponse
	(*ClearHistoryRequest)(nil),       // 2: v1.ClearHistoryRequest
	(*ListSnapshotsRequest)(nil),      // 3: v1.ListSnapshotsRequest
	(*GetOperationsRequest)(nil),      // 4: v1.GetOperationsRequest
	(*RestoreSnapshotRequest)(nil),    // 5: v1.RestoreSnapshotRequest
	(*ListSnapshotFilesRequest)(nil),  // 6: v1.ListSnapshotFilesRequest
	(*ListSnapshotFilesResponse)(nil), // 7: v1.ListSnapshotFilesResponse
	(*OperationDataRequest)(nil),      // 8: v1.OperationDataRequest
	(*LsEntry)(nil),                   // 9: v1.LsEntry
	(*emptypb.Empty)(nil),             // 10: google.protobuf.Empty
	(*Config)(nil),                    // 11: v1.Config
	(*Repo)(nil),                      // 12: v1.Repo
	(*types.StringValue)(nil),         // 13: types.StringValue
	(*types.Int64Value)(nil),          // 14: types.Int64Value
	(*OperationEvent)(nil),            // 15: v1.OperationEvent
	(*OperationList)(nil),             // 16: v1.OperationList
	(*ResticSnapshotList)(nil),        // 17: v1.ResticSnapshotList
	(*types.BytesValue)(nil),          // 18: types.BytesValue
	(*types.StringList)(nil),          // 19: types.StringList
}
var file_v1_service_proto_depIdxs = []int32{
	9,  // 0: v1.ListSnapshotFilesResponse.entries:type_name -> v1.LsEntry
	10, // 1: v1.Backrest.GetConfig:input_type -> google.protobuf.Empty
	11, // 2: v1.Backrest.SetConfig:input_type -> v1.Config
	12, // 3: v1.Backrest.AddRepo:input_type -> v1.Repo
	10, // 4: v1.Backrest.GetOperationEvents:input_type -> google.protobuf.Empty
	4,  // 5: v1.Backrest.GetOperations:input_type -> v1.GetOperationsRequest
	3,  // 6: v1.Backrest.ListSnapshots:input_type -> v1.ListSnapshotsRequest
	6,  // 7: v1.Backrest.ListSnapshotFiles:input_type -> v1.ListSnapshotFilesRequest
	13, // 8: v1.Backrest.IndexSnapshots:input_type -> types.StringValue
	13, // 9: v1.Backrest.Backup:input_type -> types.StringValue
	13, // 10: v1.Backrest.Prune:input_type -> types.StringValue
	13, // 11: v1.Backrest.Forget:input_type -> types.StringValue
	5,  // 12: v1.Backrest.Restore:input_type -> v1.RestoreSnapshotRequest
	13, // 13: v1.Backrest.Unlock:input_type -> types.StringValue
	13, // 14: v1.Backrest.Stats:input_type -> types.StringValue
	14, // 15: v1.Backrest.Cancel:input_type -> types.Int64Value
	8,  // 16: v1.Backrest.GetOperationData:input_type -> v1.OperationDataRequest
	2,  // 17: v1.Backrest.ClearHistory:input_type -> v1.ClearHistoryRequest
	13, // 18: v1.Backrest.PathAutocomplete:input_type -> types.StringValue
	11, // 19: v1.Backrest.GetConfig:output_type -> v1.Config
	11, // 20: v1.Backrest.SetConfig:output_type -> v1.Config
	11, // 21: v1.Backrest.AddRepo:output_type -> v1.Config
	15, // 22: v1.Backrest.GetOperationEvents:output_type -> v1.OperationEvent
	16, // 23: v1.Backrest.GetOperations:output_type -> v1.OperationList
	17, // 24: v1.Backrest.ListSnapshots:output_type -> v1.ResticSnapshotList
	7,  // 25: v1.Backrest.ListSnapshotFiles:output_type -> v1.ListSnapshotFilesResponse
	10, // 26: v1.Backrest.IndexSnapshots:output_type -> google.protobuf.Empty
	10, // 27: v1.Backrest.Backup:output_type -> google.protobuf.Empty
	10, // 28: v1.Backrest.Prune:output_type -> google.protobuf.Empty
	10, // 29: v1.Backrest.Forget:output_type -> google.protobuf.Empty
	10, // 30: v1.Backrest.Restore:output_type -> google.protobuf.Empty
	10, // 31: v1.Backrest.Unlock:output_type -> google.protobuf.Empty
	10, // 32: v1.Backrest.Stats:output_type -> google.protobuf.Empty
	10, // 33: v1.Backrest.Cancel:output_type -> google.protobuf.Empty
	18, // 34: v1.Backrest.GetOperationData:output_type -> types.BytesValue
	10, // 35: v1.Backrest.ClearHistory:output_type -> google.protobuf.Empty
	19, // 36: v1.Backrest.PathAutocomplete:output_type -> types.StringList
	19, // [19:37] is the sub-list for method output_type
	1,  // [1:19] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_v1_service_proto_init() }
func file_v1_service_proto_init() {
	if File_v1_service_proto != nil {
		return
	}
	file_v1_config_proto_init()
	file_v1_restic_proto_init()
	file_v1_operations_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearHistoryRequest); i {
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
		file_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSnapshotsRequest); i {
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
		file_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOperationsRequest); i {
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
		file_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestoreSnapshotRequest); i {
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
		file_v1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSnapshotFilesRequest); i {
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
		file_v1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSnapshotFilesResponse); i {
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
		file_v1_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationDataRequest); i {
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
		file_v1_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LsEntry); i {
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
			RawDescriptor: file_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_service_proto_goTypes,
		DependencyIndexes: file_v1_service_proto_depIdxs,
		MessageInfos:      file_v1_service_proto_msgTypes,
	}.Build()
	File_v1_service_proto = out.File
	file_v1_service_proto_rawDesc = nil
	file_v1_service_proto_goTypes = nil
	file_v1_service_proto_depIdxs = nil
}
