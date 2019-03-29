// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/identityserver.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/lyft/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AuthInfoResponse struct {
	// Types that are valid to be assigned to AccessMethod:
	//	*AuthInfoResponse_APIKey
	//	*AuthInfoResponse_OAuthAccessToken
	AccessMethod         isAuthInfoResponse_AccessMethod `protobuf_oneof:"access_method"`
	UniversalRights      *Rights                         `protobuf:"bytes,3,opt,name=universal_rights,json=universalRights,proto3" json:"universal_rights,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *AuthInfoResponse) Reset()      { *m = AuthInfoResponse{} }
func (*AuthInfoResponse) ProtoMessage() {}
func (*AuthInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c7e02f6181562c, []int{0}
}
func (m *AuthInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthInfoResponse.Merge(m, src)
}
func (m *AuthInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *AuthInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthInfoResponse proto.InternalMessageInfo

type isAuthInfoResponse_AccessMethod interface {
	isAuthInfoResponse_AccessMethod()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type AuthInfoResponse_APIKey struct {
	APIKey *AuthInfoResponse_APIKeyAccess `protobuf:"bytes,1,opt,name=api_key,json=apiKey,proto3,oneof"`
}
type AuthInfoResponse_OAuthAccessToken struct {
	OAuthAccessToken *OAuthAccessToken `protobuf:"bytes,2,opt,name=oauth_access_token,json=oauthAccessToken,proto3,oneof"`
}

func (*AuthInfoResponse_APIKey) isAuthInfoResponse_AccessMethod()           {}
func (*AuthInfoResponse_OAuthAccessToken) isAuthInfoResponse_AccessMethod() {}

func (m *AuthInfoResponse) GetAccessMethod() isAuthInfoResponse_AccessMethod {
	if m != nil {
		return m.AccessMethod
	}
	return nil
}

func (m *AuthInfoResponse) GetAPIKey() *AuthInfoResponse_APIKeyAccess {
	if x, ok := m.GetAccessMethod().(*AuthInfoResponse_APIKey); ok {
		return x.APIKey
	}
	return nil
}

func (m *AuthInfoResponse) GetOAuthAccessToken() *OAuthAccessToken {
	if x, ok := m.GetAccessMethod().(*AuthInfoResponse_OAuthAccessToken); ok {
		return x.OAuthAccessToken
	}
	return nil
}

func (m *AuthInfoResponse) GetUniversalRights() *Rights {
	if m != nil {
		return m.UniversalRights
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AuthInfoResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AuthInfoResponse_OneofMarshaler, _AuthInfoResponse_OneofUnmarshaler, _AuthInfoResponse_OneofSizer, []interface{}{
		(*AuthInfoResponse_APIKey)(nil),
		(*AuthInfoResponse_OAuthAccessToken)(nil),
	}
}

func _AuthInfoResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AuthInfoResponse)
	// access_method
	switch x := m.AccessMethod.(type) {
	case *AuthInfoResponse_APIKey:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.APIKey); err != nil {
			return err
		}
	case *AuthInfoResponse_OAuthAccessToken:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OAuthAccessToken); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AuthInfoResponse.AccessMethod has unexpected type %T", x)
	}
	return nil
}

func _AuthInfoResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AuthInfoResponse)
	switch tag {
	case 1: // access_method.api_key
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AuthInfoResponse_APIKeyAccess)
		err := b.DecodeMessage(msg)
		m.AccessMethod = &AuthInfoResponse_APIKey{msg}
		return true, err
	case 2: // access_method.oauth_access_token
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(OAuthAccessToken)
		err := b.DecodeMessage(msg)
		m.AccessMethod = &AuthInfoResponse_OAuthAccessToken{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AuthInfoResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AuthInfoResponse)
	// access_method
	switch x := m.AccessMethod.(type) {
	case *AuthInfoResponse_APIKey:
		s := proto.Size(x.APIKey)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AuthInfoResponse_OAuthAccessToken:
		s := proto.Size(x.OAuthAccessToken)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type AuthInfoResponse_APIKeyAccess struct {
	APIKey               `protobuf:"bytes,1,opt,name=api_key,json=apiKey,proto3,embedded=api_key" json:"api_key"`
	EntityIDs            EntityIdentifiers `protobuf:"bytes,2,opt,name=entity_ids,json=entityIds,proto3" json:"entity_ids"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AuthInfoResponse_APIKeyAccess) Reset()      { *m = AuthInfoResponse_APIKeyAccess{} }
