// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: parking/gunk/api/v1/reports/all.proto

package reports

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalEarning int32 `protobuf:"varint,1,opt,name=TotalEarning,proto3" json:"TotalEarning,omitempty"`
	TotalParked  int32 `protobuf:"varint,2,opt,name=TotalParked,proto3" json:"TotalParked,omitempty"`
	TotalHours   int32 `protobuf:"varint,3,opt,name=TotalHours,proto3" json:"TotalHours,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parking_gunk_api_v1_reports_all_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_parking_gunk_api_v1_reports_all_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_parking_gunk_api_v1_reports_all_proto_rawDescGZIP(), []int{0}
}

func (x *Report) GetTotalEarning() int32 {
	if x != nil {
		return x.TotalEarning
	}
	return 0
}

func (x *Report) GetTotalParked() int32 {
	if x != nil {
		return x.TotalParked
	}
	return 0
}

func (x *Report) GetTotalHours() int32 {
	if x != nil {
		return x.TotalHours
	}
	return 0
}

type ReportData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report *Report `protobuf:"bytes,1,opt,name=Report,json=Reports,proto3" json:"Reports,omitempty"`
}

func (x *ReportData) Reset() {
	*x = ReportData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parking_gunk_api_v1_reports_all_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportData) ProtoMessage() {}

func (x *ReportData) ProtoReflect() protoreflect.Message {
	mi := &file_parking_gunk_api_v1_reports_all_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportData.ProtoReflect.Descriptor instead.
func (*ReportData) Descriptor() ([]byte, []int) {
	return file_parking_gunk_api_v1_reports_all_proto_rawDescGZIP(), []int{1}
}

func (x *ReportData) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

type GetReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDate *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
}

func (x *GetReportRequest) Reset() {
	*x = GetReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parking_gunk_api_v1_reports_all_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportRequest) ProtoMessage() {}

func (x *GetReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parking_gunk_api_v1_reports_all_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportRequest.ProtoReflect.Descriptor instead.
func (*GetReportRequest) Descriptor() ([]byte, []int) {
	return file_parking_gunk_api_v1_reports_all_proto_rawDescGZIP(), []int{2}
}

func (x *GetReportRequest) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *GetReportRequest) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

// UnParkVehicleResponse ...
type GetReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *ReportData `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GetReportResponse) Reset() {
	*x = GetReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parking_gunk_api_v1_reports_all_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReportResponse) ProtoMessage() {}

func (x *GetReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parking_gunk_api_v1_reports_all_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReportResponse.ProtoReflect.Descriptor instead.
func (*GetReportResponse) Descriptor() ([]byte, []int) {
	return file_parking_gunk_api_v1_reports_all_proto_rawDescGZIP(), []int{3}
}

func (x *GetReportResponse) GetData() *ReportData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_parking_gunk_api_v1_reports_all_proto protoreflect.FileDescriptor

var file_parking_gunk_api_v1_reports_all_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x61, 0x6c,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9a, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x0c, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x45, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x0c, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x45, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x0a, 0x0b, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x0b, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x50, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00,
	0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x48,
	0x6f, 0x75, 0x72, 0x73, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x4a, 0x0a, 0x0a,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x06, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x0a, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x07, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0xa2, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a,
	0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0x08, 0x00,
	0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x07, 0x45, 0x6e,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x50, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50,
	0x00, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x32,
	0xb0, 0x05, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x99, 0x05, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x19, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xd0, 0x04, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x92,
	0x41, 0xb6, 0x04, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x06, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x1a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x2e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x4a, 0x53, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x4c, 0x0a, 0x1e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64,
	0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x2e, 0x12, 0x2a,
	0x0a, 0x28, 0x1a, 0x26, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x7e, 0x0a, 0x03, 0x34, 0x30,
	0x30, 0x12, 0x77, 0x0a, 0x30, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68,
	0x65, 0x6e, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x69, 0x73, 0x20, 0x69, 0x6e,
	0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x20, 0x6f, 0x72, 0x20, 0x6d, 0x61, 0x6c, 0x66, 0x6f,
	0x72, 0x6d, 0x65, 0x64, 0x2e, 0x12, 0x43, 0x0a, 0x41, 0x4a, 0x3f, 0x7b, 0x20, 0x22, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x3a, 0x20, 0x34, 0x30, 0x30, 0x2c, 0x20, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x69, 0x73,
	0x20, 0x69, 0x6e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x20, 0x6f, 0x72, 0x20, 0x6d, 0x61,
	0x6c, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x22, 0x20, 0x7d, 0x4a, 0x6f, 0x0a, 0x03, 0x34, 0x30,
	0x31, 0x12, 0x68, 0x0a, 0x34, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68,
	0x65, 0x6e, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x64, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x74, 0x68, 0x69,
	0x73, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x12, 0x30, 0x0a, 0x2e, 0x4a, 0x2c, 0x7b,
	0x20, 0x22, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x20, 0x34, 0x30, 0x31, 0x2c, 0x20, 0x22, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x6e, 0x6f, 0x74, 0x20, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x22, 0x20, 0x7d, 0x4a, 0x5b, 0x0a, 0x03, 0x34,
	0x30, 0x34, 0x12, 0x54, 0x0a, 0x18, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77,
	0x68, 0x65, 0x6e, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x12, 0x38,
	0x0a, 0x36, 0x4a, 0x34, 0x7b, 0x20, 0x22, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x20, 0x34, 0x30,
	0x34, 0x2c, 0x20, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x70,
	0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x6e, 0x6f, 0x74, 0x20,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x20, 0x7d, 0x4a, 0x5f, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12,
	0x58, 0x0a, 0x24, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e,
	0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x12, 0x30, 0x0a, 0x2e, 0x4a, 0x2c, 0x7b, 0x20, 0x22,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x20, 0x35, 0x30, 0x30, 0x2c, 0x20, 0x22, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x20, 0x7d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12,
	0x08, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x28, 0x00, 0x30, 0x00, 0x1a, 0x03, 0x88,
	0x02, 0x00, 0x42, 0x3e, 0x48, 0x01, 0x50, 0x00, 0x5a, 0x23, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x2f, 0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x3b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x80, 0x01, 0x00,
	0x88, 0x01, 0x00, 0x90, 0x01, 0x00, 0xb8, 0x01, 0x00, 0xd8, 0x01, 0x00, 0xf8, 0x01, 0x01, 0xd0,
	0x02, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_parking_gunk_api_v1_reports_all_proto_rawDescOnce sync.Once
	file_parking_gunk_api_v1_reports_all_proto_rawDescData = file_parking_gunk_api_v1_reports_all_proto_rawDesc
)

