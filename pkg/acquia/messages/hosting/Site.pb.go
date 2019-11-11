// Code generated by protoc-gen-go. DO NOT EDIT.
// source: acquia/messages/hosting/Site.proto

package hosting

import (
	fmt "fmt"
	shared "github.com/acquia/message-broker-sdk-go/pkg/acquia/messages/shared"
	proto "github.com/golang/protobuf/proto"
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

type CreateSite struct {
	Header               *shared.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	EnvUUID              string         `protobuf:"bytes,2,opt,name=envUUID,proto3" json:"envUUID,omitempty"`
	SiteUUID             string         `protobuf:"bytes,3,opt,name=siteUUID,proto3" json:"siteUUID,omitempty"`
	SiteLabel            string         `protobuf:"bytes,4,opt,name=siteLabel,proto3" json:"siteLabel,omitempty"`
	SiteName             string         `protobuf:"bytes,5,opt,name=siteName,proto3" json:"siteName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateSite) Reset()         { *m = CreateSite{} }
func (m *CreateSite) String() string { return proto.CompactTextString(m) }
func (*CreateSite) ProtoMessage()    {}
func (*CreateSite) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73279a37a5f57fc, []int{0}
}

func (m *CreateSite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSite.Unmarshal(m, b)
}
func (m *CreateSite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSite.Marshal(b, m, deterministic)
}
func (m *CreateSite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSite.Merge(m, src)
}
func (m *CreateSite) XXX_Size() int {
	return xxx_messageInfo_CreateSite.Size(m)
}
func (m *CreateSite) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSite.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSite proto.InternalMessageInfo

func (m *CreateSite) GetHeader() *shared.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CreateSite) GetEnvUUID() string {
	if m != nil {
		return m.EnvUUID
	}
	return ""
}

func (m *CreateSite) GetSiteUUID() string {
	if m != nil {
		return m.SiteUUID
	}
	return ""
}

func (m *CreateSite) GetSiteLabel() string {
	if m != nil {
		return m.SiteLabel
	}
	return ""
}

func (m *CreateSite) GetSiteName() string {
	if m != nil {
		return m.SiteName
	}
	return ""
}

type CreateSiteResponse struct {
	Header               *shared.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	SiteUUID             string         `protobuf:"bytes,2,opt,name=siteUUID,proto3" json:"siteUUID,omitempty"`
	Status               string         `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateSiteResponse) Reset()         { *m = CreateSiteResponse{} }
func (m *CreateSiteResponse) String() string { return proto.CompactTextString(m) }
func (*CreateSiteResponse) ProtoMessage()    {}
func (*CreateSiteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73279a37a5f57fc, []int{1}
}

func (m *CreateSiteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSiteResponse.Unmarshal(m, b)
}
func (m *CreateSiteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSiteResponse.Marshal(b, m, deterministic)
}
func (m *CreateSiteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSiteResponse.Merge(m, src)
}
func (m *CreateSiteResponse) XXX_Size() int {
	return xxx_messageInfo_CreateSiteResponse.Size(m)
}
func (m *CreateSiteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSiteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSiteResponse proto.InternalMessageInfo

func (m *CreateSiteResponse) GetHeader() *shared.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CreateSiteResponse) GetSiteUUID() string {
	if m != nil {
		return m.SiteUUID
	}
	return ""
}

func (m *CreateSiteResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

// NW-410
type CopyDB struct {
	Header       *shared.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	SrcSiteUUID  string         `protobuf:"bytes,2,opt,name=srcSiteUUID,proto3" json:"srcSiteUUID,omitempty"`
	DestSiteUUID string         `protobuf:"bytes,3,opt,name=destSiteUUID,proto3" json:"destSiteUUID,omitempty"`
	// Product level name, which is currently role.
	DatabaseNames        []string `protobuf:"bytes,4,rep,name=databaseNames,proto3" json:"databaseNames,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CopyDB) Reset()         { *m = CopyDB{} }
func (m *CopyDB) String() string { return proto.CompactTextString(m) }
func (*CopyDB) ProtoMessage()    {}
func (*CopyDB) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73279a37a5f57fc, []int{2}
}

func (m *CopyDB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CopyDB.Unmarshal(m, b)
}
func (m *CopyDB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CopyDB.Marshal(b, m, deterministic)
}
func (m *CopyDB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CopyDB.Merge(m, src)
}
func (m *CopyDB) XXX_Size() int {
	return xxx_messageInfo_CopyDB.Size(m)
}
func (m *CopyDB) XXX_DiscardUnknown() {
	xxx_messageInfo_CopyDB.DiscardUnknown(m)
}

var xxx_messageInfo_CopyDB proto.InternalMessageInfo

func (m *CopyDB) GetHeader() *shared.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CopyDB) GetSrcSiteUUID() string {
	if m != nil {
		return m.SrcSiteUUID
	}
	return ""
}

func (m *CopyDB) GetDestSiteUUID() string {
	if m != nil {
		return m.DestSiteUUID
	}
	return ""
}

func (m *CopyDB) GetDatabaseNames() []string {
	if m != nil {
		return m.DatabaseNames
	}
	return nil
}

type CopyFiles struct {
	Header               *shared.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	SrcSiteUUID          string         `protobuf:"bytes,2,opt,name=srcSiteUUID,proto3" json:"srcSiteUUID,omitempty"`
	DestSiteUUID         string         `protobuf:"bytes,3,opt,name=destSiteUUID,proto3" json:"destSiteUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CopyFiles) Reset()         { *m = CopyFiles{} }
func (m *CopyFiles) String() string { return proto.CompactTextString(m) }
func (*CopyFiles) ProtoMessage()    {}
func (*CopyFiles) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73279a37a5f57fc, []int{3}
}

func (m *CopyFiles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CopyFiles.Unmarshal(m, b)
}
func (m *CopyFiles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CopyFiles.Marshal(b, m, deterministic)
}
func (m *CopyFiles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CopyFiles.Merge(m, src)
}
func (m *CopyFiles) XXX_Size() int {
	return xxx_messageInfo_CopyFiles.Size(m)
}
func (m *CopyFiles) XXX_DiscardUnknown() {
	xxx_messageInfo_CopyFiles.DiscardUnknown(m)
}

var xxx_messageInfo_CopyFiles proto.InternalMessageInfo

func (m *CopyFiles) GetHeader() *shared.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CopyFiles) GetSrcSiteUUID() string {
	if m != nil {
		return m.SrcSiteUUID
	}
	return ""
}

func (m *CopyFiles) GetDestSiteUUID() string {
	if m != nil {
		return m.DestSiteUUID
	}
	return ""
}

// NW-411
type RestoreDB struct {
	Header   *shared.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	SiteUUID string         `protobuf:"bytes,2,opt,name=siteUUID,proto3" json:"siteUUID,omitempty"`
	// ?? Don't know how to refer to this.
	BackupUrl string `protobuf:"bytes,3,opt,name=backupUrl,proto3" json:"backupUrl,omitempty"`
	// Product level name, which is currently role.
	DatabaseName         string   `protobuf:"bytes,4,opt,name=databaseName,proto3" json:"databaseName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestoreDB) Reset()         { *m = RestoreDB{} }
func (m *RestoreDB) String() string { return proto.CompactTextString(m) }
func (*RestoreDB) ProtoMessage()    {}
func (*RestoreDB) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73279a37a5f57fc, []int{4}
}

func (m *RestoreDB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestoreDB.Unmarshal(m, b)
}
func (m *RestoreDB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestoreDB.Marshal(b, m, deterministic)
}
func (m *RestoreDB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestoreDB.Merge(m, src)
}
func (m *RestoreDB) XXX_Size() int {
	return xxx_messageInfo_RestoreDB.Size(m)
}
func (m *RestoreDB) XXX_DiscardUnknown() {
	xxx_messageInfo_RestoreDB.DiscardUnknown(m)
}

var xxx_messageInfo_RestoreDB proto.InternalMessageInfo

func (m *RestoreDB) GetHeader() *shared.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *RestoreDB) GetSiteUUID() string {
	if m != nil {
		return m.SiteUUID
	}
	return ""
}

func (m *RestoreDB) GetBackupUrl() string {
	if m != nil {
		return m.BackupUrl
	}
	return ""
}

func (m *RestoreDB) GetDatabaseName() string {
	if m != nil {
		return m.DatabaseName
	}
	return ""
}

// NW-413
type CreateDBBackupUrl struct {
	Header   *shared.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	SiteUUID string         `protobuf:"bytes,2,opt,name=siteUUID,proto3" json:"siteUUID,omitempty"`
	// Product level name, which is currently role.
	DatabaseName         string   `protobuf:"bytes,3,opt,name=databaseName,proto3" json:"databaseName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDBBackupUrl) Reset()         { *m = CreateDBBackupUrl{} }
func (m *CreateDBBackupUrl) String() string { return proto.CompactTextString(m) }
func (*CreateDBBackupUrl) ProtoMessage()    {}
func (*CreateDBBackupUrl) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73279a37a5f57fc, []int{5}
}

func (m *CreateDBBackupUrl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDBBackupUrl.Unmarshal(m, b)
}
func (m *CreateDBBackupUrl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDBBackupUrl.Marshal(b, m, deterministic)
}
func (m *CreateDBBackupUrl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDBBackupUrl.Merge(m, src)
}
func (m *CreateDBBackupUrl) XXX_Size() int {
	return xxx_messageInfo_CreateDBBackupUrl.Size(m)
}
func (m *CreateDBBackupUrl) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDBBackupUrl.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDBBackupUrl proto.InternalMessageInfo

func (m *CreateDBBackupUrl) GetHeader() *shared.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CreateDBBackupUrl) GetSiteUUID() string {
	if m != nil {
		return m.SiteUUID
	}
	return ""
}

