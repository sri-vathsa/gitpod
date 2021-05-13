// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.5
// source: imagespec.proto

package api

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

// ImageSpec configures the image one wishes to pull from a registry facade
type ImageSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// base_ref points to an image in another registry
	BaseRef string `protobuf:"bytes,1,opt,name=base_ref,json=baseRef,proto3" json:"base_ref,omitempty"`
	// ide_ref points to an image denoting the IDE to use
	IdeRef string `protobuf:"bytes,2,opt,name=ide_ref,json=ideRef,proto3" json:"ide_ref,omitempty"`
	// content_layer describe the last few layers which provide the workspace's content
	ContentLayer []*ContentLayer `protobuf:"bytes,3,rep,name=content_layer,json=contentLayer,proto3" json:"content_layer,omitempty"`
}

func (x *ImageSpec) Reset() {
	*x = ImageSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imagespec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageSpec) ProtoMessage() {}

func (x *ImageSpec) ProtoReflect() protoreflect.Message {
	mi := &file_imagespec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageSpec.ProtoReflect.Descriptor instead.
func (*ImageSpec) Descriptor() ([]byte, []int) {
	return file_imagespec_proto_rawDescGZIP(), []int{0}
}

func (x *ImageSpec) GetBaseRef() string {
	if x != nil {
		return x.BaseRef
	}
	return ""
}

func (x *ImageSpec) GetIdeRef() string {
	if x != nil {
		return x.IdeRef
	}
	return ""
}

func (x *ImageSpec) GetContentLayer() []*ContentLayer {
	if x != nil {
		return x.ContentLayer
	}
	return nil
}

// ContentLayer is a layer that provides a workspace's content
type ContentLayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Spec:
	//	*ContentLayer_Remote
	//	*ContentLayer_Direct
	Spec isContentLayer_Spec `protobuf_oneof:"spec"`
}

func (x *ContentLayer) Reset() {
	*x = ContentLayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imagespec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentLayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentLayer) ProtoMessage() {}

func (x *ContentLayer) ProtoReflect() protoreflect.Message {
	mi := &file_imagespec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentLayer.ProtoReflect.Descriptor instead.
func (*ContentLayer) Descriptor() ([]byte, []int) {
	return file_imagespec_proto_rawDescGZIP(), []int{1}
}

func (m *ContentLayer) GetSpec() isContentLayer_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (x *ContentLayer) GetRemote() *RemoteContentLayer {
	if x, ok := x.GetSpec().(*ContentLayer_Remote); ok {
		return x.Remote
	}
	return nil
}

func (x *ContentLayer) GetDirect() *DirectContentLayer {
	if x, ok := x.GetSpec().(*ContentLayer_Direct); ok {
		return x.Direct
	}
	return nil
}

type isContentLayer_Spec interface {
	isContentLayer_Spec()
}

type ContentLayer_Remote struct {
	Remote *RemoteContentLayer `protobuf:"bytes,1,opt,name=remote,proto3,oneof"`
}

type ContentLayer_Direct struct {
	Direct *DirectContentLayer `protobuf:"bytes,2,opt,name=direct,proto3,oneof"`
}

func (*ContentLayer_Remote) isContentLayer_Spec() {}

func (*ContentLayer_Direct) isContentLayer_Spec() {}

// RemoteContentLayer is a layer which can be downloaded from a remote URL.
// If the diff_id is empty or equals the digest the layer is expected to be uncompressed.
type RemoteContentLayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// url points to the actual content location. This must be a valid HTTPS URL pointing
	// to a tar.gz file.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// digest is the digest (content hash) of the file the URL points to.
	Digest string `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	// diff_id is the digest (content hash) of the uncompressed data the URL points to if the
	// the URL points to a compressed file. If the file is uncompressed to begin with this field
	// can either be empty or the same as digest.
	DiffId string `protobuf:"bytes,3,opt,name=diff_id,json=diffId,proto3" json:"diff_id,omitempty"`
	// media_type is the content type of the layer and is expected to be one of:
	//  application/vnd.oci.image.layer.v1.tar
	//  application/vnd.oci.image.layer.v1.tar+gzip
	//  application/vnd.oci.image.layer.v1.tar+zstd
	MediaType string `protobuf:"bytes,4,opt,name=media_type,json=mediaType,proto3" json:"media_type,omitempty"`
	// size is the size of the layer download in bytes
	Size int64 `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *RemoteContentLayer) Reset() {
	*x = RemoteContentLayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imagespec_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteContentLayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteContentLayer) ProtoMessage() {}

