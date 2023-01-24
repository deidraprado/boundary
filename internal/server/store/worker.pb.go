// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: controller/storage/servers/store/v1/worker.proto

package store

import (
	timestamp "github.com/hashicorp/boundary/internal/db/timestamp"
	_ "github.com/hashicorp/boundary/sdk/pbs/controller/protooptions"
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

// Worker contains all fields related to a Worker resource
type Worker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// public_id is a surrogate key suitable for use in a public API
	// @inject_tag: `gorm:"primary_key"`
	PublicId string `protobuf:"bytes,10,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty" gorm:"primary_key"`
	// Name of the resource (optional)
	// @inject_tag: `gorm:"default:null"`
	Name string `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty" gorm:"default:null"`
	// Description of the resource (optional)
	// @inject_tag: `gorm:"default:null"`
	Description string `protobuf:"bytes,30,opt,name=description,proto3" json:"description,omitempty" gorm:"default:null"`
	// Address for the worker. This is optional.
	// @inject_tag: `gorm:"default:null"`
	Address string `protobuf:"bytes,40,opt,name=address,proto3" json:"address,omitempty" gorm:"default:null"`
	// The create_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,50,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// The update_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,60,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty" gorm:"default:current_timestamp"`
	// The scope_id of the owning scope and must be set.
	// @inject_tag: `gorm:"not_null"`
	ScopeId string `protobuf:"bytes,70,opt,name=scope_id,json=scopeId,proto3" json:"scope_id,omitempty" gorm:"not_null"`
	// version allows optimistic locking of the resource.
	// @inject_tag: `gorm:"default:null"`
	Version uint32 `protobuf:"varint,80,opt,name=version,proto3" json:"version,omitempty" gorm:"default:null"`
	// The last_status_time is the last time a worker daemon sent a status update.
	// @inject_tag: `gorm:"default:null"`
	LastStatusTime *timestamp.Timestamp `protobuf:"bytes,110,opt,name=last_status_time,json=lastStatusTime,proto3" json:"last_status_time,omitempty" gorm:"default:null"`
	// The type of the worker, denoted by how it authenticates: pki or kms.
	// @inject_tag: `gorm:"not_null"`
	Type string `protobuf:"bytes,130,opt,name=type,proto3" json:"type,omitempty" gorm:"not_null"`
	// The version of Boundary the worker binary is running
	// @inject_tag: `gorm:"default:null"`
	ReleaseVersion string `protobuf:"bytes,140,opt,name=release_version,json=releaseVersion,proto3" json:"release_version,omitempty" gorm:"default:null"`
	// The state of the worker, to indicate if the worker is active or in shutdown.
	// @inject_tag: `gorm:"not_null"`
	OperationalState string `protobuf:"bytes,150,opt,name=operational_state,json=operationalState,proto3" json:"operational_state,omitempty" gorm:"not_null"`
}

func (x *Worker) Reset() {
	*x = Worker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_servers_store_v1_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Worker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Worker) ProtoMessage() {}

func (x *Worker) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_servers_store_v1_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Worker.ProtoReflect.Descriptor instead.
func (*Worker) Descriptor() ([]byte, []int) {
	return file_controller_storage_servers_store_v1_worker_proto_rawDescGZIP(), []int{0}
}

func (x *Worker) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *Worker) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Worker) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Worker) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Worker) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Worker) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Worker) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

func (x *Worker) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Worker) GetLastStatusTime() *timestamp.Timestamp {
	if x != nil {
		return x.LastStatusTime
	}
	return nil
}

func (x *Worker) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Worker) GetReleaseVersion() string {
	if x != nil {
		return x.ReleaseVersion
	}
	return ""
}

func (x *Worker) GetOperationalState() string {
	if x != nil {
		return x.OperationalState
	}
	return ""
}