func (*AuthInfoResponse_APIKeyAccess) ProtoMessage() {}
func (*AuthInfoResponse_APIKeyAccess) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c7e02f6181562c, []int{0, 0}
}
func (m *AuthInfoResponse_APIKeyAccess) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthInfoResponse_APIKeyAccess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthInfoResponse_APIKeyAccess.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthInfoResponse_APIKeyAccess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthInfoResponse_APIKeyAccess.Merge(m, src)
}
func (m *AuthInfoResponse_APIKeyAccess) XXX_Size() int {
	return m.Size()
}
func (m *AuthInfoResponse_APIKeyAccess) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthInfoResponse_APIKeyAccess.DiscardUnknown(m)
}

var xxx_messageInfo_AuthInfoResponse_APIKeyAccess proto.InternalMessageInfo

func (m *AuthInfoResponse_APIKeyAccess) GetEntityIDs() EntityIdentifiers {
	if m != nil {
		return m.EntityIDs
	}
	return EntityIdentifiers{}
}

func init() {
	proto.RegisterType((*AuthInfoResponse)(nil), "ttn.lorawan.v3.AuthInfoResponse")
	golang_proto.RegisterType((*AuthInfoResponse)(nil), "ttn.lorawan.v3.AuthInfoResponse")
	proto.RegisterType((*AuthInfoResponse_APIKeyAccess)(nil), "ttn.lorawan.v3.AuthInfoResponse.APIKeyAccess")
	golang_proto.RegisterType((*AuthInfoResponse_APIKeyAccess)(nil), "ttn.lorawan.v3.AuthInfoResponse.APIKeyAccess")
}

func init() {
	proto.RegisterFile("lorawan-stack/api/identityserver.proto", fileDescriptor_a1c7e02f6181562c)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/identityserver.proto", fileDescriptor_a1c7e02f6181562c)
}

