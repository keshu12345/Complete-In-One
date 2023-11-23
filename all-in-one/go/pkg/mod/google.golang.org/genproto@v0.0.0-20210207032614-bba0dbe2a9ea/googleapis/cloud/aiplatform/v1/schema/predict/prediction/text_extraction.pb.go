// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: google/cloud/aiplatform/v1/schema/predict/prediction/text_extraction.proto

package prediction

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Prediction output format for Text Extraction.
type TextExtractionPredictionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource IDs of the AnnotationSpecs that had been identified,
	// ordered by the confidence score descendingly.
	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	// The display names of the AnnotationSpecs that had been identified,
	// order matches the IDs.
	DisplayNames []string `protobuf:"bytes,2,rep,name=display_names,json=displayNames,proto3" json:"display_names,omitempty"`
	// The start offsets, inclusive, of the text segment in which the
	// AnnotationSpec has been identified. Expressed as a zero-based number
	// of characters as measured from the start of the text snippet.
	TextSegmentStartOffsets []int64 `protobuf:"varint,3,rep,packed,name=text_segment_start_offsets,json=textSegmentStartOffsets,proto3" json:"text_segment_start_offsets,omitempty"`
	// The end offsets, inclusive, of the text segment in which the
	// AnnotationSpec has been identified. Expressed as a zero-based number
	// of characters as measured from the start of the text snippet.
	TextSegmentEndOffsets []int64 `protobuf:"varint,4,rep,packed,name=text_segment_end_offsets,json=textSegmentEndOffsets,proto3" json:"text_segment_end_offsets,omitempty"`
	// The Model's confidences in correctness of the predicted IDs, higher
	// value means higher confidence. Order matches the Ids.
	Confidences []float32 `protobuf:"fixed32,5,rep,packed,name=confidences,proto3" json:"confidences,omitempty"`
}

func (x *TextExtractionPredictionResult) Reset() {
	*x = TextExtractionPredictionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextExtractionPredictionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextExtractionPredictionResult) ProtoMessage() {}

func (x *TextExtractionPredictionResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextExtractionPredictionResult.ProtoReflect.Descriptor instead.
func (*TextExtractionPredictionResult) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_rawDescGZIP(), []int{0}
}

func (x *TextExtractionPredictionResult) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *TextExtractionPredictionResult) GetDisplayNames() []string {
	if x != nil {
		return x.DisplayNames
	}
	return nil
}

func (x *TextExtractionPredictionResult) GetTextSegmentStartOffsets() []int64 {
	if x != nil {
		return x.TextSegmentStartOffsets
	}
	return nil
}

func (x *TextExtractionPredictionResult) GetTextSegmentEndOffsets() []int64 {
	if x != nil {
		return x.TextSegmentEndOffsets
	}
	return nil
}

func (x *TextExtractionPredictionResult) GetConfidences() []float32 {
	if x != nil {
		return x.Confidences
	}
	return nil
}

var File_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x34, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xef, 0x01, 0x0a, 0x1e, 0x54, 0x65, 0x78, 0x74, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x1a, 0x74, 0x65,
	0x78, 0x74, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x17,
	0x74, 0x65, 0x78, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x18, 0x74, 0x65, 0x78, 0x74, 0x5f,
	0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x15, 0x74, 0x65, 0x78, 0x74, 0x53,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x42, 0xc1, 0x01, 0x0a, 0x38, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x23, 0x54, 0x65, 0x78, 0x74, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_rawDescData = file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_goTypes = []interface{}{
	(*TextExtractionPredictionResult)(nil), // 0: google.cloud.aiplatform.v1.schema.predict.prediction.TextExtractionPredictionResult
}
var file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_init() }
func file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_init() {
	if File_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextExtractionPredictionResult); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto = out.File
	file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_schema_predict_prediction_text_extraction_proto_depIdxs = nil
}