func (m *CreateDBBackupUrl) GetDatabaseName() string {
	if m != nil {
		return m.DatabaseName
	}
	return ""
}

// NW-414
type AddSiteDomain struct {
	Header               *shared.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	SiteUUID             string         `protobuf:"bytes,2,opt,name=siteUUID,proto3" json:"siteUUID,omitempty"`
	Domain               string         `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddSiteDomain) Reset()         { *m = AddSiteDomain{} }
func (m *AddSiteDomain) String() string { return proto.CompactTextString(m) }
func (*AddSiteDomain) ProtoMessage()    {}
func (*AddSiteDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73279a37a5f57fc, []int{6}
}

func (m *AddSiteDomain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSiteDomain.Unmarshal(m, b)
}
func (m *AddSiteDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSiteDomain.Marshal(b, m, deterministic)
}
func (m *AddSiteDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSiteDomain.Merge(m, src)
}
func (m *AddSiteDomain) XXX_Size() int {
	return xxx_messageInfo_AddSiteDomain.Size(m)
}
func (m *AddSiteDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSiteDomain.DiscardUnknown(m)
}

var xxx_messageInfo_AddSiteDomain proto.InternalMessageInfo

func (m *AddSiteDomain) GetHeader() *shared.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AddSiteDomain) GetSiteUUID() string {
	if m != nil {
		return m.SiteUUID
	}
	return ""
}

func (m *AddSiteDomain) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

// NW-415
type RemoveSiteDomain struct {
	Header               *shared.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	SiteUUID             string         `protobuf:"bytes,2,opt,name=siteUUID,proto3" json:"siteUUID,omitempty"`
	Domain               string         `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RemoveSiteDomain) Reset()         { *m = RemoveSiteDomain{} }