var fileDescriptor_a1c7e02f6181562c = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x68, 0x14, 0x4d,
	0x14, 0x9e, 0x49, 0xf8, 0xf3, 0x9b, 0x31, 0x9a, 0x63, 0x91, 0x10, 0x2e, 0xfa, 0x2e, 0x46, 0x10,
	0x0b, 0x6f, 0x17, 0x92, 0xd2, 0xea, 0x4e, 0x03, 0x09, 0x29, 0x94, 0x23, 0x85, 0x08, 0x72, 0xcc,
	0xdd, 0xcd, 0xed, 0x0e, 0x77, 0x99, 0x59, 0x76, 0xe7, 0x2e, 0x5c, 0x17, 0xac, 0x82, 0x95, 0x60,
	0x63, 0x29, 0x56, 0x29, 0x83, 0x55, 0xb0, 0x4a, 0x99, 0x32, 0x60, 0x13, 0x9b, 0x23, 0x3b, 0x6b,
	0x11, 0xbb, 0x94, 0x29, 0xe5, 0x66, 0x27, 0xe7, 0xe5, 0x0e, 0xb1, 0x9b, 0x79, 0xdf, 0xf7, 0xbe,
	0xf7, 0xbe, 0xf7, 0x1e, 0x79, 0xdc, 0x96, 0x11, 0xdd, 0xa5, 0xa2, 0x18, 0x2b, 0x5a, 0x6f, 0x79,
	0x34, 0xe4, 0x1e, 0x6f, 0x30, 0xa1, 0xb8, 0xea, 0xc5, 0x2c, 0xea, 0xb2, 0xc8, 0x0d, 0x23, 0xa9,
	0xa4, 0x73, 0x57, 0x29, 0xe1, 0x5a, 0xae, 0xdb, 0x5d, 0xcb, 0x17, 0x7d, 0xae, 0x82, 0x4e, 0xcd,
	0xad, 0xcb, 0x1d, 0xcf, 0x97, 0xbe, 0xf4, 0x0c, 0xad, 0xd6, 0x69, 0x9a, 0x9f, 0xf9, 0x98, 0x57,
	0x96, 0x9e, 0x7f, 0x36, 0x42, 0x6f, 0xf7, 0x9a, 0x2a, 0xa3, 0xd7, 0x8b, 0x3e, 0x13, 0xc5, 0x2e,
	0x6d, 0xf3, 0x06, 0x55, 0xcc, 0x9b, 0x78, 0xd8, 0xe4, 0xfb, 0xbe, 0x94, 0x7e, 0x9b, 0x99, 0xe6,
	0xa8, 0x10, 0x52, 0x51, 0xc5, 0xa5, 0x88, 0x2d, 0xba, 0x64, 0xd1, 0x61, 0x03, 0x6c, 0x27, 0x54,
	0x3d, 0x0b, 0x3e, 0xfa, 0x9b, 0xbd, 0x26, 0x67, 0xd1, 0xb5, 0xc2, 0x83, 0x49, 0x92, 0xa4, 0x1d,
	0x15, 0x58, 0x18, 0x26, 0xe1, 0x88, 0xfb, 0x81, 0xb2, 0xe9, 0x2b, 0x3f, 0xa6, 0x49, 0xae, 0xd4,
	0x51, 0xc1, 0xa6, 0x68, 0xca, 0x0a, 0x8b, 0x43, 0x29, 0x62, 0xe6, 0x6c, 0x93, 0xff, 0x69, 0xc8,
	0xab, 0x2d, 0xd6, 0x5b, 0xc4, 0xcb, 0xf8, 0xc9, 0xed, 0xd5, 0xa2, 0x7b, 0x73, 0x82, 0xee, 0x78,
	0x8a, 0x5b, 0x7a, 0xb5, 0xb9, 0xc5, 0x7a, 0xa5, 0x7a, 0x9d, 0xc5, 0x71, 0x99, 0xe8, 0x7e, 0x61,
	0x26, 0x8b, 0x6c, 0xa0, 0xca, 0x0c, 0x0d, 0xf9, 0x16, 0xeb, 0x39, 0x4d, 0xe2, 0x98, 0xce, 0xaa,
	0xd4, 0xb0, 0xaa, 0x4a, 0xb6, 0x98, 0x58, 0x9c, 0x32, 0x05, 0x96, 0xc7, 0x0b, 0xbc, 0x1c, 0x54,
	0xc8, 0xe4, 0xb6, 0x07, 0xbc, 0xf2, 0x3d, 0xdd, 0x2f, 0xe4, 0xc6, 0xa3, 0x1b, 0xa8, 0x92, 0x33,
	0x9a, 0x23, 0x31, 0xa7, 0x44, 0x72, 0x1d, 0xc1, 0xbb, 0x2c, 0x8a, 0x69, 0xbb, 0x9a, 0x99, 0x5d,
	0x9c, 0x36, 0x55, 0x16, 0xc6, 0xab, 0x54, 0x0c, 0x5a, 0x99, 0x1f, 0xf2, 0xb3, 0x40, 0xfe, 0x2b,
	0x26, 0x73, 0xa3, 0x8e, 0x9c, 0xe7, 0xe3, 0x13, 0x99, 0x90, 0xca, 0xe8, 0x65, 0xe7, 0xa4, 0x5f,
	0x40, 0xa7, 0xfd, 0x02, 0xfe, 0xf6, 0xeb, 0x78, 0xfa, 0xbf, 0xf7, 0x78, 0x2a, 0x87, 0x87, 0x03,
	0x78, 0x4b, 0x48, 0x76, 0x9c, 0x55, 0xde, 0x88, 0xad, 0xf1, 0x87, 0xe3, 0x3a, 0xeb, 0x86, 0xb1,
	0xf9, 0x67, 0xcf, 0xe5, 0xa5, 0x81, 0xa4, 0xee, 0x17, 0x66, 0x2d, 0xf4, 0x22, 0x1e, 0xd1, 0x9e,
	0x65, 0x96, 0x1f, 0x97, 0xe7, 0xc9, 0x1d, 0x3b, 0xd9, 0x1d, 0xa6, 0x02, 0xd9, 0x58, 0x0d, 0xc8,
	0x5c, 0x96, 0x62, 0x4d, 0xbc, 0x26, 0xb7, 0xae, 0xf7, 0xe6, 0x2c, 0xb8, 0xd9, 0xe5, 0xb9, 0xd7,
	0x97, 0xe7, 0xae, 0x0f, 0x2e, 0x2f, 0xbf, 0xfc, 0xaf, 0x4d, 0xaf, 0x38, 0xef, 0xbe, 0xff, 0xfc,
	0x38, 0x35, 0xe7, 0x10, 0xcf, 0x2c, 0x93, 0x8b, 0xa6, 0x2c, 0x7f, 0xc1, 0x27, 0x09, 0xe0, 0xd3,
	0x04, 0xf0, 0x59, 0x02, 0xe8, 0x3c, 0x01, 0x74, 0x91, 0x00, 0xba, 0x4c, 0x00, 0x5d, 0x25, 0x80,
	0xf7, 0x34, 0xe0, 0x7d, 0x0d, 0xe8, 0x40, 0x03, 0x3e, 0xd4, 0x80, 0x8e, 0x34, 0xa0, 0x63, 0x0d,
	0xe8, 0x44, 0x03, 0x3e, 0xd5, 0x80, 0xcf, 0x34, 0xa0, 0x73, 0x0d, 0xf8, 0x42, 0x03, 0xba, 0xd4,
	0x80, 0xaf, 0x34, 0xa0, 0xbd, 0x14, 0xd0, 0x7e, 0x0a, 0xf8, 0x43, 0x0a, 0xe8, 0x53, 0x0a, 0xf8,
	0x73, 0x0a, 0xe8, 0x20, 0x05, 0x74, 0x98, 0x02, 0x3e, 0x4a, 0x01, 0x1f, 0xa7, 0x80, 0xdf, 0x3c,
	0xf5, 0xa5, 0xab, 0x02, 0xa6, 0x02, 0x2e, 0xfc, 0xd8, 0x15, 0x4c, 0xed, 0xca, 0xa8, 0xe5, 0xdd,
	0xbc, 0xf8, 0xb0, 0xe5, 0x7b, 0x4a, 0x89, 0xb0, 0x56, 0x9b, 0x31, 0x56, 0xd7, 0x7e, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xb4, 0x9b, 0x8f, 0x17, 0x36, 0x04, 0x00, 0x00,
}

