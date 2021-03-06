// Code generated by protoc-gen-go.
// source: hreg.proto
// DO NOT EDIT!

/*
Package hcr is a generated protocol buffer package.

Version 1.0

It is generated from these files:
	hreg.proto

It has these top-level messages:
	Auth
	Status
	GetMoteIDParams
	GetMoteIDResponse
	CreateInstanceParams
	CreateInstanceResponse
	CreateDeploymentParams
	CreateDeploymentResponse
	BindMoteParams
	BindMoteResponse
	MoteInfoParams
	MoteInfoResponse
	InstanceInfoParams
	InstanceInfoResponse
*/
package hcr

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

type Auth struct {
	UserSecret       string `protobuf:"bytes,1,opt,name=userSecret" json:"userSecret,omitempty"`
	DeploymentSecret string `protobuf:"bytes,2,opt,name=deploymentSecret" json:"deploymentSecret,omitempty"`
}

func (m *Auth) Reset()                    { *m = Auth{} }
func (m *Auth) String() string            { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()               {}
func (*Auth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Auth) GetUserSecret() string {
	if m != nil {
		return m.UserSecret
	}
	return ""
}

func (m *Auth) GetDeploymentSecret() string {
	if m != nil {
		return m.DeploymentSecret
	}
	return ""
}

type Status struct {
	Okay    bool   `protobuf:"varint,1,opt,name=okay" json:"okay,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Status) GetOkay() bool {
	if m != nil {
		return m.Okay
	}
	return false
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetMoteIDParams struct {
	Auth     *Auth  `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	DeviceId string `protobuf:"bytes,2,opt,name=deviceId" json:"deviceId,omitempty"`
}

func (m *GetMoteIDParams) Reset()                    { *m = GetMoteIDParams{} }
func (m *GetMoteIDParams) String() string            { return proto.CompactTextString(m) }
func (*GetMoteIDParams) ProtoMessage()               {}
func (*GetMoteIDParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetMoteIDParams) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *GetMoteIDParams) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

type GetMoteIDResponse struct {
	Status *Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Id     uint32  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Mac    string  `protobuf:"bytes,3,opt,name=mac" json:"mac,omitempty"`
}

func (m *GetMoteIDResponse) Reset()                    { *m = GetMoteIDResponse{} }
func (m *GetMoteIDResponse) String() string            { return proto.CompactTextString(m) }
func (*GetMoteIDResponse) ProtoMessage()               {}
func (*GetMoteIDResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetMoteIDResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetMoteIDResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetMoteIDResponse) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

