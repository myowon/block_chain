// Copyright 2021 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/retail/logging/error_log.proto

package logging

import (
	reflect "reflect"
	sync "sync"

	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Describes a running service that sends errors.
type ServiceContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An identifier of the service.
	// For example, "retail.googleapis.com".
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *ServiceContext) Reset() {
	*x = ServiceContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_logging_error_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceContext) ProtoMessage() {}

func (x *ServiceContext) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_logging_error_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceContext.ProtoReflect.Descriptor instead.
func (*ServiceContext) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_logging_error_log_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceContext) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

// HTTP request data that is related to a reported error.
type HttpRequestContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTP response status code for the request.
	ResponseStatusCode int32 `protobuf:"varint,1,opt,name=response_status_code,json=responseStatusCode,proto3" json:"response_status_code,omitempty"`
}

func (x *HttpRequestContext) Reset() {
	*x = HttpRequestContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_logging_error_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpRequestContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpRequestContext) ProtoMessage() {}

func (x *HttpRequestContext) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_logging_error_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpRequestContext.ProtoReflect.Descriptor instead.
func (*HttpRequestContext) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_logging_error_log_proto_rawDescGZIP(), []int{1}
}

func (x *HttpRequestContext) GetResponseStatusCode() int32 {
	if x != nil {
		return x.ResponseStatusCode
	}
	return 0
}

// Indicates a location in the source code of the service for which
// errors are reported.
type SourceLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Human-readable name of a function or method.
	// For example, "google.cloud.retail.v2.UserEventService.ImportUserEvents".
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
}

func (x *SourceLocation) Reset() {
	*x = SourceLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_logging_error_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceLocation) ProtoMessage() {}

func (x *SourceLocation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_logging_error_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceLocation.ProtoReflect.Descriptor instead.
func (*SourceLocation) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_logging_error_log_proto_rawDescGZIP(), []int{2}
}

func (x *SourceLocation) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

// A description of the context in which an error occurred.
type ErrorContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTP request which was processed when the error was triggered.
	HttpRequest *HttpRequestContext `protobuf:"bytes,1,opt,name=http_request,json=httpRequest,proto3" json:"http_request,omitempty"`
	// The location in the source code where the decision was made to
	// report the error, usually the place where it was logged.
	ReportLocation *SourceLocation `protobuf:"bytes,2,opt,name=report_location,json=reportLocation,proto3" json:"report_location,omitempty"`
}

func (x *ErrorContext) Reset() {
	*x = ErrorContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_logging_error_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorContext) ProtoMessage() {}

func (x *ErrorContext) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_logging_error_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorContext.ProtoReflect.Descriptor instead.
func (*ErrorContext) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_logging_error_log_proto_rawDescGZIP(), []int{3}
}

func (x *ErrorContext) GetHttpRequest() *HttpRequestContext {
	if x != nil {
		return x.HttpRequest
	}
	return nil
}

func (x *ErrorContext) GetReportLocation() *SourceLocation {
	if x != nil {
		return x.ReportLocation
	}
	return nil
}

// The error payload that is populated on LRO import APIs. Including:
//   "google.cloud.retail.v2.ProductService.ImportProducts"
//   "google.cloud.retail.v2.EventService.ImportUserEvents"
type ImportErrorContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The operation resource name of the LRO.
	OperationName string `protobuf:"bytes,1,opt,name=operation_name,json=operationName,proto3" json:"operation_name,omitempty"`
	// Cloud Storage file path of the import source.
	// Can be set for batch operation error.
	GcsPath string `protobuf:"bytes,2,opt,name=gcs_path,json=gcsPath,proto3" json:"gcs_path,omitempty"`
	// Line number of the content in file.
	// Should be empty for permission or batch operation error.
	LineNumber string `protobuf:"bytes,3,opt,name=line_number,json=lineNumber,proto3" json:"line_number,omitempty"`
	// Detailed content which caused the error.
	// Should be empty for permission or batch operation error.
	//
	// Types that are assignable to LineContent:
	//	*ImportErrorContext_CatalogItem
	//	*ImportErrorContext_Product
	//	*ImportErrorContext_UserEvent
	LineContent isImportErrorContext_LineContent `protobuf_oneof:"line_content"`
}

func (x *ImportErrorContext) Reset() {
	*x = ImportErrorContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_logging_error_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportErrorContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportErrorContext) ProtoMessage() {}

func (x *ImportErrorContext) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_logging_error_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportErrorContext.ProtoReflect.Descriptor instead.
func (*ImportErrorContext) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_logging_error_log_proto_rawDescGZIP(), []int{4}
}

func (x *ImportErrorContext) GetOperationName() string {
	if x != nil {
		return x.OperationName
	}
	return ""
}

func (x *ImportErrorContext) GetGcsPath() string {
	if x != nil {
		return x.GcsPath
	}
	return ""
}