func (m *RemoveSiteDomain) String() string { return proto.CompactTextString(m) }
func (*RemoveSiteDomain) ProtoMessage()    {}
func (*RemoveSiteDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73279a37a5f57fc, []int{7}
}

func (m *RemoveSiteDomain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveSiteDomain.Unmarshal(m, b)
}
func (m *RemoveSiteDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveSiteDomain.Marshal(b, m, deterministic)
}
func (m *RemoveSiteDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveSiteDomain.Merge(m, src)
}
func (m *RemoveSiteDomain) XXX_Size() int {
	return xxx_messageInfo_RemoveSiteDomain.Size(m)
}
func (m *RemoveSiteDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveSiteDomain.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveSiteDomain proto.InternalMessageInfo

func (m *RemoveSiteDomain) GetHeader() *shared.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *RemoveSiteDomain) GetSiteUUID() string {
	if m != nil {
		return m.SiteUUID
	}
	return ""
}

func (m *RemoveSiteDomain) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateSite)(nil), "acquia.messages.hosting.CreateSite")
	proto.RegisterType((*CreateSiteResponse)(nil), "acquia.messages.hosting.CreateSiteResponse")
	proto.RegisterType((*CopyDB)(nil), "acquia.messages.hosting.CopyDB")
	proto.RegisterType((*CopyFiles)(nil), "acquia.messages.hosting.CopyFiles")
	proto.RegisterType((*RestoreDB)(nil), "acquia.messages.hosting.RestoreDB")
	proto.RegisterType((*CreateDBBackupUrl)(nil), "acquia.messages.hosting.CreateDBBackupUrl")
	proto.RegisterType((*AddSiteDomain)(nil), "acquia.messages.hosting.AddSiteDomain")
	proto.RegisterType((*RemoveSiteDomain)(nil), "acquia.messages.hosting.RemoveSiteDomain")
}

