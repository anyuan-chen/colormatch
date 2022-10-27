// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: shared/v1/main.proto

package sharedv1

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

type Color struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HexCode string `protobuf:"bytes,1,opt,name=hex_code,json=hexCode,proto3" json:"hex_code,omitempty"`
}

func (x *Color) Reset() {
	*x = Color{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_v1_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Color) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Color) ProtoMessage() {}

func (x *Color) ProtoReflect() protoreflect.Message {
	mi := &file_shared_v1_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Color.ProtoReflect.Descriptor instead.
func (*Color) Descriptor() ([]byte, []int) {
	return file_shared_v1_main_proto_rawDescGZIP(), []int{0}
}

func (x *Color) GetHexCode() string {
	if x != nil {
		return x.HexCode
	}
	return ""
}

type Palette struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color []*Color `protobuf:"bytes,1,rep,name=color,proto3" json:"color,omitempty"`
}

func (x *Palette) Reset() {
	*x = Palette{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_v1_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Palette) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Palette) ProtoMessage() {}

func (x *Palette) ProtoReflect() protoreflect.Message {
	mi := &file_shared_v1_main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Palette.ProtoReflect.Descriptor instead.
func (*Palette) Descriptor() ([]byte, []int) {
	return file_shared_v1_main_proto_rawDescGZIP(), []int{1}
}

func (x *Palette) GetColor() []*Color {
	if x != nil {
		return x.Color
	}
	return nil
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width     int32  `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height    int32  `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	ImageData []byte `protobuf:"bytes,3,opt,name=image_data,json=imageData,proto3" json:"image_data,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_v1_main_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_shared_v1_main_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_shared_v1_main_proto_rawDescGZIP(), []int{2}
}

func (x *Image) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Image) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Image) GetImageData() []byte {
	if x != nil {
		return x.ImageData
	}
	return nil
}

var File_shared_v1_main_proto protoreflect.FileDescriptor

var file_shared_v1_main_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76,
	0x31, 0x22, 0x22, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x65,
	0x78, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x65,
	0x78, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x31, 0x0a, 0x07, 0x50, 0x61, 0x6c, 0x65, 0x74, 0x74, 0x65,
	0x12, 0x26, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x6f,
	0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x54, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x43,
	0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x79,
	0x75, 0x61, 0x6e, 0x2d, 0x63, 0x68, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shared_v1_main_proto_rawDescOnce sync.Once
	file_shared_v1_main_proto_rawDescData = file_shared_v1_main_proto_rawDesc
)

func file_shared_v1_main_proto_rawDescGZIP() []byte {
	file_shared_v1_main_proto_rawDescOnce.Do(func() {
		file_shared_v1_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_v1_main_proto_rawDescData)
	})
	return file_shared_v1_main_proto_rawDescData
}

var file_shared_v1_main_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_shared_v1_main_proto_goTypes = []interface{}{
	(*Color)(nil),   // 0: shared.v1.Color
	(*Palette)(nil), // 1: shared.v1.Palette
	(*Image)(nil),   // 2: shared.v1.Image
}
var file_shared_v1_main_proto_depIdxs = []int32{
	0, // 0: shared.v1.Palette.color:type_name -> shared.v1.Color
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shared_v1_main_proto_init() }
func file_shared_v1_main_proto_init() {
	if File_shared_v1_main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shared_v1_main_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Color); i {
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
		file_shared_v1_main_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Palette); i {
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
		file_shared_v1_main_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
			RawDescriptor: file_shared_v1_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shared_v1_main_proto_goTypes,
		DependencyIndexes: file_shared_v1_main_proto_depIdxs,
		MessageInfos:      file_shared_v1_main_proto_msgTypes,
	}.Build()
	File_shared_v1_main_proto = out.File
	file_shared_v1_main_proto_rawDesc = nil
	file_shared_v1_main_proto_goTypes = nil
	file_shared_v1_main_proto_depIdxs = nil
}