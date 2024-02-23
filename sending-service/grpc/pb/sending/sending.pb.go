// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: grpc/pb/sending/sending.proto

package sending_proto

import (
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

type CreateJobAssignmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListJobId []string `protobuf:"bytes,1,rep,name=list_job_id,json=listJobId,proto3" json:"list_job_id,omitempty"`
	CreatorId string   `protobuf:"bytes,2,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
}

func (x *CreateJobAssignmentRequest) Reset() {
	*x = CreateJobAssignmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_pb_sending_sending_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobAssignmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobAssignmentRequest) ProtoMessage() {}

func (x *CreateJobAssignmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_pb_sending_sending_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobAssignmentRequest.ProtoReflect.Descriptor instead.
func (*CreateJobAssignmentRequest) Descriptor() ([]byte, []int) {
	return file_grpc_pb_sending_sending_proto_rawDescGZIP(), []int{0}
}

func (x *CreateJobAssignmentRequest) GetListJobId() []string {
	if x != nil {
		return x.ListJobId
	}
	return nil
}

func (x *CreateJobAssignmentRequest) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

type CreateJobAssignmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateJobAssignmentResponse) Reset() {
	*x = CreateJobAssignmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_pb_sending_sending_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobAssignmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobAssignmentResponse) ProtoMessage() {}

func (x *CreateJobAssignmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_pb_sending_sending_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobAssignmentResponse.ProtoReflect.Descriptor instead.
func (*CreateJobAssignmentResponse) Descriptor() ([]byte, []int) {
	return file_grpc_pb_sending_sending_proto_rawDescGZIP(), []int{1}
}

func (x *CreateJobAssignmentResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CreateJobAssignmentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_pb_sending_sending_proto protoreflect.FileDescriptor

var file_grpc_pb_sending_sending_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x5b, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6a,
	0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x73,
	0x74, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a,
	0x6f, 0x62, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x71, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x70, 0x63, 0x12, 0x62,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_pb_sending_sending_proto_rawDescOnce sync.Once
	file_grpc_pb_sending_sending_proto_rawDescData = file_grpc_pb_sending_sending_proto_rawDesc
)

func file_grpc_pb_sending_sending_proto_rawDescGZIP() []byte {
	file_grpc_pb_sending_sending_proto_rawDescOnce.Do(func() {
		file_grpc_pb_sending_sending_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_pb_sending_sending_proto_rawDescData)
	})
	return file_grpc_pb_sending_sending_proto_rawDescData
}

var file_grpc_pb_sending_sending_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_pb_sending_sending_proto_goTypes = []interface{}{
	(*CreateJobAssignmentRequest)(nil),  // 0: sending.CreateJobAssignmentRequest
	(*CreateJobAssignmentResponse)(nil), // 1: sending.CreateJobAssignmentResponse
}
var file_grpc_pb_sending_sending_proto_depIdxs = []int32{
	0, // 0: sending.PricingGrpc.CreateJobAssignment:input_type -> sending.CreateJobAssignmentRequest
	1, // 1: sending.PricingGrpc.CreateJobAssignment:output_type -> sending.CreateJobAssignmentResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_pb_sending_sending_proto_init() }
func file_grpc_pb_sending_sending_proto_init() {
	if File_grpc_pb_sending_sending_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_pb_sending_sending_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobAssignmentRequest); i {
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
		file_grpc_pb_sending_sending_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobAssignmentResponse); i {
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
			RawDescriptor: file_grpc_pb_sending_sending_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_pb_sending_sending_proto_goTypes,
		DependencyIndexes: file_grpc_pb_sending_sending_proto_depIdxs,
		MessageInfos:      file_grpc_pb_sending_sending_proto_msgTypes,
	}.Build()
	File_grpc_pb_sending_sending_proto = out.File
	file_grpc_pb_sending_sending_proto_rawDesc = nil
	file_grpc_pb_sending_sending_proto_goTypes = nil
	file_grpc_pb_sending_sending_proto_depIdxs = nil
}
