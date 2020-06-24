// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: proto/todomvc/error.proto

package todomvc

import (
	proto "github.com/golang/protobuf/proto"
	code "google.golang.org/genproto/googleapis/rpc/code"
	errdetails "google.golang.org/genproto/googleapis/rpc/errdetails"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ErrorDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetryInfo           *errdetails.RetryInfo           `protobuf:"bytes,1,opt,name=retry_info,json=retryInfo,proto3" json:"retry_info,omitempty"`
	PreconditionFailure *errdetails.PreconditionFailure `protobuf:"bytes,2,opt,name=precondition_failure,json=preconditionFailure,proto3" json:"precondition_failure,omitempty"`
	BadRequest          *errdetails.BadRequest          `protobuf:"bytes,3,opt,name=bad_request,json=badRequest,proto3" json:"bad_request,omitempty"`
	LocalizedMessage    *errdetails.LocalizedMessage    `protobuf:"bytes,4,opt,name=localized_message,json=localizedMessage,proto3" json:"localized_message,omitempty"`
	ErrorInfo           *errdetails.ErrorInfo           `protobuf:"bytes,5,opt,name=error_info,json=errorInfo,proto3" json:"error_info,omitempty"`
}

func (x *ErrorDetails) Reset() {
	*x = ErrorDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todomvc_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDetails) ProtoMessage() {}

func (x *ErrorDetails) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todomvc_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDetails.ProtoReflect.Descriptor instead.
func (*ErrorDetails) Descriptor() ([]byte, []int) {
	return file_proto_todomvc_error_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorDetails) GetRetryInfo() *errdetails.RetryInfo {
	if x != nil {
		return x.RetryInfo
	}
	return nil
}

func (x *ErrorDetails) GetPreconditionFailure() *errdetails.PreconditionFailure {
	if x != nil {
		return x.PreconditionFailure
	}
	return nil
}

func (x *ErrorDetails) GetBadRequest() *errdetails.BadRequest {
	if x != nil {
		return x.BadRequest
	}
	return nil
}

func (x *ErrorDetails) GetLocalizedMessage() *errdetails.LocalizedMessage {
	if x != nil {
		return x.LocalizedMessage
	}
	return nil
}

func (x *ErrorDetails) GetErrorInfo() *errdetails.ErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    code.Code `protobuf:"varint,1,opt,name=code,proto3,enum=google.rpc.Code" json:"code,omitempty"`
	Message string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	//string request_id  = 3;
	Details *ErrorDetails `protobuf:"bytes,10,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todomvc_error_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todomvc_error_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_proto_todomvc_error_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetCode() code.Code {
	if x != nil {
		return x.Code
	}
	return code.Code_OK
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetDetails() *ErrorDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_proto_todomvc_error_proto protoreflect.FileDescriptor

var file_proto_todomvc_error_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x6d, 0x76, 0x63, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x6f, 0x64,
	0x6f, 0x6d, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2,
	0x02, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x34, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x72, 0x65, 0x74, 0x72,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x52, 0x0a, 0x14, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x50, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x52, 0x13, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x62, 0x61, 0x64,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x62, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x49, 0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x10, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a,
	0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x7b, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x6d, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x69, 0x6e, 0x75, 0x6d, 0x65, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x6d, 0x76, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x6d, 0x76, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_todomvc_error_proto_rawDescOnce sync.Once
	file_proto_todomvc_error_proto_rawDescData = file_proto_todomvc_error_proto_rawDesc
)

func file_proto_todomvc_error_proto_rawDescGZIP() []byte {
	file_proto_todomvc_error_proto_rawDescOnce.Do(func() {
		file_proto_todomvc_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_todomvc_error_proto_rawDescData)
	})
	return file_proto_todomvc_error_proto_rawDescData
}

var file_proto_todomvc_error_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_todomvc_error_proto_goTypes = []interface{}{
	(*ErrorDetails)(nil),                   // 0: todomvc.v1.ErrorDetails
	(*Error)(nil),                          // 1: todomvc.v1.Error
	(*errdetails.RetryInfo)(nil),           // 2: google.rpc.RetryInfo
	(*errdetails.PreconditionFailure)(nil), // 3: google.rpc.PreconditionFailure
	(*errdetails.BadRequest)(nil),          // 4: google.rpc.BadRequest
	(*errdetails.LocalizedMessage)(nil),    // 5: google.rpc.LocalizedMessage
	(*errdetails.ErrorInfo)(nil),           // 6: google.rpc.ErrorInfo
	(code.Code)(0),                         // 7: google.rpc.Code
}
var file_proto_todomvc_error_proto_depIdxs = []int32{
	2, // 0: todomvc.v1.ErrorDetails.retry_info:type_name -> google.rpc.RetryInfo
	3, // 1: todomvc.v1.ErrorDetails.precondition_failure:type_name -> google.rpc.PreconditionFailure
	4, // 2: todomvc.v1.ErrorDetails.bad_request:type_name -> google.rpc.BadRequest
	5, // 3: todomvc.v1.ErrorDetails.localized_message:type_name -> google.rpc.LocalizedMessage
	6, // 4: todomvc.v1.ErrorDetails.error_info:type_name -> google.rpc.ErrorInfo
	7, // 5: todomvc.v1.Error.code:type_name -> google.rpc.Code
	0, // 6: todomvc.v1.Error.details:type_name -> todomvc.v1.ErrorDetails
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_todomvc_error_proto_init() }
func file_proto_todomvc_error_proto_init() {
	if File_proto_todomvc_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_todomvc_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDetails); i {
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
		file_proto_todomvc_error_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_proto_todomvc_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_todomvc_error_proto_goTypes,
		DependencyIndexes: file_proto_todomvc_error_proto_depIdxs,
		MessageInfos:      file_proto_todomvc_error_proto_msgTypes,
	}.Build()
	File_proto_todomvc_error_proto = out.File
	file_proto_todomvc_error_proto_rawDesc = nil
	file_proto_todomvc_error_proto_goTypes = nil
	file_proto_todomvc_error_proto_depIdxs = nil
}
