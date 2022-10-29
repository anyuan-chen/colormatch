// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: images/v1/images.proto

package imagesv1

import (
	v1 "github.com/anyuan-chen/colormatch/gen/proto/go/shared/v1"
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

type GetBackgroundColorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image   *v1.Image   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Palette *v1.Palette `protobuf:"bytes,2,opt,name=palette,proto3" json:"palette,omitempty"`
}

func (x *GetBackgroundColorRequest) Reset() {
	*x = GetBackgroundColorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_images_v1_images_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBackgroundColorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBackgroundColorRequest) ProtoMessage() {}

func (x *GetBackgroundColorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_images_v1_images_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBackgroundColorRequest.ProtoReflect.Descriptor instead.
func (*GetBackgroundColorRequest) Descriptor() ([]byte, []int) {
	return file_images_v1_images_proto_rawDescGZIP(), []int{0}
}

func (x *GetBackgroundColorRequest) GetImage() *v1.Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *GetBackgroundColorRequest) GetPalette() *v1.Palette {
	if x != nil {
		return x.Palette
	}
	return nil
}

type GetBackgroundColorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color *v1.Color `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *GetBackgroundColorResponse) Reset() {
	*x = GetBackgroundColorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_images_v1_images_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBackgroundColorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBackgroundColorResponse) ProtoMessage() {}

func (x *GetBackgroundColorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_images_v1_images_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBackgroundColorResponse.ProtoReflect.Descriptor instead.
func (*GetBackgroundColorResponse) Descriptor() ([]byte, []int) {
	return file_images_v1_images_proto_rawDescGZIP(), []int{1}
}

func (x *GetBackgroundColorResponse) GetColor() *v1.Color {
	if x != nil {
		return x.Color
	}
	return nil
}

type GetHighlightColorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image   *v1.Image   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Palette *v1.Palette `protobuf:"bytes,2,opt,name=palette,proto3" json:"palette,omitempty"`
}

func (x *GetHighlightColorRequest) Reset() {
	*x = GetHighlightColorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_images_v1_images_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHighlightColorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHighlightColorRequest) ProtoMessage() {}

func (x *GetHighlightColorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_images_v1_images_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHighlightColorRequest.ProtoReflect.Descriptor instead.
func (*GetHighlightColorRequest) Descriptor() ([]byte, []int) {
	return file_images_v1_images_proto_rawDescGZIP(), []int{2}
}

func (x *GetHighlightColorRequest) GetImage() *v1.Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *GetHighlightColorRequest) GetPalette() *v1.Palette {
	if x != nil {
		return x.Palette
	}
	return nil
}

type GetHighlightColorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color *v1.Color `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *GetHighlightColorResponse) Reset() {
	*x = GetHighlightColorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_images_v1_images_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHighlightColorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHighlightColorResponse) ProtoMessage() {}

func (x *GetHighlightColorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_images_v1_images_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHighlightColorResponse.ProtoReflect.Descriptor instead.
func (*GetHighlightColorResponse) Descriptor() ([]byte, []int) {
	return file_images_v1_images_proto_rawDescGZIP(), []int{3}
}

func (x *GetHighlightColorResponse) GetColor() *v1.Color {
	if x != nil {
		return x.Color
	}
	return nil
}

var File_images_v1_images_proto protoreflect.FileDescriptor

var file_images_v1_images_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x14, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2c,
	0x0a, 0x07, 0x70, 0x61, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x6c, 0x65,
	0x74, 0x74, 0x65, 0x52, 0x07, 0x70, 0x61, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x22, 0x44, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x22, 0x70, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x48, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x70, 0x61, 0x6c, 0x65, 0x74, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x52, 0x07, 0x70, 0x61, 0x6c,
	0x65, 0x74, 0x74, 0x65, 0x22, 0x43, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x48, 0x69, 0x67, 0x68, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x32, 0xdb, 0x01, 0x0a, 0x16, 0x50, 0x61,
	0x6c, 0x65, 0x74, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x24, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x69,
	0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x23, 0x2e, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x67, 0x68,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x48, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x79, 0x75, 0x61, 0x6e, 0x2d, 0x63, 0x68, 0x65,
	0x6e, 0x2f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_images_v1_images_proto_rawDescOnce sync.Once
	file_images_v1_images_proto_rawDescData = file_images_v1_images_proto_rawDesc
)

func file_images_v1_images_proto_rawDescGZIP() []byte {
	file_images_v1_images_proto_rawDescOnce.Do(func() {
		file_images_v1_images_proto_rawDescData = protoimpl.X.CompressGZIP(file_images_v1_images_proto_rawDescData)
	})
	return file_images_v1_images_proto_rawDescData
}

var file_images_v1_images_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_images_v1_images_proto_goTypes = []interface{}{
	(*GetBackgroundColorRequest)(nil),  // 0: images.v1.GetBackgroundColorRequest
	(*GetBackgroundColorResponse)(nil), // 1: images.v1.GetBackgroundColorResponse
	(*GetHighlightColorRequest)(nil),   // 2: images.v1.GetHighlightColorRequest
	(*GetHighlightColorResponse)(nil),  // 3: images.v1.GetHighlightColorResponse
	(*v1.Image)(nil),                   // 4: shared.v1.Image
	(*v1.Palette)(nil),                 // 5: shared.v1.Palette
	(*v1.Color)(nil),                   // 6: shared.v1.Color
}
var file_images_v1_images_proto_depIdxs = []int32{
	4, // 0: images.v1.GetBackgroundColorRequest.image:type_name -> shared.v1.Image
	5, // 1: images.v1.GetBackgroundColorRequest.palette:type_name -> shared.v1.Palette
	6, // 2: images.v1.GetBackgroundColorResponse.color:type_name -> shared.v1.Color
	4, // 3: images.v1.GetHighlightColorRequest.image:type_name -> shared.v1.Image
	5, // 4: images.v1.GetHighlightColorRequest.palette:type_name -> shared.v1.Palette
	6, // 5: images.v1.GetHighlightColorResponse.color:type_name -> shared.v1.Color
	0, // 6: images.v1.PaletteMatchingService.GetBackgroundColor:input_type -> images.v1.GetBackgroundColorRequest
	2, // 7: images.v1.PaletteMatchingService.GetHighlightColor:input_type -> images.v1.GetHighlightColorRequest
	1, // 8: images.v1.PaletteMatchingService.GetBackgroundColor:output_type -> images.v1.GetBackgroundColorResponse
	3, // 9: images.v1.PaletteMatchingService.GetHighlightColor:output_type -> images.v1.GetHighlightColorResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_images_v1_images_proto_init() }
func file_images_v1_images_proto_init() {
	if File_images_v1_images_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_images_v1_images_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBackgroundColorRequest); i {
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
		file_images_v1_images_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBackgroundColorResponse); i {
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
		file_images_v1_images_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHighlightColorRequest); i {
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
		file_images_v1_images_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHighlightColorResponse); i {
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
			RawDescriptor: file_images_v1_images_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_images_v1_images_proto_goTypes,
		DependencyIndexes: file_images_v1_images_proto_depIdxs,
		MessageInfos:      file_images_v1_images_proto_msgTypes,
	}.Build()
	File_images_v1_images_proto = out.File
	file_images_v1_images_proto_rawDesc = nil
	file_images_v1_images_proto_goTypes = nil
	file_images_v1_images_proto_depIdxs = nil
}