type CreateInstanceParams struct {
	Auth       *Auth  `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	DeviceId   string `protobuf:"bytes,2,opt,name=deviceId" json:"deviceId,omitempty"`
	Repository string `protobuf:"bytes,3,opt,name=repository" json:"repository,omitempty"`
	Commit     string `protobuf:"bytes,4,opt,name=commit" json:"commit,omitempty"`
	Motetype   uint64 `protobuf:"varint,5,opt,name=motetype" json:"motetype,omitempty"`
	Extradata  string `protobuf:"bytes,6,opt,name=extradata" json:"extradata,omitempty"`
}

func (m *CreateInstanceParams) Reset()                    { *m = CreateInstanceParams{} }
func (m *CreateInstanceParams) String() string            { return proto.CompactTextString(m) }
func (*CreateInstanceParams) ProtoMessage()               {}
func (*CreateInstanceParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateInstanceParams) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *CreateInstanceParams) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *CreateInstanceParams) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *CreateInstanceParams) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func (m *CreateInstanceParams) GetMotetype() uint64 {
	if m != nil {
		return m.Motetype
	}
	return 0
}

func (m *CreateInstanceParams) GetExtradata() string {
	if m != nil {
		return m.Extradata
	}
	return ""
}

type CreateInstanceResponse struct {
	Status    *Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Id        uint32  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Mac       string  `protobuf:"bytes,3,opt,name=mac" json:"mac,omitempty"`
	Ed25519VK []byte  `protobuf:"bytes,4,opt,name=ed25519VK,proto3" json:"ed25519VK,omitempty"`
	Ed25519SK []byte  `protobuf:"bytes,5,opt,name=ed25519SK,proto3" json:"ed25519SK,omitempty"`
	AESK      []byte  `protobuf:"bytes,6,opt,name=AESK,json=aESK,proto3" json:"AESK,omitempty"`
	Instance  string  `protobuf:"bytes,7,opt,name=instance" json:"instance,omitempty"`
}

func (m *CreateInstanceResponse) Reset()                    { *m = CreateInstanceResponse{} }
func (m *CreateInstanceResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateInstanceResponse) ProtoMessage()               {}
func (*CreateInstanceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CreateInstanceResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CreateInstanceResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateInstanceResponse) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *CreateInstanceResponse) GetEd25519VK() []byte {
	if m != nil {
		return m.Ed25519VK
	}
	return nil
}

func (m *CreateInstanceResponse) GetEd25519SK() []byte {
	if m != nil {
		return m.Ed25519SK
	}
	return nil
}

func (m *CreateInstanceResponse) GetAESK() []byte {
	if m != nil {
		return m.AESK
	}
	return nil
}

func (m *CreateInstanceResponse) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

type CreateDeploymentParams struct {
	Auth *Auth  `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *CreateDeploymentParams) Reset()                    { *m = CreateDeploymentParams{} }
func (m *CreateDeploymentParams) String() string            { return proto.CompactTextString(m) }
func (*CreateDeploymentParams) ProtoMessage()               {}
func (*CreateDeploymentParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CreateDeploymentParams) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *CreateDeploymentParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateDeploymentResponse struct {
	Status   *Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	ReadKey  string  `protobuf:"bytes,2,opt,name=readKey" json:"readKey,omitempty"`
	WriteKey string  `protobuf:"bytes,3,opt,name=writeKey" json:"writeKey,omitempty"`
}

func (m *CreateDeploymentResponse) Reset()                    { *m = CreateDeploymentResponse{} }
func (m *CreateDeploymentResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateDeploymentResponse) ProtoMessage()               {}
func (*CreateDeploymentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CreateDeploymentResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CreateDeploymentResponse) GetReadKey() string {
	if m != nil {
		return m.ReadKey
	}
	return ""
}

func (m *CreateDeploymentResponse) GetWriteKey() string {
	if m != nil {
		return m.WriteKey
	}
	return ""
}

type BindMoteParams struct {
	Auth           *Auth  `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	Moteid         uint32 `protobuf:"varint,2,opt,name=moteid" json:"moteid,omitempty"`
	Deployment     string `protobuf:"bytes,3,opt,name=deployment" json:"deployment,omitempty"`
	RemoveExisting bool   `protobuf:"varint,4,opt,name=removeExisting" json:"removeExisting,omitempty"`
}

func (m *BindMoteParams) Reset()                    { *m = BindMoteParams{} }
func (m *BindMoteParams) String() string            { return proto.CompactTextString(m) }
func (*BindMoteParams) ProtoMessage()               {}
func (*BindMoteParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *BindMoteParams) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *BindMoteParams) GetMoteid() uint32 {
	if m != nil {
		return m.Moteid
	}
	return 0
}

func (m *BindMoteParams) GetDeployment() string {
	if m != nil {
		return m.Deployment
	}
	return ""
}

func (m *BindMoteParams) GetRemoveExisting() bool {
	if m != nil {
		return m.RemoveExisting
	}
	return false
}

type BindMoteResponse struct {
	Status *Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *BindMoteResponse) Reset()                    { *m = BindMoteResponse{} }
func (m *BindMoteResponse) String() string            { return proto.CompactTextString(m) }
func (*BindMoteResponse) ProtoMessage()               {}
func (*BindMoteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *BindMoteResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type MoteInfoParams struct {
	Auth   *Auth  `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	Moteid uint32 `protobuf:"varint,2,opt,name=moteid" json:"moteid,omitempty"`
}

