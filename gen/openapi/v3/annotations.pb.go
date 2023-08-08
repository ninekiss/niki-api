// Copyright 2022 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: openapi/v3/annotations.proto

package openapi_v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_openapi_v3_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*Document)(nil),
		Field:         1143,
		Name:          "openapi.v3.document",
		Tag:           "bytes,1143,opt,name=document",
		Filename:      "openapi/v3/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Operation)(nil),
		Field:         1143,
		Name:          "openapi.v3.operation",
		Tag:           "bytes,1143,opt,name=operation",
		Filename:      "openapi/v3/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*Schema)(nil),
		Field:         1143,
		Name:          "openapi.v3.schema",
		Tag:           "bytes,1143,opt,name=schema",
		Filename:      "openapi/v3/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*Schema)(nil),
		Field:         1143,
		Name:          "openapi.v3.property",
		Tag:           "bytes,1143,opt,name=property",
		Filename:      "openapi/v3/annotations.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional openapi.v3.Document document = 1143;
	E_Document = &file_openapi_v3_annotations_proto_extTypes[0]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional openapi.v3.Operation operation = 1143;
	E_Operation = &file_openapi_v3_annotations_proto_extTypes[1]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional openapi.v3.Schema schema = 1143;
	E_Schema = &file_openapi_v3_annotations_proto_extTypes[2]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional openapi.v3.Schema property = 1143;
	E_Property = &file_openapi_v3_annotations_proto_extTypes[3]
)

var File_openapi_v3_annotations_proto protoreflect.FileDescriptor

var file_openapi_v3_annotations_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x33, 0x1a, 0x18, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x4f, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xf7, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x54, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf7, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x4c, 0x0a,
	0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf7, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x3a, 0x4e, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf7, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x42, 0x5a, 0x0a, 0x0e, 0x6f,
	0x72, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x33, 0x42, 0x10, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x33, 0x3b, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f, 0x76,
	0x33, 0xa2, 0x02, 0x03, 0x4f, 0x41, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_openapi_v3_annotations_proto_goTypes = []interface{}{
	(*descriptorpb.FileOptions)(nil),    // 0: google.protobuf.FileOptions
	(*descriptorpb.MethodOptions)(nil),  // 1: google.protobuf.MethodOptions
	(*descriptorpb.MessageOptions)(nil), // 2: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 3: google.protobuf.FieldOptions
	(*Document)(nil),                    // 4: openapi.v3.Document
	(*Operation)(nil),                   // 5: openapi.v3.Operation
	(*Schema)(nil),                      // 6: openapi.v3.Schema
}
var file_openapi_v3_annotations_proto_depIdxs = []int32{
	0, // 0: openapi.v3.document:extendee -> google.protobuf.FileOptions
	1, // 1: openapi.v3.operation:extendee -> google.protobuf.MethodOptions
	2, // 2: openapi.v3.schema:extendee -> google.protobuf.MessageOptions
	3, // 3: openapi.v3.property:extendee -> google.protobuf.FieldOptions
	4, // 4: openapi.v3.document:type_name -> openapi.v3.Document
	5, // 5: openapi.v3.operation:type_name -> openapi.v3.Operation
	6, // 6: openapi.v3.schema:type_name -> openapi.v3.Schema
	6, // 7: openapi.v3.property:type_name -> openapi.v3.Schema
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	4, // [4:8] is the sub-list for extension type_name
	0, // [0:4] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_openapi_v3_annotations_proto_init() }
func file_openapi_v3_annotations_proto_init() {
	if File_openapi_v3_annotations_proto != nil {
		return
	}
	file_openapi_v3_openapi_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_openapi_v3_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_openapi_v3_annotations_proto_goTypes,
		DependencyIndexes: file_openapi_v3_annotations_proto_depIdxs,
		ExtensionInfos:    file_openapi_v3_annotations_proto_extTypes,
	}.Build()
	File_openapi_v3_annotations_proto = out.File
	file_openapi_v3_annotations_proto_rawDesc = nil
	file_openapi_v3_annotations_proto_goTypes = nil
	file_openapi_v3_annotations_proto_depIdxs = nil
}
