// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: services/proto/reviewPB/review.proto

package reviewPB

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

type Review struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Review) Reset() {
	*x = Review{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_reviewPB_review_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_reviewPB_review_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_services_proto_reviewPB_review_proto_rawDescGZIP(), []int{0}
}

var File_services_proto_reviewPB_review_proto protoreflect.FileDescriptor

var file_services_proto_reviewPB_review_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x50, 0x42, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22, 0x08,
	0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x50, 0x42, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_proto_reviewPB_review_proto_rawDescOnce sync.Once
	file_services_proto_reviewPB_review_proto_rawDescData = file_services_proto_reviewPB_review_proto_rawDesc
)

func file_services_proto_reviewPB_review_proto_rawDescGZIP() []byte {
	file_services_proto_reviewPB_review_proto_rawDescOnce.Do(func() {
		file_services_proto_reviewPB_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_proto_reviewPB_review_proto_rawDescData)
	})
	return file_services_proto_reviewPB_review_proto_rawDescData
}

var file_services_proto_reviewPB_review_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_services_proto_reviewPB_review_proto_goTypes = []interface{}{
	(*Review)(nil), // 0: review.Review
}
var file_services_proto_reviewPB_review_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_proto_reviewPB_review_proto_init() }
func file_services_proto_reviewPB_review_proto_init() {
	if File_services_proto_reviewPB_review_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_proto_reviewPB_review_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Review); i {
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
			RawDescriptor: file_services_proto_reviewPB_review_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_proto_reviewPB_review_proto_goTypes,
		DependencyIndexes: file_services_proto_reviewPB_review_proto_depIdxs,
		MessageInfos:      file_services_proto_reviewPB_review_proto_msgTypes,
	}.Build()
	File_services_proto_reviewPB_review_proto = out.File
	file_services_proto_reviewPB_review_proto_rawDesc = nil
	file_services_proto_reviewPB_review_proto_goTypes = nil
	file_services_proto_reviewPB_review_proto_depIdxs = nil
}