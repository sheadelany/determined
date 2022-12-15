// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// source: determined/notebook/v1/notebook.proto

package notebookv1

import (
	containerv1 "github.com/determined-ai/determined/proto/pkg/containerv1"
	taskv1 "github.com/determined-ai/determined/proto/pkg/taskv1"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

// Notebook is a Jupyter notebook in a containerized environment.
type Notebook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the notebook.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The description of the notebook.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The state of the notebook.
	State taskv1.State `protobuf:"varint,3,opt,name=state,proto3,enum=determined.task.v1.State" json:"state,omitempty"`
	// The time the notebook was started.
	StartTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The container running the notebook.
	Container *containerv1.Container `protobuf:"bytes,6,opt,name=container,proto3" json:"container,omitempty"`
	// The display name of the user that created the notebook.
	DisplayName string `protobuf:"bytes,15,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The id of the user that created the notebook.
	UserId int32 `protobuf:"varint,16,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// The username of the user that created the notebook.
	Username string `protobuf:"bytes,10,opt,name=username,proto3" json:"username,omitempty"`
	// The service address.
	ServiceAddress string `protobuf:"bytes,11,opt,name=service_address,json=serviceAddress,proto3" json:"service_address,omitempty"`
	// The name of the resource pool the Notebook was created in.
	ResourcePool string `protobuf:"bytes,12,opt,name=resource_pool,json=resourcePool,proto3" json:"resource_pool,omitempty"`
	// The exit status.
	ExitStatus string `protobuf:"bytes,13,opt,name=exit_status,json=exitStatus,proto3" json:"exit_status,omitempty"`
	// The associated job id.
	JobId string `protobuf:"bytes,14,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *Notebook) Reset() {
	*x = Notebook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_notebook_v1_notebook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notebook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notebook) ProtoMessage() {}

func (x *Notebook) ProtoReflect() protoreflect.Message {
	mi := &file_determined_notebook_v1_notebook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notebook.ProtoReflect.Descriptor instead.
func (*Notebook) Descriptor() ([]byte, []int) {
	return file_determined_notebook_v1_notebook_proto_rawDescGZIP(), []int{0}
}

func (x *Notebook) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notebook) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Notebook) GetState() taskv1.State {
	if x != nil {
		return x.State
	}
	return taskv1.State_STATE_UNSPECIFIED
}

func (x *Notebook) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Notebook) GetContainer() *containerv1.Container {
	if x != nil {
		return x.Container
	}
	return nil
}

func (x *Notebook) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Notebook) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Notebook) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Notebook) GetServiceAddress() string {
	if x != nil {
		return x.ServiceAddress
	}
	return ""
}

func (x *Notebook) GetResourcePool() string {
	if x != nil {
		return x.ResourcePool
	}
	return ""
}

func (x *Notebook) GetExitStatus() string {
	if x != nil {
		return x.ExitStatus
	}
	return ""
}

func (x *Notebook) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

var File_determined_notebook_v1_notebook_proto protoreflect.FileDescriptor

var file_determined_notebook_v1_notebook_proto_rawDesc = []byte{
	0x0a, 0x25, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x6e, 0x6f, 0x74,
	0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x65, 0x64, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27,
	0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x65, 0x64, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x04, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x65, 0x62,
	0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65,
	0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x40, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65,
	0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x69,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x78, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f,
	0x62, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49,
	0x64, 0x3a, 0x51, 0x92, 0x41, 0x4e, 0x0a, 0x4c, 0xd2, 0x01, 0x02, 0x69, 0x64, 0xd2, 0x01, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0xd2, 0x01, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0xd2, 0x01, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0xd2, 0x01, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0xd2, 0x01, 0x06, 0x6a, 0x6f,
	0x62, 0x5f, 0x69, 0x64, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2d, 0x61, 0x69,
	0x2f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_determined_notebook_v1_notebook_proto_rawDescOnce sync.Once
	file_determined_notebook_v1_notebook_proto_rawDescData = file_determined_notebook_v1_notebook_proto_rawDesc
)

func file_determined_notebook_v1_notebook_proto_rawDescGZIP() []byte {
	file_determined_notebook_v1_notebook_proto_rawDescOnce.Do(func() {
		file_determined_notebook_v1_notebook_proto_rawDescData = protoimpl.X.CompressGZIP(file_determined_notebook_v1_notebook_proto_rawDescData)
	})
	return file_determined_notebook_v1_notebook_proto_rawDescData
}

var file_determined_notebook_v1_notebook_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_determined_notebook_v1_notebook_proto_goTypes = []interface{}{
	(*Notebook)(nil),              // 0: determined.notebook.v1.Notebook
	(taskv1.State)(0),             // 1: determined.task.v1.State
	(*timestamp.Timestamp)(nil),   // 2: google.protobuf.Timestamp
	(*containerv1.Container)(nil), // 3: determined.container.v1.Container
}
var file_determined_notebook_v1_notebook_proto_depIdxs = []int32{
	1, // 0: determined.notebook.v1.Notebook.state:type_name -> determined.task.v1.State
	2, // 1: determined.notebook.v1.Notebook.start_time:type_name -> google.protobuf.Timestamp
	3, // 2: determined.notebook.v1.Notebook.container:type_name -> determined.container.v1.Container
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_determined_notebook_v1_notebook_proto_init() }
func file_determined_notebook_v1_notebook_proto_init() {
	if File_determined_notebook_v1_notebook_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_determined_notebook_v1_notebook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notebook); i {
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
			RawDescriptor: file_determined_notebook_v1_notebook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_determined_notebook_v1_notebook_proto_goTypes,
		DependencyIndexes: file_determined_notebook_v1_notebook_proto_depIdxs,
		MessageInfos:      file_determined_notebook_v1_notebook_proto_msgTypes,
	}.Build()
	File_determined_notebook_v1_notebook_proto = out.File
	file_determined_notebook_v1_notebook_proto_rawDesc = nil
	file_determined_notebook_v1_notebook_proto_goTypes = nil
	file_determined_notebook_v1_notebook_proto_depIdxs = nil
}