func (x *ImportErrorContext) GetLineNumber() string {
	if x != nil {
		return x.LineNumber
	}
	return ""
}

func (m *ImportErrorContext) GetLineContent() isImportErrorContext_LineContent {
	if m != nil {
		return m.LineContent
	}
	return nil
}

func (x *ImportErrorContext) GetCatalogItem() string {
	if x, ok := x.GetLineContent().(*ImportErrorContext_CatalogItem); ok {
		return x.CatalogItem
	}
	return ""
}

func (x *ImportErrorContext) GetProduct() string {
	if x, ok := x.GetLineContent().(*ImportErrorContext_Product); ok {
		return x.Product
	}
	return ""
}

func (x *ImportErrorContext) GetUserEvent() string {
	if x, ok := x.GetLineContent().(*ImportErrorContext_UserEvent); ok {
		return x.UserEvent
	}
	return ""
}

type isImportErrorContext_LineContent interface {
	isImportErrorContext_LineContent()
}

type ImportErrorContext_CatalogItem struct {
	// The detailed content which caused the error on importing a catalog item.
	CatalogItem string `protobuf:"bytes,4,opt,name=catalog_item,json=catalogItem,proto3,oneof"`
}

type ImportErrorContext_Product struct {
	// The detailed content which caused the error on importing a product.
	Product string `protobuf:"bytes,5,opt,name=product,proto3,oneof"`
}

type ImportErrorContext_UserEvent struct {
	// The detailed content which caused the error on importing a user event.
	UserEvent string `protobuf:"bytes,6,opt,name=user_event,json=userEvent,proto3,oneof"`
}

func (*ImportErrorContext_CatalogItem) isImportErrorContext_LineContent() {}

func (*ImportErrorContext_Product) isImportErrorContext_LineContent() {}

func (*ImportErrorContext_UserEvent) isImportErrorContext_LineContent() {}

// An error log which is reported to the Error Reporting system.
// This proto a superset of
// google.devtools.clouderrorreporting.v1beta1.ReportedErrorEvent.
type ErrorLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The service context in which this error has occurred.
	ServiceContext *ServiceContext `protobuf:"bytes,1,opt,name=service_context,json=serviceContext,proto3" json:"service_context,omitempty"`
	// A description of the context in which the error occurred.
	Context *ErrorContext `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	// A message describing the error.
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// The RPC status associated with the error log.
	Status *status.Status `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	// The API request payload, represented as a protocol buffer.
	//
	// Most API request types are supported. For example:
	//
	//   "type.googleapis.com/google.cloud.retail.v2.ProductService.CreateProductRequest"
	//   "type.googleapis.com/google.cloud.retail.v2.UserEventService.WriteUserEventRequest"
	RequestPayload *structpb.Struct `protobuf:"bytes,5,opt,name=request_payload,json=requestPayload,proto3" json:"request_payload,omitempty"`
	// The API response payload, represented as a protocol buffer.
	//
	// This is used to log some "soft errors", where the response is valid but we
	// consider there are some quality issues like unjoined events.
	//
	// The following API responses are supported and no PII is included:
	//   "google.cloud.retail.v2.PredictionService.Predict"
	//   "google.cloud.retail.v2.UserEventService.WriteUserEvent"
	//   "google.cloud.retail.v2.UserEventService.CollectUserEvent"
	ResponsePayload *structpb.Struct `protobuf:"bytes,6,opt,name=response_payload,json=responsePayload,proto3" json:"response_payload,omitempty"`
	// The error payload that is populated on LRO import APIs.
	ImportPayload *ImportErrorContext `protobuf:"bytes,7,opt,name=import_payload,json=importPayload,proto3" json:"import_payload,omitempty"`
}

func (x *ErrorLog) Reset() {
	*x = ErrorLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_logging_error_log_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorLog) ProtoMessage() {}

func (x *ErrorLog) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_logging_error_log_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorLog.ProtoReflect.Descriptor instead.
func (*ErrorLog) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_logging_error_log_proto_rawDescGZIP(), []int{5}
}

func (x *ErrorLog) GetServiceContext() *ServiceContext {
	if x != nil {
		return x.ServiceContext
	}
	return nil
}

func (x *ErrorLog) GetContext() *ErrorContext {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *ErrorLog) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorLog) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ErrorLog) GetRequestPayload() *structpb.Struct {
	if x != nil {
		return x.RequestPayload
	}
	return nil
}

func (x *ErrorLog) GetResponsePayload() *structpb.Struct {
	if x != nil {
		return x.ResponsePayload
	}
	return nil
}

func (x *ErrorLog) GetImportPayload() *ImportErrorContext {
	if x != nil {
		return x.ImportPayload
	}
	return nil
}

var File_google_cloud_retail_logging_error_log_proto protoreflect.FileDescriptor