func file_parking_gunk_api_v1_reports_all_proto_rawDescGZIP() []byte {
	file_parking_gunk_api_v1_reports_all_proto_rawDescOnce.Do(func() {
		file_parking_gunk_api_v1_reports_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_parking_gunk_api_v1_reports_all_proto_rawDescData)
	})
	return file_parking_gunk_api_v1_reports_all_proto_rawDescData
}

var (
	file_parking_gunk_api_v1_reports_all_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
	file_parking_gunk_api_v1_reports_all_proto_goTypes  = []interface{}{
		(*Report)(nil),                // 0: reports.Report
		(*ReportData)(nil),            // 1: reports.ReportData
		(*GetReportRequest)(nil),      // 2: reports.GetReportRequest
		(*GetReportResponse)(nil),     // 3: reports.GetReportResponse
		(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	}
)

var file_parking_gunk_api_v1_reports_all_proto_depIdxs = []int32{
	0, // 0: reports.ReportData.Report:type_name -> reports.Report
	4, // 1: reports.GetReportRequest.StartDate:type_name -> google.protobuf.Timestamp
	4, // 2: reports.GetReportRequest.EndDate:type_name -> google.protobuf.Timestamp
	1, // 3: reports.GetReportResponse.Data:type_name -> reports.ReportData
	2, // 4: reports.ReportService.GetReport:input_type -> reports.GetReportRequest
	3, // 5: reports.ReportService.GetReport:output_type -> reports.GetReportResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_parking_gunk_api_v1_reports_all_proto_init() }
func file_parking_gunk_api_v1_reports_all_proto_init() {
	if File_parking_gunk_api_v1_reports_all_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_parking_gunk_api_v1_reports_all_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
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
		file_parking_gunk_api_v1_reports_all_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportData); i {
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
		file_parking_gunk_api_v1_reports_all_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReportRequest); i {
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
		file_parking_gunk_api_v1_reports_all_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReportResponse); i {
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
			RawDescriptor: file_parking_gunk_api_v1_reports_all_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_parking_gunk_api_v1_reports_all_proto_goTypes,
		DependencyIndexes: file_parking_gunk_api_v1_reports_all_proto_depIdxs,
		MessageInfos:      file_parking_gunk_api_v1_reports_all_proto_msgTypes,
	}.Build()
	File_parking_gunk_api_v1_reports_all_proto = out.File
	file_parking_gunk_api_v1_reports_all_proto_rawDesc = nil
	file_parking_gunk_api_v1_reports_all_proto_goTypes = nil
	file_parking_gunk_api_v1_reports_all_proto_depIdxs = nil
}