func (m *MoteInfoParams) Reset()                    { *m = MoteInfoParams{} }
func (m *MoteInfoParams) String() string            { return proto.CompactTextString(m) }
func (*MoteInfoParams) ProtoMessage()               {}
func (*MoteInfoParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *MoteInfoParams) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *MoteInfoParams) GetMoteid() uint32 {
	if m != nil {
		return m.Moteid
	}
	return 0
}

type MoteInfoResponse struct {
	Status     *Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	AESK       []byte  `protobuf:"bytes,2,opt,name=AESK,json=aESK,proto3" json:"AESK,omitempty"`
	Ed25519VK  []byte  `protobuf:"bytes,3,opt,name=ed25519VK,proto3" json:"ed25519VK,omitempty"`
	Repository string  `protobuf:"bytes,4,opt,name=repository" json:"repository,omitempty"`
	Commit     string  `protobuf:"bytes,5,opt,name=commit" json:"commit,omitempty"`
	Motetype   uint64  `protobuf:"varint,6,opt,name=motetype" json:"motetype,omitempty"`
	Extradata  string  `protobuf:"bytes,7,opt,name=extradata" json:"extradata,omitempty"`
}

func (m *MoteInfoResponse) Reset()                    { *m = MoteInfoResponse{} }
func (m *MoteInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*MoteInfoResponse) ProtoMessage()               {}
func (*MoteInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *MoteInfoResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *MoteInfoResponse) GetAESK() []byte {
	if m != nil {
		return m.AESK
	}
	return nil
}

func (m *MoteInfoResponse) GetEd25519VK() []byte {
	if m != nil {
		return m.Ed25519VK
	}
	return nil
}

func (m *MoteInfoResponse) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *MoteInfoResponse) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func (m *MoteInfoResponse) GetMotetype() uint64 {
	if m != nil {
		return m.Motetype
	}
	return 0
}

func (m *MoteInfoResponse) GetExtradata() string {
	if m != nil {
		return m.Extradata
	}
	return ""
}

type InstanceInfoParams struct {
	Instance string `protobuf:"bytes,1,opt,name=instance" json:"instance,omitempty"`
}

func (m *InstanceInfoParams) Reset()                    { *m = InstanceInfoParams{} }
func (m *InstanceInfoParams) String() string            { return proto.CompactTextString(m) }
func (*InstanceInfoParams) ProtoMessage()               {}
func (*InstanceInfoParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *InstanceInfoParams) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

type InstanceInfoResponse struct {
	Status           *Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Moteid           uint32  `protobuf:"varint,2,opt,name=moteid" json:"moteid,omitempty"`
	RegistrationDate int64   `protobuf:"varint,3,opt,name=registrationDate" json:"registrationDate,omitempty"`
	Repository       string  `protobuf:"bytes,4,opt,name=repository" json:"repository,omitempty"`
	Commit           string  `protobuf:"bytes,5,opt,name=commit" json:"commit,omitempty"`
	Motetype         uint64  `protobuf:"varint,6,opt,name=motetype" json:"motetype,omitempty"`
	Extradata        string  `protobuf:"bytes,7,opt,name=extradata" json:"extradata,omitempty"`
	InstanceDate     int64   `protobuf:"varint,8,opt,name=instanceDate" json:"instanceDate,omitempty"`
	Ed25519VK        string  `protobuf:"bytes,9,opt,name=ed25519VK" json:"ed25519VK,omitempty"`
	LatestInstance   string  `protobuf:"bytes,10,opt,name=latestInstance" json:"latestInstance,omitempty"`
}

