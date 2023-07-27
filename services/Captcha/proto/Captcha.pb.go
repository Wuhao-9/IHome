// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.22.5
// source: Captcha.proto

package Captcha

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

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Captcha_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Captcha_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_Captcha_proto_rawDescGZIP(), []int{0}
}

type Img struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaptchaImage []byte `protobuf:"bytes,1,opt,name=CaptchaImage,proto3" json:"CaptchaImage,omitempty"`
}

func (x *Img) Reset() {
	*x = Img{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Captcha_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Img) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Img) ProtoMessage() {}

func (x *Img) ProtoReflect() protoreflect.Message {
	mi := &file_Captcha_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Img.ProtoReflect.Descriptor instead.
func (*Img) Descriptor() ([]byte, []int) {
	return file_Captcha_proto_rawDescGZIP(), []int{1}
}

func (x *Img) GetCaptchaImage() []byte {
	if x != nil {
		return x.CaptchaImage
	}
	return nil
}

var File_Captcha_proto protoreflect.FileDescriptor

var file_Captcha_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x22, 0x0e, 0x0a, 0x0c, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x12,
	0x22, 0x0a, 0x0c, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x32, 0x3c, 0x0a, 0x07, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x31,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x15, 0x2e, 0x43,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x69, 0x6d,
	0x67, 0x42, 0x11, 0x5a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x43, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Captcha_proto_rawDescOnce sync.Once
	file_Captcha_proto_rawDescData = file_Captcha_proto_rawDesc
)

func file_Captcha_proto_rawDescGZIP() []byte {
	file_Captcha_proto_rawDescOnce.Do(func() {
		file_Captcha_proto_rawDescData = protoimpl.X.CompressGZIP(file_Captcha_proto_rawDescData)
	})
	return file_Captcha_proto_rawDescData
}

var file_Captcha_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Captcha_proto_goTypes = []interface{}{
	(*EmptyRequest)(nil), // 0: Captcha.emptyRequest
	(*Img)(nil),          // 1: Captcha.img
}
var file_Captcha_proto_depIdxs = []int32{
	0, // 0: Captcha.Captcha.GetCaptcha:input_type -> Captcha.emptyRequest
	1, // 1: Captcha.Captcha.GetCaptcha:output_type -> Captcha.img
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Captcha_proto_init() }
func file_Captcha_proto_init() {
	if File_Captcha_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Captcha_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_Captcha_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Img); i {
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
			RawDescriptor: file_Captcha_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Captcha_proto_goTypes,
		DependencyIndexes: file_Captcha_proto_depIdxs,
		MessageInfos:      file_Captcha_proto_msgTypes,
	}.Build()
	File_Captcha_proto = out.File
	file_Captcha_proto_rawDesc = nil
	file_Captcha_proto_goTypes = nil
	file_Captcha_proto_depIdxs = nil
}
