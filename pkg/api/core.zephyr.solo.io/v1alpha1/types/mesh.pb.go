// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/core/v1alpha1/mesh.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MeshGroupStatus_ConfigStatus int32

const (
	// config was processed properly, and no error(s) were found
	MeshGroupStatus_VALID MeshGroupStatus_ConfigStatus = 0
	// config was processed properly, and error(s) were found
	MeshGroupStatus_INVALID MeshGroupStatus_ConfigStatus = 1
	// config could not be processed properly
	MeshGroupStatus_PROCESSING_ERROR MeshGroupStatus_ConfigStatus = 2
)

var MeshGroupStatus_ConfigStatus_name = map[int32]string{
	0: "VALID",
	1: "INVALID",
	2: "PROCESSING_ERROR",
}

var MeshGroupStatus_ConfigStatus_value = map[string]int32{
	"VALID":            0,
	"INVALID":          1,
	"PROCESSING_ERROR": 2,
}

func (x MeshGroupStatus_ConfigStatus) String() string {
	return proto.EnumName(MeshGroupStatus_ConfigStatus_name, int32(x))
}

func (MeshGroupStatus_ConfigStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d4089e58fa1f863, []int{8, 0}
}

//
//Meshes represent a currently registered service mesh.
type MeshSpec struct {
	// Types that are valid to be assigned to MeshType:
	//	*MeshSpec_Istio
	//	*MeshSpec_AwsAppMesh
	//	*MeshSpec_Linkerd
	//	*MeshSpec_ConsulConnect
	MeshType             isMeshSpec_MeshType `protobuf_oneof:"mesh_type"`
	Cluster              *ResourceRef        `protobuf:"bytes,5,opt,name=cluster,proto3" json:"cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MeshSpec) Reset()         { *m = MeshSpec{} }
func (m *MeshSpec) String() string { return proto.CompactTextString(m) }
func (*MeshSpec) ProtoMessage()    {}
func (*MeshSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d4089e58fa1f863, []int{0}
}
func (m *MeshSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshSpec.Unmarshal(m, b)
}
func (m *MeshSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshSpec.Marshal(b, m, deterministic)
}
func (m *MeshSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshSpec.Merge(m, src)
}
func (m *MeshSpec) XXX_Size() int {
	return xxx_messageInfo_MeshSpec.Size(m)
}
func (m *MeshSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshSpec.DiscardUnknown(m)
}

var xxx_messageInfo_MeshSpec proto.InternalMessageInfo

type isMeshSpec_MeshType interface {
	isMeshSpec_MeshType()
	Equal(interface{}) bool
}

type MeshSpec_Istio struct {
	Istio *IstioMesh `protobuf:"bytes,1,opt,name=istio,proto3,oneof" json:"istio,omitempty"`
}
type MeshSpec_AwsAppMesh struct {
	AwsAppMesh *AwsAppMesh `protobuf:"bytes,2,opt,name=aws_app_mesh,json=awsAppMesh,proto3,oneof" json:"aws_app_mesh,omitempty"`
}
type MeshSpec_Linkerd struct {
	Linkerd *LinkerdMesh `protobuf:"bytes,3,opt,name=linkerd,proto3,oneof" json:"linkerd,omitempty"`
}
type MeshSpec_ConsulConnect struct {
	ConsulConnect *ConsulConnectMesh `protobuf:"bytes,4,opt,name=consul_connect,json=consulConnect,proto3,oneof" json:"consul_connect,omitempty"`
}

func (*MeshSpec_Istio) isMeshSpec_MeshType()         {}
func (*MeshSpec_AwsAppMesh) isMeshSpec_MeshType()    {}
func (*MeshSpec_Linkerd) isMeshSpec_MeshType()       {}
func (*MeshSpec_ConsulConnect) isMeshSpec_MeshType() {}

func (m *MeshSpec) GetMeshType() isMeshSpec_MeshType {
	if m != nil {
		return m.MeshType
	}
	return nil
}

func (m *MeshSpec) GetIstio() *IstioMesh {
	if x, ok := m.GetMeshType().(*MeshSpec_Istio); ok {
		return x.Istio
	}
	return nil
}

func (m *MeshSpec) GetAwsAppMesh() *AwsAppMesh {
	if x, ok := m.GetMeshType().(*MeshSpec_AwsAppMesh); ok {
		return x.AwsAppMesh
	}
	return nil
}

func (m *MeshSpec) GetLinkerd() *LinkerdMesh {
	if x, ok := m.GetMeshType().(*MeshSpec_Linkerd); ok {
		return x.Linkerd
	}
	return nil
}

func (m *MeshSpec) GetConsulConnect() *ConsulConnectMesh {
	if x, ok := m.GetMeshType().(*MeshSpec_ConsulConnect); ok {
		return x.ConsulConnect
	}
	return nil
}

func (m *MeshSpec) GetCluster() *ResourceRef {
	if m != nil {
		return m.Cluster
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MeshSpec) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MeshSpec_Istio)(nil),
		(*MeshSpec_AwsAppMesh)(nil),
		(*MeshSpec_Linkerd)(nil),
		(*MeshSpec_ConsulConnect)(nil),
	}
}

type MeshStatus struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshStatus) Reset()         { *m = MeshStatus{} }
func (m *MeshStatus) String() string { return proto.CompactTextString(m) }
func (*MeshStatus) ProtoMessage()    {}
func (*MeshStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d4089e58fa1f863, []int{1}
}
func (m *MeshStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshStatus.Unmarshal(m, b)
}
func (m *MeshStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshStatus.Marshal(b, m, deterministic)
}
func (m *MeshStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshStatus.Merge(m, src)
}
func (m *MeshStatus) XXX_Size() int {
	return xxx_messageInfo_MeshStatus.Size(m)
}
func (m *MeshStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshStatus.DiscardUnknown(m)
}

var xxx_messageInfo_MeshStatus proto.InternalMessageInfo

type MeshInstallation struct {
	// where the control plane has been installed
	InstallationNamespace string `protobuf:"bytes,1,opt,name=installation_namespace,json=installationNamespace,proto3" json:"installation_namespace,omitempty"`
	// version of the mesh which has been installed
	// Note that the version may be "latest"
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshInstallation) Reset()         { *m = MeshInstallation{} }
func (m *MeshInstallation) String() string { return proto.CompactTextString(m) }
func (*MeshInstallation) ProtoMessage()    {}
func (*MeshInstallation) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d4089e58fa1f863, []int{2}
}
func (m *MeshInstallation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshInstallation.Unmarshal(m, b)
}
func (m *MeshInstallation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshInstallation.Marshal(b, m, deterministic)
}
func (m *MeshInstallation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshInstallation.Merge(m, src)
}
func (m *MeshInstallation) XXX_Size() int {
	return xxx_messageInfo_MeshInstallation.Size(m)
}
func (m *MeshInstallation) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshInstallation.DiscardUnknown(m)
}

var xxx_messageInfo_MeshInstallation proto.InternalMessageInfo

func (m *MeshInstallation) GetInstallationNamespace() string {
	if m != nil {
		return m.InstallationNamespace
	}
	return ""
}

func (m *MeshInstallation) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Mesh object representing an installed Istio control plane
type IstioMesh struct {
	Installation         *MeshInstallation `protobuf:"bytes,1,opt,name=installation,proto3" json:"installation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IstioMesh) Reset()         { *m = IstioMesh{} }
func (m *IstioMesh) String() string { return proto.CompactTextString(m) }
func (*IstioMesh) ProtoMessage()    {}
func (*IstioMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d4089e58fa1f863, []int{3}
}
func (m *IstioMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IstioMesh.Unmarshal(m, b)
}
func (m *IstioMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IstioMesh.Marshal(b, m, deterministic)
}
func (m *IstioMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioMesh.Merge(m, src)
}
func (m *IstioMesh) XXX_Size() int {
	return xxx_messageInfo_IstioMesh.Size(m)
}
func (m *IstioMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioMesh.DiscardUnknown(m)
}

var xxx_messageInfo_IstioMesh proto.InternalMessageInfo

func (m *IstioMesh) GetInstallation() *MeshInstallation {
	if m != nil {
		return m.Installation
	}
	return nil
}

// Mesh object representing AWS App Mesh
type AwsAppMesh struct {
	Installation *MeshInstallation `protobuf:"bytes,1,opt,name=installation,proto3" json:"installation,omitempty"`
	// The AWS region the AWS App Mesh control plane resources exist in.
	Region               string   `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsAppMesh) Reset()         { *m = AwsAppMesh{} }
func (m *AwsAppMesh) String() string { return proto.CompactTextString(m) }
func (*AwsAppMesh) ProtoMessage()    {}
func (*AwsAppMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d4089e58fa1f863, []int{4}
}
func (m *AwsAppMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsAppMesh.Unmarshal(m, b)
}
func (m *AwsAppMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsAppMesh.Marshal(b, m, deterministic)
}
func (m *AwsAppMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsAppMesh.Merge(m, src)
}
func (m *AwsAppMesh) XXX_Size() int {
	return xxx_messageInfo_AwsAppMesh.Size(m)
}
func (m *AwsAppMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsAppMesh.DiscardUnknown(m)
}

var xxx_messageInfo_AwsAppMesh proto.InternalMessageInfo

func (m *AwsAppMesh) GetInstallation() *MeshInstallation {
	if m != nil {
		return m.Installation
	}
	return nil
}

func (m *AwsAppMesh) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

// Mesh object representing an installed Linkerd control plane
type LinkerdMesh struct {
	Installation         *MeshInstallation `protobuf:"bytes,1,opt,name=installation,proto3" json:"installation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LinkerdMesh) Reset()         { *m = LinkerdMesh{} }
func (m *LinkerdMesh) String() string { return proto.CompactTextString(m) }
func (*LinkerdMesh) ProtoMessage()    {}
func (*LinkerdMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d4089e58fa1f863, []int{5}
}
func (m *LinkerdMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkerdMesh.Unmarshal(m, b)
}
func (m *LinkerdMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkerdMesh.Marshal(b, m, deterministic)
}
func (m *LinkerdMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkerdMesh.Merge(m, src)
}
func (m *LinkerdMesh) XXX_Size() int {
	return xxx_messageInfo_LinkerdMesh.Size(m)
}
func (m *LinkerdMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkerdMesh.DiscardUnknown(m)
}

var xxx_messageInfo_LinkerdMesh proto.InternalMessageInfo

func (m *LinkerdMesh) GetInstallation() *MeshInstallation {
	if m != nil {
		return m.Installation
	}
	return nil
}

type ConsulConnectMesh struct {
	Installation         *MeshInstallation `protobuf:"bytes,1,opt,name=installation,proto3" json:"installation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ConsulConnectMesh) Reset()         { *m = ConsulConnectMesh{} }
func (m *ConsulConnectMesh) String() string { return proto.CompactTextString(m) }
func (*ConsulConnectMesh) ProtoMessage()    {}
func (*ConsulConnectMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d4089e58fa1f863, []int{6}
}
func (m *ConsulConnectMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsulConnectMesh.Unmarshal(m, b)
}
func (m *ConsulConnectMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsulConnectMesh.Marshal(b, m, deterministic)
}
func (m *ConsulConnectMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsulConnectMesh.Merge(m, src)
}
func (m *ConsulConnectMesh) XXX_Size() int {
	return xxx_messageInfo_ConsulConnectMesh.Size(m)
}
func (m *ConsulConnectMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsulConnectMesh.DiscardUnknown(m)
}

var xxx_messageInfo_ConsulConnectMesh proto.InternalMessageInfo

func (m *ConsulConnectMesh) GetInstallation() *MeshInstallation {
	if m != nil {
		return m.Installation
	}
	return nil
}

type MeshGroupSpec struct {
	// User-provided display name for the mesh group.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The meshes contained in this group.
	Meshes []*ResourceRef `protobuf:"bytes,2,rep,name=meshes,proto3" json:"meshes,omitempty"`
	//
	//optional: if set will use this resource ref to determine the secret from which the
	//root trust bundle should be grabbed. Otherwise it will default to
	//
	//trust_bundle:
	//name: cacerts
	//namespace: service-mesh-hub (default write namespace)
	TrustBundleRef *ResourceRef `protobuf:"bytes,3,opt,name=trust_bundle_ref,json=trustBundleRef,proto3" json:"trust_bundle_ref,omitempty"`
	// Types that are valid to be assigned to TrustModel:
	//	*MeshGroupSpec_Shared
	//	*MeshGroupSpec_Limited
	TrustModel           isMeshGroupSpec_TrustModel `protobuf_oneof:"trust_model"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MeshGroupSpec) Reset()         { *m = MeshGroupSpec{} }
func (m *MeshGroupSpec) String() string { return proto.CompactTextString(m) }
func (*MeshGroupSpec) ProtoMessage()    {}
func (*MeshGroupSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d4089e58fa1f863, []int{7}
}
func (m *MeshGroupSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshGroupSpec.Unmarshal(m, b)
}
func (m *MeshGroupSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshGroupSpec.Marshal(b, m, deterministic)
}
func (m *MeshGroupSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshGroupSpec.Merge(m, src)
}
func (m *MeshGroupSpec) XXX_Size() int {
	return xxx_messageInfo_MeshGroupSpec.Size(m)
}
func (m *MeshGroupSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshGroupSpec.DiscardUnknown(m)
}

var xxx_messageInfo_MeshGroupSpec proto.InternalMessageInfo

type isMeshGroupSpec_TrustModel interface {
	isMeshGroupSpec_TrustModel()
	Equal(interface{}) bool
}

type MeshGroupSpec_Shared struct {
	Shared *MeshGroupSpec_SharedTrust `protobuf:"bytes,4,opt,name=shared,proto3,oneof" json:"shared,omitempty"`
}
type MeshGroupSpec_Limited struct {
	Limited *MeshGroupSpec_LimitedTrust `protobuf:"bytes,5,opt,name=limited,proto3,oneof" json:"limited,omitempty"`
}

func (*MeshGroupSpec_Shared) isMeshGroupSpec_TrustModel()  {}
func (*MeshGroupSpec_Limited) isMeshGroupSpec_TrustModel() {}

func (m *MeshGroupSpec) GetTrustModel() isMeshGroupSpec_TrustModel {
	if m != nil {
		return m.TrustModel
	}
	return nil
}

func (m *MeshGroupSpec) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *MeshGroupSpec) GetMeshes() []*ResourceRef {
	if m != nil {
		return m.Meshes
	}
	return nil
}

func (m *MeshGroupSpec) GetTrustBundleRef() *ResourceRef {
	if m != nil {
		return m.TrustBundleRef
	}
	return nil
}

func (m *MeshGroupSpec) GetShared() *MeshGroupSpec_SharedTrust {
	if x, ok := m.GetTrustModel().(*MeshGroupSpec_Shared); ok {
		return x.Shared
	}
	return nil
}

func (m *MeshGroupSpec) GetLimited() *MeshGroupSpec_LimitedTrust {
	if x, ok := m.GetTrustModel().(*MeshGroupSpec_Limited); ok {
		return x.Limited
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MeshGroupSpec) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MeshGroupSpec_Shared)(nil),
		(*MeshGroupSpec_Limited)(nil),
	}
}