func (m *InstanceInfoResponse) Reset()                    { *m = InstanceInfoResponse{} }
func (m *InstanceInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*InstanceInfoResponse) ProtoMessage()               {}
func (*InstanceInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *InstanceInfoResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *InstanceInfoResponse) GetMoteid() uint32 {
	if m != nil {
		return m.Moteid
	}
	return 0
}

func (m *InstanceInfoResponse) GetRegistrationDate() int64 {
	if m != nil {
		return m.RegistrationDate
	}
	return 0
}

func (m *InstanceInfoResponse) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *InstanceInfoResponse) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func (m *InstanceInfoResponse) GetMotetype() uint64 {
	if m != nil {
		return m.Motetype
	}
	return 0
}

func (m *InstanceInfoResponse) GetExtradata() string {
	if m != nil {
		return m.Extradata
	}
	return ""
}

func (m *InstanceInfoResponse) GetInstanceDate() int64 {
	if m != nil {
		return m.InstanceDate
	}
	return 0
}

func (m *InstanceInfoResponse) GetEd25519VK() string {
	if m != nil {
		return m.Ed25519VK
	}
	return ""
}

func (m *InstanceInfoResponse) GetLatestInstance() string {
	if m != nil {
		return m.LatestInstance
	}
	return ""
}