// WorkerTag is a tag for a worker.  The primary key is comprised of the
// worker_id, key, value, and source.
type WorkerTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// worker_id is the public key that key of the worker this tag is for.
	// @inject_tag: `gorm:"primary_key"`
	WorkerId string `protobuf:"bytes,10,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty" gorm:"primary_key"`
	// key is the key of the tag. This must be set.
	// @inject_tag: `gorm:"primary_key"`
	Key string `protobuf:"bytes,20,opt,name=key,proto3" json:"key,omitempty" gorm:"primary_key"`
	// value is the value
	// @inject_tag: `gorm:"primary_key"`
	Value string `protobuf:"bytes,30,opt,name=value,proto3" json:"value,omitempty" gorm:"primary_key"`
	// source is the source of the tag. Either 'configuration' or 'api'.
	// @inject_tag: `gorm:"primary_key"`
	Source string `protobuf:"bytes,40,opt,name=source,proto3" json:"source,omitempty" gorm:"primary_key"`
}

func (x *WorkerTag) Reset() {
	*x = WorkerTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_servers_store_v1_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerTag) ProtoMessage() {}

func (x *WorkerTag) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_servers_store_v1_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerTag.ProtoReflect.Descriptor instead.
func (*WorkerTag) Descriptor() ([]byte, []int) {
	return file_controller_storage_servers_store_v1_worker_proto_rawDescGZIP(), []int{1}
}

func (x *WorkerTag) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *WorkerTag) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *WorkerTag) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *WorkerTag) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

var File_controller_storage_servers_store_v1_worker_proto protoreflect.FileDescriptor

var file_controller_storage_servers_store_v1_worker_proto_rawDesc = []byte{
	0x0a, 0x30, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x04, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xc2, 0xdd, 0x29, 0x0c,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xc2, 0xdd, 0x29, 0x1a, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x4b,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x32, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x50,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a,
	0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x13, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x82, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x8c, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x22, 0x68, 0x0a, 0x09, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x54, 0x61, 0x67, 0x12, 0x1b, 0x0a,
	0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_storage_servers_store_v1_worker_proto_rawDescOnce sync.Once
	file_controller_storage_servers_store_v1_worker_proto_rawDescData = file_controller_storage_servers_store_v1_worker_proto_rawDesc
)

func file_controller_storage_servers_store_v1_worker_proto_rawDescGZIP() []byte {
	file_controller_storage_servers_store_v1_worker_proto_rawDescOnce.Do(func() {
		file_controller_storage_servers_store_v1_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_storage_servers_store_v1_worker_proto_rawDescData)
	})
	return file_controller_storage_servers_store_v1_worker_proto_rawDescData
}

var file_controller_storage_servers_store_v1_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_controller_storage_servers_store_v1_worker_proto_goTypes = []interface{}{
	(*Worker)(nil),              // 0: controller.storage.servers.store.v1.Worker
	(*WorkerTag)(nil),           // 1: controller.storage.servers.store.v1.WorkerTag
	(*timestamp.Timestamp)(nil), // 2: controller.storage.timestamp.v1.Timestamp
}
var file_controller_storage_servers_store_v1_worker_proto_depIdxs = []int32{
	2, // 0: controller.storage.servers.store.v1.Worker.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // 1: controller.storage.servers.store.v1.Worker.update_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // 2: controller.storage.servers.store.v1.Worker.last_status_time:type_name -> controller.storage.timestamp.v1.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_controller_storage_servers_store_v1_worker_proto_init() }
func file_controller_storage_servers_store_v1_worker_proto_init() {
	if File_controller_storage_servers_store_v1_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_storage_servers_store_v1_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Worker); i {
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
		file_controller_storage_servers_store_v1_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerTag); i {
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
			RawDescriptor: file_controller_storage_servers_store_v1_worker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_servers_store_v1_worker_proto_goTypes,
		DependencyIndexes: file_controller_storage_servers_store_v1_worker_proto_depIdxs,
		MessageInfos:      file_controller_storage_servers_store_v1_worker_proto_msgTypes,
	}.Build()
	File_controller_storage_servers_store_v1_worker_proto = out.File
	file_controller_storage_servers_store_v1_worker_proto_rawDesc = nil
	file_controller_storage_servers_store_v1_worker_proto_goTypes = nil
	file_controller_storage_servers_store_v1_worker_proto_depIdxs = nil
}