//
//Shared trust is a mesh group trust model requiring a shared root certificate, as well as shared identity
//between all entities which wish to communicate within the mesh group.
//
//The best current example of this would be the replicated control planes example from istio:
//https://preliminary.istio.io/docs/setup/install/multicluster/gateways/
type MeshGroupSpec_SharedTrust struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshGroupSpec_SharedTrust) Reset()         { *m = MeshGroupSpec_SharedTrust{} }
func (m *MeshGroupSpec_SharedTrust) String() string { return proto.CompactTextString(m) }
func (*MeshGroupSpec_SharedTrust) ProtoMessage()    {}
func (*MeshGroupSpec_SharedTrust) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d4089e58fa1f863, []int{7, 0}
}
func (m *MeshGroupSpec_SharedTrust) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshGroupSpec_SharedTrust.Unmarshal(m, b)
}
func (m *MeshGroupSpec_SharedTrust) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshGroupSpec_SharedTrust.Marshal(b, m, deterministic)
}
func (m *MeshGroupSpec_SharedTrust) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshGroupSpec_SharedTrust.Merge(m, src)
}
func (m *MeshGroupSpec_SharedTrust) XXX_Size() int {
	return xxx_messageInfo_MeshGroupSpec_SharedTrust.Size(m)
}
func (m *MeshGroupSpec_SharedTrust) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshGroupSpec_SharedTrust.DiscardUnknown(m)
}