func (this *AuthInfoResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AuthInfoResponse)
	if !ok {
		that2, ok := that.(AuthInfoResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.AccessMethod == nil {
		if this.AccessMethod != nil {
			return false
		}
	} else if this.AccessMethod == nil {
		return false
	} else if !this.AccessMethod.Equal(that1.AccessMethod) {
		return false
	}
	if !this.UniversalRights.Equal(that1.UniversalRights) {
		return false
	}
	return true
}
func (this *AuthInfoResponse_APIKey) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AuthInfoResponse_APIKey)
	if !ok {
		that2, ok := that.(AuthInfoResponse_APIKey)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.APIKey.Equal(that1.APIKey) {
		return false
	}
	return true
}
func (this *AuthInfoResponse_OAuthAccessToken) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AuthInfoResponse_OAuthAccessToken)
	if !ok {
		that2, ok := that.(AuthInfoResponse_OAuthAccessToken)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.OAuthAccessToken.Equal(that1.OAuthAccessToken) {
		return false
	}
	return true
}
func (this *AuthInfoResponse_APIKeyAccess) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AuthInfoResponse_APIKeyAccess)
	if !ok {
		that2, ok := that.(AuthInfoResponse_APIKeyAccess)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.APIKey.Equal(&that1.APIKey) {
		return false
	}
	if !this.EntityIDs.Equal(&that1.EntityIDs) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EntityAccessClient is the client API for EntityAccess service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EntityAccessClient interface {
	// AuthInfo returns information about the authentication that is used on the request.
	AuthInfo(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*AuthInfoResponse, error)
}

type entityAccessClient struct {
	cc *grpc.ClientConn
}

func NewEntityAccessClient(cc *grpc.ClientConn) EntityAccessClient {
	return &entityAccessClient{cc}
}

func (c *entityAccessClient) AuthInfo(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*AuthInfoResponse, error) {
	out := new(AuthInfoResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.EntityAccess/AuthInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntityAccessServer is the server API for EntityAccess service.
type EntityAccessServer interface {
	// AuthInfo returns information about the authentication that is used on the request.
	AuthInfo(context.Context, *types.Empty) (*AuthInfoResponse, error)
}

func RegisterEntityAccessServer(s *grpc.Server, srv EntityAccessServer) {
	s.RegisterService(&_EntityAccess_serviceDesc, srv)
}

func _EntityAccess_AuthInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityAccessServer).AuthInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.EntityAccess/AuthInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityAccessServer).AuthInfo(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _EntityAccess_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.EntityAccess",
	HandlerType: (*EntityAccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthInfo",
			Handler:    _EntityAccess_AuthInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/identityserver.proto",
}

func (m *AuthInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AccessMethod != nil {
		nn1, err := m.AccessMethod.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.UniversalRights != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintIdentityserver(dAtA, i, uint64(m.UniversalRights.Size()))
		n2, err := m.UniversalRights.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *AuthInfoResponse_APIKey) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.APIKey != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintIdentityserver(dAtA, i, uint64(m.APIKey.Size()))
		n3, err := m.APIKey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *AuthInfoResponse_OAuthAccessToken) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.OAuthAccessToken != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintIdentityserver(dAtA, i, uint64(m.OAuthAccessToken.Size()))
		n4, err := m.OAuthAccessToken.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *AuthInfoResponse_APIKeyAccess) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthInfoResponse_APIKeyAccess) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintIdentityserver(dAtA, i, uint64(m.APIKey.Size()))
	n5, err := m.APIKey.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	dAtA[i] = 0x12
	i++
	i = encodeVarintIdentityserver(dAtA, i, uint64(m.EntityIDs.Size()))
	n6, err := m.EntityIDs.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	return i, nil
}

