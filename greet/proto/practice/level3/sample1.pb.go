// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: sample1.proto

package level3

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

type SampleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsActive bool     `protobuf:"varint,2,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Name     string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	AltNames []string `protobuf:"bytes,4,rep,name=alt_names,json=altNames,proto3" json:"alt_names,omitempty"`
}

func (x *SampleMessage) Reset() {
	*x = SampleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleMessage) ProtoMessage() {}

func (x *SampleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_sample1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleMessage.ProtoReflect.Descriptor instead.
func (*SampleMessage) Descriptor() ([]byte, []int) {
	return file_sample1_proto_rawDescGZIP(), []int{0}
}

func (x *SampleMessage) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SampleMessage) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *SampleMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SampleMessage) GetAltNames() []string {
	if x != nil {
		return x.AltNames
	}
	return nil
}

type ComplexMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Single   *SampleMessage   `protobuf:"bytes,1,opt,name=single,proto3" json:"single,omitempty"`
	Multiple []*SampleMessage `protobuf:"bytes,2,rep,name=multiple,proto3" json:"multiple,omitempty"`
}

func (x *ComplexMessage) Reset() {
	*x = ComplexMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexMessage) ProtoMessage() {}

func (x *ComplexMessage) ProtoReflect() protoreflect.Message {
	mi := &file_sample1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexMessage.ProtoReflect.Descriptor instead.
func (*ComplexMessage) Descriptor() ([]byte, []int) {
	return file_sample1_proto_rawDescGZIP(), []int{1}
}

func (x *ComplexMessage) GetSingle() *SampleMessage {
	if x != nil {
		return x.Single
	}
	return nil
}

func (x *ComplexMessage) GetMultiple() []*SampleMessage {
	if x != nil {
		return x.Multiple
	}
	return nil
}

var File_sample1_proto protoreflect.FileDescriptor

var file_sample1_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x2e,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x33, 0x22, 0x6d, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6c, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6c, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x2e, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x33,
	0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x06,
	0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x2e, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x33,
	0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6d, 0x6f, 0x67, 0x68, 0x32, 0x30, 0x31, 0x39,
	0x2f, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x5f, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x2f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x33, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sample1_proto_rawDescOnce sync.Once
	file_sample1_proto_rawDescData = file_sample1_proto_rawDesc
)

func file_sample1_proto_rawDescGZIP() []byte {
	file_sample1_proto_rawDescOnce.Do(func() {
		file_sample1_proto_rawDescData = protoimpl.X.CompressGZIP(file_sample1_proto_rawDescData)
	})
	return file_sample1_proto_rawDescData
}

var file_sample1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sample1_proto_goTypes = []interface{}{
	(*SampleMessage)(nil),  // 0: greet.practice.level3.SampleMessage
	(*ComplexMessage)(nil), // 1: greet.practice.level3.ComplexMessage
}
var file_sample1_proto_depIdxs = []int32{
	0, // 0: greet.practice.level3.ComplexMessage.single:type_name -> greet.practice.level3.SampleMessage
	0, // 1: greet.practice.level3.ComplexMessage.multiple:type_name -> greet.practice.level3.SampleMessage
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sample1_proto_init() }
func file_sample1_proto_init() {
	if File_sample1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sample1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleMessage); i {
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
		file_sample1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexMessage); i {
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
			RawDescriptor: file_sample1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sample1_proto_goTypes,
		DependencyIndexes: file_sample1_proto_depIdxs,
		MessageInfos:      file_sample1_proto_msgTypes,
	}.Build()
	File_sample1_proto = out.File
	file_sample1_proto_rawDesc = nil
	file_sample1_proto_goTypes = nil
	file_sample1_proto_depIdxs = nil
}