var file_google_cloud_retail_logging_error_log_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x46, 0x0a,
	0x12, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x12, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x35, 0x0a, 0x0e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb8, 0x01, 0x0a,
	0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x52, 0x0a,
	0x0c, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x54, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe9, 0x01, 0x0a, 0x12, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x63, 0x73, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x63, 0x73, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x23, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x1f, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0xc9, 0x03, 0x0a, 0x08, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4c, 0x6f, 0x67,
	0x12, 0x54, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x43, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x40, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x42, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x56, 0x0a, 0x0e, 0x69, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x52, 0x0d, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42,
	0xdc, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x42, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4c, 0x6f, 0x67, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x3b, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0xa2, 0x02, 0x06, 0x52, 0x45, 0x54, 0x41, 0x49,
	0x4c, 0xaa, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0xca,
	0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x52,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x5c, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0xea, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x52,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x3a, 0x3a, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_retail_logging_error_log_proto_rawDescOnce sync.Once
	file_google_cloud_retail_logging_error_log_proto_rawDescData = file_google_cloud_retail_logging_error_log_proto_rawDesc
)

func file_google_cloud_retail_logging_error_log_proto_rawDescGZIP() []byte {
	file_google_cloud_retail_logging_error_log_proto_rawDescOnce.Do(func() {
		file_google_cloud_retail_logging_error_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_retail_logging_error_log_proto_rawDescData)
	})
	return file_google_cloud_retail_logging_error_log_proto_rawDescData
}

var file_google_cloud_retail_logging_error_log_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_retail_logging_error_log_proto_goTypes = []interface{}{
	(*ServiceContext)(nil),     // 0: google.cloud.retail.logging.ServiceContext
	(*HttpRequestContext)(nil), // 1: google.cloud.retail.logging.HttpRequestContext
	(*SourceLocation)(nil),     // 2: google.cloud.retail.logging.SourceLocation
	(*ErrorContext)(nil),       // 3: google.cloud.retail.logging.ErrorContext
	(*ImportErrorContext)(nil), // 4: google.cloud.retail.logging.ImportErrorContext
	(*ErrorLog)(nil),           // 5: google.cloud.retail.logging.ErrorLog
	(*status.Status)(nil),      // 6: google.rpc.Status
	(*structpb.Struct)(nil),    // 7: google.protobuf.Struct
}
var file_google_cloud_retail_logging_error_log_proto_depIdxs = []int32{
	1, // 0: google.cloud.retail.logging.ErrorContext.http_request:type_name -> google.cloud.retail.logging.HttpRequestContext
	2, // 1: google.cloud.retail.logging.ErrorContext.report_location:type_name -> google.cloud.retail.logging.SourceLocation
	0, // 2: google.cloud.retail.logging.ErrorLog.service_context:type_name -> google.cloud.retail.logging.ServiceContext
	3, // 3: google.cloud.retail.logging.ErrorLog.context:type_name -> google.cloud.retail.logging.ErrorContext
	6, // 4: google.cloud.retail.logging.ErrorLog.status:type_name -> google.rpc.Status
	7, // 5: google.cloud.retail.logging.ErrorLog.request_payload:type_name -> google.protobuf.Struct
	7, // 6: google.cloud.retail.logging.ErrorLog.response_payload:type_name -> google.protobuf.Struct
	4, // 7: google.cloud.retail.logging.ErrorLog.import_payload:type_name -> google.cloud.retail.logging.ImportErrorContext
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_cloud_retail_logging_error_log_proto_init() }
func file_google_cloud_retail_logging_error_log_proto_init() {
	if File_google_cloud_retail_logging_error_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_retail_logging_error_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceContext); i {
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
		file_google_cloud_retail_logging_error_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpRequestContext); i {
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
		file_google_cloud_retail_logging_error_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceLocation); i {
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
		file_google_cloud_retail_logging_error_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorContext); i {
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
		file_google_cloud_retail_logging_error_log_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportErrorContext); i {
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
		file_google_cloud_retail_logging_error_log_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorLog); i {
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
	file_google_cloud_retail_logging_error_log_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ImportErrorContext_CatalogItem)(nil),
		(*ImportErrorContext_Product)(nil),
		(*ImportErrorContext_UserEvent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_retail_logging_error_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_retail_logging_error_log_proto_goTypes,
		DependencyIndexes: file_google_cloud_retail_logging_error_log_proto_depIdxs,
		MessageInfos:      file_google_cloud_retail_logging_error_log_proto_msgTypes,
	}.Build()
	File_google_cloud_retail_logging_error_log_proto = out.File
	file_google_cloud_retail_logging_error_log_proto_rawDesc = nil
	file_google_cloud_retail_logging_error_log_proto_goTypes = nil
	file_google_cloud_retail_logging_error_log_proto_depIdxs = nil
}
