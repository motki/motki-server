// Code generated by protoc-gen-go. DO NOT EDIT.
// source: motki.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A Status indicates success or failure.
type Status int32

const (
	Status_FAILURE Status = 0
	Status_SUCCESS Status = 1
)

var Status_name = map[int32]string{
	0: "FAILURE",
	1: "SUCCESS",
}
var Status_value = map[string]int32{
	"FAILURE": 0,
	"SUCCESS": 1,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_motki_d276ca21445293e2, []int{0}
}

// A Result contains a status and an optional description.
type Result struct {
	Status Status `protobuf:"varint,1,opt,name=status,enum=motki.Status" json:"status,omitempty"`
	// Description contains some text about a failure in most cases.
	Description          string   `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_motki_d276ca21445293e2, []int{0}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_FAILURE
}

func (m *Result) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// A Token contains a session identifier representing a valid user session.
type Token struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_motki_d276ca21445293e2, []int{1}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (dst *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(dst, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type AuthenticateRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateRequest) Reset()         { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()    {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_motki_d276ca21445293e2, []int{2}
}
func (m *AuthenticateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateRequest.Unmarshal(m, b)
}
func (m *AuthenticateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateRequest.Marshal(b, m, deterministic)
}
func (dst *AuthenticateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateRequest.Merge(dst, src)
}
func (m *AuthenticateRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticateRequest.Size(m)
}
func (m *AuthenticateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateRequest proto.InternalMessageInfo

func (m *AuthenticateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthenticateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthenticateResponse struct {
	Result               *Result  `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	Token                *Token   `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateResponse) Reset()         { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()    {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_motki_d276ca21445293e2, []int{3}
}
func (m *AuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateResponse.Unmarshal(m, b)
}
func (m *AuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateResponse.Marshal(b, m, deterministic)
}
func (dst *AuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateResponse.Merge(dst, src)
}
func (m *AuthenticateResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticateResponse.Size(m)
}
func (m *AuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateResponse proto.InternalMessageInfo

func (m *AuthenticateResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *AuthenticateResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func init() {
	proto.RegisterType((*Result)(nil), "motki.Result")
	proto.RegisterType((*Token)(nil), "motki.Token")
	proto.RegisterType((*AuthenticateRequest)(nil), "motki.AuthenticateRequest")
	proto.RegisterType((*AuthenticateResponse)(nil), "motki.AuthenticateResponse")
	proto.RegisterEnum("motki.Status", Status_name, Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
}

type authenticationServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationServiceClient(cc *grpc.ClientConn) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/motki.AuthenticationService/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
type AuthenticationServiceServer interface {
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
}

func RegisterAuthenticationServiceServer(s *grpc.Server, srv AuthenticationServiceServer) {
	s.RegisterService(&_AuthenticationService_serviceDesc, srv)
}

func _AuthenticationService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/motki.AuthenticationService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthenticationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "motki.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _AuthenticationService_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "motki.proto",
}

func init() { proto.RegisterFile("motki.proto", fileDescriptor_motki_d276ca21445293e2) }

var fileDescriptor_motki_d276ca21445293e2 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xff, 0xf9, 0x43, 0x52, 0x3b, 0xa9, 0x52, 0x56, 0x85, 0x12, 0x41, 0xc2, 0x82, 0x58,
	0x3c, 0xf4, 0x10, 0x3f, 0x41, 0x2d, 0x55, 0x04, 0x3d, 0xb8, 0xb1, 0x17, 0x4f, 0xc6, 0x64, 0xc4,
	0xa5, 0x76, 0x37, 0xee, 0x4e, 0xf4, 0xeb, 0x4b, 0x36, 0x4b, 0x88, 0xd0, 0x53, 0x78, 0xf3, 0x26,
	0x6f, 0xe6, 0x37, 0x0b, 0xf1, 0x4e, 0xd3, 0x56, 0x2e, 0x6a, 0xa3, 0x49, 0xb3, 0xd0, 0x09, 0xfe,
	0x04, 0x91, 0x40, 0xdb, 0x7c, 0x12, 0xbb, 0x80, 0xc8, 0x52, 0x41, 0x8d, 0x9d, 0x05, 0x69, 0x30,
	0x3f, 0xca, 0x0e, 0x17, 0x5d, 0x7b, 0xee, 0x8a, 0xc2, 0x9b, 0x2c, 0x85, 0xb8, 0x42, 0x5b, 0x1a,
	0x59, 0x93, 0xd4, 0x6a, 0xf6, 0x3f, 0x0d, 0xe6, 0x63, 0x31, 0x2c, 0xf1, 0x4b, 0x08, 0x9f, 0xf5,
	0x16, 0x15, 0x3b, 0x07, 0x90, 0x15, 0x2a, 0x92, 0xef, 0x12, 0x8d, 0x4b, 0x1d, 0x8b, 0x41, 0x85,
	0x3f, 0xc2, 0xf1, 0xb2, 0xa1, 0x8f, 0x56, 0x97, 0x05, 0xa1, 0xc0, 0xaf, 0x06, 0x2d, 0xb1, 0x04,
	0x0e, 0x1a, 0x8b, 0x46, 0x15, 0x3b, 0xf4, 0x3f, 0xf5, 0xba, 0xf5, 0xea, 0xc2, 0xda, 0x1f, 0x6d,
	0x2a, 0x3f, 0xba, 0xd7, 0xbc, 0x80, 0x93, 0xbf, 0x71, 0xb6, 0xd6, 0xca, 0x62, 0x0b, 0x66, 0x1c,
	0xa2, 0x4b, 0x8b, 0x7b, 0xb0, 0x8e, 0x5b, 0x78, 0x93, 0x71, 0x08, 0xa9, 0x5d, 0xdb, 0xe5, 0xc6,
	0xd9, 0xc4, 0x77, 0x39, 0x14, 0xd1, 0x59, 0x57, 0x1c, 0xa2, 0xee, 0x1c, 0x2c, 0x86, 0xd1, 0xed,
	0xf2, 0xfe, 0x61, 0x23, 0xd6, 0xd3, 0x7f, 0xad, 0xc8, 0x37, 0xab, 0xd5, 0x3a, 0xcf, 0xa7, 0x41,
	0xf6, 0x0a, 0xa7, 0x83, 0x35, 0xa4, 0x56, 0x39, 0x9a, 0x6f, 0x59, 0x22, 0xbb, 0x83, 0xc9, 0x70,
	0x3f, 0x96, 0xf8, 0x09, 0x7b, 0x6e, 0x90, 0x9c, 0xed, 0xf5, 0x3a, 0xa0, 0x9b, 0xd1, 0x4b, 0xe8,
	0xde, 0xf0, 0x2d, 0x72, 0x9f, 0xeb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x5b, 0x42, 0xb7,
	0xd9, 0x01, 0x00, 0x00,
}
