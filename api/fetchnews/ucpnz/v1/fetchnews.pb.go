// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: api/fetchnews/ucpnz/v1/fetchnews.proto

package v1

import (
	v1 "github.com/hi20160616/fetchnews-api/proto/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_fetchnews_ucpnz_v1_fetchnews_proto protoreflect.FileDescriptor

var file_api_fetchnews_ucpnz_v1_fetchnews_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x6e, 0x65, 0x77, 0x73, 0x2f,
	0x75, 0x63, 0x70, 0x6e, 0x7a, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x6e, 0x65,
	0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x66, 0x65, 0x74, 0x63, 0x68, 0x6e,
	0x65, 0x77, 0x73, 0x2e, 0x75, 0x63, 0x70, 0x6e, 0x7a, 0x2e, 0x76, 0x31, 0x1a, 0x3c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x32, 0x30, 0x31, 0x36, 0x30,
	0x36, 0x31, 0x36, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x6e, 0x65, 0x77, 0x73, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68,
	0x6e, 0x65, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x87, 0x02, 0x0a, 0x05, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x12, 0x57, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x6e, 0x65, 0x77, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x6e,
	0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x66, 0x65,
	0x74, 0x63, 0x68, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x66,
	0x65, 0x74, 0x63, 0x68, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x6e,
	0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x66,
	0x65, 0x74, 0x63, 0x68, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x32, 0x30, 0x31, 0x36, 0x30, 0x36, 0x31, 0x36, 0x2f, 0x6d, 0x73,
	0x2d, 0x75, 0x63, 0x70, 0x6e, 0x7a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68,
	0x6e, 0x65, 0x77, 0x73, 0x2f, 0x75, 0x63, 0x70, 0x6e, 0x7a, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_fetchnews_ucpnz_v1_fetchnews_proto_goTypes = []interface{}{
	(*v1.ListArticlesRequest)(nil),    // 0: fetchnews.v1.ListArticlesRequest
	(*v1.GetArticleRequest)(nil),      // 1: fetchnews.v1.GetArticleRequest
	(*v1.SearchArticlesRequest)(nil),  // 2: fetchnews.v1.SearchArticlesRequest
	(*v1.ListArticlesResponse)(nil),   // 3: fetchnews.v1.ListArticlesResponse
	(*v1.Article)(nil),                // 4: fetchnews.v1.Article
	(*v1.SearchArticlesResponse)(nil), // 5: fetchnews.v1.SearchArticlesResponse
}
var file_api_fetchnews_ucpnz_v1_fetchnews_proto_depIdxs = []int32{
	0, // 0: fetchnews.ucpnz.v1.Fetch.ListArticles:input_type -> fetchnews.v1.ListArticlesRequest
	1, // 1: fetchnews.ucpnz.v1.Fetch.GetArticle:input_type -> fetchnews.v1.GetArticleRequest
	2, // 2: fetchnews.ucpnz.v1.Fetch.SearchArticles:input_type -> fetchnews.v1.SearchArticlesRequest
	3, // 3: fetchnews.ucpnz.v1.Fetch.ListArticles:output_type -> fetchnews.v1.ListArticlesResponse
	4, // 4: fetchnews.ucpnz.v1.Fetch.GetArticle:output_type -> fetchnews.v1.Article
	5, // 5: fetchnews.ucpnz.v1.Fetch.SearchArticles:output_type -> fetchnews.v1.SearchArticlesResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_fetchnews_ucpnz_v1_fetchnews_proto_init() }
func file_api_fetchnews_ucpnz_v1_fetchnews_proto_init() {
	if File_api_fetchnews_ucpnz_v1_fetchnews_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_fetchnews_ucpnz_v1_fetchnews_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_fetchnews_ucpnz_v1_fetchnews_proto_goTypes,
		DependencyIndexes: file_api_fetchnews_ucpnz_v1_fetchnews_proto_depIdxs,
	}.Build()
	File_api_fetchnews_ucpnz_v1_fetchnews_proto = out.File
	file_api_fetchnews_ucpnz_v1_fetchnews_proto_rawDesc = nil
	file_api_fetchnews_ucpnz_v1_fetchnews_proto_goTypes = nil
	file_api_fetchnews_ucpnz_v1_fetchnews_proto_depIdxs = nil
}