func init() { proto.RegisterFile("acquia/messages/hosting/Site.proto", fileDescriptor_f73279a37a5f57fc) }

var fileDescriptor_f73279a37a5f57fc = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x99, 0xed, 0x5a, 0xcd, 0x5b, 0x17, 0x74, 0x0e, 0x1a, 0x96, 0x45, 0xca, 0xe8, 0xa1,
	0x97, 0x26, 0xa0, 0xe0, 0xdd, 0xb6, 0x8a, 0x82, 0x78, 0x48, 0xe9, 0xc5, 0xdb, 0xa4, 0xf3, 0x48,
	0x43, 0x9b, 0x4e, 0x9c, 0x37, 0x5d, 0x10, 0x41, 0xbc, 0x8a, 0x9f, 0xc2, 0xbb, 0x47, 0x3f, 0xa0,
	0x64, 0x32, 0xd9, 0x34, 0x41, 0x3d, 0xec, 0xb2, 0xec, 0xa9, 0x7d, 0xff, 0x79, 0xff, 0xf7, 0x7e,
	0xbc, 0x3c, 0x1e, 0x08, 0xb9, 0xfa, 0xb4, 0xcf, 0x65, 0x5c, 0x20, 0x91, 0xcc, 0x90, 0xe2, 0xb5,
	0x26, 0x9b, 0xef, 0xb2, 0x78, 0x91, 0x5b, 0x8c, 0x4a, 0xa3, 0xad, 0xe6, 0x8f, 0xeb, 0x9c, 0xa8,
	0xc9, 0x89, 0x7c, 0xce, 0xd9, 0xd3, 0xbe, 0x99, 0xd6, 0xd2, 0xa0, 0xf2, 0x3f, 0xb5, 0x5b, 0xfc,
	0x66, 0x00, 0x33, 0x83, 0xd2, 0x62, 0x55, 0x92, 0xbf, 0x84, 0xe1, 0x1a, 0xa5, 0x42, 0x13, 0xb2,
	0x11, 0x1b, 0x9f, 0x3c, 0x7f, 0x12, 0xf5, 0xab, 0x7b, 0xf7, 0x5b, 0x97, 0x95, 0xf8, 0x6c, 0x1e,
	0xc2, 0x5d, 0xdc, 0x5d, 0x2c, 0x97, 0xef, 0xe6, 0xe1, 0xd1, 0x88, 0x8d, 0x83, 0xa4, 0x09, 0xf9,
	0x19, 0xdc, 0xa3, 0xdc, 0xa2, 0x7b, 0x1a, 0xb8, 0xa7, 0xcb, 0x98, 0x9f, 0x43, 0x50, 0xfd, 0x7f,
	0x2f, 0x53, 0xdc, 0x86, 0xc7, 0xee, 0xb1, 0x15, 0x1a, 0xe7, 0x07, 0x59, 0x60, 0x78, 0xa7, 0x75,
	0x56, 0xb1, 0xf8, 0xc6, 0x80, 0xb7, 0xd8, 0x09, 0x52, 0xa9, 0x77, 0x74, 0x75, 0xfc, 0x43, 0xc8,
	0xa3, 0x1e, 0xe4, 0x23, 0x18, 0x92, 0x95, 0x76, 0x4f, 0x1e, 0xdf, 0x47, 0xe2, 0x17, 0x83, 0xe1,
	0x4c, 0x97, 0x9f, 0xe7, 0xd3, 0x2b, 0xb7, 0x1d, 0xc1, 0x09, 0x99, 0xd5, 0xa2, 0xdb, 0xf9, 0x50,
	0xe2, 0x02, 0xee, 0x2b, 0x24, 0xbb, 0xe8, 0x4e, 0xb0, 0xa3, 0xf1, 0x67, 0x70, 0xaa, 0xa4, 0x95,
	0xa9, 0x24, 0x37, 0x1b, 0x0a, 0x8f, 0x47, 0x83, 0x71, 0x90, 0x74, 0x45, 0xf1, 0x9d, 0x41, 0x50,
	0xe1, 0xbe, 0xc9, 0xb7, 0x48, 0xb7, 0x4b, 0x2c, 0x7e, 0x32, 0x08, 0x12, 0x24, 0xab, 0x0d, 0x5e,
	0x63, 0x7a, 0xff, 0xfb, 0x68, 0xe7, 0x10, 0xa4, 0x72, 0xb5, 0xd9, 0x97, 0x4b, 0xb3, 0xf5, 0x08,
	0xad, 0xe0, 0x18, 0x0f, 0x86, 0xe3, 0x57, 0xaf, 0xa3, 0x89, 0x1f, 0x0c, 0x1e, 0xd6, 0x1b, 0x36,
	0x9f, 0x4e, 0x2f, 0x9d, 0x37, 0xc1, 0xda, 0xa7, 0x19, 0xfc, 0x85, 0xe6, 0x0b, 0x9c, 0xbe, 0x52,
	0xaa, 0x1a, 0xe0, 0x5c, 0x17, 0x32, 0xdf, 0xdd, 0xd4, 0xa6, 0x2b, 0x57, 0xbd, 0xd9, 0xf4, 0x3a,
	0x12, 0x5f, 0xe1, 0x41, 0x82, 0x85, 0xbe, 0xc0, 0xdb, 0xe9, 0x3f, 0x7d, 0xfd, 0x71, 0x96, 0xe5,
	0x76, 0xbd, 0x4f, 0xa3, 0x95, 0x2e, 0xe2, 0xee, 0x55, 0x9b, 0xa4, 0x46, 0x6f, 0xd0, 0x4c, 0x48,
	0x6d, 0x26, 0x99, 0x8e, 0xcb, 0x4d, 0x16, 0xff, 0xe3, 0x66, 0xa6, 0x43, 0x77, 0xf1, 0x5e, 0xfc,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xc5, 0xf1, 0xe1, 0x55, 0x05, 0x00, 0x00,
}