var xxx_messageInfo_MeshGroupSpec_SharedTrust proto.InternalMessageInfo

//
//Limited trust is a mesh group trust model which does not require all meshes sharing the same root certificate
//or identity model. But rather, the limited trust creates trust between meshes running on different clusters
//by connecting their ingress/egress gateways with a common cert/identity. In this model all requests
//between different have the following request path when communicating between clusters
//
//cluster 1 MTLS               shared MTLS                  cluster 2 MTLS
//client/workload <-----------> egress gateway <----------> ingress gateway <--------------> server
//
//This approach has the downside of not maintaining identity from client to server, but allows for ad-hoc
//addition of additional clusters into a mesh group.
type MeshGroupSpec_LimitedTrust struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshGroupSpec_LimitedTrust) Reset()         { *m = MeshGroupSpec_LimitedTrust{} }
func (m *MeshGroupSpec_LimitedTrust) String() string { return proto.CompactTextString(m) }
func (*MeshGroupSpec_LimitedTrust) ProtoMessage()    {}
func (*MeshGroupSpec_LimitedTrust) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d4089e58fa1f863, []int{7, 1}
}
func (m *MeshGroupSpec_LimitedTrust) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshGroupSpec_LimitedTrust.Unmarshal(m, b)
}
func (m *MeshGroupSpec_LimitedTrust) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshGroupSpec_LimitedTrust.Marshal(b, m, deterministic)
}
func (m *MeshGroupSpec_LimitedTrust) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshGroupSpec_LimitedTrust.Merge(m, src)
}
func (m *MeshGroupSpec_LimitedTrust) XXX_Size() int {
	return xxx_messageInfo_MeshGroupSpec_LimitedTrust.Size(m)
}
func (m *MeshGroupSpec_LimitedTrust) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshGroupSpec_LimitedTrust.DiscardUnknown(m)
}

