// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package stakepoolrpc is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	GetAddedLowFeeTicketsRequest
	GetAddedLowFeeTicketsResponse
	GetIgnoredLowFeeTicketsRequest
	GetIgnoredLowFeeTicketsResponse
	GetLiveTicketsRequest
	GetLiveTicketsResponse
	PingRequest
	PingResponse
	SetAddedLowFeeTicketsRequest
	SetAddedLowFeeTicketsResponse
	SetUserVotingPrefsResponse
	SetUserVotingPrefsRequest
	TicketEntry
	UserVotingConfigEntry
	VersionRequest
	VersionResponse
*/
package stakepoolrpc

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

type GetAddedLowFeeTicketsRequest struct {
}

func (m *GetAddedLowFeeTicketsRequest) Reset()                    { *m = GetAddedLowFeeTicketsRequest{} }
func (m *GetAddedLowFeeTicketsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAddedLowFeeTicketsRequest) ProtoMessage()               {}
func (*GetAddedLowFeeTicketsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetAddedLowFeeTicketsResponse struct {
	Tickets []*TicketEntry `protobuf:"bytes,1,rep,name=tickets" json:"tickets,omitempty"`
}

func (m *GetAddedLowFeeTicketsResponse) Reset()                    { *m = GetAddedLowFeeTicketsResponse{} }
func (m *GetAddedLowFeeTicketsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAddedLowFeeTicketsResponse) ProtoMessage()               {}
func (*GetAddedLowFeeTicketsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetAddedLowFeeTicketsResponse) GetTickets() []*TicketEntry {
	if m != nil {
		return m.Tickets
	}
	return nil
}

type GetIgnoredLowFeeTicketsRequest struct {
}

func (m *GetIgnoredLowFeeTicketsRequest) Reset()                    { *m = GetIgnoredLowFeeTicketsRequest{} }
func (m *GetIgnoredLowFeeTicketsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetIgnoredLowFeeTicketsRequest) ProtoMessage()               {}
func (*GetIgnoredLowFeeTicketsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GetIgnoredLowFeeTicketsResponse struct {
	Tickets []*TicketEntry `protobuf:"bytes,1,rep,name=tickets" json:"tickets,omitempty"`
}

func (m *GetIgnoredLowFeeTicketsResponse) Reset()                    { *m = GetIgnoredLowFeeTicketsResponse{} }
func (m *GetIgnoredLowFeeTicketsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetIgnoredLowFeeTicketsResponse) ProtoMessage()               {}
func (*GetIgnoredLowFeeTicketsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetIgnoredLowFeeTicketsResponse) GetTickets() []*TicketEntry {
	if m != nil {
		return m.Tickets
	}
	return nil
}

type GetLiveTicketsRequest struct {
}

func (m *GetLiveTicketsRequest) Reset()                    { *m = GetLiveTicketsRequest{} }
func (m *GetLiveTicketsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLiveTicketsRequest) ProtoMessage()               {}
func (*GetLiveTicketsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type GetLiveTicketsResponse struct {
	Tickets []*TicketEntry `protobuf:"bytes,1,rep,name=tickets" json:"tickets,omitempty"`
}

func (m *GetLiveTicketsResponse) Reset()                    { *m = GetLiveTicketsResponse{} }
func (m *GetLiveTicketsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetLiveTicketsResponse) ProtoMessage()               {}
func (*GetLiveTicketsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetLiveTicketsResponse) GetTickets() []*TicketEntry {
	if m != nil {
		return m.Tickets
	}
	return nil
}

type PingRequest struct {
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type PingResponse struct {
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type SetAddedLowFeeTicketsRequest struct {
	Tickets []*TicketEntry `protobuf:"bytes,1,rep,name=tickets" json:"tickets,omitempty"`
}

func (m *SetAddedLowFeeTicketsRequest) Reset()                    { *m = SetAddedLowFeeTicketsRequest{} }
func (m *SetAddedLowFeeTicketsRequest) String() string            { return proto.CompactTextString(m) }
func (*SetAddedLowFeeTicketsRequest) ProtoMessage()               {}
func (*SetAddedLowFeeTicketsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SetAddedLowFeeTicketsRequest) GetTickets() []*TicketEntry {
	if m != nil {
		return m.Tickets
	}
	return nil
}

type SetAddedLowFeeTicketsResponse struct {
}

func (m *SetAddedLowFeeTicketsResponse) Reset()                    { *m = SetAddedLowFeeTicketsResponse{} }
func (m *SetAddedLowFeeTicketsResponse) String() string            { return proto.CompactTextString(m) }
func (*SetAddedLowFeeTicketsResponse) ProtoMessage()               {}
func (*SetAddedLowFeeTicketsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type SetUserVotingPrefsResponse struct {
}

func (m *SetUserVotingPrefsResponse) Reset()                    { *m = SetUserVotingPrefsResponse{} }
func (m *SetUserVotingPrefsResponse) String() string            { return proto.CompactTextString(m) }
func (*SetUserVotingPrefsResponse) ProtoMessage()               {}
func (*SetUserVotingPrefsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type SetUserVotingPrefsRequest struct {
	UserVotingConfig []*UserVotingConfigEntry `protobuf:"bytes,1,rep,name=user_voting_config,json=userVotingConfig" json:"user_voting_config,omitempty"`
}

func (m *SetUserVotingPrefsRequest) Reset()                    { *m = SetUserVotingPrefsRequest{} }
func (m *SetUserVotingPrefsRequest) String() string            { return proto.CompactTextString(m) }
func (*SetUserVotingPrefsRequest) ProtoMessage()               {}
func (*SetUserVotingPrefsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *SetUserVotingPrefsRequest) GetUserVotingConfig() []*UserVotingConfigEntry {
	if m != nil {
		return m.UserVotingConfig
	}
	return nil
}

type TicketEntry struct {
	TicketAddress string `protobuf:"bytes,1,opt,name=TicketAddress" json:"TicketAddress,omitempty"`
	TicketHash    []byte `protobuf:"bytes,2,opt,name=TicketHash,proto3" json:"TicketHash,omitempty"`
}

func (m *TicketEntry) Reset()                    { *m = TicketEntry{} }
func (m *TicketEntry) String() string            { return proto.CompactTextString(m) }
func (*TicketEntry) ProtoMessage()               {}
func (*TicketEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *TicketEntry) GetTicketAddress() string {
	if m != nil {
		return m.TicketAddress
	}
	return ""
}

func (m *TicketEntry) GetTicketHash() []byte {
	if m != nil {
		return m.TicketHash
	}
	return nil
}

type UserVotingConfigEntry struct {
	UserId          int64  `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	MultiSigAddress string `protobuf:"bytes,2,opt,name=MultiSigAddress" json:"MultiSigAddress,omitempty"`
	VoteBits        int64  `protobuf:"varint,3,opt,name=VoteBits" json:"VoteBits,omitempty"`
	VoteBitsVersion int64  `protobuf:"varint,4,opt,name=VoteBitsVersion" json:"VoteBitsVersion,omitempty"`
}

func (m *UserVotingConfigEntry) Reset()                    { *m = UserVotingConfigEntry{} }
func (m *UserVotingConfigEntry) String() string            { return proto.CompactTextString(m) }
func (*UserVotingConfigEntry) ProtoMessage()               {}
func (*UserVotingConfigEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *UserVotingConfigEntry) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserVotingConfigEntry) GetMultiSigAddress() string {
	if m != nil {
		return m.MultiSigAddress
	}
	return ""
}

func (m *UserVotingConfigEntry) GetVoteBits() int64 {
	if m != nil {
		return m.VoteBits
	}
	return 0
}

func (m *UserVotingConfigEntry) GetVoteBitsVersion() int64 {
	if m != nil {
		return m.VoteBitsVersion
	}
	return 0
}

type VersionRequest struct {
}

func (m *VersionRequest) Reset()                    { *m = VersionRequest{} }
func (m *VersionRequest) String() string            { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()               {}
func (*VersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type VersionResponse struct {
	VersionString string `protobuf:"bytes,1,opt,name=version_string,json=versionString" json:"version_string,omitempty"`
	Major         uint32 `protobuf:"varint,2,opt,name=major" json:"major,omitempty"`
	Minor         uint32 `protobuf:"varint,3,opt,name=minor" json:"minor,omitempty"`
	Patch         uint32 `protobuf:"varint,4,opt,name=patch" json:"patch,omitempty"`
	Prerelease    string `protobuf:"bytes,5,opt,name=prerelease" json:"prerelease,omitempty"`
	BuildMetadata string `protobuf:"bytes,6,opt,name=build_metadata,json=buildMetadata" json:"build_metadata,omitempty"`
}

func (m *VersionResponse) Reset()                    { *m = VersionResponse{} }
func (m *VersionResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()               {}
func (*VersionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *VersionResponse) GetVersionString() string {
	if m != nil {
		return m.VersionString
	}
	return ""
}

func (m *VersionResponse) GetMajor() uint32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *VersionResponse) GetMinor() uint32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *VersionResponse) GetPatch() uint32 {
	if m != nil {
		return m.Patch
	}
	return 0
}

func (m *VersionResponse) GetPrerelease() string {
	if m != nil {
		return m.Prerelease
	}
	return ""
}

func (m *VersionResponse) GetBuildMetadata() string {
	if m != nil {
		return m.BuildMetadata
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAddedLowFeeTicketsRequest)(nil), "stakepoolrpc.GetAddedLowFeeTicketsRequest")
	proto.RegisterType((*GetAddedLowFeeTicketsResponse)(nil), "stakepoolrpc.GetAddedLowFeeTicketsResponse")
	proto.RegisterType((*GetIgnoredLowFeeTicketsRequest)(nil), "stakepoolrpc.GetIgnoredLowFeeTicketsRequest")
	proto.RegisterType((*GetIgnoredLowFeeTicketsResponse)(nil), "stakepoolrpc.GetIgnoredLowFeeTicketsResponse")
	proto.RegisterType((*GetLiveTicketsRequest)(nil), "stakepoolrpc.GetLiveTicketsRequest")
	proto.RegisterType((*GetLiveTicketsResponse)(nil), "stakepoolrpc.GetLiveTicketsResponse")
	proto.RegisterType((*PingRequest)(nil), "stakepoolrpc.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "stakepoolrpc.PingResponse")
	proto.RegisterType((*SetAddedLowFeeTicketsRequest)(nil), "stakepoolrpc.SetAddedLowFeeTicketsRequest")
	proto.RegisterType((*SetAddedLowFeeTicketsResponse)(nil), "stakepoolrpc.SetAddedLowFeeTicketsResponse")
	proto.RegisterType((*SetUserVotingPrefsResponse)(nil), "stakepoolrpc.SetUserVotingPrefsResponse")
	proto.RegisterType((*SetUserVotingPrefsRequest)(nil), "stakepoolrpc.SetUserVotingPrefsRequest")
	proto.RegisterType((*TicketEntry)(nil), "stakepoolrpc.TicketEntry")
	proto.RegisterType((*UserVotingConfigEntry)(nil), "stakepoolrpc.UserVotingConfigEntry")
	proto.RegisterType((*VersionRequest)(nil), "stakepoolrpc.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "stakepoolrpc.VersionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for StakepooldService service

type StakepooldServiceClient interface {
	GetAddedLowFeeTickets(ctx context.Context, in *GetAddedLowFeeTicketsRequest, opts ...grpc.CallOption) (*GetAddedLowFeeTicketsResponse, error)
	GetIgnoredLowFeeTickets(ctx context.Context, in *GetIgnoredLowFeeTicketsRequest, opts ...grpc.CallOption) (*GetIgnoredLowFeeTicketsResponse, error)
	GetLiveTickets(ctx context.Context, in *GetLiveTicketsRequest, opts ...grpc.CallOption) (*GetLiveTicketsResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	SetAddedLowFeeTickets(ctx context.Context, in *SetAddedLowFeeTicketsRequest, opts ...grpc.CallOption) (*SetAddedLowFeeTicketsResponse, error)
	SetUserVotingPrefs(ctx context.Context, in *SetUserVotingPrefsRequest, opts ...grpc.CallOption) (*SetUserVotingPrefsResponse, error)
}

type stakepooldServiceClient struct {
	cc *grpc.ClientConn
}

func NewStakepooldServiceClient(cc *grpc.ClientConn) StakepooldServiceClient {
	return &stakepooldServiceClient{cc}
}

func (c *stakepooldServiceClient) GetAddedLowFeeTickets(ctx context.Context, in *GetAddedLowFeeTicketsRequest, opts ...grpc.CallOption) (*GetAddedLowFeeTicketsResponse, error) {
	out := new(GetAddedLowFeeTicketsResponse)
	err := grpc.Invoke(ctx, "/stakepoolrpc.StakepooldService/GetAddedLowFeeTickets", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stakepooldServiceClient) GetIgnoredLowFeeTickets(ctx context.Context, in *GetIgnoredLowFeeTicketsRequest, opts ...grpc.CallOption) (*GetIgnoredLowFeeTicketsResponse, error) {
	out := new(GetIgnoredLowFeeTicketsResponse)
	err := grpc.Invoke(ctx, "/stakepoolrpc.StakepooldService/GetIgnoredLowFeeTickets", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stakepooldServiceClient) GetLiveTickets(ctx context.Context, in *GetLiveTicketsRequest, opts ...grpc.CallOption) (*GetLiveTicketsResponse, error) {
	out := new(GetLiveTicketsResponse)
	err := grpc.Invoke(ctx, "/stakepoolrpc.StakepooldService/GetLiveTickets", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stakepooldServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/stakepoolrpc.StakepooldService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stakepooldServiceClient) SetAddedLowFeeTickets(ctx context.Context, in *SetAddedLowFeeTicketsRequest, opts ...grpc.CallOption) (*SetAddedLowFeeTicketsResponse, error) {
	out := new(SetAddedLowFeeTicketsResponse)
	err := grpc.Invoke(ctx, "/stakepoolrpc.StakepooldService/SetAddedLowFeeTickets", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stakepooldServiceClient) SetUserVotingPrefs(ctx context.Context, in *SetUserVotingPrefsRequest, opts ...grpc.CallOption) (*SetUserVotingPrefsResponse, error) {
	out := new(SetUserVotingPrefsResponse)
	err := grpc.Invoke(ctx, "/stakepoolrpc.StakepooldService/SetUserVotingPrefs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StakepooldService service

type StakepooldServiceServer interface {
	GetAddedLowFeeTickets(context.Context, *GetAddedLowFeeTicketsRequest) (*GetAddedLowFeeTicketsResponse, error)
	GetIgnoredLowFeeTickets(context.Context, *GetIgnoredLowFeeTicketsRequest) (*GetIgnoredLowFeeTicketsResponse, error)
	GetLiveTickets(context.Context, *GetLiveTicketsRequest) (*GetLiveTicketsResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	SetAddedLowFeeTickets(context.Context, *SetAddedLowFeeTicketsRequest) (*SetAddedLowFeeTicketsResponse, error)
	SetUserVotingPrefs(context.Context, *SetUserVotingPrefsRequest) (*SetUserVotingPrefsResponse, error)
}

func RegisterStakepooldServiceServer(s *grpc.Server, srv StakepooldServiceServer) {
	s.RegisterService(&_StakepooldService_serviceDesc, srv)
}

func _StakepooldService_GetAddedLowFeeTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddedLowFeeTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakepooldServiceServer).GetAddedLowFeeTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stakepoolrpc.StakepooldService/GetAddedLowFeeTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakepooldServiceServer).GetAddedLowFeeTickets(ctx, req.(*GetAddedLowFeeTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StakepooldService_GetIgnoredLowFeeTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIgnoredLowFeeTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakepooldServiceServer).GetIgnoredLowFeeTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stakepoolrpc.StakepooldService/GetIgnoredLowFeeTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakepooldServiceServer).GetIgnoredLowFeeTickets(ctx, req.(*GetIgnoredLowFeeTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StakepooldService_GetLiveTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLiveTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakepooldServiceServer).GetLiveTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stakepoolrpc.StakepooldService/GetLiveTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakepooldServiceServer).GetLiveTickets(ctx, req.(*GetLiveTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StakepooldService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakepooldServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stakepoolrpc.StakepooldService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakepooldServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StakepooldService_SetAddedLowFeeTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAddedLowFeeTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakepooldServiceServer).SetAddedLowFeeTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stakepoolrpc.StakepooldService/SetAddedLowFeeTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakepooldServiceServer).SetAddedLowFeeTickets(ctx, req.(*SetAddedLowFeeTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StakepooldService_SetUserVotingPrefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserVotingPrefsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakepooldServiceServer).SetUserVotingPrefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stakepoolrpc.StakepooldService/SetUserVotingPrefs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakepooldServiceServer).SetUserVotingPrefs(ctx, req.(*SetUserVotingPrefsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StakepooldService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stakepoolrpc.StakepooldService",
	HandlerType: (*StakepooldServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAddedLowFeeTickets",
			Handler:    _StakepooldService_GetAddedLowFeeTickets_Handler,
		},
		{
			MethodName: "GetIgnoredLowFeeTickets",
			Handler:    _StakepooldService_GetIgnoredLowFeeTickets_Handler,
		},
		{
			MethodName: "GetLiveTickets",
			Handler:    _StakepooldService_GetLiveTickets_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _StakepooldService_Ping_Handler,
		},
		{
			MethodName: "SetAddedLowFeeTickets",
			Handler:    _StakepooldService_SetAddedLowFeeTickets_Handler,
		},
		{
			MethodName: "SetUserVotingPrefs",
			Handler:    _StakepooldService_SetUserVotingPrefs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// Client API for VersionService service

type VersionServiceClient interface {
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type versionServiceClient struct {
	cc *grpc.ClientConn
}

func NewVersionServiceClient(cc *grpc.ClientConn) VersionServiceClient {
	return &versionServiceClient{cc}
}

func (c *versionServiceClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := grpc.Invoke(ctx, "/stakepoolrpc.VersionService/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VersionService service

type VersionServiceServer interface {
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
}

func RegisterVersionServiceServer(s *grpc.Server, srv VersionServiceServer) {
	s.RegisterService(&_VersionService_serviceDesc, srv)
}

func _VersionService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stakepoolrpc.VersionService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServiceServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VersionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stakepoolrpc.VersionService",
	HandlerType: (*VersionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _VersionService_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdd, 0x4e, 0xdb, 0x4c,
	0x10, 0x55, 0xf8, 0xfd, 0x18, 0x48, 0xe0, 0x5b, 0x15, 0x30, 0x56, 0x12, 0x22, 0x43, 0xd5, 0xa8,
	0x3f, 0xb9, 0x08, 0xd7, 0xbd, 0xa0, 0x55, 0x49, 0x91, 0x88, 0x44, 0x6d, 0x1a, 0x55, 0xea, 0x45,
	0x64, 0xe2, 0xc1, 0x6c, 0x09, 0x5e, 0x77, 0x77, 0x93, 0xaa, 0x4f, 0xd3, 0x67, 0xe9, 0x3b, 0xf5,
	0x01, 0x2a, 0xef, 0xda, 0x21, 0xb6, 0x13, 0x37, 0x2d, 0x77, 0x9e, 0x33, 0x93, 0x73, 0x8e, 0xd7,
	0xb3, 0x27, 0xb0, 0xe1, 0x86, 0xb4, 0x15, 0x72, 0x26, 0x19, 0xd9, 0x12, 0xd2, 0xbd, 0xc3, 0x90,
	0xb1, 0x21, 0x0f, 0x07, 0x56, 0x1d, 0xaa, 0x1d, 0x94, 0xa7, 0x9e, 0x87, 0xde, 0x05, 0xfb, 0x76,
	0x86, 0x78, 0x45, 0x07, 0x77, 0x28, 0x85, 0x8d, 0x5f, 0x47, 0x28, 0xa4, 0x75, 0x05, 0xb5, 0x39,
	0x7d, 0x11, 0xb2, 0x40, 0x20, 0x39, 0x81, 0x75, 0xa9, 0x21, 0xa3, 0xd4, 0x58, 0x6e, 0x6e, 0xb6,
	0x0f, 0x5a, 0xd3, 0x02, 0x2d, 0x3d, 0xff, 0x2e, 0x90, 0xfc, 0xbb, 0x9d, 0x4c, 0x5a, 0x0d, 0xa8,
	0x77, 0x50, 0x9e, 0xfb, 0x01, 0xe3, 0x73, 0x74, 0x7b, 0x70, 0x38, 0x77, 0xe2, 0x31, 0xca, 0xfb,
	0xb0, 0xdb, 0x41, 0x79, 0x41, 0xc7, 0x59, 0xc1, 0x2e, 0xec, 0x65, 0x1b, 0x8f, 0xd1, 0x29, 0xc3,
	0xe6, 0x25, 0x0d, 0xfc, 0x84, 0xbd, 0x02, 0x5b, 0xba, 0xd4, 0x9c, 0x96, 0x03, 0x55, 0xa7, 0xe0,
	0xd8, 0xff, 0x4d, 0xf3, 0x10, 0x6a, 0x4e, 0xd1, 0xb7, 0xb2, 0xaa, 0x60, 0x3a, 0x28, 0x3f, 0x0a,
	0xe4, 0x3d, 0x26, 0x69, 0xe0, 0x5f, 0x72, 0xbc, 0x79, 0xe8, 0x06, 0x70, 0x30, 0xab, 0xab, 0x0d,
	0x7d, 0x00, 0x32, 0x12, 0xc8, 0xfb, 0x63, 0xd5, 0xea, 0x0f, 0x58, 0x70, 0x43, 0xfd, 0xd8, 0xdb,
	0x51, 0xda, 0xdb, 0x03, 0xc3, 0x5b, 0x35, 0xa5, 0x5d, 0xee, 0x8c, 0x32, 0xb0, 0xe5, 0xc0, 0xe6,
	0xd4, 0x6b, 0x90, 0x63, 0x28, 0xeb, 0xf2, 0xd4, 0xf3, 0x38, 0x8a, 0xe8, 0xc5, 0x4b, 0xcd, 0x0d,
	0x3b, 0x0d, 0x92, 0x3a, 0x80, 0x06, 0xde, 0xbb, 0xe2, 0xd6, 0x58, 0x6a, 0x94, 0x9a, 0x5b, 0xf6,
	0x14, 0x62, 0xfd, 0x28, 0xc1, 0xee, 0x4c, 0x03, 0x64, 0x0f, 0xd6, 0xa2, 0xc6, 0xb9, 0xa7, 0x88,
	0x97, 0xed, 0xb8, 0x22, 0x4d, 0xd8, 0xee, 0x8e, 0x86, 0x92, 0x3a, 0xd4, 0x4f, 0x94, 0x97, 0x94,
	0x72, 0x16, 0x26, 0x26, 0xfc, 0xd7, 0x63, 0x12, 0xdf, 0x50, 0x29, 0x8c, 0x65, 0xc5, 0x31, 0xa9,
	0x23, 0x96, 0xe4, 0xb9, 0x87, 0x5c, 0x50, 0x16, 0x18, 0x2b, 0x6a, 0x24, 0x0b, 0x5b, 0x3b, 0x50,
	0x89, 0x1f, 0x93, 0xe5, 0xf8, 0x59, 0x82, 0xed, 0x09, 0x14, 0x2f, 0xdd, 0x53, 0xa8, 0x8c, 0x35,
	0xd4, 0x17, 0x92, 0xd3, 0xc0, 0x4f, 0x8e, 0x23, 0x46, 0x1d, 0x05, 0x92, 0x27, 0xb0, 0x7a, 0xef,
	0x7e, 0x61, 0x5c, 0x59, 0x2e, 0xdb, 0xba, 0x50, 0x28, 0x0d, 0x18, 0x57, 0x2e, 0x23, 0x34, 0x2a,
	0x22, 0x34, 0x74, 0xe5, 0xe0, 0x56, 0x19, 0x2b, 0xdb, 0xba, 0x88, 0x0e, 0x34, 0xe4, 0xc8, 0x71,
	0x88, 0xae, 0x40, 0x63, 0x55, 0x89, 0x4c, 0x21, 0x91, 0x91, 0xeb, 0x11, 0x1d, 0x7a, 0xfd, 0x7b,
	0x94, 0xae, 0xe7, 0x4a, 0xd7, 0x58, 0xd3, 0x46, 0x14, 0xda, 0x8d, 0xc1, 0xf6, 0xaf, 0x15, 0xf8,
	0xdf, 0x49, 0xb6, 0xc0, 0x73, 0x90, 0x8f, 0xe9, 0x00, 0x49, 0xa8, 0x6e, 0x5b, 0x7e, 0x23, 0xc9,
	0xf3, 0xf4, 0xca, 0x14, 0x45, 0x90, 0xf9, 0x62, 0xa1, 0xd9, 0xf8, 0xdc, 0xc6, 0xb0, 0x3f, 0x27,
	0x37, 0xc8, 0xcb, 0x1c, 0x4f, 0x41, 0x00, 0x99, 0xaf, 0x16, 0x9c, 0x8e, 0x75, 0x3f, 0x43, 0x25,
	0x1d, 0x1f, 0xe4, 0x28, 0x47, 0x90, 0x4f, 0x1d, 0xf3, 0xb8, 0x78, 0x28, 0x26, 0x7f, 0x0d, 0x2b,
	0x51, 0x7a, 0x90, 0x4c, 0x08, 0x4c, 0x05, 0x8c, 0x69, 0xce, 0x6a, 0xc5, 0x3f, 0x0f, 0x61, 0xd7,
	0x59, 0xe4, 0x2b, 0x38, 0x7f, 0xf1, 0x15, 0x0a, 0x83, 0x86, 0xf8, 0x40, 0xf2, 0x51, 0x42, 0x9e,
	0xe5, 0x28, 0x66, 0x87, 0x8d, 0xd9, 0xfc, 0xf3, 0xa0, 0x16, 0x6a, 0x7f, 0x9a, 0x5c, 0xa6, 0x64,
	0xe5, 0xce, 0x60, 0x3d, 0x46, 0x48, 0x35, 0x4d, 0x93, 0xbe, 0x75, 0x66, 0x6d, 0x4e, 0x57, 0x33,
	0x5f, 0xaf, 0xa9, 0x7f, 0xcb, 0x93, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x57, 0xe8, 0x76,
	0x3a, 0x07, 0x00, 0x00,
}
