// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sensors.proto

package pb

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

// The request message containing the user's name.
type SensorRequest struct {
	DeviceId             string   `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	RecordId             string   `protobuf:"bytes,2,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	SensorType           string   `protobuf:"bytes,3,opt,name=sensor_type,json=sensorType,proto3" json:"sensor_type,omitempty"`
	DataStr              []string `protobuf:"bytes,4,rep,name=data_str,json=dataStr,proto3" json:"data_str,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SensorRequest) Reset()         { *m = SensorRequest{} }
func (m *SensorRequest) String() string { return proto.CompactTextString(m) }
func (*SensorRequest) ProtoMessage()    {}
func (*SensorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b96d375cab19813b, []int{0}
}

func (m *SensorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SensorRequest.Unmarshal(m, b)
}
func (m *SensorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SensorRequest.Marshal(b, m, deterministic)
}
func (m *SensorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SensorRequest.Merge(m, src)
}
func (m *SensorRequest) XXX_Size() int {
	return xxx_messageInfo_SensorRequest.Size(m)
}
func (m *SensorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SensorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SensorRequest proto.InternalMessageInfo

func (m *SensorRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *SensorRequest) GetRecordId() string {
	if m != nil {
		return m.RecordId
	}
	return ""
}

func (m *SensorRequest) GetSensorType() string {
	if m != nil {
		return m.SensorType
	}
	return ""
}

func (m *SensorRequest) GetDataStr() []string {
	if m != nil {
		return m.DataStr
	}
	return nil
}

// The response message containing the greetings
type SensorReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SensorReply) Reset()         { *m = SensorReply{} }
func (m *SensorReply) String() string { return proto.CompactTextString(m) }
func (*SensorReply) ProtoMessage()    {}
func (*SensorReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b96d375cab19813b, []int{1}
}

func (m *SensorReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SensorReply.Unmarshal(m, b)
}
func (m *SensorReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SensorReply.Marshal(b, m, deterministic)
}
func (m *SensorReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SensorReply.Merge(m, src)
}
func (m *SensorReply) XXX_Size() int {
	return xxx_messageInfo_SensorReply.Size(m)
}
func (m *SensorReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SensorReply.DiscardUnknown(m)
}

var xxx_messageInfo_SensorReply proto.InternalMessageInfo

func (m *SensorReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SensorRequest)(nil), "sensors.SensorRequest")
	proto.RegisterType((*SensorReply)(nil), "sensors.SensorReply")
}

func init() { proto.RegisterFile("sensors.proto", fileDescriptor_b96d375cab19813b) }

var fileDescriptor_b96d375cab19813b = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0xcd, 0x2b,
	0xce, 0x2f, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x5a, 0x19,
	0xb9, 0x78, 0x83, 0xc1, 0xec, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x69, 0x2e, 0xce,
	0x94, 0xd4, 0xb2, 0xcc, 0xe4, 0xd4, 0xf8, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x0e, 0x88, 0x80, 0x67, 0x0a, 0x48, 0xb2, 0x28, 0x35, 0x39, 0xbf, 0x28, 0x05, 0x24, 0xc9, 0x04,
	0x91, 0x84, 0x08, 0x78, 0xa6, 0x08, 0xc9, 0x73, 0x71, 0x43, 0x8c, 0x8d, 0x2f, 0xa9, 0x2c, 0x48,
	0x95, 0x60, 0x06, 0x4b, 0x73, 0x41, 0x84, 0x42, 0x2a, 0x0b, 0x52, 0x85, 0x24, 0xb9, 0x38, 0x52,
	0x12, 0x4b, 0x12, 0xe3, 0x8b, 0x4b, 0x8a, 0x24, 0x58, 0x14, 0x98, 0x35, 0x38, 0x83, 0xd8, 0x41,
	0xfc, 0xe0, 0x92, 0x22, 0x25, 0x75, 0x2e, 0x6e, 0x98, 0x33, 0x0a, 0x72, 0x2a, 0x85, 0x24, 0xb8,
	0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0xa1, 0x4e, 0x80, 0x71, 0x8d, 0x5c, 0xb9, 0xd8,
	0x21, 0x0a, 0x8b, 0x85, 0xac, 0xb8, 0x38, 0x5c, 0x12, 0x4b, 0x12, 0x03, 0x4a, 0x8b, 0x33, 0x84,
	0xc4, 0xf4, 0x60, 0x1e, 0x44, 0xf1, 0x8d, 0x94, 0x08, 0x86, 0x78, 0x41, 0x4e, 0xa5, 0x12, 0x83,
	0x93, 0x02, 0x17, 0x47, 0x72, 0x7e, 0xae, 0x5e, 0x7a, 0x51, 0x41, 0xb2, 0x13, 0x0f, 0xd4, 0xc0,
	0x00, 0x50, 0xd0, 0x04, 0x30, 0x2e, 0x62, 0x62, 0xf6, 0xf0, 0x09, 0x4f, 0x62, 0x03, 0x87, 0x94,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x59, 0xfd, 0x13, 0xa4, 0x3a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SensorsClient is the client API for Sensors service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SensorsClient interface {
	// Sends a greeting
	DataPush(ctx context.Context, in *SensorRequest, opts ...grpc.CallOption) (*SensorReply, error)
}

type sensorsClient struct {
	cc *grpc.ClientConn
}

func NewSensorsClient(cc *grpc.ClientConn) SensorsClient {
	return &sensorsClient{cc}
}

func (c *sensorsClient) DataPush(ctx context.Context, in *SensorRequest, opts ...grpc.CallOption) (*SensorReply, error) {
	out := new(SensorReply)
	err := c.cc.Invoke(ctx, "/sensors.Sensors/DataPush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SensorsServer is the server API for Sensors service.
type SensorsServer interface {
	// Sends a greeting
	DataPush(context.Context, *SensorRequest) (*SensorReply, error)
}

// UnimplementedSensorsServer can be embedded to have forward compatible implementations.
type UnimplementedSensorsServer struct {
}

func (*UnimplementedSensorsServer) DataPush(ctx context.Context, req *SensorRequest) (*SensorReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataPush not implemented")
}

func RegisterSensorsServer(s *grpc.Server, srv SensorsServer) {
	s.RegisterService(&_Sensors_serviceDesc, srv)
}

func _Sensors_DataPush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SensorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorsServer).DataPush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sensors.Sensors/DataPush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorsServer).DataPush(ctx, req.(*SensorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sensors_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sensors.Sensors",
	HandlerType: (*SensorsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DataPush",
			Handler:    _Sensors_DataPush_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sensors.proto",
}