var xxx_messageInfo_MeshGroupSpec_LimitedTrust proto.InternalMessageInfo

type MeshGroupStatus struct {
	//
	//0:  The state of the mesh group is unknown, this most likely means that the state is being calculated.
	//1:  The state of the mesh group is Good. All operations have been completed successfully, and all routing
	//is enabled.
	//2:  The state of the mesh group is bad. Either a mesh group operation has failed, or some portion of
	//inter-mesh routing is down.
	//
	//TODO: Enumerate further information below on the following states in a description message
	Health               HealthStatus                 `protobuf:"varint,1,opt,name=health,proto3,enum=core.zephyr.solo.io.HealthStatus" json:"health,omitempty"`
	Config               MeshGroupStatus_ConfigStatus `protobuf:"varint,2,opt,name=config,proto3,enum=core.zephyr.solo.io.MeshGroupStatus_ConfigStatus" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *MeshGroupStatus) Reset()         { *m = MeshGroupStatus{} }
func (m *MeshGroupStatus) String() string { return proto.CompactTextString(m) }
func (*MeshGroupStatus) ProtoMessage()    {}
func (*MeshGroupStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d4089e58fa1f863, []int{8}
}
func (m *MeshGroupStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshGroupStatus.Unmarshal(m, b)
}
func (m *MeshGroupStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshGroupStatus.Marshal(b, m, deterministic)
}
func (m *MeshGroupStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshGroupStatus.Merge(m, src)
}
func (m *MeshGroupStatus) XXX_Size() int {
	return xxx_messageInfo_MeshGroupStatus.Size(m)
}
func (m *MeshGroupStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshGroupStatus.DiscardUnknown(m)
}

var xxx_messageInfo_MeshGroupStatus proto.InternalMessageInfo

func (m *MeshGroupStatus) GetHealth() HealthStatus {
	if m != nil {
		return m.Health
	}
	return HealthStatus_UNKNOWN
}

func (m *MeshGroupStatus) GetConfig() MeshGroupStatus_ConfigStatus {
	if m != nil {
		return m.Config
	}
	return MeshGroupStatus_VALID
}

func init() {
	proto.RegisterEnum("core.zephyr.solo.io.MeshGroupStatus_ConfigStatus", MeshGroupStatus_ConfigStatus_name, MeshGroupStatus_ConfigStatus_value)
	proto.RegisterType((*MeshSpec)(nil), "core.zephyr.solo.io.MeshSpec")
	proto.RegisterType((*MeshStatus)(nil), "core.zephyr.solo.io.MeshStatus")
	proto.RegisterType((*MeshInstallation)(nil), "core.zephyr.solo.io.MeshInstallation")
	proto.RegisterType((*IstioMesh)(nil), "core.zephyr.solo.io.IstioMesh")
	proto.RegisterType((*AwsAppMesh)(nil), "core.zephyr.solo.io.AwsAppMesh")
	proto.RegisterType((*LinkerdMesh)(nil), "core.zephyr.solo.io.LinkerdMesh")
	proto.RegisterType((*ConsulConnectMesh)(nil), "core.zephyr.solo.io.ConsulConnectMesh")
	proto.RegisterType((*MeshGroupSpec)(nil), "core.zephyr.solo.io.MeshGroupSpec")
	proto.RegisterType((*MeshGroupSpec_SharedTrust)(nil), "core.zephyr.solo.io.MeshGroupSpec.SharedTrust")
	proto.RegisterType((*MeshGroupSpec_LimitedTrust)(nil), "core.zephyr.solo.io.MeshGroupSpec.LimitedTrust")
	proto.RegisterType((*MeshGroupStatus)(nil), "core.zephyr.solo.io.MeshGroupStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/mesh-projects/api/core/v1alpha1/mesh.proto", fileDescriptor_4d4089e58fa1f863)
}

var fileDescriptor_4d4089e58fa1f863 = []byte{
	// 691 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x5d, 0x4f, 0x13, 0x4d,
	0x14, 0xc7, 0x69, 0x79, 0x68, 0x9f, 0x9e, 0xbe, 0x58, 0x47, 0x24, 0x4d, 0x2f, 0x10, 0x36, 0xd1,
	0x68, 0x22, 0xdb, 0x80, 0xd1, 0xa8, 0xc1, 0x0b, 0x5a, 0x09, 0xad, 0x20, 0xe8, 0xd4, 0x10, 0xe3,
	0x85, 0x9b, 0x61, 0x3b, 0x6d, 0x47, 0xb6, 0x3b, 0x93, 0x99, 0x59, 0x08, 0x7e, 0x22, 0xbf, 0x90,
	0xd7, 0x26, 0x7e, 0x0f, 0x13, 0x33, 0xb3, 0xdb, 0x76, 0xc5, 0x4a, 0xb9, 0xe0, 0x6e, 0xcf, 0xcc,
	0xf9, 0xfd, 0xf7, 0xbc, 0xcc, 0x99, 0x81, 0x57, 0x03, 0xa6, 0x87, 0xd1, 0x89, 0xeb, 0xf3, 0x51,
	0x43, 0xf1, 0x80, 0x6f, 0x30, 0xde, 0x18, 0x51, 0x35, 0xdc, 0x10, 0x92, 0x7f, 0xa1, 0xbe, 0x56,
	0x0d, 0x22, 0x58, 0xc3, 0xe7, 0x92, 0x36, 0xce, 0x36, 0x49, 0x20, 0x86, 0x64, 0xd3, 0xee, 0xbb,
	0x42, 0x72, 0xcd, 0xd1, 0x1d, 0xb3, 0xe3, 0x7e, 0xa5, 0x62, 0x78, 0x21, 0x5d, 0xc3, 0xbb, 0x8c,
	0xd7, 0x1f, 0xcd, 0x13, 0x90, 0xb4, 0x1f, 0xf3, 0xf5, 0xc7, 0xf3, 0x5c, 0x87, 0x94, 0x04, 0x3a,
	0xf9, 0x5b, 0x7d, 0x79, 0xc0, 0x07, 0xdc, 0x7e, 0x36, 0xcc, 0x57, 0xbc, 0xea, 0xfc, 0xc8, 0xc2,
	0xff, 0x6f, 0xa9, 0x1a, 0x76, 0x05, 0xf5, 0xd1, 0x33, 0x58, 0x62, 0x4a, 0x33, 0x5e, 0xcb, 0xac,
	0x65, 0x1e, 0x16, 0xb7, 0x56, 0xdd, 0x19, 0x01, 0xba, 0x1d, 0xe3, 0x61, 0x90, 0xf6, 0x02, 0x8e,
	0xdd, 0x51, 0x0b, 0x4a, 0xe4, 0x5c, 0x79, 0x44, 0x08, 0xcf, 0x84, 0x54, 0xcb, 0x5a, 0xfc, 0xde,
	0x4c, 0x7c, 0xe7, 0x5c, 0xed, 0x08, 0x91, 0xf0, 0x40, 0x26, 0x16, 0xda, 0x86, 0x7c, 0xc0, 0xc2,
	0x53, 0x2a, 0x7b, 0xb5, 0x45, 0xcb, 0xaf, 0xcd, 0xe4, 0x0f, 0x62, 0x9f, 0x44, 0x60, 0x8c, 0xa0,
	0x23, 0xa8, 0xf8, 0x3c, 0x54, 0x51, 0xe0, 0xf9, 0x3c, 0x0c, 0xa9, 0xaf, 0x6b, 0xff, 0x59, 0x91,
	0x07, 0x33, 0x45, 0x5a, 0xd6, 0xb5, 0x15, 0x7b, 0x26, 0x52, 0x65, 0x3f, 0xbd, 0x88, 0x5e, 0x42,
	0xde, 0x0f, 0x22, 0xa5, 0xa9, 0xac, 0x2d, 0x5d, 0x11, 0x0e, 0xa6, 0x8a, 0x47, 0xd2, 0xa7, 0x98,
	0xf6, 0xf1, 0x18, 0x68, 0x16, 0xa1, 0x60, 0xea, 0xe0, 0xe9, 0x0b, 0x41, 0x9d, 0x12, 0x80, 0x2d,
	0xb0, 0x26, 0x3a, 0x52, 0x8e, 0x0f, 0x55, 0x63, 0x75, 0x42, 0xa5, 0x49, 0x10, 0x10, 0xcd, 0x78,
	0x88, 0x9e, 0xc2, 0x0a, 0x4b, 0xd9, 0x5e, 0x48, 0x46, 0x54, 0x09, 0xe2, 0x53, 0xdb, 0x87, 0x02,
	0xbe, 0x9b, 0xde, 0x3d, 0x1c, 0x6f, 0xa2, 0x1a, 0xe4, 0xcf, 0xa8, 0x54, 0x8c, 0x87, 0xb6, 0xe0,
	0x05, 0x3c, 0x36, 0x9d, 0x63, 0x28, 0x4c, 0xba, 0x84, 0x3a, 0x50, 0x4a, 0xf3, 0x49, 0x6f, 0xef,
	0xcf, 0xcc, 0xe6, 0x72, 0x68, 0xf8, 0x0f, 0xd4, 0xe1, 0x00, 0xd3, 0xf6, 0xdd, 0xa0, 0x30, 0x5a,
	0x81, 0x9c, 0xa4, 0x83, 0x69, 0x26, 0x89, 0xe5, 0x7c, 0x84, 0x62, 0xaa, 0xdf, 0x37, 0x99, 0xca,
	0x67, 0xb8, 0xfd, 0xd7, 0x21, 0xb8, 0x49, 0xfd, 0x5f, 0x59, 0x28, 0x1b, 0x97, 0x3d, 0xc9, 0x23,
	0x61, 0x87, 0x6b, 0x1d, 0x4a, 0x3d, 0xa6, 0x44, 0x40, 0x2e, 0x6c, 0x83, 0x93, 0xde, 0x16, 0x93,
	0x35, 0xd3, 0x56, 0xf4, 0x1c, 0x72, 0xe6, 0xdc, 0x50, 0x55, 0xcb, 0xae, 0x2d, 0x5e, 0xeb, 0xc8,
	0x25, 0xfe, 0xe8, 0x0d, 0x54, 0xb5, 0x8c, 0x94, 0xf6, 0x4e, 0xa2, 0xb0, 0x17, 0x50, 0x4f, 0xd2,
	0xfe, 0x95, 0x53, 0x94, 0xd6, 0xa8, 0x58, 0xb2, 0x69, 0x41, 0x4c, 0xfb, 0xa8, 0x0d, 0x39, 0x35,
	0x24, 0x92, 0xf6, 0x92, 0x11, 0x72, 0xff, 0x99, 0xff, 0x24, 0x39, 0xb7, 0x6b, 0x81, 0x0f, 0x46,
	0xa8, 0xbd, 0x80, 0x13, 0x1e, 0xed, 0x9b, 0x91, 0x1e, 0x31, 0x4d, 0x7b, 0xc9, 0x0c, 0x35, 0xae,
	0x21, 0x75, 0x10, 0x13, 0x63, 0xad, 0xb1, 0x42, 0xbd, 0x0c, 0xc5, 0xd4, 0x5f, 0xea, 0x15, 0x28,
	0xa5, 0x3d, 0x9b, 0x65, 0x28, 0xc6, 0x15, 0x18, 0xf1, 0x1e, 0x0d, 0x9c, 0xef, 0x19, 0xb8, 0x35,
	0xd5, 0xb5, 0xb3, 0x87, 0x5e, 0x40, 0x2e, 0xbe, 0x11, 0x6d, 0xed, 0x2b, 0x5b, 0xeb, 0x33, 0xa3,
	0x69, 0x5b, 0x97, 0x18, 0xc1, 0x09, 0x80, 0x3a, 0x90, 0xf3, 0x79, 0xd8, 0x67, 0x03, 0x7b, 0x40,
	0x2b, 0x5b, 0x9b, 0x73, 0x12, 0xb1, 0xb4, 0xb9, 0x66, 0xfa, 0x6c, 0x30, 0x96, 0x8a, 0x05, 0x9c,
	0x6d, 0x28, 0xa5, 0xd7, 0x51, 0x01, 0x96, 0x8e, 0x77, 0x0e, 0x3a, 0xaf, 0xab, 0x0b, 0xa8, 0x08,
	0xf9, 0xce, 0x61, 0x6c, 0x64, 0xd0, 0x32, 0x54, 0xdf, 0xe1, 0xa3, 0xd6, 0x6e, 0xb7, 0xdb, 0x39,
	0xdc, 0xf3, 0x76, 0x31, 0x3e, 0xc2, 0xd5, 0x6c, 0xf3, 0xfd, 0xb7, 0x9f, 0xab, 0x99, 0x4f, 0xfb,
	0x73, 0x1f, 0x1e, 0x71, 0x3a, 0x98, 0x3c, 0x08, 0x97, 0xc2, 0x9c, 0xbe, 0x0f, 0xe6, 0x7e, 0x52,
	0x27, 0x39, 0xfb, 0x12, 0x3c, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x5a, 0x7b, 0x1c, 0xce,
	0x06, 0x00, 0x00,
}

func (this *MeshSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshSpec)
	if !ok {
		that2, ok := that.(MeshSpec)
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
	if that1.MeshType == nil {
		if this.MeshType != nil {
			return false
		}
	} else if this.MeshType == nil {
		return false
	} else if !this.MeshType.Equal(that1.MeshType) {
		return false
	}
	if !this.Cluster.Equal(that1.Cluster) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshSpec_Istio) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshSpec_Istio)
	if !ok {
		that2, ok := that.(MeshSpec_Istio)
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
	if !this.Istio.Equal(that1.Istio) {
		return false
	}
	return true
}
func (this *MeshSpec_AwsAppMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshSpec_AwsAppMesh)
	if !ok {
		that2, ok := that.(MeshSpec_AwsAppMesh)
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
	if !this.AwsAppMesh.Equal(that1.AwsAppMesh) {
		return false
	}
	return true
}
func (this *MeshSpec_Linkerd) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshSpec_Linkerd)
	if !ok {
		that2, ok := that.(MeshSpec_Linkerd)
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
	if !this.Linkerd.Equal(that1.Linkerd) {
		return false
	}
	return true
}
func (this *MeshSpec_ConsulConnect) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshSpec_ConsulConnect)
	if !ok {
		that2, ok := that.(MeshSpec_ConsulConnect)
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
	if !this.ConsulConnect.Equal(that1.ConsulConnect) {
		return false
	}
	return true
}
func (this *MeshStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshStatus)
	if !ok {
		that2, ok := that.(MeshStatus)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshInstallation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshInstallation)
	if !ok {
		that2, ok := that.(MeshInstallation)
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
	if this.InstallationNamespace != that1.InstallationNamespace {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *IstioMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IstioMesh)
	if !ok {
		that2, ok := that.(IstioMesh)
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
	if !this.Installation.Equal(that1.Installation) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AwsAppMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AwsAppMesh)
	if !ok {
		that2, ok := that.(AwsAppMesh)
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
	if !this.Installation.Equal(that1.Installation) {
		return false
	}
	if this.Region != that1.Region {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LinkerdMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LinkerdMesh)
	if !ok {
		that2, ok := that.(LinkerdMesh)
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
	if !this.Installation.Equal(that1.Installation) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ConsulConnectMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConsulConnectMesh)
	if !ok {
		that2, ok := that.(ConsulConnectMesh)
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
	if !this.Installation.Equal(that1.Installation) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshGroupSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshGroupSpec)
	if !ok {
		that2, ok := that.(MeshGroupSpec)
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
	if this.DisplayName != that1.DisplayName {
		return false
	}
	if len(this.Meshes) != len(that1.Meshes) {
		return false
	}
	for i := range this.Meshes {
		if !this.Meshes[i].Equal(that1.Meshes[i]) {
			return false
		}
	}
	if !this.TrustBundleRef.Equal(that1.TrustBundleRef) {
		return false
	}
	if that1.TrustModel == nil {
		if this.TrustModel != nil {
			return false
		}
	} else if this.TrustModel == nil {
		return false
	} else if !this.TrustModel.Equal(that1.TrustModel) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshGroupSpec_Shared) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshGroupSpec_Shared)
	if !ok {
		that2, ok := that.(MeshGroupSpec_Shared)
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
	if !this.Shared.Equal(that1.Shared) {
		return false
	}
	return true
}
func (this *MeshGroupSpec_Limited) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshGroupSpec_Limited)
	if !ok {
		that2, ok := that.(MeshGroupSpec_Limited)
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
	if !this.Limited.Equal(that1.Limited) {
		return false
	}
	return true
}
func (this *MeshGroupSpec_SharedTrust) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshGroupSpec_SharedTrust)
	if !ok {
		that2, ok := that.(MeshGroupSpec_SharedTrust)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshGroupSpec_LimitedTrust) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshGroupSpec_LimitedTrust)
	if !ok {
		that2, ok := that.(MeshGroupSpec_LimitedTrust)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshGroupStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshGroupStatus)
	if !ok {
		that2, ok := that.(MeshGroupStatus)
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
	if this.Health != that1.Health {
		return false
	}
	if this.Config != that1.Config {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