func (x *RemoteContentLayer) ProtoReflect() protoreflect.Message {
	mi := &file_imagespec_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteContentLayer.ProtoReflect.Descriptor instead.
func (*RemoteContentLayer) Descriptor() ([]byte, []int) {
	return file_imagespec_proto_rawDescGZIP(), []int{2}
}

func (x *RemoteContentLayer) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *RemoteContentLayer) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *RemoteContentLayer) GetDiffId() string {
	if x != nil {
		return x.DiffId
	}
	return ""
}

func (x *RemoteContentLayer) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *RemoteContentLayer) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

// DirectContentLayer is an uncompressed tar file which is directly added as layer
type DirectContentLayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the bytes of the uncompressed tar file which is served as layer
	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *DirectContentLayer) Reset() {
	*x = DirectContentLayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imagespec_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectContentLayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectContentLayer) ProtoMessage() {}

func (x *DirectContentLayer) ProtoReflect() protoreflect.Message {
	mi := &file_imagespec_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectContentLayer.ProtoReflect.Descriptor instead.
func (*DirectContentLayer) Descriptor() ([]byte, []int) {
	return file_imagespec_proto_rawDescGZIP(), []int{3}
}

func (x *DirectContentLayer) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_imagespec_proto protoreflect.FileDescriptor

var file_imagespec_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x66, 0x61, 0x63, 0x61, 0x64,
	0x65, 0x22, 0x82, 0x01, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x19, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x66, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64,
	0x65, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x65,
	0x52, 0x65, 0x66, 0x12, 0x41, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x66, 0x61, 0x63, 0x61, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x92, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x66, 0x61, 0x63, 0x61, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x66, 0x61, 0x63, 0x61, 0x64, 0x65, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x8a, 0x01, 0x0a, 0x12,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x64, 0x69, 0x66, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x69, 0x66, 0x66, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2d, 0x69, 0x6f,
	0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2d, 0x66, 0x61, 0x63, 0x61, 0x64, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_imagespec_proto_rawDescOnce sync.Once
	file_imagespec_proto_rawDescData = file_imagespec_proto_rawDesc
)

func file_imagespec_proto_rawDescGZIP() []byte {
	file_imagespec_proto_rawDescOnce.Do(func() {
		file_imagespec_proto_rawDescData = protoimpl.X.CompressGZIP(file_imagespec_proto_rawDescData)
	})
	return file_imagespec_proto_rawDescData
}

var file_imagespec_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_imagespec_proto_goTypes = []interface{}{
	(*ImageSpec)(nil),          // 0: registryfacade.ImageSpec
	(*ContentLayer)(nil),       // 1: registryfacade.ContentLayer
	(*RemoteContentLayer)(nil), // 2: registryfacade.RemoteContentLayer
	(*DirectContentLayer)(nil), // 3: registryfacade.DirectContentLayer
}
var file_imagespec_proto_depIdxs = []int32{
	1, // 0: registryfacade.ImageSpec.content_layer:type_name -> registryfacade.ContentLayer
	2, // 1: registryfacade.ContentLayer.remote:type_name -> registryfacade.RemoteContentLayer
	3, // 2: registryfacade.ContentLayer.direct:type_name -> registryfacade.DirectContentLayer
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_imagespec_proto_init() }
func file_imagespec_proto_init() {
	if File_imagespec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_imagespec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageSpec); i {
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
		file_imagespec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentLayer); i {
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
		file_imagespec_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteContentLayer); i {
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
		file_imagespec_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectContentLayer); i {
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
	file_imagespec_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ContentLayer_Remote)(nil),
		(*ContentLayer_Direct)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_imagespec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_imagespec_proto_goTypes,
		DependencyIndexes: file_imagespec_proto_depIdxs,
		MessageInfos:      file_imagespec_proto_msgTypes,
	}.Build()
	File_imagespec_proto = out.File
	file_imagespec_proto_rawDesc = nil
	file_imagespec_proto_goTypes = nil
	file_imagespec_proto_depIdxs = nil
}