func encodeVarintIdentityserver(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedAuthInfoResponse(r randyIdentityserver, easy bool) *AuthInfoResponse {
	this := &AuthInfoResponse{}
	oneofNumber_AccessMethod := []int32{1, 2}[r.Intn(2)]
	switch oneofNumber_AccessMethod {
	case 1:
		this.AccessMethod = NewPopulatedAuthInfoResponse_APIKey(r, easy)
	case 2:
		this.AccessMethod = NewPopulatedAuthInfoResponse_OAuthAccessToken(r, easy)
	}
	if r.Intn(10) != 0 {
		this.UniversalRights = NewPopulatedRights(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedAuthInfoResponse_APIKey(r randyIdentityserver, easy bool) *AuthInfoResponse_APIKey {
	this := &AuthInfoResponse_APIKey{}
	this.APIKey = NewPopulatedAuthInfoResponse_APIKeyAccess(r, easy)
	return this
}
func NewPopulatedAuthInfoResponse_OAuthAccessToken(r randyIdentityserver, easy bool) *AuthInfoResponse_OAuthAccessToken {
	this := &AuthInfoResponse_OAuthAccessToken{}
	this.OAuthAccessToken = NewPopulatedOAuthAccessToken(r, easy)
	return this
}
func NewPopulatedAuthInfoResponse_APIKeyAccess(r randyIdentityserver, easy bool) *AuthInfoResponse_APIKeyAccess {
	this := &AuthInfoResponse_APIKeyAccess{}
	v1 := NewPopulatedAPIKey(r, easy)
	this.APIKey = *v1
	v2 := NewPopulatedEntityIdentifiers(r, easy)
	this.EntityIDs = *v2
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyIdentityserver interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneIdentityserver(r randyIdentityserver) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringIdentityserver(r randyIdentityserver) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneIdentityserver(r)
	}
	return string(tmps)
}
func randUnrecognizedIdentityserver(r randyIdentityserver, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldIdentityserver(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldIdentityserver(dAtA []byte, r randyIdentityserver, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateIdentityserver(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateIdentityserver(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateIdentityserver(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateIdentityserver(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateIdentityserver(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateIdentityserver(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateIdentityserver(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *AuthInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AccessMethod != nil {
		n += m.AccessMethod.Size()
	}
	if m.UniversalRights != nil {
		l = m.UniversalRights.Size()
		n += 1 + l + sovIdentityserver(uint64(l))
	}
	return n
}

func (m *AuthInfoResponse_APIKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.APIKey != nil {
		l = m.APIKey.Size()
		n += 1 + l + sovIdentityserver(uint64(l))
	}
	return n
}
func (m *AuthInfoResponse_OAuthAccessToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OAuthAccessToken != nil {
		l = m.OAuthAccessToken.Size()
		n += 1 + l + sovIdentityserver(uint64(l))
	}
	return n
}
func (m *AuthInfoResponse_APIKeyAccess) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.APIKey.Size()
	n += 1 + l + sovIdentityserver(uint64(l))
	l = m.EntityIDs.Size()
	n += 1 + l + sovIdentityserver(uint64(l))
	return n
}

func sovIdentityserver(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozIdentityserver(x uint64) (n int) {
	return sovIdentityserver((x << 1) ^ uint64((int64(x) >> 63)))
}
func (this *AuthInfoResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AuthInfoResponse{`,
		`AccessMethod:` + fmt.Sprintf("%v", this.AccessMethod) + `,`,
		`UniversalRights:` + strings.Replace(fmt.Sprintf("%v", this.UniversalRights), "Rights", "Rights", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AuthInfoResponse_APIKey) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AuthInfoResponse_APIKey{`,
		`APIKey:` + strings.Replace(fmt.Sprintf("%v", this.APIKey), "AuthInfoResponse_APIKeyAccess", "AuthInfoResponse_APIKeyAccess", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AuthInfoResponse_OAuthAccessToken) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AuthInfoResponse_OAuthAccessToken{`,
		`OAuthAccessToken:` + strings.Replace(fmt.Sprintf("%v", this.OAuthAccessToken), "OAuthAccessToken", "OAuthAccessToken", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AuthInfoResponse_APIKeyAccess) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AuthInfoResponse_APIKeyAccess{`,
		`APIKey:` + strings.Replace(strings.Replace(this.APIKey.String(), "APIKey", "APIKey", 1), `&`, ``, 1) + `,`,
		`EntityIDs:` + strings.Replace(strings.Replace(this.EntityIDs.String(), "EntityIdentifiers", "EntityIdentifiers", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringIdentityserver(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AuthInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentityserver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuthInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field APIKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentityserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIdentityserver
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdentityserver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &AuthInfoResponse_APIKeyAccess{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.AccessMethod = &AuthInfoResponse_APIKey{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OAuthAccessToken", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentityserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIdentityserver
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdentityserver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &OAuthAccessToken{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.AccessMethod = &AuthInfoResponse_OAuthAccessToken{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UniversalRights", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentityserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIdentityserver
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdentityserver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UniversalRights == nil {
				m.UniversalRights = &Rights{}
			}
			if err := m.UniversalRights.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentityserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIdentityserver
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIdentityserver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthInfoResponse_APIKeyAccess) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentityserver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: APIKeyAccess: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: APIKeyAccess: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field APIKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentityserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIdentityserver
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdentityserver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.APIKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityIDs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentityserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIdentityserver
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdentityserver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EntityIDs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentityserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIdentityserver
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIdentityserver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipIdentityserver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIdentityserver
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIdentityserver
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIdentityserver
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthIdentityserver
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthIdentityserver
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIdentityserver
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipIdentityserver(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthIdentityserver
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthIdentityserver = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIdentityserver   = fmt.Errorf("proto: integer overflow")
)