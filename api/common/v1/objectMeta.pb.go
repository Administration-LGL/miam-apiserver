// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: common/v1/objectMeta.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type ObjectMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	InstanceID string                 `protobuf:"bytes,2,opt,name=instanceID,proto3" json:"instanceID,omitempty"`
	Name       string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Extend     map[string]*anypb.Any  `protobuf:"bytes,4,rep,name=extend,proto3" json:"extend,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreateAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createAt,proto3" json:"createAt,omitempty"`
	UpdateAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
}

func (x *ObjectMeta) Reset() {
	*x = ObjectMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_objectMeta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectMeta) ProtoMessage() {}

func (x *ObjectMeta) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_objectMeta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectMeta.ProtoReflect.Descriptor instead.
func (*ObjectMeta) Descriptor() ([]byte, []int) {
	return file_common_v1_objectMeta_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectMeta) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ObjectMeta) GetInstanceID() string {
	if x != nil {
		return x.InstanceID
	}
	return ""
}

func (x *ObjectMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ObjectMeta) GetExtend() map[string]*anypb.Any {
	if x != nil {
		return x.Extend
	}
	return nil
}

func (x *ObjectMeta) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

func (x *ObjectMeta) GetUpdateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateAt
	}
	return nil
}

type RequestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Reason  string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RequestReply) Reset() {
	*x = RequestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_objectMeta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestReply) ProtoMessage() {}

func (x *RequestReply) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_objectMeta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestReply.ProtoReflect.Descriptor instead.
func (*RequestReply) Descriptor() ([]byte, []int) {
	return file_common_v1_objectMeta_proto_rawDescGZIP(), []int{1}
}

func (x *RequestReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RequestReply) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *RequestReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64 `protobuf:"varint,1,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
}

func (x *ListMeta) Reset() {
	*x = ListMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_objectMeta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMeta) ProtoMessage() {}

func (x *ListMeta) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_objectMeta_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMeta.ProtoReflect.Descriptor instead.
func (*ListMeta) Descriptor() ([]byte, []int) {
	return file_common_v1_objectMeta_proto_rawDescGZIP(), []int{2}
}

func (x *ListMeta) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_common_v1_objectMeta_proto protoreflect.FileDescriptor

var file_common_v1_objectMeta_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcc, 0x02, 0x0a, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x12,
	0x36, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x1a,
	0x4f, 0x0a, 0x0b, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x54, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x42, 0x4c, 0x0a, 0x18, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0d,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a,
	0x1f, 0x6d, 0x69, 0x61, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_v1_objectMeta_proto_rawDescOnce sync.Once
	file_common_v1_objectMeta_proto_rawDescData = file_common_v1_objectMeta_proto_rawDesc
)

func file_common_v1_objectMeta_proto_rawDescGZIP() []byte {
	file_common_v1_objectMeta_proto_rawDescOnce.Do(func() {
		file_common_v1_objectMeta_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_v1_objectMeta_proto_rawDescData)
	})
	return file_common_v1_objectMeta_proto_rawDescData
}

var file_common_v1_objectMeta_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_common_v1_objectMeta_proto_goTypes = []interface{}{
	(*ObjectMeta)(nil),            // 0: common.v1.objectMeta
	(*RequestReply)(nil),          // 1: common.v1.requestReply
	(*ListMeta)(nil),              // 2: common.v1.ListMeta
	nil,                           // 3: common.v1.objectMeta.ExtendEntry
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 5: google.protobuf.Any
}
var file_common_v1_objectMeta_proto_depIdxs = []int32{
	3, // 0: common.v1.objectMeta.extend:type_name -> common.v1.objectMeta.ExtendEntry
	4, // 1: common.v1.objectMeta.createAt:type_name -> google.protobuf.Timestamp
	4, // 2: common.v1.objectMeta.updateAt:type_name -> google.protobuf.Timestamp
	5, // 3: common.v1.objectMeta.ExtendEntry.value:type_name -> google.protobuf.Any
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_common_v1_objectMeta_proto_init() }
func file_common_v1_objectMeta_proto_init() {
	if File_common_v1_objectMeta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_v1_objectMeta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectMeta); i {
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
		file_common_v1_objectMeta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestReply); i {
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
		file_common_v1_objectMeta_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMeta); i {
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
			RawDescriptor: file_common_v1_objectMeta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_v1_objectMeta_proto_goTypes,
		DependencyIndexes: file_common_v1_objectMeta_proto_depIdxs,
		MessageInfos:      file_common_v1_objectMeta_proto_msgTypes,
	}.Build()
	File_common_v1_objectMeta_proto = out.File
	file_common_v1_objectMeta_proto_rawDesc = nil
	file_common_v1_objectMeta_proto_goTypes = nil
	file_common_v1_objectMeta_proto_depIdxs = nil
}