func init() {
	proto.RegisterType((*Auth)(nil), "hcr.Auth")
	proto.RegisterType((*Status)(nil), "hcr.Status")
	proto.RegisterType((*GetMoteIDParams)(nil), "hcr.GetMoteIDParams")
	proto.RegisterType((*GetMoteIDResponse)(nil), "hcr.GetMoteIDResponse")
	proto.RegisterType((*CreateInstanceParams)(nil), "hcr.CreateInstanceParams")
	proto.RegisterType((*CreateInstanceResponse)(nil), "hcr.CreateInstanceResponse")
	proto.RegisterType((*CreateDeploymentParams)(nil), "hcr.CreateDeploymentParams")
	proto.RegisterType((*CreateDeploymentResponse)(nil), "hcr.CreateDeploymentResponse")
	proto.RegisterType((*BindMoteParams)(nil), "hcr.BindMoteParams")
	proto.RegisterType((*BindMoteResponse)(nil), "hcr.BindMoteResponse")
	proto.RegisterType((*MoteInfoParams)(nil), "hcr.MoteInfoParams")
	proto.RegisterType((*MoteInfoResponse)(nil), "hcr.MoteInfoResponse")
	proto.RegisterType((*InstanceInfoParams)(nil), "hcr.InstanceInfoParams")
	proto.RegisterType((*InstanceInfoResponse)(nil), "hcr.InstanceInfoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HamiltonRegistry service

type HamiltonRegistryClient interface {
	// Create or get a mote ID given its device ID
	// Create requires admin
	// Get is always permitted
	GetMoteID(ctx context.Context, in *GetMoteIDParams, opts ...grpc.CallOption) (*GetMoteIDResponse, error)
	// Create a new instance (program)
	// Requires deployment key
	CreateInstance(ctx context.Context, in *CreateInstanceParams, opts ...grpc.CallOption) (*CreateInstanceResponse, error)
	// Create a new deployment (keyring)
	// Requires admin
	CreateDeployment(ctx context.Context, in *CreateDeploymentParams, opts ...grpc.CallOption) (*CreateDeploymentResponse, error)
	// Bind a mote into a deployment
	// Requires admin
	BindMote(ctx context.Context, in *BindMoteParams, opts ...grpc.CallOption) (*BindMoteResponse, error)
	// Get the key for a mote id
	// Requires a deployment key
	MoteInfo(ctx context.Context, in *MoteInfoParams, opts ...grpc.CallOption) (*MoteInfoResponse, error)
	// Get the info for an instance code
	InstanceInfo(ctx context.Context, in *InstanceInfoParams, opts ...grpc.CallOption) (*InstanceInfoResponse, error)
}

type hamiltonRegistryClient struct {
	cc *grpc.ClientConn
}

func NewHamiltonRegistryClient(cc *grpc.ClientConn) HamiltonRegistryClient {
	return &hamiltonRegistryClient{cc}
}

func (c *hamiltonRegistryClient) GetMoteID(ctx context.Context, in *GetMoteIDParams, opts ...grpc.CallOption) (*GetMoteIDResponse, error) {
	out := new(GetMoteIDResponse)
	err := grpc.Invoke(ctx, "/hcr.HamiltonRegistry/GetMoteID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hamiltonRegistryClient) CreateInstance(ctx context.Context, in *CreateInstanceParams, opts ...grpc.CallOption) (*CreateInstanceResponse, error) {
	out := new(CreateInstanceResponse)
	err := grpc.Invoke(ctx, "/hcr.HamiltonRegistry/CreateInstance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hamiltonRegistryClient) CreateDeployment(ctx context.Context, in *CreateDeploymentParams, opts ...grpc.CallOption) (*CreateDeploymentResponse, error) {
	out := new(CreateDeploymentResponse)
	err := grpc.Invoke(ctx, "/hcr.HamiltonRegistry/CreateDeployment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hamiltonRegistryClient) BindMote(ctx context.Context, in *BindMoteParams, opts ...grpc.CallOption) (*BindMoteResponse, error) {
	out := new(BindMoteResponse)
	err := grpc.Invoke(ctx, "/hcr.HamiltonRegistry/BindMote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hamiltonRegistryClient) MoteInfo(ctx context.Context, in *MoteInfoParams, opts ...grpc.CallOption) (*MoteInfoResponse, error) {
	out := new(MoteInfoResponse)
	err := grpc.Invoke(ctx, "/hcr.HamiltonRegistry/MoteInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hamiltonRegistryClient) InstanceInfo(ctx context.Context, in *InstanceInfoParams, opts ...grpc.CallOption) (*InstanceInfoResponse, error) {
	out := new(InstanceInfoResponse)
	err := grpc.Invoke(ctx, "/hcr.HamiltonRegistry/InstanceInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HamiltonRegistry service

type HamiltonRegistryServer interface {
	// Create or get a mote ID given its device ID
	// Create requires admin
	// Get is always permitted
	GetMoteID(context.Context, *GetMoteIDParams) (*GetMoteIDResponse, error)
	// Create a new instance (program)
	// Requires deployment key
	CreateInstance(context.Context, *CreateInstanceParams) (*CreateInstanceResponse, error)
	// Create a new deployment (keyring)
	// Requires admin
	CreateDeployment(context.Context, *CreateDeploymentParams) (*CreateDeploymentResponse, error)
	// Bind a mote into a deployment
	// Requires admin
	BindMote(context.Context, *BindMoteParams) (*BindMoteResponse, error)
	// Get the key for a mote id
	// Requires a deployment key
	MoteInfo(context.Context, *MoteInfoParams) (*MoteInfoResponse, error)
	// Get the info for an instance code
	InstanceInfo(context.Context, *InstanceInfoParams) (*InstanceInfoResponse, error)
}

func RegisterHamiltonRegistryServer(s *grpc.Server, srv HamiltonRegistryServer) {
	s.RegisterService(&_HamiltonRegistry_serviceDesc, srv)
}

func _HamiltonRegistry_GetMoteID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMoteIDParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HamiltonRegistryServer).GetMoteID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hcr.HamiltonRegistry/GetMoteID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HamiltonRegistryServer).GetMoteID(ctx, req.(*GetMoteIDParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _HamiltonRegistry_CreateInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstanceParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HamiltonRegistryServer).CreateInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hcr.HamiltonRegistry/CreateInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HamiltonRegistryServer).CreateInstance(ctx, req.(*CreateInstanceParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _HamiltonRegistry_CreateDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeploymentParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HamiltonRegistryServer).CreateDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hcr.HamiltonRegistry/CreateDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HamiltonRegistryServer).CreateDeployment(ctx, req.(*CreateDeploymentParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _HamiltonRegistry_BindMote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindMoteParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HamiltonRegistryServer).BindMote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hcr.HamiltonRegistry/BindMote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HamiltonRegistryServer).BindMote(ctx, req.(*BindMoteParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _HamiltonRegistry_MoteInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoteInfoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HamiltonRegistryServer).MoteInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hcr.HamiltonRegistry/MoteInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HamiltonRegistryServer).MoteInfo(ctx, req.(*MoteInfoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _HamiltonRegistry_InstanceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstanceInfoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HamiltonRegistryServer).InstanceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hcr.HamiltonRegistry/InstanceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HamiltonRegistryServer).InstanceInfo(ctx, req.(*InstanceInfoParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _HamiltonRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hcr.HamiltonRegistry",
	HandlerType: (*HamiltonRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMoteID",
			Handler:    _HamiltonRegistry_GetMoteID_Handler,
		},
		{
			MethodName: "CreateInstance",
			Handler:    _HamiltonRegistry_CreateInstance_Handler,
		},
		{
			MethodName: "CreateDeployment",
			Handler:    _HamiltonRegistry_CreateDeployment_Handler,
		},
		{
			MethodName: "BindMote",
			Handler:    _HamiltonRegistry_BindMote_Handler,
		},
		{
			MethodName: "MoteInfo",
			Handler:    _HamiltonRegistry_MoteInfo_Handler,
		},
		{
			MethodName: "InstanceInfo",
			Handler:    _HamiltonRegistry_InstanceInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hreg.proto",
}

func init() { proto.RegisterFile("hreg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 727 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x13, 0xd7, 0x49, 0xa6, 0x21, 0x98, 0xa5, 0x04, 0x37, 0x50, 0x54, 0x19, 0xa9, 0xaa,
	0x7a, 0xa8, 0xa0, 0xa8, 0x45, 0x3d, 0xb6, 0xb4, 0x6a, 0x2b, 0x03, 0x42, 0x8e, 0xc4, 0x81, 0xdb,
	0x62, 0x0f, 0x89, 0x45, 0xed, 0x8d, 0xd6, 0x9b, 0xd2, 0x3c, 0x05, 0x0f, 0xc0, 0xe3, 0x70, 0xe4,
	0x11, 0xfa, 0x32, 0xc8, 0x9b, 0x5d, 0xc7, 0x76, 0xda, 0xaa, 0x51, 0x11, 0x37, 0xcf, 0x37, 0x3b,
	0x9f, 0xe7, 0xe7, 0x9b, 0xb5, 0x01, 0x86, 0x1c, 0x07, 0xdb, 0x23, 0xce, 0x04, 0x23, 0xf5, 0x61,
	0xc0, 0x5d, 0x1f, 0xcc, 0x83, 0xb1, 0x18, 0x92, 0x17, 0x00, 0xe3, 0x14, 0x79, 0x1f, 0x03, 0x8e,
	0xc2, 0x31, 0xd6, 0x8d, 0xcd, 0x96, 0x5f, 0x40, 0xc8, 0x16, 0xd8, 0x21, 0x8e, 0xce, 0xd9, 0x24,
	0xc6, 0x44, 0xa8, 0x53, 0x35, 0x79, 0x6a, 0x0e, 0x77, 0xf7, 0xc0, 0xea, 0x0b, 0x2a, 0xc6, 0x29,
	0x21, 0x60, 0xb2, 0xef, 0x74, 0x22, 0xf9, 0x9a, 0xbe, 0x7c, 0x26, 0x0e, 0x34, 0x62, 0x4c, 0x53,
	0x3a, 0x40, 0x45, 0xa0, 0x4d, 0xf7, 0x3d, 0x3c, 0x3c, 0x41, 0xf1, 0x81, 0x09, 0x3c, 0x3b, 0xfa,
	0x44, 0x39, 0x8d, 0x53, 0xb2, 0x06, 0x26, 0x1d, 0x8b, 0xa1, 0x24, 0x58, 0xde, 0x69, 0x6d, 0x0f,
	0x03, 0xbe, 0x9d, 0xe5, 0xeb, 0x4b, 0x98, 0xf4, 0xa0, 0x19, 0xe2, 0x45, 0x14, 0xe0, 0x59, 0xa8,
	0xc8, 0x72, 0xdb, 0xfd, 0x02, 0x8f, 0x72, 0x36, 0x1f, 0xd3, 0x11, 0x4b, 0x52, 0x24, 0x2f, 0xc1,
	0x4a, 0x65, 0x6a, 0x8a, 0x71, 0x59, 0x32, 0x4e, 0xb3, 0xf5, 0x95, 0x8b, 0x74, 0xa0, 0x16, 0x4d,
	0xf9, 0x1e, 0xf8, 0xb5, 0x28, 0x24, 0x36, 0xd4, 0x63, 0x1a, 0x38, 0x75, 0xf9, 0x82, 0xec, 0xd1,
	0xfd, 0x6d, 0xc0, 0xca, 0x3b, 0x8e, 0x54, 0xe0, 0x59, 0x92, 0x0a, 0x9a, 0x04, 0x78, 0xef, 0x7c,
	0xb3, 0x09, 0x70, 0x1c, 0xb1, 0x34, 0x12, 0x8c, 0x4f, 0xd4, 0xcb, 0x0a, 0x08, 0xe9, 0x82, 0x15,
	0xb0, 0x38, 0x8e, 0x84, 0x63, 0x4a, 0x9f, 0xb2, 0x32, 0xce, 0x98, 0x09, 0x14, 0x93, 0x11, 0x3a,
	0x4b, 0xeb, 0xc6, 0xa6, 0xe9, 0xe7, 0x36, 0x79, 0x0e, 0x2d, 0xbc, 0x14, 0x9c, 0x86, 0x54, 0x50,
	0xc7, 0x92, 0x61, 0x33, 0xc0, 0xfd, 0x63, 0x40, 0xb7, 0x5c, 0xc5, 0x3f, 0xee, 0x93, 0x7c, 0x7f,
	0xb8, 0xb3, 0xbb, 0xfb, 0x7a, 0xff, 0xb3, 0x27, 0xd3, 0x6e, 0xfb, 0x33, 0xa0, 0xe0, 0xed, 0x7b,
	0x32, 0xf5, 0x99, 0xb7, 0xef, 0x65, 0xda, 0x39, 0x38, 0xee, 0x7b, 0x32, 0xed, 0xb6, 0x6f, 0xd2,
	0xe3, 0xbe, 0x97, 0xd5, 0x1a, 0xa9, 0x54, 0x9d, 0xc6, 0xb4, 0x7f, 0xda, 0x76, 0x3d, 0x5d, 0xcc,
	0x51, 0xae, 0xc7, 0xbb, 0x0d, 0x85, 0x80, 0x99, 0xd0, 0x58, 0xab, 0x51, 0x3e, 0xbb, 0x63, 0x70,
	0xaa, 0x64, 0x8b, 0xf5, 0xc6, 0x81, 0x06, 0x47, 0x1a, 0x7a, 0x38, 0xd1, 0x2a, 0x57, 0x66, 0x56,
	0xc3, 0x0f, 0x1e, 0x09, 0xcc, 0x5c, 0xd3, 0x56, 0xe5, 0xb6, 0xfb, 0xd3, 0x80, 0xce, 0x61, 0x94,
	0x84, 0x99, 0x6a, 0xef, 0x96, 0x7c, 0x17, 0xac, 0x6c, 0xda, 0xf9, 0x1c, 0x94, 0x95, 0xa9, 0x69,
	0xb6, 0x97, 0x5a, 0x4d, 0x33, 0x84, 0x6c, 0x40, 0x87, 0x63, 0xcc, 0x2e, 0xf0, 0xf8, 0x32, 0x4a,
	0x45, 0x94, 0x0c, 0xe4, 0x78, 0x9a, 0x7e, 0x05, 0x75, 0xdf, 0x82, 0xad, 0x13, 0x5a, 0xa8, 0x01,
	0xee, 0x09, 0x74, 0xe4, 0xee, 0x25, 0xdf, 0xd8, 0xbd, 0x2a, 0x71, 0xaf, 0x0c, 0xb0, 0x35, 0xd3,
	0x62, 0x33, 0xd0, 0x0a, 0xaa, 0x15, 0x14, 0x54, 0x52, 0x64, 0xbd, 0xaa, 0xc8, 0xf2, 0x0e, 0x9a,
	0xb7, 0xec, 0xe0, 0xd2, 0x8d, 0x3b, 0x68, 0xdd, 0xb6, 0x83, 0x8d, 0xea, 0x0e, 0xbe, 0x02, 0xa2,
	0x97, 0xaf, 0xd0, 0xaa, 0xa2, 0xce, 0x8d, 0x8a, 0xce, 0xaf, 0x6a, 0xb0, 0x52, 0x0c, 0x59, 0xac,
	0x27, 0x37, 0xe9, 0x65, 0x0b, 0x6c, 0x8e, 0x83, 0x28, 0x15, 0x9c, 0x8a, 0x88, 0x25, 0x47, 0x54,
	0xa0, 0x6c, 0x4f, 0xdd, 0x9f, 0xc3, 0xff, 0x7f, 0x97, 0x88, 0x0b, 0x6d, 0x5d, 0xbf, 0xcc, 0xac,
	0x29, 0x33, 0x2b, 0x61, 0xe5, 0xc9, 0xb6, 0x14, 0x43, 0x3e, 0xd9, 0x0d, 0xe8, 0x9c, 0x53, 0x81,
	0xa9, 0xd0, 0xad, 0x73, 0x40, 0x1e, 0xa9, 0xa0, 0x3b, 0xbf, 0xea, 0x60, 0x9f, 0xd2, 0x38, 0x3a,
	0x17, 0x2c, 0xf1, 0xa7, 0x85, 0x4f, 0xc8, 0x3e, 0xb4, 0xf2, 0x4f, 0x09, 0x59, 0x91, 0x6d, 0xad,
	0x7c, 0xa8, 0x7a, 0xdd, 0x32, 0x9a, 0x0f, 0xe5, 0x14, 0x3a, 0xe5, 0x2b, 0x96, 0xac, 0xca, 0x93,
	0xd7, 0x7d, 0x3d, 0x7a, 0xcf, 0xae, 0x71, 0xe5, 0x4c, 0x1f, 0xc1, 0xae, 0x5e, 0x49, 0xa4, 0x18,
	0x50, 0xbd, 0xf6, 0x7a, 0x6b, 0xd7, 0x3a, 0x73, 0xbe, 0x3d, 0x68, 0xea, 0xcd, 0x26, 0x8f, 0xe5,
	0xd1, 0xf2, 0xcd, 0xd3, 0x7b, 0x52, 0x02, 0x8b, 0x71, 0x7a, 0x1d, 0x55, 0x5c, 0x79, 0xcf, 0x55,
	0xdc, 0xdc, 0xca, 0x1e, 0x42, 0xbb, 0x28, 0x5b, 0xf2, 0x54, 0x1e, 0x9b, 0x17, 0x7f, 0x6f, 0x75,
	0xce, 0xa1, 0x39, 0xbe, 0x5a, 0xf2, 0xcf, 0xe5, 0xcd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3,
	0x1b, 0x40, 0x46, 0xc7, 0x08, 0x00, 0x00,
}
