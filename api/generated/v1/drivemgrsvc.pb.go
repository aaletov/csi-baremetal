// Code generated by protoc-gen-go. DO NOT EDIT.
// source: drivemgrsvc.proto

package v1api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DrivesRequest struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DrivesRequest) Reset()         { *m = DrivesRequest{} }
func (m *DrivesRequest) String() string { return proto.CompactTextString(m) }
func (*DrivesRequest) ProtoMessage()    {}
func (*DrivesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65bf77650f5c7dcf, []int{0}
}

func (m *DrivesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DrivesRequest.Unmarshal(m, b)
}
func (m *DrivesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DrivesRequest.Marshal(b, m, deterministic)
}
func (m *DrivesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DrivesRequest.Merge(m, src)
}
func (m *DrivesRequest) XXX_Size() int {
	return xxx_messageInfo_DrivesRequest.Size(m)
}
func (m *DrivesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DrivesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DrivesRequest proto.InternalMessageInfo

func (m *DrivesRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type DrivesResponse struct {
	Disks                []*Drive `protobuf:"bytes,1,rep,name=disks,proto3" json:"disks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DrivesResponse) Reset()         { *m = DrivesResponse{} }
func (m *DrivesResponse) String() string { return proto.CompactTextString(m) }
func (*DrivesResponse) ProtoMessage()    {}
func (*DrivesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65bf77650f5c7dcf, []int{1}
}

func (m *DrivesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DrivesResponse.Unmarshal(m, b)
}
func (m *DrivesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DrivesResponse.Marshal(b, m, deterministic)
}
func (m *DrivesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DrivesResponse.Merge(m, src)
}
func (m *DrivesResponse) XXX_Size() int {
	return xxx_messageInfo_DrivesResponse.Size(m)
}
func (m *DrivesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DrivesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DrivesResponse proto.InternalMessageInfo

func (m *DrivesResponse) GetDisks() []*Drive {
	if m != nil {
		return m.Disks
	}
	return nil
}

type DriveLocateRequest struct {
	DriveSerialNumber    string   `protobuf:"bytes,1,opt,name=driveSerialNumber,proto3" json:"driveSerialNumber,omitempty"`
	Action               int32    `protobuf:"varint,2,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DriveLocateRequest) Reset()         { *m = DriveLocateRequest{} }
func (m *DriveLocateRequest) String() string { return proto.CompactTextString(m) }
func (*DriveLocateRequest) ProtoMessage()    {}
func (*DriveLocateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65bf77650f5c7dcf, []int{2}
}

func (m *DriveLocateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriveLocateRequest.Unmarshal(m, b)
}
func (m *DriveLocateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriveLocateRequest.Marshal(b, m, deterministic)
}
func (m *DriveLocateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriveLocateRequest.Merge(m, src)
}
func (m *DriveLocateRequest) XXX_Size() int {
	return xxx_messageInfo_DriveLocateRequest.Size(m)
}
func (m *DriveLocateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DriveLocateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DriveLocateRequest proto.InternalMessageInfo

func (m *DriveLocateRequest) GetDriveSerialNumber() string {
	if m != nil {
		return m.DriveSerialNumber
	}
	return ""
}

func (m *DriveLocateRequest) GetAction() int32 {
	if m != nil {
		return m.Action
	}
	return 0
}

type DriveLocateResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DriveLocateResponse) Reset()         { *m = DriveLocateResponse{} }
func (m *DriveLocateResponse) String() string { return proto.CompactTextString(m) }
func (*DriveLocateResponse) ProtoMessage()    {}
func (*DriveLocateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65bf77650f5c7dcf, []int{3}
}

func (m *DriveLocateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriveLocateResponse.Unmarshal(m, b)
}
func (m *DriveLocateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriveLocateResponse.Marshal(b, m, deterministic)
}
func (m *DriveLocateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriveLocateResponse.Merge(m, src)
}
func (m *DriveLocateResponse) XXX_Size() int {
	return xxx_messageInfo_DriveLocateResponse.Size(m)
}
func (m *DriveLocateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DriveLocateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DriveLocateResponse proto.InternalMessageInfo

func (m *DriveLocateResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*DrivesRequest)(nil), "v1api.DrivesRequest")
	proto.RegisterType((*DrivesResponse)(nil), "v1api.DrivesResponse")
	proto.RegisterType((*DriveLocateRequest)(nil), "v1api.DriveLocateRequest")
	proto.RegisterType((*DriveLocateResponse)(nil), "v1api.DriveLocateResponse")
}

func init() {
	proto.RegisterFile("drivemgrsvc.proto", fileDescriptor_65bf77650f5c7dcf)
}

var fileDescriptor_65bf77650f5c7dcf = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xb6, 0x4a, 0x2b, 0xbe, 0x6d, 0x82, 0x51, 0x47, 0xed, 0xa9, 0xe4, 0x62, 0x0f, 0x5a, 0x70,
	0x7a, 0x16, 0x94, 0x81, 0x08, 0xc3, 0x43, 0xbc, 0xed, 0x96, 0xb5, 0x0f, 0x09, 0xba, 0xa6, 0xe6,
	0xa5, 0x05, 0xff, 0x0c, 0xff, 0x63, 0x59, 0x92, 0xc2, 0xa6, 0x1e, 0xbf, 0x7c, 0x1f, 0xdf, 0x8f,
	0x3c, 0x38, 0xa9, 0x8d, 0xea, 0x71, 0xfd, 0x66, 0xa8, 0xaf, 0xca, 0xd6, 0x68, 0xab, 0x59, 0xdc,
	0xdf, 0xc8, 0x56, 0x65, 0x23, 0xfb, 0xd5, 0x22, 0xf9, 0x37, 0x7e, 0x09, 0x93, 0xf9, 0x46, 0x48,
	0x02, 0x3f, 0x3b, 0x24, 0xcb, 0xa6, 0x90, 0x34, 0xba, 0xc6, 0xe7, 0x3a, 0x8d, 0xf2, 0xa8, 0x38,
	0x12, 0x01, 0xf1, 0x3b, 0x38, 0x1e, 0x84, 0xd4, 0xea, 0x86, 0x90, 0x71, 0x88, 0x6b, 0x45, 0xef,
	0x94, 0x46, 0xf9, 0x41, 0x31, 0x9a, 0x8d, 0x4b, 0x67, 0x5f, 0x3a, 0x95, 0xf0, 0x14, 0x5f, 0x02,
	0x73, 0x78, 0xa1, 0x2b, 0x69, 0x71, 0xc8, 0xb8, 0x0a, 0xed, 0x5e, 0xd1, 0x28, 0xf9, 0xf1, 0xd2,
	0xad, 0x57, 0x68, 0x42, 0xdc, 0x5f, 0x62, 0xd3, 0x48, 0x56, 0x56, 0xe9, 0x26, 0xdd, 0xcf, 0xa3,
	0x22, 0x16, 0x01, 0xf1, 0x6b, 0x38, 0xdd, 0xf1, 0x0e, 0xb5, 0xa6, 0x90, 0x90, 0x95, 0xb6, 0x23,
	0xe7, 0x18, 0x8b, 0x80, 0x66, 0xdf, 0x11, 0x8c, 0xe7, 0xc1, 0xbc, 0x57, 0x15, 0xb2, 0x7b, 0x98,
	0x3c, 0xa1, 0xf5, 0xa3, 0x16, 0x8a, 0x2c, 0x3b, 0xdb, 0x5e, 0x30, 0x7c, 0x48, 0x76, 0xfe, 0xeb,
	0xd5, 0xc7, 0xf0, 0x3d, 0xf6, 0x00, 0x89, 0x8f, 0x66, 0x17, 0xdb, 0x92, 0x9d, 0xa9, 0x59, 0xf6,
	0x1f, 0x35, 0x58, 0x3c, 0x1e, 0x2e, 0xfd, 0x4d, 0x56, 0x89, 0xbb, 0xc6, 0xed, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xa9, 0x6e, 0x7f, 0xab, 0xb6, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DriveServiceClient is the client API for DriveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DriveServiceClient interface {
	GetDrivesList(ctx context.Context, in *DrivesRequest, opts ...grpc.CallOption) (*DrivesResponse, error)
	Locate(ctx context.Context, in *DriveLocateRequest, opts ...grpc.CallOption) (*DriveLocateResponse, error)
}

type driveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDriveServiceClient(cc grpc.ClientConnInterface) DriveServiceClient {
	return &driveServiceClient{cc}
}

func (c *driveServiceClient) GetDrivesList(ctx context.Context, in *DrivesRequest, opts ...grpc.CallOption) (*DrivesResponse, error) {
	out := new(DrivesResponse)
	err := c.cc.Invoke(ctx, "/v1api.DriveService/GetDrivesList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driveServiceClient) Locate(ctx context.Context, in *DriveLocateRequest, opts ...grpc.CallOption) (*DriveLocateResponse, error) {
	out := new(DriveLocateResponse)
	err := c.cc.Invoke(ctx, "/v1api.DriveService/Locate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriveServiceServer is the server API for DriveService service.
type DriveServiceServer interface {
	GetDrivesList(context.Context, *DrivesRequest) (*DrivesResponse, error)
	Locate(context.Context, *DriveLocateRequest) (*DriveLocateResponse, error)
}

// UnimplementedDriveServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDriveServiceServer struct {
}

func (*UnimplementedDriveServiceServer) GetDrivesList(ctx context.Context, req *DrivesRequest) (*DrivesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDrivesList not implemented")
}
func (*UnimplementedDriveServiceServer) Locate(ctx context.Context, req *DriveLocateRequest) (*DriveLocateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Locate not implemented")
}

func RegisterDriveServiceServer(s *grpc.Server, srv DriveServiceServer) {
	s.RegisterService(&_DriveService_serviceDesc, srv)
}

func _DriveService_GetDrivesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DrivesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriveServiceServer).GetDrivesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1api.DriveService/GetDrivesList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriveServiceServer).GetDrivesList(ctx, req.(*DrivesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriveService_Locate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriveLocateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriveServiceServer).Locate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1api.DriveService/Locate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriveServiceServer).Locate(ctx, req.(*DriveLocateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DriveService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1api.DriveService",
	HandlerType: (*DriveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDrivesList",
			Handler:    _DriveService_GetDrivesList_Handler,
		},
		{
			MethodName: "Locate",
			Handler:    _DriveService_Locate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "drivemgrsvc.proto",
}